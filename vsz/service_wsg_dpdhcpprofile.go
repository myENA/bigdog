package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPProfileService struct {
    c *Client
}

func NewWSGDPDHCPProfileService (c *Client) *WSGDPDHCPProfileService {
    s := new(WSGDPDHCPProfileService)
    s.c = c
    return s
}

