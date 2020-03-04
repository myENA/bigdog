package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type SwitchMConfigurationCloneService struct {
	apiClient *APIClient
}

func NewSwitchMConfigurationCloneService(c *APIClient) *SwitchMConfigurationCloneService {
	s := new(SwitchMConfigurationCloneService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationCloneService() *SwitchMConfigurationCloneService {
	return NewSwitchMConfigurationCloneService(ss.apiClient)
}

// AddCloneConfiguration
//
// Use this API command to Get Switch Config.
//
// Request Body:
//	 - body *SwitchMGroupGetConfigBySwitch
func (s *SwitchMConfigurationCloneService) AddCloneConfiguration(ctx context.Context, body *SwitchMGroupGetConfigBySwitch) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.switchMPath, RouteSwitchMAddCloneConfiguration), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddCloneConfigurationByGroup
//
// Use this API command to Clone Switch Group Config.
//
// Request Body:
//	 - body *SwitchMGroupCloneConfigByGroup
func (s *SwitchMConfigurationCloneService) AddCloneConfigurationByGroup(ctx context.Context, body *SwitchMGroupCloneConfigByGroup) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.switchMPath, RouteSwitchMAddCloneConfigurationByGroup), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// UpdateCloneConfiguration
//
// Use this API command to Clone Switch Config.
//
// Request Body:
//	 - body *SwitchMGroupCloneConfigBySwitch
func (s *SwitchMConfigurationCloneService) UpdateCloneConfiguration(ctx context.Context, body *SwitchMGroupCloneConfigBySwitch) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
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
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("/%s%s", s.apiClient.switchMPath, RouteSwitchMUpdateCloneConfiguration), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
