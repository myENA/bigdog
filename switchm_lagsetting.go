package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type SwitchMLAGSettingService struct {
	apiClient *VSZClient
}

func NewSwitchMLAGSettingService(c *VSZClient) *SwitchMLAGSettingService {
	s := new(SwitchMLAGSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMLAGSettingService() *SwitchMLAGSettingService {
	return NewSwitchMLAGSettingService(ss.apiClient)
}

// AddLagConfigs
//
// Use this API command to Create LAG Config.
//
// Operation ID: addLagConfigs
// Operation path: /lagConfigs
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMLAGConfigCreate
func (s *SwitchMLAGSettingService) AddLagConfigs(ctx context.Context, body *SwitchMLAGConfigCreate, mutators ...RequestMutator) (*SwitchMLAGConfigCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddLagConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigCreateResultAPIResponse), err
}

// DeleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Operation ID: deleteLagConfigs
// Operation path: /lagConfigs
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMLAGSettingService) DeleteLagConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteLagConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Operation ID: deleteLagConfigsById
// Operation path: /lagConfigs/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) DeleteLagConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteLagConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
//
// Operation ID: findLagConfigs
// Operation path: /lagConfigs
// Success code: 200 (OK)
func (s *SwitchMLAGSettingService) FindLagConfigs(ctx context.Context, mutators ...RequestMutator) (*SwitchMLAGConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindLagConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigListAPIResponse), err
}

// FindLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
//
// Operation ID: findLagConfigsById
// Operation path: /lagConfigs/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) FindLagConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMLAGConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindLagConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigAPIResponse), err
}

// FindLagConfigsByQueryCriteria
//
// Use this API command to Retrieve LAG Config list.
//
// Operation ID: findLagConfigsByQueryCriteria
// Operation path: /lagConfigs/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMLAGConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindLagConfigsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigListAPIResponse), err
}

// UpdateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Operation ID: updateLagConfigsById
// Operation path: /lagConfigs/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMLAGConfigModify
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *SwitchMLAGConfigModify, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateLagConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
