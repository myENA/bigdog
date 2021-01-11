package bigdog

import (
	"encoding/json"
	"errors"
	"fmt"
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

type APIResponseMeta struct {
	RequestMethod string
	RequestURI    string

	SuccessCode int

	ResponseCode   int
	ResponseStatus string
	ResponseHeader http.Header
}

func newAPIResponseMeta(req *APIRequest, successCode int, httpResp *http.Response) APIResponseMeta {
	rm := APIResponseMeta{
		RequestMethod: req.Method,
		RequestURI:    req.CompiledURI(),
		SuccessCode:   successCode,
	}
	if httpResp != nil {
		rm.ResponseCode = httpResp.StatusCode
		rm.ResponseStatus = http.StatusText(httpResp.StatusCode)
		// make copy of response headers so gc can clean up response after body has been read
		// todo: always read the body, dingus.
		rm.ResponseHeader = make(http.Header, len(httpResp.Header))
		for k, vs := range httpResp.Header {
			l := len(vs)
			rm.ResponseHeader[k] = make([]string, l, l)
			copy(rm.ResponseHeader[k], vs)
		}
	}
	return rm
}

func (rm APIResponseMeta) ContentType() string {
	return rm.ResponseHeader.Get(headerKeyContentType)
}

func (rm APIResponseMeta) ContentEncoding() string {
	return rm.ResponseHeader.Get(headerKeyContentEncoding)
}

func (rm APIResponseMeta) ContentLength() int {
	if v := rm.ResponseHeader.Get(headerKeyContentLength); v != "" {
		l, _ := strconv.Atoi(v)
		return l
	}
	return 0
}

func (rm APIResponseMeta) ContentDisposition() string {
	return rm.ResponseHeader.Get(headerKeyContentDisposition)
}

// Completed indicates whether the request described by this meta type completed, but does not indicate success
func (rm APIResponseMeta) Completed() bool {
	return rm.ResponseCode != 0
}

func (rm APIResponseMeta) String() string {
	var msg string
	if rm.SuccessCode == rm.ResponseCode {
		msg = "Success"
	} else {
		msg = "Error"
	}
	return fmt.Sprintf("%s response from request %s %s", msg, rm.RequestMethod, rm.RequestURI)
}

// APIResponseMetaContainer describes any type containing metadata about an API call
type APIResponseMetaContainer interface {
	ResponseMeta() APIResponseMeta
}

// APIResponse is implemented by each response model, allowing you to either unmarshal the resulting bytes into a model
// or receive the raw bytes themselves.
//
// This can be useful for middlewares that perform proxy duties, where there is nothing to be gained by unmarshalling
// and marshalling the returned data and it is fine to simply ship the bytes on along to the receiver.
type APIResponse interface {
	APIResponseMetaContainer

	// Source must return the source of this API response
	Source() APISource

	// Err must return the current state error, or nil if beginning state.
	Err() error

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
	Hydrate() (interface{}, error)
}

// APIResponseFactory is used by the internal response handling mechanism to construct each response type
type APIResponseFactory func(source APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse

// RawAPIResponse is the base implementation of APIResponse. Each api either returns this or returns a type that embeds
// this type.
type RawAPIResponse struct {
	mu   sync.Mutex
	src  APISource
	body io.ReadCloser
	err  error

	meta APIResponseMeta
}

func newRawAPIResponse(source APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	b := new(RawAPIResponse)
	b.src = source
	b.meta = meta
	if body != nil {
		b.body = body
	}
	return b
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

// ResponseMeta returns a portable meta type
func (b *RawAPIResponse) ResponseMeta() APIResponseMeta {
	return b.meta
}

// Source returns the upstream source of this response (VSZ or SCI)
func (b *RawAPIResponse) Source() APISource {
	return b.src
}

// Err returns the current state error, if there is one.
func (b *RawAPIResponse) Err() error {
	b.mu.Lock()
	err := b.err
	b.mu.Unlock()
	return err
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

// EmptyAPIResponse is returned by API calls that do not return a body, i.e. those with an http response code of 204
// or 201.  Calls to Read on this type will immediately return an io.EOF
type EmptyAPIResponse struct {
	*RawAPIResponse
}

func newEmptyAPIResponse(source APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(EmptyAPIResponse)
	cleanupReadCloser(body)
	r.RawAPIResponse = newRawAPIResponse(source, meta, nil).(*RawAPIResponse)
	return r
}

func (*EmptyAPIResponse) Read(_ []byte) (int, error) {
	return 0, io.EOF
}

type FileAPIResponse struct {
	*RawAPIResponse
}

func newFileAPIResponse(source APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(FileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(source, meta, body).(*RawAPIResponse)
	return r
}
