package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGClientIsolationWhitelistService struct {
    c *Client
}

func NewWSGClientIsolationWhitelistService (c *Client) *WSGClientIsolationWhitelistService {
    s := new(WSGClientIsolationWhitelistService)
    s.c = c
    return s
}

