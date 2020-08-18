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
func (s *WSGDevicePolicyInDomainLevelService) AddDevicePolicy(ctx context.Context, body *WSGDomainDevicePolicyCreateDomainDevicePolicy, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDevicePolicy, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDevicePolicy
//
// Operation ID: deleteDevicePolicy
//
// Use this API command to delete a list of Device Policy profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDevicePolicyInDomainLevelService) DeleteDevicePolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDevicePolicy, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGDevicePolicyInDomainLevelService) DeleteDevicePolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDevicePolicyById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicy(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDomainDevicePolicyProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDevicePolicy, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDomainDevicePolicyProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDomainDevicePolicyProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDevicePolicyById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDomainDevicePolicyProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDevicePolicyByQueryCriteria
//
// Operation ID: findDevicePolicyByQueryCriteria
//
// Query Device Policy Profile with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDevicePolicyInDomainLevelService) FindDevicePolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGDomainDevicePolicyProfileByQueryCriteria, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDomainDevicePolicyProfileByQueryCriteria
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindDevicePolicyByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDomainDevicePolicyProfileByQueryCriteria()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGDevicePolicyInDomainLevelService) UpdateDevicePolicyById(ctx context.Context, body *WSGDomainDevicePolicyModifyDomainDevicePolicy, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDevicePolicyById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
