package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMVLANSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMVLANSettingsService(c *VSZClient) *SwitchMVLANSettingsService {
	s := new(SwitchMVLANSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVLANSettingsService() *SwitchMVLANSettingsService {
	return NewSwitchMVLANSettingsService(ss.apiClient)
}

// AddVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *SwitchMSwitchVLANSettingCreateVlanConfig
func (s *SwitchMVLANSettingsService) AddVlans(ctx context.Context, body *SwitchMSwitchVLANSettingCreateVlanConfig) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddVlans, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVLANSettingsService) DeleteVlans(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVlans, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingsService) DeleteVlansById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVlansById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindVlans
//
// Use this API command to Retrieve the VLAN Config List.
func (s *SwitchMVLANSettingsService) FindVlans(ctx context.Context) (*SwitchMSwitchVLANSettingVlanConfigQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchVLANSettingVlanConfigQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVlans, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchVLANSettingVlanConfigQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindVlansById
//
// Use this API command to Retrieve the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingsService) FindVlansById(ctx context.Context, id string) (*SwitchMSwitchVLANSettingVlanConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchVLANSettingVlanConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVlansById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchVLANSettingVlanConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindVlansByQueryCriteria
//
// Use this API command to Retrieve the VLAN Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVLANSettingsService) FindVlansByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchVLANSettingVlanConfigQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchVLANSettingVlanConfigQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindVlansByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchVLANSettingVlanConfigQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *SwitchMSwitchVLANSettingUpdateVlanConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingsService) UpdateVlansById(ctx context.Context, body *SwitchMSwitchVLANSettingUpdateVlanConfig, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateVlansById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
