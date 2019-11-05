package vsz

import (
	"fmt"
)

type APIError struct {
	SourceError error `json:"sourceError"`

	RequestMethod string `json:"requestMethod"`
	RequestURI    string `json:"requestURI"`

	ResponseCode   int    `json:"responseCode"`
	ResponseStatus string `json:"responseStatus"`
	ResponseBody   []byte `json:"responseBody"`
}

func newAPIError(sourceError error, method, uri string, code int, status string, body []byte) error {
	return APIError{
		SourceError:    sourceError,
		RequestMethod:  method,
		RequestURI:     uri,
		ResponseCode:   code,
		ResponseStatus: status,
		ResponseBody:   body,
	}
}

func (e APIError) Error() string {
	if e.SourceError != nil {
		return e.SourceError.Error()
	}
	return fmt.Sprintf("%d: %s", e.ResponseCode, e.ResponseStatus)
}

func (e APIError) Unwrap() error {
	return e.SourceError
}
