package bigdog

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var (
	ErrResponseNoBody    = errors.New("response did not have a body")
	ErrResponseBeingRead = errors.New("response bytes have been either partially or fully read")
	ErrResponseClosed    = errors.New("response has been closed")
	ErrResponseHydrated  = errors.New("response bytes have been read via Hydrate call")
)

// APIResponse is implemented by each response model, allowing you to either unmarshal the resulting bytes into a model
// or receive the raw bytes themselves.
//
// This can be useful for middlewares that perform proxy duties, where there is nothing to be gained by unmarshalling
// and marshalling the returned data and it is fine to simply ship the bytes on along to the receiver.
//
// NOTE: No concurrency guarantee is made.  It is expected that will either immediately call Raw() or Hydrate() after
// receiving a model.
type APIResponse interface {
	// RequestMethod must return the HTTP method used.
	RequestMethod() string
	// RequestURI must return the fully qualified URL used.
	RequestURI() string
	// SuccessCode must return the expected HTTP response code.
	SuccessCode() int
	// ResponseCode must return the actual HTTP response code, if there is one.  If no response was received, this may
	// be zero.
	ResponseCode() int
	// ResponseStatus must return the actual HTTP response status, if there is one.  If no response was received, this
	// may be empty.
	ResponseStatus() string
	// ResponseHeader must return the actual HTTP header present in the response, if any.  If no response was received,
	// this may be empty
	ResponseHeader() http.Header

	// Read returns the raw bytes from the HTTP response body.  Once called
	//
	// This action is mutually exclusive with Hydrate.
	Read(p []byte) (n int, err error)

	// Close renders the internal response body defunct.  Once called, further calls to Read or Hydrate will
	Close() error
}

type ModeledAPIResponse interface {
	APIResponse

	// Hydrate must attempt to perform whatever action is necessary to ingest any returned bytes into a specific model
	// This action is mutually exclusive with Hydrate, and any attempt to use this as a reader after having called
	// Hydrate will result in an error
	Hydrate() error
}

type ResponseFactoryFunc func(req *APIRequest, successCode int, httpResp *http.Response) APIResponse

type RawAPIResponse struct {
	mu   sync.Mutex
	body io.ReadCloser
	err  error

	reqMethod   string
	reqURI      string
	successCode int
	respCode    int
	respHeader  http.Header
}

func newRawAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	b := new(RawAPIResponse)
	b.reqMethod = req.Method
	b.reqURI = req.CompiledURI()
	b.successCode = successCode
	if httpResp != nil {
		b.respCode = httpResp.StatusCode
		b.body = httpResp.Body
		// make copy of response headers so gc can clean up response after body has been read
		// todo: always read the body, dingus.
		b.respHeader = make(http.Header, len(httpResp.Header))
		for k, vs := range httpResp.Header {
			l := len(vs)
			b.respHeader[k] = make([]string, l, l)
			copy(b.respHeader[k], vs)
		}
	}
	return b
}

func (b *RawAPIResponse) ContentType() string {
	return b.respHeader.Get(headerKeyContentType)
}

func (b *RawAPIResponse) ContentLength() int {
	if v := b.respHeader.Get(headerKeyContentLength); v != "" {
		l, _ := strconv.Atoi(v)
		return l
	}
	return 0
}

func (b *RawAPIResponse) ContentDisposition() string {
	return b.respHeader.Get(headerKeyContentDisposition)
}

func (b *RawAPIResponse) cleanupBody() error {
	if b.body == nil {
		return nil
	}
	_, _ = io.Copy(ioutil.Discard, b.body)
	err := b.body.Close()
	b.body = nil
	return err
}

// RequestMethod returns name of HTTP request method used
func (b *RawAPIResponse) RequestMethod() string {
	return b.reqMethod
}

// RequestURI returns fully qualified URL
func (b *RawAPIResponse) RequestURI() string {
	return b.reqURI
}

// SuccessCode returns the expected "success" response code
func (b *RawAPIResponse) SuccessCode() int {
	return b.successCode
}

// ResponseCode returns the actual response code, if there was one
func (b *RawAPIResponse) ResponseCode() int {
	return b.respCode
}

// ResponseStatus returns the actual status response, if there was one
func (b *RawAPIResponse) ResponseStatus() string {
	return http.StatusText(b.respCode)
}

// ResponseHeader returns the raw header present in the response, if there was one.
func (b *RawAPIResponse) ResponseHeader() http.Header {
	return b.respHeader
}

// Read performs a read from the response body. Once this has been called, any Hydrate implementation will fail with an
// error
func (b *RawAPIResponse) Read(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// test for prior action
	if b.err != nil && !errors.Is(b.err, ErrResponseBeingRead) {
		return 0, b.err
	}

	// check for body being nil
	if b.body == nil {
		b.err = ErrResponseNoBody
		return 0, ErrResponseNoBody
	}

	var (
		n   int
		err error
	)

	// always set state error
	b.err = ErrResponseBeingRead

	// perform read
	if n, err = b.body.Read(p); err != nil {
		// if any error is encountered, set state error before returning
		b.err = err
	}

	// if no error encountered, simply return
	return n, err
}

// Close closes the upstream http response body and removes its reference from the local type, enabling response gc
func (b *RawAPIResponse) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	// if this response was previous hydrated, move on
	if b.err != nil && errors.Is(b.err, ErrResponseHydrated) {
		return nil
	}
	// in all other cases, perform full cleanupBody and return any close error
	b.err = ErrResponseClosed
	return b.cleanupBody()
}

func (b *RawAPIResponse) doHydrate(ptr interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// test for an action already having been taken
	if b.err != nil {
		// test for a prior Hydrate action
		if errors.Is(b.err, ErrResponseHydrated) {
			return nil
		}
		// if err is anything else, return directly
		return b.err
	}

	var (
		data []byte
		err  error
	)

	// queue up body cleanup
	defer func() { _ = b.cleanupBody() }()

	// consume body
	if data, err = ioutil.ReadAll(b.body); err != nil {
		b.err = err
		return err
	}
	// attempt unmarshal
	if err = json.Unmarshal(data, ptr); err != nil {
		b.err = err
		return err
	}
	// specify Hydrate error for subsequent action attempts
	b.err = ErrResponseHydrated
	return nil
}
