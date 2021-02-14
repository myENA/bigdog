package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGDevicePolicyInDomainLevelService struct {
	apiClient *VSZClient
}

func NewWSGDevicePolicyInDomainLevelService(c *VSZClient) *WSGDevicePolicyInDomainLevelService {
	s := new(WSGDevicePolicyInDomainLevelService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDevicePolicyInDomainLevelService() *WSGDevicePolicyInDomainLevelService {
	return NewWSGDevicePolicyInDomainLevelService(ss.apiClient)
}

// AddDevicePolicy
//
// Use this API command to create a Device Policy profile.
//
// Operation ID: addDevicePolicy
// Operation path: /devicePolicy
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGDomainDevicePolicyCreateDomainDevicePolicy
func (s *WSGDevicePolicyInDomainLevelService) AddDevicePolicy(ctx context.Context, body *WSGDomainDevicePolicyCreateDomainDevicePolicy, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDevicePolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteDevicePolicy
//
// Use this API command to delete a list of Device Policy profile.
//
// Operation ID: deleteDevicePolicy
// Operation path: /devicePolicy
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDevicePolicyInDomainLevelService) DeleteDevicePolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDevicePolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteDevicePolicyById
//
// Use this API command to delete a Device Policy profile.
//
// Operation ID: deleteDevicePolicyById
// Operation path: /devicePolicy/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDevicePolicyInDomainLevelService) DeleteDevicePolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindDevicePolicy
//
// Use this API command to retrieve list of Device Policy profiles.
//
// Operation ID: findDevicePolicy
// Operation path: /devicePolicy
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicy(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainDevicePolicyProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDevicePolicy, true)
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
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainDevicePolicyProfileListAPIResponse), err
}

// FindDevicePolicyById
//
// Use this API command to retrieve a Device Policy profile.
//
// Operation ID: findDevicePolicyById
// Operation path: /devicePolicy/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainDevicePolicyProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainDevicePolicyProfileAPIResponse), err
}

// FindDevicePolicyByQueryCriteria
//
// Query Device Policy Profile with specified filters.
//
// Operation ID: findDevicePolicyByQueryCriteria
// Operation path: /devicePolicy/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindDevicePolicyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse), err
}

// UpdateDevicePolicyById
//
// Use this API command to update a Device Policy profile.
//
// Operation ID: updateDevicePolicyById
// Operation path: /devicePolicy/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDomainDevicePolicyModifyDomainDevicePolicy
//
// Required parameters:
// - id string
//		- required
func (s *WSGDevicePolicyInDomainLevelService) UpdateDevicePolicyById(ctx context.Context, body *WSGDomainDevicePolicyModifyDomainDevicePolicy, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
