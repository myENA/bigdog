package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGIPSECProfileService struct {
    c *Client
}

func NewWSGIPSECProfileService (c *Client) *WSGIPSECProfileService {
    s := new(WSGIPSECProfileService)
    s.c = c
    return s
}

