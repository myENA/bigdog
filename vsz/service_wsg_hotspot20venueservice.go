package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20VenueServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspot20VenueServiceService(c *APIClient) *WSGHotspot20VenueServiceService {
	s := new(WSGHotspot20VenueServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20VenueServiceService() *WSGHotspot20VenueServiceService {
	serv := new(WSGHotspot20VenueServiceService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesHs20VenuesByZoneId
//
// Use this API command to create a new Hotspot 2.0 venue profile of a zone.
//
// Request Body:
//	 - body *portalservice.CreateHotspot20VenueProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) AddRkszonesHs20VenuesByZoneId(ctx context.Context, body *portalservice.CreateHotspot20VenueProfile, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesHs20VenuesById
//
// Use this API command to delete Hotspot 2.0 venue profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) DeleteRkszonesHs20VenuesById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
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

// PartialUpdateRkszonesHs20VenuesById
//
// Use this API command to modify the basic information on Hotspot 2.0 venue profile of a zone.
//
// Request Body:
//	 - body *portalservice.ModifyHotspot20VenueProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) PartialUpdateRkszonesHs20VenuesById(ctx context.Context, body *portalservice.ModifyHotspot20VenueProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
}
