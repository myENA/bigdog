package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRuckusGRETunnelProfileService struct {
    c *Client
}

func NewWSGRuckusGRETunnelProfileService (c *Client) *WSGRuckusGRETunnelProfileService {
    s := new(WSGRuckusGRETunnelProfileService)
    s.c = c
    return s
}

