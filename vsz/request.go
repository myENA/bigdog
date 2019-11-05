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

var requestID uint64

type Request struct {
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

func NewRequest(method, uri string, authenticated bool) *Request {
	r := &Request{
		id:              atomic.AddUint64(&requestID, 1),
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

func (r *Request) ID() uint64 {
	return r.id
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) URI() string {
	return r.uri
}

func (r *Request) Authenticated() bool {
	return r.authenticated
}

func (r *Request) SetHeaders(headers url.Values) {
	r.headers = headers
}

func (r *Request) AddHeader(name, value string) {
	r.headers.Add(name, value)
}

func (r *Request) SetHeader(name, value string) {
	r.headers.Set(name, value)
}

func (r *Request) RemoveHeader(name string) {
	r.headers.Del(name)
}

func (r *Request) Headers() url.Values {
	return r.headers
}

func (r *Request) SetCookies(cookies []*http.Cookie) {
	r.cookies = cookies
}

func (r *Request) AddCookie(cookie *http.Cookie) {
	r.cookies = append(r.cookies, cookie)
}

func (r *Request) SetCookie(cookie *http.Cookie) {
	for i, cc := range r.cookies {
		if cc.Name == cookie.Name {
			r.cookies[i] = cookie
			return
		}
	}
	r.cookies = append(r.cookies, cookie)
}

func (r *Request) RemoveCookie(name string) {
	nc := make([]*http.Cookie, 0)
	for _, cc := range r.cookies {
		if cc.Name != name {
			nc = append(nc, cc)
		}
	}
	r.cookies = nc
}

func (r *Request) Cookies() []*http.Cookie {
	return r.cookies
}

// SetQueryParameters will override any / all existing query parameters
func (r *Request) SetQueryParameters(params map[string][]string) {
	r.queryParameters = params
	r.compiledURI = ""
}

// SetQueryParameter will set a query param to a specific value, overriding any previously set value
func (r *Request) SetQueryParameter(param, value string) {
	r.queryParameters[param] = []string{value}
	r.compiledURI = ""
}

// AddQueryParameter will add a value to the specified param
func (r *Request) AddQueryParameter(param, value string) {
	if _, ok := r.queryParameters[param]; !ok {
		r.queryParameters[param] = make([]string, 0)
	}
	r.queryParameters[param] = append(r.queryParameters[param], value)
	r.compiledURI = ""
}

// RemoveQueryParameter will attempt to delete all values for a specific query parameter from this request.
func (r *Request) RemoveQueryParameter(param string) {
	delete(r.queryParameters, param)
	r.compiledURI = ""
}

// QueryParameters will return all values of currently set query parameters
func (r *Request) QueryParameters() map[string][]string {
	return r.queryParameters
}

// SetPathParameters will re-define all path parameters, overriding any / all existing ones
func (r *Request) SetPathParameters(params map[string]string) {
	r.pathParameters = params
	r.compiledURI = ""
}

// SetPathParameter will define a path parameter value, overriding any existing value
func (r *Request) SetPathParameter(param, value string) {
	r.pathParameters[param] = value
	r.compiledURI = ""
}

// RemovePathParameter will attempt to remove a single parameter from the current list of path parameters
func (r *Request) RemovePathParameter(param string) {
	delete(r.pathParameters, param)
	r.compiledURI = ""
}

func (r *Request) PathParameters() map[string]string {
	return r.pathParameters
}

func (r *Request) SetBody(body []byte) {
	r.body = body
}

func (r *Request) SetBodyModel(model interface{}) error {
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

func (r *Request) Body() []byte {
	return r.body
}

func (r *Request) CompiledURI() string {
	if r.compiledURI == "" {
		r.compileURI()
	}
	return r.compiledURI
}

func (r *Request) compileURI() {
	pathParams := r.PathParameters()
	queryParams := r.QueryParameters()
	uri := r.uri
	if len(pathParams) > 0 {
		for k, v := range pathParams {
			uri = strings.Replace(uri, fmt.Sprintf("{%s}", k), v, 1)
		}
	}
	// TODO: could probably be made more efficient.
	if len(queryParams) > 0 {
		uri = fmt.Sprintf("%s?", uri)
		for param, values := range queryParams {
			for _, value := range values {
				if value == "" {
					uri = fmt.Sprintf("%s%s&", uri, param)
				} else {
					uri = fmt.Sprintf("%s%s=%s&", uri, param, value)
				}
			}
		}
		uri = strings.TrimRight(uri, "&")
	}
	r.compiledURI = uri
}

// toHTTP will attempt to construct an executable http.request
func (r *Request) toHTTP(ctx context.Context, addr string, port int, debug bool) (*http.Request, error) {
	var err error
	var httpRequest *http.Request

	body := r.Body()
	bodyLen := len(body)
	compiledURL := fmt.Sprintf("https://%s:%d%s", addr, port, r.CompiledURI())

	// if debug mode is enabled, prepare a big'ol log statement.
	if debug {
		logMsg := fmt.Sprintf("[request-%d] Preparing request \"%s %s\"", r.id, r.method, compiledURL)

		if bodyLen == 0 {
			logMsg = fmt.Sprintf("%s without body", logMsg)
		} else {
			logMsg = fmt.Sprintf("%s with body: %s", logMsg, string(body))
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
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	httpRequest.Header.Set("Accept", "application/json")

	return httpRequest.WithContext(ctx), nil
}
