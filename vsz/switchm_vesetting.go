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
		req      *APIRequest
		resp     SwitchMVeConfigCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddVeConfigs, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSwitchMVeConfigCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, resp, err)
}

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVESettingService) DeleteVeConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVeConfigs, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVESettingService) DeleteVeConfigsById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVeConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindVeConfigs
//
// Use this API command to Retrieve VE Config List.
func (s *SwitchMVESettingService) FindVeConfigs(ctx context.Context) (*SwitchMVeConfigList, error) {
	var (
		req      *APIRequest
		resp     *SwitchMVeConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVeConfigs, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfigList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
		req      *APIRequest
		resp     *SwitchMVeConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVeConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindVeConfigsByQueryCriteria
//
// Use this API command to Retrieve VE Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVeConfigList, error) {
	var (
		req      *APIRequest
		resp     *SwitchMVeConfigList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindVeConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfigList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
		req      *APIRequest
		resp     *SwitchMVeConfigEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateVeConfigsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfigEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
