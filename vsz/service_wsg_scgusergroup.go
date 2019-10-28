package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserGroupService struct {
	client *Client
}

func NewWSGSCGUserGroupService(client *Client) *WSGSCGUserGroupService {
	s := new(WSGSCGUserGroupService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSCGUserGroupService() *WSGSCGUserGroupService {
	serv := new(WSGSCGUserGroupService)
	serv.client = ss.client
	return serv
}

// AddUserGroups
//
// Add SCG user group.
//
// Request Body:
//	 - body *scguser.ScgUserGroup
func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *scguser.ScgUserGroup) (*scguser.ScgUserGroupAuditId, error) {
}

// FindUserGroupsByQueryCriteria
//
// Query user groups.
//
// Request Body:
//	 - body *scguser.QueryCriteria
func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *scguser.QueryCriteria) (*scguser.ScgUserGroupList, error) {
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
func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, pUserGroupId string, qIncludeUsers string) (*scguser.ScgUserGroup, error) {
}

// FindUserGroupsCurrentUserPermissionCategories
//
// Get permitted categories of current user.
func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context) (*scguser.ScgUserGroupPermissionList, error) {
}

// FindUserGroupsRoles
//
// Get pre-defined roles.
func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context) (*scguser.ScgUserGroupRoleLabelValueList, error) {
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
func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, pRole string, qDomainId string) (*scguser.ScgUserGroupPermissionList, error) {
}
