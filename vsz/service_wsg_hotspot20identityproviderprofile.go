package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGHotspot20IdentityProviderProfileService struct {
    c *Client
}

func NewWSGHotspot20IdentityProviderProfileService (c *Client) *WSGHotspot20IdentityProviderProfileService {
    s := new(WSGHotspot20IdentityProviderProfileService)
    s.c = c
    return s
}

