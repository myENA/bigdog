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
		req  *APIRequest
		resp *SwitchMCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddAccessControls, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMAccessControlListService) DeleteAccessControls(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteAccessControls, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
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
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteAccessControlsById, true)
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
		req  *APIRequest
		resp *SwitchMACLConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindAccessControlsById, true)
}

// FindAccessControlsByQueryCriteria
//
// Use this API command to Retrieve the Access Control Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMACLConfigsQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMACLConfigsQueryResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindAccessControlsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
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
		req  *APIRequest
		resp *SwitchMACLConfigEmptyResult
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
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateAccessControlsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
