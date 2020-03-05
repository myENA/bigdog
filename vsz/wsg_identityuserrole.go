package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
func (s *WSGIdentityUserRoleService) AddIdentityUserrole(ctx context.Context, body *WSGIdentityCreateIdentityUserRole) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddIdentityUserrole), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityList
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddIdentityUserRoleList), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteIdentityUserrole
//
// Use this API command to delete multiple identity user roles.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserRoleService) DeleteIdentityUserrole(ctx context.Context, body *WSGIdentityDeleteBulk) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteIdentityUserrole), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteIdentityUserroleById
//
// Use this API command to delete identity user role.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) DeleteIdentityUserroleById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteIdentityUserroleById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindIdentityUserrole
//
// Use this API command to retrieve a list of identity user role.
func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context) (*WSGIdentityList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindIdentityUserrole), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindIdentityUserroleById
//
// Use this API command to retrieve identity user role by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) FindIdentityUserroleById(ctx context.Context, id string) (*WSGIdentityUserRole, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityUserRole
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindIdentityUserroleById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserRole()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateIdentityUserroleById
//
// Use this API command to modify the configuration of identity user role.
//
// Request Body:
//	 - body *WSGIdentityModifyIdentityUserRole
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) PartialUpdateIdentityUserroleById(ctx context.Context, body *WSGIdentityModifyIdentityUserRole, id string) (*APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateIdentityUserroleById), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
