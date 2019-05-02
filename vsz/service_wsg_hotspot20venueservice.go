package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20VenueServiceService struct {
    c *Client
}

func NewWSGHotspot20VenueServiceService (c *Client) *WSGHotspot20VenueServiceService {
    s := new(WSGHotspot20VenueServiceService)
    s.c = c
    return s
}

