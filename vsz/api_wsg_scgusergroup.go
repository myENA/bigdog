package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGScgusergroupService struct {
	apiClient *APIClient
}

func NewWSGScgusergroupService(c *APIClient) *WSGScgusergroupService {
	s := new(WSGScgusergroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGScgusergroupService() *WSGScgusergroupService {
	return NewWSGScgusergroupService(ss.apiClient)
}

// AddUserGroups
//
// Add SCG user group.
//
// Request Body:
//	 - body *WSGSCGUserGroup
func (s *WSGScgusergroupService) AddUserGroups(ctx context.Context, body *WSGSCGUserGroup) (*WSGSCGUserGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUserGroups, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupAuditId()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteUserGroups
//
// Delete multiple SCG user group.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGScgusergroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUserGroups, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// DeleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGScgusergroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, userGroupId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUserGroupsByUserGroupId, true)
	req.SetPathParameter("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGScgusergroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria) (*WSGSCGUserGroupList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGroupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUserGroupsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGScgusergroupService) FindUserGroupsByUserGroupId(ctx context.Context, userGroupId string, optionalParams map[string][]string) (*WSGSCGUserGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsByUserGroupId, true)
	req.SetPathParameter("userGroupId", userGroupId)
	if v, ok := optionalParams["includeUsers"]; ok && len(v) > 0 {
		req.SetQueryParameter("includeUsers", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
func (s *WSGScgusergroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context) (*WSGSCGUserGroupPermissionList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGroupPermissionList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsCurrentUserPermissionCategories, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupPermissionList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
func (s *WSGScgusergroupService) FindUserGroupsRoles(ctx context.Context) (*WSGSCGUserGroupRoleLabelValueList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGroupRoleLabelValueList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsRoles, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupRoleLabelValueList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGScgusergroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, role string, optionalParams map[string][]string) (*WSGSCGUserGroupPermissionList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGroupPermissionList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUserGroupsRolesPermissionsByRole, true)
	req.SetPathParameter("role", role)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGroupPermissionList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGScgusergroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, userGroupId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUserGroupsByUserGroupId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
