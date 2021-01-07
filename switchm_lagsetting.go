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
// Operation ID: addLagConfigs
//
// Use this API command to Create LAG Config.
//
// Request Body:
//	 - body *SwitchMLAGConfigCreate
func (s *SwitchMLAGSettingService) AddLagConfigs(ctx context.Context, body *SwitchMLAGConfigCreate, mutators ...RequestMutator) (*SwitchMLAGConfigCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMLAGConfigCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddLagConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMLAGConfigCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigCreateResultAPIResponse), err
}

// DeleteLagConfigs
//
// Operation ID: deleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMLAGSettingService) DeleteLagConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteLagConfigs, true)
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

// DeleteLagConfigsById
//
// Operation ID: deleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) DeleteLagConfigsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteLagConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindLagConfigs
//
// Operation ID: findLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
func (s *SwitchMLAGSettingService) FindLagConfigs(ctx context.Context, mutators ...RequestMutator) (*SwitchMLAGConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMLAGConfigListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindLagConfigs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigListAPIResponse), err
}

// FindLagConfigsById
//
// Operation ID: findLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMLAGConfigAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindLagConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigAPIResponse), err
}

// FindLagConfigsByQueryCriteria
//
// Operation ID: findLagConfigsByQueryCriteria
//
// Use this API command to Retrieve LAG Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMLAGConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMLAGConfigListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMLAGConfigListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindLagConfigsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMLAGConfigListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMLAGConfigListAPIResponse), err
}

// UpdateLagConfigsById
//
// Operation ID: updateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Request Body:
//	 - body *SwitchMLAGConfigModify
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *SwitchMLAGConfigModify, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateLagConfigsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
