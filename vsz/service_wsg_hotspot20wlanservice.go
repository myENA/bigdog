package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20WLANServiceService struct {
    c *Client
}

func NewWSGHotspot20WLANServiceService (c *Client) *WSGHotspot20WLANServiceService {
    s := new(WSGHotspot20WLANServiceService)
    s.c = c
    return s
}

