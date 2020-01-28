package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMAccessControlListService struct {
	apiClient *APIClient
}

func NewSwitchMAccessControlListService(c *APIClient) *SwitchMAccessControlListService {
	s := new(SwitchMAccessControlListService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAccessControlListService() *SwitchMAccessControlListService {
	return NewSwitchMAccessControlListService(ss.apiClient)
}

// AddAccessControls
//
// Use this API command to Create the Access Control Config.
//
// Request Body:
//	 - body *SwitchMACLConfigCreateACLConfig
func (s *SwitchMAccessControlListService) AddAccessControls(ctx context.Context, body *SwitchMACLConfigCreateACLConfig) (*SwitchMCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddAccessControls, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// DeleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMAccessControlListService) DeleteAccessControls(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteAccessControls, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// DeleteAccessControlsById
//
// Use this API command to Delete the Access Control Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) DeleteAccessControlsById(ctx context.Context, id string) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteAccessControlsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindAccessControlsById
//
// Use this API command to Retrieve the Access Control Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) FindAccessControlsById(ctx context.Context, id string) (*SwitchMACLConfig, error) {
	var (
		req      *APIRequest
		resp     *SwitchMACLConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindAccessControlsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMACLConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAccessControlsByQueryCriteria
//
// Use this API command to Retrieve the Access Control Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMACLConfigsQueryResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMACLConfigsQueryResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindAccessControlsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMACLConfigsQueryResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateAccessControlsById
//
// Use this API command to Update the Access Control Config.
//
// Request Body:
//	 - body *SwitchMACLConfigUpdateACLConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) UpdateAccessControlsById(ctx context.Context, body *SwitchMACLConfigUpdateACLConfig, id string) (*SwitchMACLConfigEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMACLConfigEmptyResult
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
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateAccessControlsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMACLConfigEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
