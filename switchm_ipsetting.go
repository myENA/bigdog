package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type SwitchMIPSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMIPSettingsService(c *VSZClient) *SwitchMIPSettingsService {
	s := new(SwitchMIPSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMIPSettingsService() *SwitchMIPSettingsService {
	return NewSwitchMIPSettingsService(ss.apiClient)
}

// AddIpConfigs
//
// Use this API command to Create IP Config.
//
// Operation ID: addIpConfigs
// Operation path: /ipConfigs
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMIPConfigCreate
func (s *SwitchMIPSettingsService) AddIpConfigs(ctx context.Context, body *SwitchMIPConfigCreate, mutators ...RequestMutator) (*SwitchMIPConfigCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMIPConfigCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMIPConfigCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddIpConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMIPConfigCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMIPConfigCreateResultAPIResponse), err
}

// DeleteIpConfigs
//
// Use this API command to Delete IP Config by Id list.
//
// Operation ID: deleteIpConfigs
// Operation path: /ipConfigs
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMIPSettingsService) DeleteIpConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteIpConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteIpConfigsById
//
// Use this API command to Delete IP Config.
//
// Operation ID: deleteIpConfigsById
// Operation path: /ipConfigs/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMIPSettingsService) DeleteIpConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteIpConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindIpConfigs
//
// Use this API command to Retrieve IP Config List.
//
// Operation ID: findIpConfigs
// Operation path: /ipConfigs
// Success code: 200 (OK)
func (s *SwitchMIPSettingsService) FindIpConfigs(ctx context.Context, mutators ...RequestMutator) (*SwitchMIPConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMIPConfigListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMIPConfigListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindIpConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMIPConfigListAPIResponse), err
}

// FindIpConfigsById
//
// Use this API command to Retrieve IP Config.
//
// Operation ID: findIpConfigsById
// Operation path: /ipConfigs/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMIPSettingsService) FindIpConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMIPConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMIPConfigAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMIPConfigAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindIpConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMIPConfigAPIResponse), err
}

// FindIpConfigsByQueryCriteria
//
// Use this API command to Retrieve IP Config list.
//
// Operation ID: findIpConfigsByQueryCriteria
// Operation path: /ipConfigs/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMIPSettingsService) FindIpConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMIPConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMIPConfigListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMIPConfigListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindIpConfigsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMIPConfigListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMIPConfigListAPIResponse), err
}

// UpdateIpConfigsById
//
// Use this API command to Update IP Config.
//
// Operation ID: updateIpConfigsById
// Operation path: /ipConfigs/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMIPConfigModify
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMIPSettingsService) UpdateIpConfigsById(ctx context.Context, body *SwitchMIPConfigModify, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateIpConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
