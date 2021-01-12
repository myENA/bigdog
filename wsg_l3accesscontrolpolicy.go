package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGL3AccessControlPolicyService struct {
	apiClient *VSZClient
}

func NewWSGL3AccessControlPolicyService(c *VSZClient) *WSGL3AccessControlPolicyService {
	s := new(WSGL3AccessControlPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL3AccessControlPolicyService() *WSGL3AccessControlPolicyService {
	return NewWSGL3AccessControlPolicyService(ss.apiClient)
}

// AddL3AccessControlPolicies
//
// Create a L3 Access Control Policy.
//
// Operation ID: addL3AccessControlPolicies
// Operation path: /l3AccessControlPolicies
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateL3AccessControlPolicy
func (s *WSGL3AccessControlPolicyService) AddL3AccessControlPolicies(ctx context.Context, body *WSGProfileCreateL3AccessControlPolicy, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddL3AccessControlPolicies, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteL3AccessControlPolicies
//
// Use this API command to delete Bulk L3 Access Control Policies.
//
// Operation ID: deleteL3AccessControlPolicies
// Operation path: /l3AccessControlPolicies
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL3AccessControlPolicyService) DeleteL3AccessControlPolicies(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteL3AccessControlPolicies, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteL3AccessControlPoliciesById
//
// Delete a L3 Access Control Policy.
//
// Operation ID: deleteL3AccessControlPoliciesById
// Operation path: /l3AccessControlPolicies/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGL3AccessControlPolicyService) DeleteL3AccessControlPoliciesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteL3AccessControlPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindL3AccessControlPolicies
//
// Retrieve L3 Access Control Policy list.
//
// Operation ID: findL3AccessControlPolicies
// Operation path: /l3AccessControlPolicies
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGL3AccessControlPolicyService) FindL3AccessControlPolicies(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileIdListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileIdListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindL3AccessControlPolicies, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileIdListAPIResponse), err
}

// FindL3AccessControlPoliciesById
//
// Retrieve a L3 Access Control Policy.
//
// Operation ID: findL3AccessControlPoliciesById
// Operation path: /l3AccessControlPolicies/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGL3AccessControlPolicyService) FindL3AccessControlPoliciesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileL3AccessControlPolicyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileL3AccessControlPolicyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindL3AccessControlPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileL3AccessControlPolicyAPIResponse), err
}

// FindL3AccessControlPoliciesByQueryCriteria
//
// Retrieve a list of L3 Access Control Policy.
//
// Operation ID: findL3AccessControlPoliciesByQueryCriteria
// Operation path: /l3AccessControlPolicies/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL3AccessControlPolicyService) FindL3AccessControlPoliciesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileL3AccessControlPolicyArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileL3AccessControlPolicyArrayAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindL3AccessControlPoliciesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileL3AccessControlPolicyArrayAPIResponse), err
}

// UpdateL3AccessControlPoliciesById
//
// Modify a L3 Access Control Policy.
//
// Operation ID: updateL3AccessControlPoliciesById
// Operation path: /l3AccessControlPolicies/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyL3AccessControlPolicy
//
// Required parameters:
// - id string
//		- required
func (s *WSGL3AccessControlPolicyService) UpdateL3AccessControlPoliciesById(ctx context.Context, body *WSGProfileModifyL3AccessControlPolicy, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateL3AccessControlPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
