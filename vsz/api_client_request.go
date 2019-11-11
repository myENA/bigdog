package vsz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
)

const (
	uriPathParameterSearchFormat      = "{%s}"
	uriQueryParameterPrefixFormat     = "%s?"
	uriQueryParameterAddNoValueFormat = "%s%s&"
	uriQueryParameterAddValueFormat   = "%s%s=%s&"
	uriQueryParameterCutset           = "&"

	apiRequestURLFormat = "https://%s:%d%s"

	headerKeyContentType       = "Content-Type"
	headerKeyAccept            = "Accept"
	headerValueApplicationJSON = "application/json"

	logDebugAPIRequestPrepFormat     = "[request-%d] Preparing api request \"%s %s\""
	logDebugAPIRequestNoBodyFormat   = "%s without body"
	logDebugAPIRequestWithBodyFormat = "%s with body: %s"
)

var apiRequestID uint64

type APIRequest struct {
	id            uint64
	method        string
	uri           string
	authenticated bool

	compiledURI string

	queryParameters map[string][]string
	pathParameters  map[string]string
	headers         url.Values
	cookies         []*http.Cookie
	body            []byte
}

func NewAPIRequest(method, uri string, authenticated bool) *APIRequest {
	r := &APIRequest{
		id:              atomic.AddUint64(&apiRequestID, 1),
		method:          method,
		uri:             uri,
		authenticated:   authenticated,
		queryParameters: make(map[string][]string),
		pathParameters:  make(map[string]string),
		headers:         make(url.Values),
		cookies:         make([]*http.Cookie, 0),
	}
	return r
}

func (r *APIRequest) ID() uint64 {
	return r.id
}

func (r *APIRequest) Method() string {
	return r.method
}

func (r *APIRequest) URI() string {
	return r.uri
}

func (r *APIRequest) Authenticated() bool {
	return r.authenticated
}

func (r *APIRequest) SetHeaders(headers url.Values) {
	r.headers = headers
}

func (r *APIRequest) AddHeader(name, value string) {
	r.headers.Add(name, value)
}

func (r *APIRequest) SetHeader(name, value string) {
	r.headers.Set(name, value)
}

func (r *APIRequest) RemoveHeader(name string) {
	r.headers.Del(name)
}

func (r *APIRequest) Headers() url.Values {
	return r.headers
}

func (r *APIRequest) SetCookies(cookies []*http.Cookie) {
	r.cookies = cookies
}

func (r *APIRequest) AddCookie(cookie *http.Cookie) {
	r.cookies = append(r.cookies, cookie)
}

func (r *APIRequest) SetCookie(cookie *http.Cookie) {
	for i, cc := range r.cookies {
		if cc.Name == cookie.Name {
			r.cookies[i] = cookie
			return
		}
	}
	r.cookies = append(r.cookies, cookie)
}

func (r *APIRequest) RemoveCookie(name string) {
	nc := make([]*http.Cookie, 0)
	for _, cc := range r.cookies {
		if cc.Name != name {
			nc = append(nc, cc)
		}
	}
	r.cookies = nc
}

func (r *APIRequest) Cookies() []*http.Cookie {
	return r.cookies
}

// SetQueryParameters will override any / all existing query parameters
func (r *APIRequest) SetQueryParameters(params map[string][]string) {
	r.queryParameters = params
	r.compiledURI = ""
}

// SetQueryParameter will set a query param to a specific value, overriding any previously set value
func (r *APIRequest) SetQueryParameter(param, value string) {
	r.queryParameters[param] = []string{value}
	r.compiledURI = ""
}

// AddQueryParameter will add a value to the specified param
func (r *APIRequest) AddQueryParameter(param, value string) {
	if _, ok := r.queryParameters[param]; !ok {
		r.queryParameters[param] = make([]string, 0)
	}
	r.queryParameters[param] = append(r.queryParameters[param], value)
	r.compiledURI = ""
}

