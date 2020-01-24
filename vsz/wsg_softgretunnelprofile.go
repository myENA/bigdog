package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGSoftGRETunnelProfileService struct {
	apiClient *APIClient
}

func NewWSGSoftGRETunnelProfileService(c *APIClient) *WSGSoftGRETunnelProfileService {
	s := new(WSGSoftGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSoftGRETunnelProfileService() *WSGSoftGRETunnelProfileService {
	return NewWSGSoftGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelSoftgre
//
// Use this API command to create SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileCreateSoftGREProfile
func (s *WSGSoftGRETunnelProfileService) AddProfilesTunnelSoftgre(ctx context.Context, body *WSGProfileCreateSoftGREProfile) (*WSGCommonCreateResult, error) {
	var (
		resp *WSGCommonCreateResult
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTunnelSoftgre, true)
}

// DeleteProfilesTunnelSoftgre
//
// Use this API command to delete multiple SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelSoftgre, true)
}

// DeleteProfilesTunnelSoftgreById
//
// Use this API command to delete SoftGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgreById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelSoftgreById, true)
}

// FindProfilesTunnelSoftgre
//
// Use this API command to retrieve a list of SoftGRE tunnel profile.
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgre(ctx context.Context) (*WSGProfileList, error) {
	var (
		resp *WSGProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelSoftgre, true)
}

// FindProfilesTunnelSoftgreById
//
// Use this API command to retrieve SoftGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreById(ctx context.Context, id string) (*WSGProfileSoftGREProfile, error) {
	var (
		resp *WSGProfileSoftGREProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelSoftgreById, true)
}

// FindProfilesTunnelSoftgreByQueryCriteria
//
// Use this API command to query a list of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileSoftGREProfileList, error) {
	var (
		resp *WSGProfileSoftGREProfileList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelSoftgreByQueryCriteria, true)
}

// PartialUpdateProfilesTunnelSoftgreById
//
// Use this API command to modify the basic information of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileModifySoftGREProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) PartialUpdateProfilesTunnelSoftgreById(ctx context.Context, body *WSGProfileModifySoftGREProfile, id string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelSoftgreById, true)
}
