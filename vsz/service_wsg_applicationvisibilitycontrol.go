package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/avc"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGApplicationVisibilityControlService struct {
    c *Client
}

func NewWSGApplicationVisibilityControlService (c *Client) *WSGApplicationVisibilityControlService {
    s := new(WSGApplicationVisibilityControlService)
    s.c = c
    return s
}