// RemoveQueryParameter will attempt to delete all values for a specific query parameter from this request.
func (r *APIRequest) RemoveQueryParameter(param string) {
	delete(r.queryParameters, param)
	r.compiledURI = ""
}

// QueryParameters will return all values of currently set query parameters
func (r *APIRequest) QueryParameters() map[string][]string {
	return r.queryParameters
}

// SetPathParameters will re-define all path parameters, overriding any / all existing ones
func (r *APIRequest) SetPathParameters(params map[string]string) {
	r.pathParameters = params
	r.compiledURI = ""
}

// SetPathParameter will define a path parameter value, overriding any existing value
func (r *APIRequest) SetPathParameter(param, value string) {
	r.pathParameters[param] = value
	r.compiledURI = ""
}

// RemovePathParameter will attempt to remove a single parameter from the current list of path parameters
func (r *APIRequest) RemovePathParameter(param string) {
	delete(r.pathParameters, param)
	r.compiledURI = ""
}

func (r *APIRequest) PathParameters() map[string]string {
	return r.pathParameters
}

func (r *APIRequest) SetBody(body []byte) {
	r.body = body
}

func (r *APIRequest) SetBodyModel(model interface{}) error {
	if model == nil {
		r.body = nil
		return nil
	}
	var err error
	if r.body, err = json.Marshal(model); err != nil {
		return err
	}
	return nil
}

func (r *APIRequest) Body() []byte {
	return r.body
}

// CompiledURI will return to you the full request URI, not including scheme, hostname, and port.  This method is not
// thread safe, as you shouldn't be calling this asynchronously anyway.
func (r *APIRequest) CompiledURI() string {
	if r.compiledURI == "" {
		pathParams := r.PathParameters()
		queryParams := r.QueryParameters()
		uri := r.uri
		if len(pathParams) > 0 {
			for k, v := range pathParams {
				uri = strings.Replace(uri, fmt.Sprintf(uriPathParameterSearchFormat, k), v, 1)
			}
		}
		// TODO: could probably be made more efficient.
		if len(queryParams) > 0 {
			uri = fmt.Sprintf(uriQueryParameterPrefixFormat, uri)
			for param, values := range queryParams {
				for _, value := range values {
					if value == "" {
						uri = fmt.Sprintf(uriQueryParameterAddNoValueFormat, uri, param)
					} else {
						uri = fmt.Sprintf(uriQueryParameterAddValueFormat, uri, param, value)
					}
				}
			}
			uri = strings.TrimRight(uri, uriQueryParameterCutset)
		}
		r.compiledURI = uri
	}
	return r.compiledURI
}

// toHTTP will attempt to construct an executable http.request
func (r *APIRequest) toHTTP(ctx context.Context, addr string, port int, debug bool) (*http.Request, error) {
	var err error
	var httpRequest *http.Request

	body := r.Body()
	bodyLen := len(body)
	compiledURL := fmt.Sprintf(apiRequestURLFormat, addr, port, r.CompiledURI())

	// if debug mode is enabled, prepare a big'ol log statement.
	if debug {
		logMsg := fmt.Sprintf(logDebugAPIRequestPrepFormat, r.id, r.method, compiledURL)

		if bodyLen == 0 {
			logMsg = fmt.Sprintf(logDebugAPIRequestNoBodyFormat, logMsg)
		} else {
			logMsg = fmt.Sprintf(logDebugAPIRequestWithBodyFormat, logMsg, string(body))
		}

		log.Print(logMsg)
	}

	if bodyLen == 0 {
		httpRequest, err = http.NewRequest(r.method, compiledURL, nil)
	} else {
		httpRequest, err = http.NewRequest(r.method, compiledURL, bytes.NewBuffer(body))
	}

	if err != nil {
		return nil, err
	}

	if bodyLen != 0 {
		httpRequest.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	}

	httpRequest.Header.Set(headerKeyAccept, headerValueApplicationJSON)

	return httpRequest.WithContext(ctx), nil
}
