package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspotServiceService struct {
    c *Client
}

func NewWSGHotspotServiceService (c *Client) *WSGHotspotServiceService {
    s := new(WSGHotspotServiceService)
    s.c = c
    return s
}

