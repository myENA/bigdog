package vsz

// API Version: v9_0

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
func (s *SwitchMCommonSettingsService) AddDnsConfig(ctx context.Context, body *SwitchMDnsConfigCreateDnsConfig) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddDnsConfig, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDnsConfigBySwitchGroupId
//
// Use this API command to Delete DNS Config.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMCommonSettingsService) DeleteDnsConfigBySwitchGroupId(ctx context.Context, switchGroupId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, switchGroupId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteDnsConfigBySwitchGroupId, true)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindDnsConfigBySwitchGroupId
//
// Use this API command to Retrieve DNS Config.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMCommonSettingsService) FindDnsConfigBySwitchGroupId(ctx context.Context, switchGroupId string) (*SwitchMDnsConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDnsConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, switchGroupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindDnsConfigBySwitchGroupId, true)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDnsConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SwitchMCommonSettingsService) UpdateDnsConfigBySwitchGroupId(ctx context.Context, body *SwitchMDnsConfigUpdateDnsConfig, switchGroupId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, switchGroupId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateDnsConfigBySwitchGroupId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
