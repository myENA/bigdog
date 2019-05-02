package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPNATProfileService struct {
    c *Client
}

func NewWSGDPNATProfileService (c *Client) *WSGDPNATProfileService {
    s := new(WSGDPNATProfileService)
    s.c = c
    return s
}

