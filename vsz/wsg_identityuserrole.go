package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGIdentityUserRoleService struct {
	apiClient *APIClient
}

func NewWSGIdentityUserRoleService(c *APIClient) *WSGIdentityUserRoleService {
	s := new(WSGIdentityUserRoleService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityUserRoleService() *WSGIdentityUserRoleService {
	return NewWSGIdentityUserRoleService(ss.apiClient)
}

// AddIdentityUserrole
//
// Use this API command to create identity user role.
//
// Request Body:
//	 - body *WSGIdentityCreateIdentityUserRole
func (s *WSGIdentityUserRoleService) AddIdentityUserrole(ctx context.Context, body *WSGIdentityCreateIdentityUserRole) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserrole, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserRoleList, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteIdentityUserrole
//
// Use this API command to delete multiple identity user roles.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserRoleService) DeleteIdentityUserrole(ctx context.Context, body *WSGIdentityDeleteBulk) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUserrole, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}

// DeleteIdentityUserroleById
//
// Use this API command to delete identity user role.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) DeleteIdentityUserroleById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUserroleById, true)
}

// FindIdentityUserrole
//
// Use this API command to retrieve a list of identity user role.
func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context) (*WSGIdentityList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUserrole, true)
}

// FindIdentityUserroleById
//
// Use this API command to retrieve identity user role by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) FindIdentityUserroleById(ctx context.Context, id string) (*WSGIdentityUserRole, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityUserRole
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUserroleById, true)
}

// PartialUpdateIdentityUserroleById
//
// Use this API command to modify the basic information of identity user role.
//
// Request Body:
//	 - body *WSGIdentityModifyIdentityUserRole
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) PartialUpdateIdentityUserroleById(ctx context.Context, body *WSGIdentityModifyIdentityUserRole, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityUserroleById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
