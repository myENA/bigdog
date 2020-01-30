package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSCGUserGroupService struct {
	apiClient *APIClient
}

func NewWSGSCGUserGroupService(c *APIClient) *WSGSCGUserGroupService {
	s := new(WSGSCGUserGroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCGUserGroupService() *WSGSCGUserGroupService {
	return NewWSGSCGUserGroupService(ss.apiClient)
}

// AddUserGroups
//
// Add SCG user group.
//
// Request Body:
//	 - body *WSGSCGUserGroup
func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *WSGSCGUserGroup) (*WSGSCGUserGroupAuditId, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupAuditId
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUserGroups, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupAuditId()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteUserGroups
//
// Delete multiple SCG user group.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserGroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUserGroups, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// DeleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, userGroupId string) (*WSGSCGUserGroupAuditId, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, userGroupId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUserGroupsByUserGroupId, true)
	req.SetPathParameter("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupAuditId()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria) (*WSGSCGUserGroupList, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUserGroupsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUserGroupsByUserGroupId
//
// Get SCG user group.
//
// Required Parameters:
// - userGroupId string
//		- required
//
// Optional Parameters:
// - includeUsers string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, userGroupId string, optionalParams map[string][]string) (*WSGSCGUserGroup, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, userGroupId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsByUserGroupId, true)
	req.SetPathParameter("userGroupId", userGroupId)
	if v, ok := optionalParams["includeUsers"]; ok {
		req.AddQueryParameter("includeUsers", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroup()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context) (*WSGSCGUserGroupPermissionList, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupPermissionList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsCurrentUserPermissionCategories, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupPermissionList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context) (*WSGSCGUserGroupRoleLabelValueList, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupRoleLabelValueList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsRoles, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupRoleLabelValueList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUserGroupsRolesPermissionsByRole
//
// Get permission items of role.
//
// Required Parameters:
// - role string
//		- required
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, role string, optionalParams map[string][]string) (*WSGSCGUserGroupPermissionList, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupPermissionList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, role, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsRolesPermissionsByRole, true)
	req.SetPathParameter("role", role)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupPermissionList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateUserGroupsByUserGroupId
//
// Update user groups.
//
// Request Body:
//	 - body *WSGSCGUserPatchScgUserGroup
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, userGroupId string) (*WSGSCGUserGroupAuditId, error) {
	var (
		req      *APIRequest
		resp     *WSGSCGUserGroupAuditId
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
	if err = pkgValidator.VarCtx(ctx, userGroupId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUserGroupsByUserGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupAuditId()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
