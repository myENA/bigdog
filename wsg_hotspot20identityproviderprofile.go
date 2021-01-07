package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGHotspot20IdentityProviderProfileService struct {
	apiClient *VSZClient
}

func NewWSGHotspot20IdentityProviderProfileService(c *VSZClient) *WSGHotspot20IdentityProviderProfileService {
	s := new(WSGHotspot20IdentityProviderProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20IdentityProviderProfileService() *WSGHotspot20IdentityProviderProfileService {
	return NewWSGHotspot20IdentityProviderProfileService(ss.apiClient)
}

// AddProfilesHs20Identityproviders
//
// Operation ID: addProfilesHs20Identityproviders
//
// Use this API command to create a new Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGProfileHs20Provider
func (s *WSGHotspot20IdentityProviderProfileService) AddProfilesHs20Identityproviders(ctx context.Context, body *WSGProfileHs20Provider, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesHs20Identityproviders, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesHs20Identityproviders
//
// Operation ID: deleteProfilesHs20Identityproviders
//
// Use this API command to delete multiple Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20Identityproviders(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20Identityproviders, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteProfilesHs20IdentityprovidersAccountingsById
//
// Operation ID: deleteProfilesHs20IdentityprovidersAccountingsById
//
// Use this API command to disable accountings of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersAccountingsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersAccountingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteProfilesHs20IdentityprovidersById
//
// Operation ID: deleteProfilesHs20IdentityprovidersById
//
// Use this API command to delete a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteProfilesHs20IdentityprovidersOsuById
//
// Operation ID: deleteProfilesHs20IdentityprovidersOsuById
//
// Use this API command to disable online signup & provisioning of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersOsuById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20IdentityprovidersOsuById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindProfilesHs20Identityproviders
//
// Operation ID: findProfilesHs20Identityproviders
//
// Use this API command to retrieve list of Hotspot 2.0 identity providers.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20Identityproviders(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileHs20ProviderListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileHs20ProviderListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileHs20ProviderListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesHs20Identityproviders, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileHs20ProviderListAPIResponse), err
}

// FindProfilesHs20IdentityprovidersById
//
// Operation ID: findProfilesHs20IdentityprovidersById
//
// Use this API command to retrieve a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileHs20ProviderAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileHs20ProviderAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileHs20ProviderAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesHs20IdentityprovidersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileHs20ProviderAPIResponse), err
}

// FindProfilesHs20IdentityprovidersByQueryCriteria
//
// Operation ID: findProfilesHs20IdentityprovidersByQueryCriteria
//
// Query hotspot 2.0 identity providers.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileHs20ProviderListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileHs20ProviderListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileHs20ProviderListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindProfilesHs20IdentityprovidersByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileHs20ProviderListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileHs20ProviderListAPIResponse), err
}

// PartialUpdateProfilesHs20IdentityprovidersById
//
// Operation ID: partialUpdateProfilesHs20IdentityprovidersById
//
// Use this API command to modify the configuration of a Hotspot 2.0 identity provider.
//
// Request Body:
//	 - body *WSGProfileHs20Provider
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20IdentityProviderProfileService) PartialUpdateProfilesHs20IdentityprovidersById(ctx context.Context, body *WSGProfileHs20Provider, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesHs20IdentityprovidersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
