package bigdog

import (
	"errors"
	"fmt"
)

var (
	ErrServiceTicketRequiresRefresh = errors.New("service ticket requires refresh")
	ErrServiceTicketCASInvalid      = errors.New("invalid CAS provided")
	ErrServiceTicketClientNil       = errors.New("client cannot be nil")
	ErrServiceTicketResponseEmpty   = errors.New("empty response from login request")

	serviceTicketErrors = []error{
		ErrServiceTicketRequiresRefresh,
		ErrServiceTicketCASInvalid,
		ErrServiceTicketClientNil,
		ErrServiceTicketResponseEmpty,
	}

	serviceTicketFatalErrors = []error{
		ErrServiceTicketCASInvalid,
		ErrServiceTicketClientNil,
		ErrServiceTicketResponseEmpty,
	}
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

func IsServiceTicketProviderError(err error) bool {
	for _, sterr := range serviceTicketErrors {
		if errors.Is(err, sterr) {
			return true
		}
	}
	return false
}

func IsFatalServiceTicketProviderError(err error) bool {
	for _, sterr := range serviceTicketFatalErrors {
		if errors.Is(err, sterr) {
			return true
		}
	}
	return false
}
