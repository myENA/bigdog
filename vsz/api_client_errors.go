package vsz

import (
	"errors"
	"fmt"
)

var (
	ErrServiceTicketRequiresRefresh = errors.New("service ticket requires refresh")
	ErrServiceTicketCASInvalid      = errors.New("invalid CAS provided")
	ErrServiceTicketClientNil       = errors.New("client cannot be nil")
	ErrServiceTicketResponseEmpty   = errors.New("empty response from login request')")
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

type APIError struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
	ErrorType string `json:"errorType"`

	Err error `json:"err"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("success=%t; errorCode=%d; errorType=%s; message=%s; err=%v", e.Success, e.ErrorCode, e.ErrorType, e.Message, e.Err)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

func IsAPIError(err error) bool {
	for err != nil {
		if aerr, ok := err.(*APIError); ok && aerr != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

type ServiceTicketProviderError struct {
	rm  *APIResponseMeta
	err error
}

func NewServiceTicketProviderError(rm *APIResponseMeta, err error) *ServiceTicketProviderError {
	e := new(ServiceTicketProviderError)
	e.rm = rm
	e.err = err
	return e
}

func (e *ServiceTicketProviderError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("service ticket provider error: %v", e.err)
	}
	return ""
}

func (e *ServiceTicketProviderError) ResponseMeta() *APIResponseMeta {
	return e.rm
}

func (e *ServiceTicketProviderError) Unwrap() error {
	return e.err
}

func IsServiceTicketProviderError(err error) bool {
	for err != nil {
		if serr, ok := err.(*ServiceTicketProviderError); ok && serr != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}
