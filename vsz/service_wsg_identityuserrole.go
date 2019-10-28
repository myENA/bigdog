package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserRoleService struct {
	client *Client
}

func NewWSGIdentityUserRoleService(client *Client) *WSGIdentityUserRoleService {
	s := new(WSGIdentityUserRoleService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIdentityUserRoleService() *WSGIdentityUserRoleService {
	serv := new(WSGIdentityUserRoleService)
	serv.client = ss.client
	return serv
}

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *identity.QueryCriteria) (*identity.IdentityList, error) {
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
