package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
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
	serv := new(WSGIdentityUserRoleService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddIdentityUserrole
//
// Use this API command to create identity user role.
//
// Request Body:
//	 - body *identity.CreateIdentityUserRole
func (s *WSGIdentityUserRoleService) AddIdentityUserrole(ctx context.Context, body *identity.CreateIdentityUserRole) (*common.CreateResult, error) {
}

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *identity.QueryCriteria) (*identity.IdentityList, error) {
}

// DeleteIdentityUserrole
//
// Use this API command to delete multiple identity user roles.
//
// Request Body:
//	 - body *identity.DeleteBulk
func (s *WSGIdentityUserRoleService) DeleteIdentityUserrole(ctx context.Context, body *identity.DeleteBulk) error {
}

// DeleteIdentityUserroleById
//
// Use this API command to delete identity user role.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserRoleService) DeleteIdentityUserroleById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindIdentityUserrole
//
// Use this API command to retrieve a list of identity user role.
func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context) (*identity.IdentityList, error) {
}

// FindIdentityUserroleById
//
// Use this API command to retrieve identity user role by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserRoleService) FindIdentityUserroleById(ctx context.Context, pId string) (*identity.IdentityUserRole, error) {
}

// PartialUpdateIdentityUserroleById
//
// Use this API command to modify the basic information of identity user role.
//
// Request Body:
//	 - body *identity.ModifyIdentityUserRole
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserRoleService) PartialUpdateIdentityUserroleById(ctx context.Context, body *identity.ModifyIdentityUserRole, pId string) (*common.EmptyResult, error) {
}
