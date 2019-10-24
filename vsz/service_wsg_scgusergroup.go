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

func (s *WSGSCGUserGroupService) AddUserGroups(ctx context.Context, body *scguser.ScgUserGroup) (*scguser.ScgUserGroupAuditId, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsByQueryCriteria(ctx context.Context, body *scguser.QueryCriteria) (*scguser.ScgUserGroupList, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsByUserGroupId(ctx context.Context, pUserGroupId string, qIncludeUsers string) (*scguser.ScgUserGroup, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsCurrentUserPermissionCategories(ctx context.Context) (*scguser.ScgUserGroupPermissionList, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsRoles(ctx context.Context) (*scguser.ScgUserGroupRoleLabelValueList, error) {
}

func (s *WSGSCGUserGroupService) FindUserGroupsRolesPermissionsByRole(ctx context.Context, pRole string, qDomainId string) (*scguser.ScgUserGroupPermissionList, error) {
}
