package vsz

// API Version: v8_1

import (
	"context"
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

func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById (ctx context.Context, id string, zoneId string) (*portalservice.Hotspot20WLANProfile, error) {
}

func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId (ctx context.Context, zoneId string) (*portalservice.PortalServiceList, error) {
}

