package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMSwitchLAGSettingService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchLAGSettingService(c *VSZClient) *SwitchMSwitchLAGSettingService {
	s := new(SwitchMSwitchLAGSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchLAGSettingService() *SwitchMSwitchLAGSettingService {
	return NewSwitchMSwitchLAGSettingService(ss.apiClient)
}

// AddLagConfigs
//
// Use this API command to Create LAG Config.
//
// Request Body:
//	 - body *SwitchMSwitchLAGConfigCreate
func (s *SwitchMSwitchLAGSettingService) AddLagConfigs(ctx context.Context, body *SwitchMSwitchLAGConfigCreate) (*SwitchMSwitchLAGConfigCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchLAGConfigCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddLagConfigs, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchLAGConfigCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMSwitchLAGSettingService) DeleteLagConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteLagConfigs, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchLAGSettingService) DeleteLagConfigsById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteLagConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
func (s *SwitchMSwitchLAGSettingService) FindLagConfigs(ctx context.Context) (*SwitchMSwitchLAGConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchLAGConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindLagConfigs, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchLAGConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchLAGSettingService) FindLagConfigsById(ctx context.Context, id string) (*SwitchMSwitchLAGConfigLagConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchLAGConfigLagConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindLagConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchLAGConfigLagConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLagConfigsByQueryCriteria
//
// Use this API command to Retrieve LAG Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchLAGConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchLAGConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindLagConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchLAGConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Request Body:
//	 - body *SwitchMSwitchLAGConfigModify
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *SwitchMSwitchLAGConfigModify, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateLagConfigsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
