package vsz

// API Version: v9_0

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
func (s *WSGHotspot20IdentityProviderProfileService) AddProfilesHs20Identityproviders(ctx context.Context, body *WSGProfileHs20Provider) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesHs20Identityproviders, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesHs20Identityproviders
//
// Use this API command to delete multiple Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20Identityproviders(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20Identityproviders, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesHs20IdentityprovidersAccountingsById
//
// Use this API command to disable accountings of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersAccountingsById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersAccountingsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesHs20IdentityprovidersById
//
// Use this API command to delete a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesHs20IdentityprovidersOsuById
//
// Use this API command to disable online signup & provisioning of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersOsuById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersOsuById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20Identityproviders(ctx context.Context, optionalParams map[string][]string) (*WSGProfileHs20ProviderList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileHs20ProviderList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20Identityproviders, true)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileHs20ProviderList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesHs20IdentityprovidersById
//
// Use this API command to retrieve a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersById(ctx context.Context, id string) (*WSGProfileHs20Provider, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileHs20Provider
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20IdentityprovidersById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileHs20Provider()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesHs20IdentityprovidersByQueryCriteria
//
// Query hotspot 2.0 identity providers.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileHs20ProviderList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileHs20ProviderList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesHs20IdentityprovidersByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileHs20ProviderList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesHs20IdentityprovidersById
//
// Use this API command to modify the configuration of a Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGProfileHs20Provider
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) PartialUpdateProfilesHs20IdentityprovidersById(ctx context.Context, body *WSGProfileHs20Provider, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesHs20IdentityprovidersById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
