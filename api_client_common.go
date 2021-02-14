package bigdog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
)

// todo: export this?
type APIAuthProviderError struct {
	meta APIResponseMeta
	err  error
}

func NewAPIAuthProviderError(meta APIResponseMeta, err error) *APIAuthProviderError {
	ape := new(APIAuthProviderError)
	ape.meta = meta
	ape.err = err
	return ape
}

func (ape *APIAuthProviderError) ResponseMeta() APIResponseMeta {
	if ape == nil {
		return APIResponseMeta{}
	}
	return ape.meta
}

func (ape *APIAuthProviderError) Is(err error) bool {
	return ape != nil && ape.err != nil && errors.Is(err, ape.err)
}

func (ape *APIAuthProviderError) Unwrap() error {
	if ape == nil || ape.err == nil {
		return nil
	}
	return ape.err
}

func (ape *APIAuthProviderError) Error() string {
	if ape == nil || ape.err == nil {
		return ""
	}
	return ape.err.Error()
}

type baseClient struct {
	log   *log.Logger
	debug bool

	addr       string
	pathPrefix string

	client *http.Client

	autoHydrate bool
}

func newBaseClient(addr, pathPrefix string, dbg bool, disableAutoHydrate bool, logger *log.Logger, hclient *http.Client) *baseClient {
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
		// pooled transport config shamelessly borrowed from
		// https://github.com/hashicorp/go-cleanhttp/blob/v0.5.1/cleanhttp.go
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			},
		}
	}

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

func handleAPIResponse(req *APIRequest, successCode int, httpResp *http.Response, execDur time.Duration, respFact APIResponseFactory, autoHydrate bool, sourceErr error) (APIResponse, error) {
	if sourceErr != nil {
		// handle some context errors

		// for deadline exceeded errors, return a timeout status code
		if errors.Is(sourceErr, context.DeadlineExceeded) {
			if httpResp != nil {
				cleanupReadCloser(httpResp.Body)
			}
			return respFact(req.Source, newAPIResponseMetaWithCode(req, successCode, http.StatusRequestTimeout, execDur), nil), sourceErr
		}

		// for context cancelled errors, return nginx's 499 Client Closed Request status
		// 499 : https://httpstatuses.com/499
		if errors.Is(sourceErr, context.Canceled) {
			if httpResp != nil {
				cleanupReadCloser(httpResp.Body)
			}
			return respFact(req.Source, newAPIResponseMetaWithCode(req, successCode, 499, execDur), nil), sourceErr
		}

		// if the incoming error is from an auth provider, return as-is
		if aerr, ok := sourceErr.(*APIAuthProviderError); ok && aerr != nil {
			if httpResp != nil {
				cleanupReadCloser(httpResp.Body)
			}
			return respFact(req.Source, aerr.ResponseMeta(), nil), aerr
		}

		// perform one final check here as we sometimes see an http resp with an error carrying an invalid status code
		if httpResp != nil && httpResp.StatusCode < 100 {
			cleanupReadCloser(httpResp.Body)
			return respFact(req.Source, newAPIResponseMetaWithCode(req, successCode, http.StatusInternalServerError, execDur), nil),
				fmt.Errorf("received invalid response code %d: %w", http.StatusInternalServerError, sourceErr)
		}

		// otherwise, pass on constructed response and incoming error
		return respFact(req.Source, newAPIResponseMeta(req, successCode, httpResp, execDur), nil), sourceErr
	}

	// this _should_ never happen, but check for it anyway.
	if httpResp == nil {
		panic(fmt.Sprintf("severe problem: nil *http.Response seen with nil error. meta: %v", newAPIResponseMeta(req, successCode, httpResp, execDur)))
	}

	// construct response
	apiResp := respFact(req.Source, newAPIResponseMeta(req, successCode, httpResp, execDur), httpResp.Body)

	// if the response code matches the expected "success" code...
	if httpResp.StatusCode == successCode {
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

	if err := json.NewDecoder(httpResp.Body).Decode(apiErr); err != nil {
		return apiResp, fmt.Errorf("error unmarshalling error body into %T: %w", apiErr, err)
	}

	return apiResp, apiErr
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
