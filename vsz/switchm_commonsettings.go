package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMCommonSettingsService struct {
	apiClient *APIClient
}

func NewSwitchMCommonSettingsService(c *APIClient) *SwitchMCommonSettingsService {
	s := new(SwitchMCommonSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMCommonSettingsService() *SwitchMCommonSettingsService {
	return NewSwitchMCommonSettingsService(ss.apiClient)
}

// AddDnsConfig
//
// Use this API command to Create DNS Config.
//
// Request Body:
//	 - body *SwitchMDnsConfigCreateDnsConfig
func (s *SwitchMCommonSettingsService) AddDnsConfig(ctx context.Context, body *SwitchMDnsConfigCreateDnsConfig) (*SwitchMCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddDnsConfig, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteDnsConfigBySwitchGroupId
//
// Use this API command to Delete DNS Config.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMCommonSettingsService) DeleteDnsConfigBySwitchGroupId(ctx context.Context, switchGroupId string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, switchGroupId, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteDnsConfigBySwitchGroupId, true)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindDnsConfigBySwitchGroupId
//
// Use this API command to Retrieve DNS Config.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMCommonSettingsService) FindDnsConfigBySwitchGroupId(ctx context.Context, switchGroupId string) (*SwitchMDnsConfig, error) {
	var (
		req      *APIRequest
		resp     *SwitchMDnsConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, switchGroupId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindDnsConfigBySwitchGroupId, true)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDnsConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateDnsConfigBySwitchGroupId
//
// Use this API command to Update DNS Config.
//
// Request Body:
//	 - body *SwitchMDnsConfigUpdateDnsConfig
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMCommonSettingsService) UpdateDnsConfigBySwitchGroupId(ctx context.Context, body *SwitchMDnsConfigUpdateDnsConfig, switchGroupId string) (*SwitchMDnsConfigEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMDnsConfigEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, switchGroupId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateDnsConfigBySwitchGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDnsConfigEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
