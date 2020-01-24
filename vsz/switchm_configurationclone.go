package vsz

// API Version: v8_1

import (
	"context"
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
func (s *SwitchMConfigurationCloneService) AddCloneConfiguration(ctx context.Context, body *SwitchMGroupGetConfigBySwitch) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddCloneConfiguration, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddCloneConfigurationByGroup
//
// Use this API command to Clone Switch Group Config.
//
// Request Body:
//	 - body *SwitchMGroupCloneConfigByGroup
func (s *SwitchMConfigurationCloneService) AddCloneConfigurationByGroup(ctx context.Context, body *SwitchMGroupCloneConfigByGroup) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddCloneConfigurationByGroup, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateCloneConfiguration
//
// Use this API command to Clone Switch Config.
//
// Request Body:
//	 - body *SwitchMGroupCloneConfigBySwitch
func (s *SwitchMConfigurationCloneService) UpdateCloneConfiguration(ctx context.Context, body *SwitchMGroupCloneConfigBySwitch) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateCloneConfiguration, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
