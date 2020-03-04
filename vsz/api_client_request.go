package vsz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	uriQueryParameterCutSet           = "&"

	apiRequestURLFormat = "%s%s"

	headerKeyContentType       = "Content-Type"
	headerKeyAccept            = "Accept"
	headerValueApplicationJSON = "application/json"
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

func (r *APIRequest) AddHeader(name, value string) {
	r.headers.Add(name, value)
}

func (r *APIRequest) SetHeader(name, value string) {
	r.headers.Set(name, value)
}

func (r *APIRequest) SetHeaders(headers url.Values) {
	r.headers = make(url.Values)
	for k, vs := range headers {
		for _, v := range vs {
			r.AddHeader(k, v)
		}
	}
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

// AddQueryParameter will add a value to the specified param
func (r *APIRequest) AddQueryParameter(param string, value []string) {
	r.compiledURI = ""
	if _, ok := r.queryParameters[param]; !ok {
		r.queryParameters[param] = make([]string, 0)
	}
	r.queryParameters[param] = append(r.queryParameters[param], value...)
}

// SetQueryParameter will set a query param to a specific value, overriding any previously set value
func (r *APIRequest) SetQueryParameter(param string, value []string) {
	delete(r.queryParameters, param)
	r.AddQueryParameter(param, value)
}

// SetQueryParameters will override any / all existing query parameters
func (r *APIRequest) SetQueryParameters(params map[string][]string) {
	r.queryParameters = make(map[string][]string)
	for k, v := range params {
		r.AddQueryParameter(k, v)
	}
}

// RemoveQueryParameter will attempt to delete all values for a specific query parameter from this request.
func (r *APIRequest) RemoveQueryParameter(param string) {
	r.compiledURI = ""
	delete(r.queryParameters, param)
}

// QueryParameters will return all values of currently set query parameters
func (r *APIRequest) QueryParameters() map[string][]string {
	return r.queryParameters
}

// SetPathParameter will define a path parameter value, overriding any existing value
func (r *APIRequest) SetPathParameter(param, value string) {
	r.compiledURI = ""
	r.pathParameters[param] = value
}

// SetPathParameters will re-define all path parameters, overriding any / all existing ones
func (r *APIRequest) SetPathParameters(params map[string]string) {
	r.pathParameters = make(map[string]string)
	for k, v := range params {
		r.SetPathParameter(k, v)
	}
}

// RemovePathParameter will attempt to remove a single parameter from the current list of path parameters
func (r *APIRequest) RemovePathParameter(param string) {
	delete(r.pathParameters, param)
	r.compiledURI = ""
}

func (r *APIRequest) PathParameters() map[string]string {
	return r.pathParameters
}

func (r *APIRequest) SetBody(body interface{}) error {
	// first, is this already a byte slice?
	if b, ok := body.([]byte); ok {
		r.body = b
		return nil
	}

	// next, is an io.Reader being passed?
	if asReader, ok := body.(io.Reader); ok {
		if b, err := ioutil.ReadAll(asReader); err != nil {
			return err
		} else {
			r.body = b
			return nil
		}
	}

	// finally, assume we gotta json serialize this thing.
	if b, err := json.Marshal(body); err != nil {
		return err
	} else {
		r.body = b
		return nil
	}
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
			uri = strings.TrimRight(uri, uriQueryParameterCutSet)
		}
		r.compiledURI = uri
	}
	return r.compiledURI
}

// toHTTP will attempt to construct an executable http.request
func (r *APIRequest) toHTTP(ctx context.Context, addr string) (*http.Request, error) {
	var err error
	var httpRequest *http.Request

	body := r.Body()
	bodyLen := len(body)
	compiledURL := fmt.Sprintf(apiRequestURLFormat, addr, r.CompiledURI())

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
