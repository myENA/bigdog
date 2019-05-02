package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAuthenticationProfileService struct {
    c *Client
}

func NewWSGAuthenticationProfileService (c *Client) *WSGAuthenticationProfileService {
    s := new(WSGAuthenticationProfileService)
    s.c = c
    return s
}

