package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteUserGroups
//
// Delete multiple SCG user group.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserGroupService) DeleteUserGroups(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteUserGroupsByUserGroupId
//
// Delete SCG user group.
//
// Path Parameters:
// - pUserGroupId string
//		- required
func (s *WSGSCGUserGroupService) DeleteUserGroupsByUserGroupId(ctx context.Context, pUserGroupId string) (*WSGSCGUserGroupAuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria) (*WSGSCGUserGroupList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUserGroupsByUserGroupId
//
// Get SCG user group.
//
// Path Parameters:
// - pUserGroupId string
//		- required
//
// Query Parameters:
// - qIncludeUsers string
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, pUserGroupId string, qIncludeUsers string) (*WSGSCGUserGroup, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context) (*WSGSCGUserGroupPermissionList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context) (*WSGSCGUserGroupRoleLabelValueList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUserGroupsRolesPermissionsByRole
//
// Get permission items of role.
//
// Path Parameters:
// - pRole string
//		- required
//
// Query Parameters:
// - qDomainId string
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, pRole string, qDomainId string) (*WSGSCGUserGroupPermissionList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateUserGroupsByUserGroupId
//
// Update user groups.
//
// Request Body:
//	 - body *WSGSCGUserPatchScgUserGroup
//
// Path Parameters:
// - pUserGroupId string
//		- required
func (s *WSGSCGUserGroupService) PartialUpdateUserGroupsByUserGroupId(ctx context.Context, body *WSGSCGUserPatchScgUserGroup, pUserGroupId string) (*WSGSCGUserGroupAuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
