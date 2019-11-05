package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspotServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspotServiceService(c *APIClient) *WSGHotspotServiceService {
	s := new(WSGHotspotServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspotServiceService() *WSGHotspotServiceService {
	serv := new(WSGHotspotServiceService)
	serv.apiClient = ss.apiClient
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

// DeleteRkszonesPortalsHotspotById
//
// Use this API command to delete a Hotspot (WISPr) of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) DeleteRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
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

// PartialUpdateRkszonesPortalsHotspotById
//
// Use this API command to modify the basic information on Hotspot (WISPr) of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.ModifyHotspot
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) PartialUpdateRkszonesPortalsHotspotById(ctx context.Context, body *portalservice.ModifyHotspot, pId string, pZoneId string) (*common.EmptyResult, error) {
}
