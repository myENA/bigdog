package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGZoneAffinityProfileService struct {
    c *Client
}

func NewWSGZoneAffinityProfileService (c *Client) *WSGZoneAffinityProfileService {
    s := new(WSGZoneAffinityProfileService)
    s.c = c
    return s
}

