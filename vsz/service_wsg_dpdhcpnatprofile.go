package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPNATProfileService struct {
    c *Client
}

func NewWSGDPDHCPNATProfileService (c *Client) *WSGDPDHCPNATProfileService {
    s := new(WSGDPDHCPNATProfileService)
    s.c = c
    return s
}

