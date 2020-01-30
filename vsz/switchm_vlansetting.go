package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMVLANSettingService struct {
	apiClient *APIClient
}

func NewSwitchMVLANSettingService(c *APIClient) *SwitchMVLANSettingService {
	s := new(SwitchMVLANSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVLANSettingService() *SwitchMVLANSettingService {
	return NewSwitchMVLANSettingService(ss.apiClient)
}

// AddVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVlanConfigCreateVlanConfig
func (s *SwitchMVLANSettingService) AddVlans(ctx context.Context, body *SwitchMVlanConfigCreateVlanConfig) (*SwitchMCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddVlans, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVLANSettingService) DeleteVlans(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVlans, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// DeleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingService) DeleteVlansById(ctx context.Context, id string) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVlansById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindVlans
//
// Use this API command to Retrieve the VLAN Config List.
func (s *SwitchMVLANSettingService) FindVlans(ctx context.Context) (*SwitchMVlanConfigQueryResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMVlanConfigQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVlans, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVlanConfigQueryResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindVlansById
//
// Use this API command to Retrieve the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingService) FindVlansById(ctx context.Context, id string) (*SwitchMVlanConfig, error) {
	var (
		req      *APIRequest
		resp     *SwitchMVlanConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVlansById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVlanConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindVlansByQueryCriteria
//
// Use this API command to Retrieve the VLAN Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVLANSettingService) FindVlansByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVlanConfigQueryResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMVlanConfigQueryResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindVlansByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVlanConfigQueryResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVlanConfigUpdateVlanConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingService) UpdateVlansById(ctx context.Context, body *SwitchMVlanConfigUpdateVlanConfig, id string) (*SwitchMVlanConfigEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMVlanConfigEmptyResult
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
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateVlansById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMVlanConfigEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
