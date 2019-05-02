package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserService struct {
    c *Client
}

func NewWSGIdentityUserService (c *Client) *WSGIdentityUserService {
    s := new(WSGIdentityUserService)
    s.c = c
    return s
}

