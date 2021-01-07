package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addDevicePolicy
//
// Use this API command to create a Device Policy profile.
//
// Request Body:
//	 - body *WSGDomainDevicePolicyCreateDomainDevicePolicy
func (s *WSGDevicePolicyInDomainLevelService) AddDevicePolicy(ctx context.Context, body *WSGDomainDevicePolicyCreateDomainDevicePolicy, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDevicePolicy, true)
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

// DeleteDevicePolicy
//
// Operation ID: deleteDevicePolicy
//
// Use this API command to delete a list of Device Policy profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDevicePolicyInDomainLevelService) DeleteDevicePolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDevicePolicy, true)
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

// DeleteDevicePolicyById
//
// Operation ID: deleteDevicePolicyById
//
// Use this API command to delete a Device Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDevicePolicyInDomainLevelService) DeleteDevicePolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindDevicePolicy
//
// Operation ID: findDevicePolicy
//
// Use this API command to retrieve list of Device Policy profiles.
//
// Optional Parameters:
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
		resp     APIResponse
		err      error

		respFn = newWSGDomainDevicePolicyProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDomainDevicePolicyProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDevicePolicy, true)
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
	return resp.(*WSGDomainDevicePolicyProfileListAPIResponse), err
}

// FindDevicePolicyById
//
// Operation ID: findDevicePolicyById
//
// Use this API command to retrieve a Device Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDomainDevicePolicyProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDomainDevicePolicyProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainDevicePolicyProfileAPIResponse), err
}

// FindDevicePolicyByQueryCriteria
//
// Operation ID: findDevicePolicyByQueryCriteria
//
// Query Device Policy Profile with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindDevicePolicyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainDevicePolicyProfileByQueryCriteriaAPIResponse), err
}

// UpdateDevicePolicyById
//
// Operation ID: updateDevicePolicyById
//
// Use this API command to update a Device Policy profile.
//
// Request Body:
//	 - body *WSGDomainDevicePolicyModifyDomainDevicePolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDevicePolicyInDomainLevelService) UpdateDevicePolicyById(ctx context.Context, body *WSGDomainDevicePolicyModifyDomainDevicePolicy, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDevicePolicyById, true)
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
