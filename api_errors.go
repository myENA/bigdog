package bigdog

import (
	"errors"
)

type APIAuthProviderError struct {
	meta APIResponseMeta
	err  error
}

func NewAPIAuthProviderError(meta APIResponseMeta, err error) *APIAuthProviderError {
	ape := new(APIAuthProviderError)
	ape.meta = meta
	ape.err = err
	return ape
}

func (ape *APIAuthProviderError) ResponseMeta() APIResponseMeta {
	if ape == nil {
		return APIResponseMeta{}
	}
	return ape.meta
}

func (ape *APIAuthProviderError) Is(err error) bool {
	return ape != nil && ape.err != nil && errors.Is(err, ape.err)
}

func (ape *APIAuthProviderError) Unwrap() error {
	if ape == nil || ape.err == nil {
		return nil
	}
	return ape.err
}

func (ape *APIAuthProviderError) Error() string {
	if ape == nil || ape.err == nil {
		return ""
	}
	return ape.err.Error()
}
