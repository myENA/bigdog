package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserGroupService struct {
    client *Client
}

func NewWSGSCGUserGroupService (client *Client) *WSGSCGUserGroupService {
    s := new(WSGSCGUserGroupService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSCGUserGroupService () *WSGSCGUserGroupService {
    serv := new(WSGSCGUserGroupService)
    serv.client = ss.client
    return serv
}

func (s *WSGSCGUserGroupService) AddUserGroups (ctx context.Context) (*scguser.SCGUserGroupAuditID, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria (ctx context.Context) (*scguser.SCGUserGroupList, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId (ctx context.Context, userGroupId string) (*scguser.SCGUserGroup, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories (ctx context.Context) (*scguser.SCGUserGroupPermissionList, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsRoles (ctx context.Context) (*scguser.SCGUserGroupRoleLabelValueList, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole (ctx context.Context, role string) (*scguser.SCGUserGroupPermissionList, error) {
}

