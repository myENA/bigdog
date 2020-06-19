package bigdog

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrServiceTicketRequiresRefresh = errors.New("service ticket requires refresh")
	ErrServiceTicketCASInvalid      = errors.New("invalid CAS provided")
	ErrServiceTicketClientNil       = errors.New("client cannot be nil")
	ErrServiceTicketResponseEmpty   = errors.New("empty response from login request")
)

func IsServiceTicketRequiresRefreshErr(err error) bool {
	return errors.Is(err, ErrServiceTicketRequiresRefresh)
}

func IsServiceTicketCASInvalidErr(err error) bool {
	return errors.Is(err, ErrServiceTicketCASInvalid)
}

func IsServiceTicketClientNilErr(err error) bool {
	return errors.Is(err, ErrServiceTicketClientNil)
}

func IsServiceTicketResponseEmptyErr(err error) bool {
	return errors.Is(err, ErrServiceTicketResponseEmpty)
}

var (
	ErrAccessTokenRequiresRefresh = errors.New("access token requires refresh")
	ErrAccessTokenCASInvalid      = errors.New("invalid CAS provided")
	ErrAccessTokenClientNil       = errors.New("client cannot be nil")
	ErrAccessTokenResponseEmpty   = errors.New("empty response from login request")
)

func IsAccessTokenRequiresRefreshErr(err error) bool {
	return errors.Is(err, ErrAccessTokenRequiresRefresh)
}

func IsAccessTokenCasInvalidErr(err error) bool {
	return errors.Is(err, ErrAccessTokenCASInvalid)
}

func IsAccessTokenClientNilErr(err error) bool {
	return errors.Is(err, ErrAccessTokenClientNil)
}

func IsAccessTokenResponseEmptyErr(err error) bool {
	return errors.Is(err, ErrAccessTokenResponseEmpty)
}

type VSZAPIError struct {
	Success bool `json:"success"`

	Message      string      `json:"message,omitempty"`
	NoSession    string      `json:"noSession,omitempty"`
	ErrorDetails interface{} `json:"error,omitempty"` // probably can be map[string]interface{}
	ErrorCode    int         `json:"errorCode,omitempty"`
	ErrorType    string      `json:"errorType,omitempty"`
	Extra        interface{} `json:"extra,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`
	Data         interface{} `json:"data,omitempty"`

	Err error `json:"err,omitempty"`
}

func (e *VSZAPIError) Error() string {
	if e.ErrorDetails != nil {
		return fmt.Sprintf("success=%t; errorDetails=%v", e.Success, e.ErrorDetails)
	}
	return fmt.Sprintf("success=%t; errorCode=%q; errorType=%q; message=%q; err=%v", e.Success, e.ErrorCode, e.ErrorType, e.Message, e.Err)
}

func (e *VSZAPIError) Unwrap() error {
	return e.Err
}

func IsVSZAPIError(err error) bool {
	for err != nil {
		if verr, ok := err.(*VSZAPIError); ok && verr != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

type VSZServiceTicketProviderError struct {
	rm  *APIResponseMeta
	err error
}

func NewVSZServiceTicketProviderError(rm *APIResponseMeta, err error) *VSZServiceTicketProviderError {
	e := new(VSZServiceTicketProviderError)
	if rm == nil {
		rm = &APIResponseMeta{ResponseCode: http.StatusInternalServerError}
	}
	e.rm = rm
	e.err = err
	return e
}

func (e *VSZServiceTicketProviderError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("service ticket provider error: %v", e.err)
	}
	return ""
}

func (e *VSZServiceTicketProviderError) ResponseMeta() *APIResponseMeta {
	return e.rm
}

func (e *VSZServiceTicketProviderError) Unwrap() error {
	return e.err
}

func IsServiceTicketProviderError(err error) bool {
	for err != nil {
		if e, ok := err.(*VSZServiceTicketProviderError); ok && e != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

type SCIAPIErrorDetails struct {
	Name       string `json:"name"`
	Status     int    `json:"status"`
	Message    string `json:"message"`
	Length     int    `json:"length"`
	Severity   string `json:"severity"`
	Code       string `json:"code"`
	StatusCode int    `json:"statusCode"`
	File       string `json:"file"`
	Line       int    `json:"line"`
	Routine    string `json:"routine"`
}

type SCIAPIError struct {
	ErrorDetails *SCIAPIErrorDetails `json:"error,omitempty"`

	Err error `json:"err,omitempty"`
}

func (e *SCIAPIError) Error() string {
	if e.ErrorDetails != nil {
		if e.ErrorDetails.Code == "" {
			return fmt.Sprintf(
				"name=%q; status=%q; statusCode=%q; message=%q",
				e.ErrorDetails.Name,
				e.ErrorDetails.Status,
				e.ErrorDetails.StatusCode,
				e.ErrorDetails.Message,
			)
		}
		return fmt.Sprintf(
			"name=%q; status=%q; length=%q; severity=%q; code=%q; statusCode=%q; file=%q; line=%q; routine=%q; message=%q",
			e.ErrorDetails.Name,
			e.ErrorDetails.Status,
			e.ErrorDetails.Length,
			e.ErrorDetails.Severity,
			e.ErrorDetails.Code,
			e.ErrorDetails.StatusCode,
			e.ErrorDetails.File,
			e.ErrorDetails.Line,
			e.ErrorDetails.Routine,
			e.ErrorDetails.Message,
		)
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func (e *SCIAPIError) Unwrap() error {
	return e.Err
}

func IsSCIAPIError(err error) bool {
	for err != nil {
		if serr, ok := err.(*SCIAPIError); ok && serr != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

type SCIAccessTokenProviderError struct {
	rm  *APIResponseMeta
	err error
}

func NewSCIAccessTokenProviderError(rm *APIResponseMeta, err error) *SCIAccessTokenProviderError {
	e := new(SCIAccessTokenProviderError)
	if rm == nil {
		rm = &APIResponseMeta{ResponseCode: http.StatusInternalServerError}
	}
	e.rm = rm
	e.err = err
	return e
}

func (e *SCIAccessTokenProviderError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("access token provider error: %v", e.err)
	}
	return "'"
}

func (e *SCIAccessTokenProviderError) ResponseMeta() *APIResponseMeta {
	return e.rm
}

func (e *SCIAccessTokenProviderError) Unwrap() error {
	return e.err
}

func IsSCIAccessTokenProviderError(err error) bool {
	for err != nil {
		if e, ok := err.(*SCIAccessTokenProviderError); ok && e != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}
