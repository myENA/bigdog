package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGRuckusGRETunnelProfileService struct {
	apiClient *VSZClient
}

func NewWSGRuckusGRETunnelProfileService(c *VSZClient) *WSGRuckusGRETunnelProfileService {
	s := new(WSGRuckusGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRuckusGRETunnelProfileService() *WSGRuckusGRETunnelProfileService {
	return NewWSGRuckusGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelRuckusgre
//
// Use this API command to create RuckusGRE tunnel profile.
//
// Operation ID: addProfilesTunnelRuckusgre
// Operation path: /profiles/tunnel/ruckusgre
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateRuckusGREProfile
func (s *WSGRuckusGRETunnelProfileService) AddProfilesTunnelRuckusgre(ctx context.Context, body *WSGProfileCreateRuckusGREProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesTunnelRuckusgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesTunnelRuckusgre
//
// Use this API command to delete multiple RuckusGRE tunnel profile.
//
// Operation ID: deleteProfilesTunnelRuckusgre
// Operation path: /profiles/tunnel/ruckusgre
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgre(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTunnelRuckusgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesTunnelRuckusgreById
//
// Use this API command to delete RuckusGRE tunnel profile.
//
// Operation ID: deleteProfilesTunnelRuckusgreById
// Operation path: /profiles/tunnel/ruckusgre/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgreById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTunnelRuckusgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesTunnelRuckusgre
//
// Use this API command to retrieve a list of RuckusGRE tunnel profile.
//
// Operation ID: findProfilesTunnelRuckusgre
// Operation path: /profiles/tunnel/ruckusgre
// Success code: 200 (OK)
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgre(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTunnelRuckusgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesTunnelRuckusgreById
//
// Use this API command to retrieve RuckusGRE tunnel profile.
//
// Operation ID: findProfilesTunnelRuckusgreById
// Operation path: /profiles/tunnel/ruckusgre/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileRuckusGREProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileRuckusGREProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTunnelRuckusgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileRuckusGREProfileAPIResponse), err
}

// FindProfilesTunnelRuckusgreByQueryCriteria
//
// Use this API command to query a list of RuckusGRE tunnel profile.
//
// Operation ID: findProfilesTunnelRuckusgreByQueryCriteria
// Operation path: /profiles/tunnel/ruckusgre/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileRuckusGREProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileRuckusGREProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesTunnelRuckusgreByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileRuckusGREProfileListAPIResponse), err
}

// PartialUpdateProfilesTunnelRuckusgreById
//
// Use this API command to modify the configuration of RuckusGRE tunnel profile.
//
// Operation ID: partialUpdateProfilesTunnelRuckusgreById
// Operation path: /profiles/tunnel/ruckusgre/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyRuckusGREProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) PartialUpdateProfilesTunnelRuckusgreById(ctx context.Context, body *WSGProfileModifyRuckusGREProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelRuckusgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
