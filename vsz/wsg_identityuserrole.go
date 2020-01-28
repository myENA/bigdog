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
		req      *APIRequest
		resp     *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserrole, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityList, error) {
	var (
		req      *APIRequest
		resp     *WSGIdentityList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserRoleList, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// DeleteIdentityUserrole
//
// Use this API command to delete multiple identity user roles.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserRoleService) DeleteIdentityUserrole(ctx context.Context, body *WSGIdentityDeleteBulk) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUserrole, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
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
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUserroleById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindIdentityUserrole
//
// Use this API command to retrieve a list of identity user role.
func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context) (*WSGIdentityList, error) {
	var (
		req      *APIRequest
		resp     *WSGIdentityList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUserrole, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
		req      *APIRequest
		resp     *WSGIdentityUserRole
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUserroleById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserRole()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
		req      *APIRequest
		resp     *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityUserroleById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
