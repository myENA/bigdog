package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
func (s *SwitchMVESettingService) AddVeConfigs(ctx context.Context, body *SwitchMVeConfigCreate) (SwitchMVeConfigCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SwitchMVeConfigCreateResult
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMAddVeConfigs), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSwitchMVeConfigCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVESettingService) DeleteVeConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMDeleteVeConfigs), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVESettingService) DeleteVeConfigsById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMDeleteVeConfigsById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindVeConfigs
//
// Use this API command to Retrieve VE Config List.
func (s *SwitchMVESettingService) FindVeConfigs(ctx context.Context) (*SwitchMVeConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMVeConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindVeConfigs), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindVeConfigsById
//
// Use this API command to Retrieve VE Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVESettingService) FindVeConfigsById(ctx context.Context, id string) (*SwitchMVeConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMVeConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindVeConfigsById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindVeConfigsByQueryCriteria
//
// Use this API command to Retrieve VE Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVeConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMVeConfigList
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindVeConfigsByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SwitchMVESettingService) UpdateVeConfigsById(ctx context.Context, body *SwitchMVeConfigModify, id string) (*SwitchMVeConfigEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMVeConfigEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMUpdateVeConfigsById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVeConfigEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
