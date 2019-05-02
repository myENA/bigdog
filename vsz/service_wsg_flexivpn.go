package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGFlexiVPNService struct {
    c *Client
}

func NewWSGFlexiVPNService (c *Client) *WSGFlexiVPNService {
    s := new(WSGFlexiVPNService)
    s.c = c
    return s
}

