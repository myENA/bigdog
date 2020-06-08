package ruckus

// API Version: v9_0

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
// Request Body:
//	 - body *SwitchMSwitchIPConfigCreate
func (s *SwitchMIPSettingsService) AddIpConfigs(ctx context.Context, body *SwitchMSwitchIPConfigCreate) (*SwitchMSwitchIPConfigCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchIPConfigCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddIpConfigs, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchIPConfigCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteIpConfigs
//
// Use this API command to Delete IP Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMIPSettingsService) DeleteIpConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteIpConfigs, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteIpConfigsById
//
// Use this API command to Delete IP Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMIPSettingsService) DeleteIpConfigsById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteIpConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindIpConfigs
//
// Use this API command to Retrieve IP Config List.
func (s *SwitchMIPSettingsService) FindIpConfigs(ctx context.Context) (*SwitchMSwitchIPConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchIPConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindIpConfigs, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchIPConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIpConfigsById
//
// Use this API command to Retrieve IP Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMIPSettingsService) FindIpConfigsById(ctx context.Context, id string) (*SwitchMSwitchIPConfigIpConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchIPConfigIpConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindIpConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchIPConfigIpConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIpConfigsByQueryCriteria
//
// Use this API command to Retrieve IP Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMIPSettingsService) FindIpConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchIPConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchIPConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindIpConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchIPConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateIpConfigsById
//
// Use this API command to Update IP Config.
//
// Request Body:
//	 - body *SwitchMSwitchIPConfigModify
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMIPSettingsService) UpdateIpConfigsById(ctx context.Context, body *SwitchMSwitchIPConfigModify, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateIpConfigsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
