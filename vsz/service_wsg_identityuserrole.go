package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserRoleService struct {
    client *Client
}

func NewWSGIdentityUserRoleService (client *Client) *WSGIdentityUserRoleService {
    s := new(WSGIdentityUserRoleService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIdentityUserRoleService () *WSGIdentityUserRoleService {
    serv := new(WSGIdentityUserRoleService)
    serv.client = ss.client
    return serv
}

