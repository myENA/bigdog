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

func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *identity.QueryCriteria) (*identity.IdentityList, error) {
}

func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context) (*identity.IdentityList, error) {
}

func (s *WSGIdentityUserRoleService) FindIdentityUserroleById(ctx context.Context, pId string) (*identity.IdentityUserRole, error) {
}
