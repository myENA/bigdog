package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20VenueServiceService struct {
    client *Client
}

func NewWSGHotspot20VenueServiceService (client *Client) *WSGHotspot20VenueServiceService {
    s := new(WSGHotspot20VenueServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGHotspot20VenueServiceService () *WSGHotspot20VenueServiceService {
    serv := new(WSGHotspot20VenueServiceService)
    serv.client = ss.client
    return serv
}

