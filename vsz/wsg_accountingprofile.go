package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGAccountingProfileService struct {
	apiClient *APIClient
}

func NewWSGAccountingProfileService(c *APIClient) *WSGAccountingProfileService {
	s := new(WSGAccountingProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountingProfileService() *WSGAccountingProfileService {
	return NewWSGAccountingProfileService(ss.apiClient)
}

// AddProfilesAcct
//
// Use this API command to create a new accounting profile.
//
// Request Body:
//	 - body *WSGProfileCreateAccountingProfile
func (s *WSGAccountingProfileService) AddProfilesAcct(ctx context.Context, body *WSGProfileCreateAccountingProfile) (*WSGCommonCreateResult, error) {
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAcct, true)
}

// AddProfilesAcctCloneById
//
// Use this API command to clone an accounting profile.
//
// Request Body:
//	 - body *WSGProfileCloneRequest
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) AddProfilesAcctCloneById(ctx context.Context, body *WSGProfileCloneRequest, id string) (*WSGProfileCloneResponse, error) {
	var (
		resp *WSGProfileCloneResponse
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAcctCloneById, true)
}

// DeleteProfilesAcct
//
// Use this API command to delete a list of accounting profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkAccountingProfile
func (s *WSGAccountingProfileService) DeleteProfilesAcct(ctx context.Context, body *WSGProfileDeleteBulkAccountingProfile) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAcct, true)
}

// DeleteProfilesAcctById
//
// Use this API command to delete an accounting profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) DeleteProfilesAcctById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAcctById, true)
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context) (*WSGProfileAccountingProfileList, error) {
	var (
		resp *WSGProfileAccountingProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAcct, true)
}

// FindProfilesAcctById
//
// Use this API command to retrieve an accounting profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) FindProfilesAcctById(ctx context.Context, id string) (*WSGProfileAccountingProfile, error) {
	var (
		resp *WSGProfileAccountingProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAcctById, true)
}

// FindProfilesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting profiles by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileAccountingProfileList, error) {
	var (
		resp *WSGProfileAccountingProfileList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindProfilesAcctByQueryCriteria, true)
}

// PartialUpdateProfilesAcctById
//
// Use this API command to modify the basic information of an accounting profile.
//
// Request Body:
//	 - body *WSGProfileModifyAccountingProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) PartialUpdateProfilesAcctById(ctx context.Context, body *WSGProfileModifyAccountingProfile, id string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesAcctById, true)
}
