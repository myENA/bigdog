package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGSCGUserGroupService struct {
	apiClient *VSZClient
}

func NewWSGSCGUserGroupService(c *VSZClient) *WSGSCGUserGroupService {
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
// Operation ID: addUserGroups
// Operation path: /userGroups
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGSCGUserGroup
func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *WSGSCGUserGroup, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddUserGroups, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}

// DeleteUserGroups
//
// Delete multiple SCG user group.
//
// Operation ID: deleteUserGroups
// Operation path: /userGroups
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserGroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUserGroups, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Operation ID: deleteUserGroupsByUserGroupId
// Operation path: /userGroups/{userGroupId}
// Success code: 204 (No Content)
//
// Required parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, userGroupId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("userGroupId", userGroupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Operation ID: findUserGroupsByQueryCriteria
// Operation path: /userGroups/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria, mutators ...RequestMutator) (*WSGSCGUserGroupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindUserGroupsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCGUserGroupListAPIResponse), err
}

// FindUserGroupsByUserGroupId
//
// Get SCG user group.
//
// Operation ID: findUserGroupsByUserGroupId
// Operation path: /userGroups/{userGroupId}
// Success code: 200 (OK)
//
// Required parameters:
// - userGroupId string
//		- required
//
// Optional parameters:
// - includeUsers string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, userGroupId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSCGUserGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("userGroupId", userGroupId)
	if v, ok := optionalParams["includeUsers"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("includeUsers", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCGUserGroupAPIResponse), err
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
//
// Operation ID: findUserGroupsCurrentUserPermissionCategories
// Operation path: /userGroups/currentUser/permissionCategories
// Success code: 200 (OK)
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context, mutators ...RequestMutator) (*WSGSCGUserGroupPermissionListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupPermissionListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsCurrentUserPermissionCategories, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
//
// Operation ID: findUserGroupsRoles
// Operation path: /userGroups/roles
// Success code: 200 (OK)
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context, mutators ...RequestMutator) (*WSGSCGUserGroupRoleLabelValueListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupRoleLabelValueListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsRoles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCGUserGroupRoleLabelValueListAPIResponse), err
}

// FindUserGroupsRolesPermissionsByRole
//
// Get permission items of role.
//
// Operation ID: findUserGroupsRolesPermissionsByRole
// Operation path: /userGroups/roles/{role}/permissions
// Success code: 200 (OK)
//
// Required parameters:
// - role string
//		- required
//
// Optional parameters:
// - domainId string
//		- nullable
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, role string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSCGUserGroupPermissionListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupPermissionListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUserGroupsRolesPermissionsByRole, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("role", role)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
}

// PartialUpdateUserGroupsByUserGroupId
//
// Update user groups.
//
// Operation ID: partialUpdateUserGroupsByUserGroupId
// Operation path: /userGroups/{userGroupId}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSCGUserPatchScgUserGroup
//
// Required parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, userGroupId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.SetBody(body)
	req.PathParams.Set("userGroupId", userGroupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
