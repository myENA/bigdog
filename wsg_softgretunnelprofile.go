package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGSoftGRETunnelProfileService struct {
	apiClient *VSZClient
}

func NewWSGSoftGRETunnelProfileService(c *VSZClient) *WSGSoftGRETunnelProfileService {
	s := new(WSGSoftGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSoftGRETunnelProfileService() *WSGSoftGRETunnelProfileService {
	return NewWSGSoftGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelSoftgre
//
// Use this API command to create SoftGRE tunnel profile.
//
// Operation ID: addProfilesTunnelSoftgre
// Operation path: /profiles/tunnel/softgre
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateSoftGREProfile
func (s *WSGSoftGRETunnelProfileService) AddProfilesTunnelSoftgre(ctx context.Context, body *WSGProfileCreateSoftGREProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesTunnelSoftgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesTunnelSoftgre
//
// Use this API command to delete multiple SoftGRE tunnel profile.
//
// Operation ID: deleteProfilesTunnelSoftgre
// Operation path: /profiles/tunnel/softgre
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgre(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTunnelSoftgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesTunnelSoftgreById
//
// Use this API command to delete SoftGRE tunnel profile.
//
// Operation ID: deleteProfilesTunnelSoftgreById
// Operation path: /profiles/tunnel/softgre/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgreById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTunnelSoftgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesTunnelSoftgre
//
// Use this API command to retrieve a list of SoftGRE tunnel profile.
//
// Operation ID: findProfilesTunnelSoftgre
// Operation path: /profiles/tunnel/softgre
// Success code: 200 (OK)
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgre(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTunnelSoftgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesTunnelSoftgreById
//
// Use this API command to retrieve SoftGRE tunnel profile.
//
// Operation ID: findProfilesTunnelSoftgreById
// Operation path: /profiles/tunnel/softgre/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileSoftGREProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileSoftGREProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTunnelSoftgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileSoftGREProfileAPIResponse), err
}

// FindProfilesTunnelSoftgreByQueryCriteria
//
// Use this API command to query a list of SoftGRE tunnel profile.
//
// Operation ID: findProfilesTunnelSoftgreByQueryCriteria
// Operation path: /profiles/tunnel/softgre/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileSoftGREProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileSoftGREProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesTunnelSoftgreByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileSoftGREProfileListAPIResponse), err
}

// PartialUpdateProfilesTunnelSoftgreById
//
// Use this API command to modify the configuration of SoftGRE tunnel profile.
//
// Operation ID: partialUpdateProfilesTunnelSoftgreById
// Operation path: /profiles/tunnel/softgre/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifySoftGREProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) PartialUpdateProfilesTunnelSoftgreById(ctx context.Context, body *WSGProfileModifySoftGREProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelSoftgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
