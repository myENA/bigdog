package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGAuthenticationProfileService struct {
	apiClient *VSZClient
}

func NewWSGAuthenticationProfileService(c *VSZClient) *WSGAuthenticationProfileService {
	s := new(WSGAuthenticationProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationProfileService() *WSGAuthenticationProfileService {
	return NewWSGAuthenticationProfileService(ss.apiClient)
}

// AddProfilesAuth
//
// Use this API command to create a new authentication profile.
//
// Operation ID: addProfilesAuth
// Operation path: /profiles/auth
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateAuthenticationProfile
func (s *WSGAuthenticationProfileService) AddProfilesAuth(ctx context.Context, body *WSGProfileCreateAuthenticationProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesAuth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddProfilesAuthCloneById
//
// Use this API command to clone an authentication profile.
//
// Operation ID: addProfilesAuthCloneById
// Operation path: /profiles/auth/clone/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGProfileClone
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) AddProfilesAuthCloneById(ctx context.Context, body *WSGProfileClone, id string, mutators ...RequestMutator) (*WSGProfileCloneAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileCloneAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesAuthCloneById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileCloneAPIResponse), err
}

// DeleteProfilesAuth
//
// Use this API command to delete a list of authentication profile.
//
// Operation ID: deleteProfilesAuth
// Operation path: /profiles/auth
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileDeleteBulkAuthenticationProfile
func (s *WSGAuthenticationProfileService) DeleteProfilesAuth(ctx context.Context, body *WSGProfileDeleteBulkAuthenticationProfile, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesAuth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesAuthById
//
// Use this API command to delete an authentication profile.
//
// Operation ID: deleteProfilesAuthById
// Operation path: /profiles/auth/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) DeleteProfilesAuthById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesAuthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesAuth
//
// Use this API command to retrieve a list of authentication profiles.
//
// Operation ID: findProfilesAuth
// Operation path: /profiles/auth
// Success code: 200 (OK)
func (s *WSGAuthenticationProfileService) FindProfilesAuth(ctx context.Context, mutators ...RequestMutator) (*WSGProfileAuthenticationProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileAuthenticationProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesAuth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileAuthenticationProfileListAPIResponse), err
}

// FindProfilesAuthAuthorizationList
//
// Use this API command to retrieve a list of authorization profiles.
//
// Operation ID: findProfilesAuthAuthorizationList
// Operation path: /profiles/auth/authorizationList
// Success code: 200 (OK)
//
// Required parameters:
// - type_ string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthorizationList(ctx context.Context, type_ string, mutators ...RequestMutator) (*WSGProfileBaseServiceInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBaseServiceInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesAuthAuthorizationList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBaseServiceInfoListAPIResponse), err
}

// FindProfilesAuthAuthServiceListByQueryCriteria
//
// Use this API command to retrieve a list of authentication service.
//
// Operation ID: findProfilesAuthAuthServiceListByQueryCriteria
// Operation path: /profiles/auth/authServiceList/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthServiceListByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileBaseServiceInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBaseServiceInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesAuthAuthServiceListByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBaseServiceInfoListAPIResponse), err
}

// FindProfilesAuthById
//
// Use this API command to retrieve an authentication profile.
//
// Operation ID: findProfilesAuthById
// Operation path: /profiles/auth/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileAuthenticationProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileAuthenticationProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesAuthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileAuthenticationProfileAPIResponse), err
}

// FindProfilesAuthByQueryCriteria
//
// Use this API command to retrieve a list of authentication profiles by query criteria.
//
// Operation ID: findProfilesAuthByQueryCriteria
// Operation path: /profiles/auth/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileAuthenticationProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileAuthenticationProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesAuthByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileAuthenticationProfileListAPIResponse), err
}

// PartialUpdateProfilesAuthById
//
// Use this API command to modify the configuration of an authentication profile.
//
// Operation ID: partialUpdateProfilesAuthById
// Operation path: /profiles/auth/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyAuthenticationProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) PartialUpdateProfilesAuthById(ctx context.Context, body *WSGProfileModifyAuthenticationProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesAuthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
