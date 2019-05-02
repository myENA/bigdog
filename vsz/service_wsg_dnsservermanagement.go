package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGDNSServerManagementService struct {
    c *Client
}

func NewWSGDNSServerManagementService (c *Client) *WSGDNSServerManagementService {
    s := new(WSGDNSServerManagementService)
    s.c = c
    return s
}

