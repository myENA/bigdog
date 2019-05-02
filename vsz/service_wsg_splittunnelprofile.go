package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/splittunnel"
)

type WSGSplitTunnelProfileService struct {
    c *Client
}

func NewWSGSplitTunnelProfileService (c *Client) *WSGSplitTunnelProfileService {
    s := new(WSGSplitTunnelProfileService)
    s.c = c
    return s
}

