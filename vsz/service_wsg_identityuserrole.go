package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserRoleService struct {
    c *Client
}

func NewWSGIdentityUserRoleService (c *Client) *WSGIdentityUserRoleService {
    s := new(WSGIdentityUserRoleService)
    s.c = c
    return s
}

