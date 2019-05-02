package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGLBSprofileService struct {
    c *Client
}

func NewWSGLBSprofileService (c *Client) *WSGLBSprofileService {
    s := new(WSGLBSprofileService)
    s.c = c
    return s
}

