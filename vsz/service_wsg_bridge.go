package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBridgeService struct {
    c *Client
}

func NewWSGBridgeService (c *Client) *WSGBridgeService {
    s := new(WSGBridgeService)
    s.c = c
    return s
}

