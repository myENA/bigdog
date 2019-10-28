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

// FindRkszonesHs20VenuesById
//
// Use this API command to retrieve a Hotspot 2.0 venue profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot20VeuneProfile, error) {
}

// FindRkszonesHs20VenuesByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 venue profile of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
}
