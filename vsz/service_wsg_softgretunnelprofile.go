package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGSoftGRETunnelProfileService struct {
    c *Client
}

func NewWSGSoftGRETunnelProfileService (c *Client) *WSGSoftGRETunnelProfileService {
    s := new(WSGSoftGRETunnelProfileService)
    s.c = c
    return s
}

