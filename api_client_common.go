package bigdog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type baseClient struct {
	log   *log.Logger
	debug bool

	addr       string
	pathPrefix string

	client *http.Client

	autoHydrate bool
	ev          APIClientEventListener
}

func newBaseClient(addr, pathPrefix string, dbg bool, disableAutoHydrate bool, logger *log.Logger, hclient *http.Client, ev APIClientEventListener) *baseClient {
	c := new(baseClient)

	if logger == nil {
		c.log = log.New(ioutil.Discard, "", 0)
	} else {
		c.log = logger
	}

	c.debug = dbg
	c.addr = addr
	c.pathPrefix = pathPrefix
	c.autoHydrate = !disableAutoHydrate

	if hclient != nil {
		c.client = hclient
	} else {
		c.client = defaultHTTPClient()
	}

	c.ev = ev

	return c
}

// VSZAddress returns the address of the VSZ node currently being connected to
func (c *baseClient) Address() string {
	return c.addr
}

// VSZPathPrefix returns the prefix that will be applied to all api calls.  It is valid for this to be blank, and is by
// default
func (c *baseClient) PathPrefix() string {
	return c.pathPrefix
}

func (c *baseClient) do(ctx context.Context, request *APIRequest, authParamName, authParamValue string, mutators ...RequestMutator) (*http.Response, error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		err          error
	)

	if c.debug {
		logMsg := fmt.Sprintf(logDebugAPIRequestPrepFormat, request.ID(), request.Method, request.CompiledURI())

		if request.Body() != nil {
			logMsg = fmt.Sprintf(logDebugAPIRequestWithBodyFormat, logMsg)
		} else {
			logMsg = fmt.Sprintf(logDebugAPIRequestNoBodyFormat, logMsg)
		}

		c.log.Print(logMsg)
	}
	if httpRequest, err = request.ToHTTP(ctx, c.addr, c.pathPrefix, authParamName, authParamValue); err != nil {
		return nil, err
	}

	if ml := len(mutators); ml > 0 {
		if c.debug {
			c.log.Printf("Executing %d mutators...", ml)
		}
		for _, fn := range mutators {
			fn(httpRequest)
		}
	}

	httpResponse, err = c.client.Do(httpRequest)
	return httpResponse, err
}

func cleanupReadCloser(rc io.ReadCloser) {
	if rc == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, rc)
	_ = rc.Close()
}

func handleFailedRequest(req *APIRequest, successCode int, httpResp *http.Response, execDur time.Duration, sourceErr error) (APIResponseMeta, error) {
	defer func() {
		if httpResp != nil {
			cleanupReadCloser(httpResp.Body)
		}
	}()

	// for deadline exceeded errors, return a timeout status code
	if errors.Is(sourceErr, context.DeadlineExceeded) {
		return newAPIResponseMetaWithCode(req, successCode, http.StatusRequestTimeout, execDur), sourceErr
	}

	// for context cancelled errors, return nginx's 499 Client Closed Request status
	// 499 : https://httpstatuses.com/499
	if errors.Is(sourceErr, context.Canceled) {
		return newAPIResponseMetaWithCode(req, successCode, 499, execDur), sourceErr
	}

	// if the incoming error is from an auth provider, return as-is
	if aerr, ok := sourceErr.(*APIAuthProviderError); ok && aerr != nil {
		return aerr.ResponseMeta(), aerr
	}

	// perform one final check here as we sometimes see an http resp with an error carrying an invalid status code
	if httpResp != nil && httpResp.StatusCode < 100 {
		return newAPIResponseMetaWithCode(req, successCode, http.StatusInternalServerError, execDur),
			fmt.Errorf("received invalid response code %d: %w", http.StatusInternalServerError, sourceErr)
	}

	// otherwise, pass on constructed response and incoming error
	return newAPIResponseMeta(req, successCode, httpResp, execDur), sourceErr
}

func handleCompletedRequest(
	req *APIRequest,
	successCode int,
	httpResp *http.Response,
	execDur time.Duration,
	respFact apiResponseFactory,
	autoHydrate bool,
) (APIResponse, error) {
	// if the response code matches the expected "success" code...
	if httpResp.StatusCode == successCode {
		// construct response
		apiResp := respFact(req.Source, newAPIResponseMeta(req, successCode, httpResp, execDur), httpResp.Body)

		// test for a modeled response and whether it needs to be automatically handled
		if hdr, ok := apiResp.(ModeledAPIResponse); ok && autoHydrate {
			if _, err := hdr.Hydrate(); err != nil {
				return apiResp, fmt.Errorf("error hydrating response: %w", err)
			}
		}
		// if not, simply return response as-is
		return apiResp, nil
	}

	// otherwise, handle api error response
	var apiErr error

	// todo: do something better here.
	if req.Source == APISourceVSZ {
		apiErr = new(VSZAPIError)
	} else {
		apiErr = new(SCIAPIError)
	}

	// attempt to unmarshal response body into error type
	err := json.NewDecoder(httpResp.Body).Decode(apiErr)
	if err != nil {
		apiErr = fmt.Errorf("error unmarshalling error body into %T: %w", apiErr, err)
	}

	// return api response type and error
	// todo: not the biggest fan of passing off a now-defunct body...
	return respFact(req.Source, newAPIResponseMeta(req, successCode, httpResp, execDur), httpResp.Body), apiErr
}

// handleAPIResponse does just that, but i dunno that i dig it.
func handleAPIResponse(
	req *APIRequest,
	successCode int,
	httpResp *http.Response,
	execDur time.Duration,
	respFact apiResponseFactory,
	autoHydrate bool,
	ev APIClientEventListener,
	sourceErr error,
) (APIResponse, error) {
	var (
		apiResp APIResponse
		meta    APIResponseMeta
		err     error
	)

	if sourceErr != nil {
		// handle failed request errors
		meta, err = handleFailedRequest(req, successCode, httpResp, execDur, sourceErr)
		apiResp = respFact(req.Source, meta, nil)
	} else if httpResp == nil {
		// this _should_ never happen, but check for it anyway.
		panic(fmt.Sprintf("severe problem: nil *http.Response seen with nil error. meta: %v", newAPIResponseMeta(req, successCode, httpResp, execDur)))
	} else {
		// handle completed, but not necessarily successful, requests
		apiResp, err = handleCompletedRequest(req, successCode, httpResp, execDur, respFact, autoHydrate)
		//goland:noinspection GoNilness
		meta = apiResp.ResponseMeta()
	}

	// if an event handler was provided...
	if ev != nil {
		go ev.Push(meta)
	}

	return apiResp, err
}

type WSGService struct {
	apiClient *VSZClient
}

func NewWSGService(c *VSZClient) *WSGService {
	s := new(WSGService)
	s.apiClient = c
	return s
}

type SwitchMService struct {
	apiClient *VSZClient
}

func NewSwitchMService(c *VSZClient) *SwitchMService {
	s := new(SwitchMService)
	s.apiClient = c
	return s
}

type SCIService struct {
	apiClient *SCIClient
}

func NewSCIService(c *SCIClient) *SCIService {
	s := new(SCIService)
	s.apiClient = c
	return s
}

func SCIFilterFormatTimestamp(t time.Time) string {
	return t.Format(SCIFilterTimestampFormat)
}
