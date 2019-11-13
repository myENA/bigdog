package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGHotspot20IdentityProviderProfileService struct {
	apiClient *APIClient
}

func NewWSGHotspot20IdentityProviderProfileService(c *APIClient) *WSGHotspot20IdentityProviderProfileService {
	s := new(WSGHotspot20IdentityProviderProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20IdentityProviderProfileService() *WSGHotspot20IdentityProviderProfileService {
	return NewWSGHotspot20IdentityProviderProfileService(ss.apiClient)
}

// AddProfilesHs20Identityproviders
//
// Use this API command to create a new Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGProfileHs20Provider
func (s *WSGHotspot20IdentityProviderProfileService) AddProfilesHs20Identityproviders(ctx context.Context, body *WSGProfileHs20Provider) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20Identityproviders
//
// Use this API command to delete multiple Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20Identityproviders(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20IdentityprovidersAccountingsById
//
// Use this API command to disable accountings of a Hotspot 2.0 identity provider.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersAccountingsById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20IdentityprovidersById
//
// Use this API command to delete a Hotspot 2.0 identity provider.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20IdentityprovidersOsuById
//
// Use this API command to disable online signup & provisioning of a Hotspot 2.0 identity provider.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersOsuById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesHs20Identityproviders
//
// Use this API command to retrieve list of Hotspot 2.0 identity providers.
//
// Query Parameters:
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20Identityproviders(ctx context.Context, qIndex string, qListSize string) (*WSGProfileHs20ProviderList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesHs20IdentityprovidersById
//
// Use this API command to retrieve a Hotspot 2.0 identity provider.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersById(ctx context.Context, pId string) (*WSGProfileHs20Provider, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesHs20IdentityprovidersByQueryCriteria
//
// Query hotspot 2.0 identity providers.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileHs20ProviderList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesHs20IdentityprovidersById
//
// Use this API command to modify the basic information of a Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGProfileHs20Provider
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) PartialUpdateProfilesHs20IdentityprovidersById(ctx context.Context, body *WSGProfileHs20Provider, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
