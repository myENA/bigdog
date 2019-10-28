package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspotServiceService struct {
	client *Client
}

func NewWSGHotspotServiceService(client *Client) *WSGHotspotServiceService {
	s := new(WSGHotspotServiceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGHotspotServiceService() *WSGHotspotServiceService {
	serv := new(WSGHotspotServiceService)
	serv.client = ss.client
	return serv
}

// AddRkszonesPortalsHotspotExternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with external logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.CreateHotspotExternal
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *portalservice.CreateHotspotExternal, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesPortalsHotspotInternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with internal logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.CreateHotspotInternal
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *portalservice.CreateHotspotInternal, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesPortalsHotspotSmartClientOnlyByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with smart client only of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.CreateHotspotSmartClientOnly
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *portalservice.CreateHotspotSmartClientOnly, pZoneId string) (*common.CreateResult, error) {
}

// FindRkszonesPortalsHotspotById
//
// Use this API command to retrieve a Hotspot (WISPr) of zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot, error) {
}

// FindRkszonesPortalsHotspotByZoneId
//
// Use this API command to retrieve a list of Hotspot (WISPr) of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
}
