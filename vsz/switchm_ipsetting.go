package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMIPSettingService struct {
	apiClient *APIClient
}

func NewSwitchMIPSettingService(c *APIClient) *SwitchMIPSettingService {
	s := new(SwitchMIPSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMIPSettingService() *SwitchMIPSettingService {
	return NewSwitchMIPSettingService(ss.apiClient)
}

// AddIpConfigs
//
// Use this API command to Create IP Config.
//
// Request Body:
//	 - body *SwitchMIpConfigCreate
func (s *SwitchMIPSettingService) AddIpConfigs(ctx context.Context, body *SwitchMIpConfigCreate) (SwitchMIpConfigCreateResult, error) {
	var (
		req  *APIRequest
		resp SwitchMIpConfigCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddIpConfigs, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteIpConfigs
//
// Use this API command to Delete IP Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMIPSettingService) DeleteIpConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteIpConfigs, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}

// DeleteIpConfigsById
//
// Use this API command to Delete IP Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMIPSettingService) DeleteIpConfigsById(ctx context.Context, id string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteIpConfigsById, true)
}

// FindIpConfigs
//
// Use this API command to Retrieve IP Config List.
func (s *SwitchMIPSettingService) FindIpConfigs(ctx context.Context) (*SwitchMIpConfigList, error) {
	var (
		req  *APIRequest
		resp *SwitchMIpConfigList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindIpConfigs, true)
}

// FindIpConfigsById
//
// Use this API command to Retrieve IP Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMIPSettingService) FindIpConfigsById(ctx context.Context, id string) (*SwitchMIpConfig, error) {
	var (
		req  *APIRequest
		resp *SwitchMIpConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindIpConfigsById, true)
}

// FindIpConfigsByQueryCriteria
//
// Use this API command to Retrieve IP Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMIPSettingService) FindIpConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMIpConfigList, error) {
	var (
		req  *APIRequest
		resp *SwitchMIpConfigList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindIpConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateIpConfigsById
//
// Use this API command to Update IP Config.
//
// Request Body:
//	 - body *SwitchMIpConfigModify
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMIPSettingService) UpdateIpConfigsById(ctx context.Context, body *SwitchMIpConfigModify, id string) (*SwitchMIpConfigEmptyResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMIpConfigEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateIpConfigsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
