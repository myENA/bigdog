package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addUserGroups
//
// Add SCG user group.
//
// Request Body:
//	 - body *WSGSCGUserGroup
func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *WSGSCGUserGroup, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddUserGroups, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}

// DeleteUserGroups
//
// Operation ID: deleteUserGroups
//
// Delete multiple SCG user group.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserGroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteUserGroups, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteUserGroupsByUserGroupId
//
// Operation ID: deleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, userGroupId string, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}

// FindUserGroupsByQueryCriteria
//
// Operation ID: findUserGroupsByQueryCriteria
//
// Query user groups.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria, mutators ...RequestMutator) (*WSGSCGUserGroupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindUserGroupsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSCGUserGroupListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupListAPIResponse), err
}

// FindUserGroupsByUserGroupId
//
// Operation ID: findUserGroupsByUserGroupId
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
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, userGroupId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSCGUserGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("userGroupId", userGroupId)
	if v, ok := optionalParams["includeUsers"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("includeUsers", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAPIResponse), err
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Operation ID: findUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context, mutators ...RequestMutator) (*WSGSCGUserGroupPermissionListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupPermissionListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUserGroupsCurrentUserPermissionCategories, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
}

// FindUserGroupsRoles
//
// Operation ID: findUserGroupsRoles
//
// Get pre-defined roles.
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context, mutators ...RequestMutator) (*WSGSCGUserGroupRoleLabelValueListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupRoleLabelValueListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupRoleLabelValueListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUserGroupsRoles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupRoleLabelValueListAPIResponse), err
}

// FindUserGroupsRolesPermissionsByRole
//
// Operation ID: findUserGroupsRolesPermissionsByRole
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
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, role string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSCGUserGroupPermissionListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupPermissionListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUserGroupsRolesPermissionsByRole, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("role", role)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupPermissionListAPIResponse), err
}

// PartialUpdateUserGroupsByUserGroupId
//
// Operation ID: partialUpdateUserGroupsByUserGroupId
//
// Update user groups.
//
// Request Body:
//	 - body *WSGSCGUserPatchScgUserGroup
//
// Required Parameters:
// - userGroupId string
//		- required
func (s *WSGSCGUserGroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, userGroupId string, mutators ...RequestMutator) (*WSGSCGUserGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGroupAuditIdAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateUserGroupsByUserGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
	}
	req.PathParams.Set("userGroupId", userGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGroupAuditIdAPIResponse), err
}
