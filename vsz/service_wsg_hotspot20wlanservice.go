package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20WLANServiceService struct {
    client *Client
}

func NewWSGHotspot20WLANServiceService (client *Client) *WSGHotspot20WLANServiceService {
    s := new(WSGHotspot20WLANServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGHotspot20WLANServiceService () *WSGHotspot20WLANServiceService {
    serv := new(WSGHotspot20WLANServiceService)
    serv.client = ss.client
    return serv
}

