package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20WLANServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspot20WLANServiceService(c *APIClient) *WSGHotspot20WLANServiceService {
	s := new(WSGHotspot20WLANServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20WLANServiceService() *WSGHotspot20WLANServiceService {
	serv := new(WSGHotspot20WLANServiceService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesHs20sByZoneId
//
// Use this API command to create a new Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *portalservice.CreateHotspot20WlanProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) AddRkszonesHs20sByZoneId(ctx context.Context, body *portalservice.CreateHotspot20WlanProfile, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesHs20sById
//
// Use this API command to delete a Hotspot 2.0 WLAN Profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) DeleteRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesHs20sById
//
// Use this API command to retrieve a Hotspot 2.0 WLAN profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot20WlanProfile, error) {
}

// FindRkszonesHs20sByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 WLAN profiles of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
}

// PartialUpdateRkszonesHs20sById
//
// Use this API command to modify the basic information on Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *portalservice.ModifyHotspot20WlanProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) PartialUpdateRkszonesHs20sById(ctx context.Context, body *portalservice.ModifyHotspot20WlanProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
}
