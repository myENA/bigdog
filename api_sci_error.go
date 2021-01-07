package bigdog

import (
	"errors"
	"fmt"
)

var (
	ErrAccessTokenRequiresRefresh = errors.New("access token requires refresh")
	ErrAccessTokenCASInvalid      = errors.New("invalid CAS provided")
	ErrAccessTokenClientNil       = errors.New("client cannot be nil")
	ErrAccessTokenResponseEmpty   = errors.New("empty response from login request")

	accessTokenErrors = []error{
		ErrAccessTokenRequiresRefresh,
		ErrAccessTokenCASInvalid,
		ErrAccessTokenClientNil,
		ErrAccessTokenResponseEmpty,
	}

	accessTokenFatalErrors = []error{
		ErrAccessTokenCASInvalid,
		ErrAccessTokenClientNil,
		ErrAccessTokenResponseEmpty,
	}
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

func IsAccessTokenProviderError(err error) bool {
	for _, aerr := range accessTokenErrors {
		if errors.Is(err, aerr) {
			return true
		}
	}
	return false
}

func IsFatalAccessTokenProviderError(err error) bool {
	for _, aerr := range accessTokenFatalErrors {
		if errors.Is(err, aerr) {
			return true
		}
	}
	return false
}
