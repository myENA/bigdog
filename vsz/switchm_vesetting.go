package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMVESettingService struct {
	apiClient *APIClient
}

func NewSwitchMVESettingService(c *APIClient) *SwitchMVESettingService {
	s := new(SwitchMVESettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVESettingService() *SwitchMVESettingService {
	return NewSwitchMVESettingService(ss.apiClient)
}

// AddVeConfigs
//
// Use this API command to Create VE Config.
//
// Request Body:
//	 - body *SwitchMVeConfigCreate
func (s *SwitchMVESettingService) AddVeConfigs(ctx context.Context, body *SwitchMVeConfigCreate) (SwitchMVeConfigCreateResult, error) {
	var (
		resp SwitchMVeConfigCreateResult
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
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddVeConfigs, true)
}

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVESettingService) DeleteVeConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVeConfigs, true)
}

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVESettingService) DeleteVeConfigsById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVeConfigsById, true)
}

// FindVeConfigs
//
// Use this API command to Retrieve VE Config List.
func (s *SwitchMVESettingService) FindVeConfigs(ctx context.Context) (*SwitchMVeConfigList, error) {
	var (
		resp *SwitchMVeConfigList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindVeConfigs, true)
}

// FindVeConfigsById
//
// Use this API command to Retrieve VE Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVESettingService) FindVeConfigsById(ctx context.Context, id string) (*SwitchMVeConfig, error) {
	var (
		resp *SwitchMVeConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindVeConfigsById, true)
}

// FindVeConfigsByQueryCriteria
//
// Use this API command to Retrieve VE Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVeConfigList, error) {
	var (
		resp *SwitchMVeConfigList
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
	req := NewAPIRequest(http.MethodPost, RouteSwitchMFindVeConfigsByQueryCriteria, true)
}

// UpdateVeConfigsById
//
// Use this API command to Update VE Config.
//
// Request Body:
//	 - body *SwitchMVeConfigModify
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVESettingService) UpdateVeConfigsById(ctx context.Context, body *SwitchMVeConfigModify, id string) (*SwitchMVeConfigEmptyResult, error) {
	var (
		resp *SwitchMVeConfigEmptyResult
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
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateVeConfigsById, true)
}
