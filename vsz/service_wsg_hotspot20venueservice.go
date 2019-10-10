package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20VenueServiceService struct {
	client *Client
}

func NewWSGHotspot20VenueServiceService(client *Client) *WSGHotspot20VenueServiceService {
	s := new(WSGHotspot20VenueServiceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGHotspot20VenueServiceService() *WSGHotspot20VenueServiceService {
	serv := new(WSGHotspot20VenueServiceService)
	serv.client = ss.client
	return serv
}

func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesById(ctx context.Context, id string, zoneId string) (*portalservice.Hotspot20VeuneProfile, error) {
}

func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesByZoneId(ctx context.Context, zoneId string) (*portalservice.PortalServiceList, error) {
}
