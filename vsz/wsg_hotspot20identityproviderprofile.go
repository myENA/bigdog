package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesHs20Identityproviders, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesHs20Identityproviders
//
// Use this API command to delete multiple Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20Identityproviders(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20Identityproviders, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesHs20IdentityprovidersAccountingsById
//
// Use this API command to disable accountings of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersAccountingsById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersAccountingsById, true)
}

// DeleteProfilesHs20IdentityprovidersById
//
// Use this API command to delete a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersById, true)
}

// DeleteProfilesHs20IdentityprovidersOsuById
//
// Use this API command to disable online signup & provisioning of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersOsuById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersOsuById, true)
}

// FindProfilesHs20Identityproviders
//
// Use this API command to retrieve list of Hotspot 2.0 identity providers.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20Identityproviders(ctx context.Context, optionalParams map[string]interface{}) (*WSGProfileHs20ProviderList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileHs20ProviderList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20Identityproviders, true)
}

// FindProfilesHs20IdentityprovidersById
//
// Use this API command to retrieve a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersById(ctx context.Context, id string) (*WSGProfileHs20Provider, error) {
	var (
		req  *APIRequest
		resp *WSGProfileHs20Provider
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20IdentityprovidersById, true)
}

// FindProfilesHs20IdentityprovidersByQueryCriteria
//
// Query hotspot 2.0 identity providers.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileHs20ProviderList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileHs20ProviderList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesHs20IdentityprovidersByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateProfilesHs20IdentityprovidersById
//
// Use this API command to modify the basic information of a Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGProfileHs20Provider
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) PartialUpdateProfilesHs20IdentityprovidersById(ctx context.Context, body *WSGProfileHs20Provider, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesHs20IdentityprovidersById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
