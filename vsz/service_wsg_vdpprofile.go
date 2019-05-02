package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGVDPProfileService struct {
    c *Client
}

func NewWSGVDPProfileService (c *Client) *WSGVDPProfileService {
    s := new(WSGVDPProfileService)
    s.c = c
    return s
}

