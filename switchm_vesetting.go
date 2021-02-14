package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type SwitchMVESettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMVESettingsService(c *VSZClient) *SwitchMVESettingsService {
	s := new(SwitchMVESettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVESettingsService() *SwitchMVESettingsService {
	return NewSwitchMVESettingsService(ss.apiClient)
}

// AddVeConfigs
//
// Use this API command to Create VE Config.
//
// Operation ID: addVeConfigs
// Operation path: /veConfigs
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMVEConfigCreate
func (s *SwitchMVESettingsService) AddVeConfigs(ctx context.Context, body *SwitchMVEConfigCreate, mutators ...RequestMutator) (*SwitchMVEConfigCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMVEConfigCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddVeConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVEConfigCreateResultAPIResponse), err
}

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Operation ID: deleteVeConfigs
// Operation path: /veConfigs
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVESettingsService) DeleteVeConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteVeConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Operation ID: deleteVeConfigsById
// Operation path: /veConfigs/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMVESettingsService) DeleteVeConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteVeConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindVeConfigs
//
// Use this API command to Retrieve VE Config List.
//
// Operation ID: findVeConfigs
// Operation path: /veConfigs
// Success code: 200 (OK)
func (s *SwitchMVESettingsService) FindVeConfigs(ctx context.Context, mutators ...RequestMutator) (*SwitchMVEConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMVEConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindVeConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVEConfigListAPIResponse), err
}

// FindVeConfigsById
//
// Use this API command to Retrieve VE Config.
//
// Operation ID: findVeConfigsById
// Operation path: /veConfigs/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMVESettingsService) FindVeConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMVEConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMVEConfigAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindVeConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVEConfigAPIResponse), err
}

// FindVeConfigsByQueryCriteria
//
// Use this API command to Retrieve VE Config list.
//
// Operation ID: findVeConfigsByQueryCriteria
// Operation path: /veConfigs/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVESettingsService) FindVeConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMVEConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMVEConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindVeConfigsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVEConfigListAPIResponse), err
}

// UpdateVeConfigsById
//
// Use this API command to Update VE Config.
//
// Operation ID: updateVeConfigsById
// Operation path: /veConfigs/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMVEConfigModify
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMVESettingsService) UpdateVeConfigsById(ctx context.Context, body *SwitchMVEConfigModify, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateVeConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
