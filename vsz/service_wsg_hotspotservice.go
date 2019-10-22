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

func NewWSGHotspotServiceService (client *Client) *WSGHotspotServiceService {
    s := new(WSGHotspotServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGHotspotServiceService () *WSGHotspotServiceService {
    serv := new(WSGHotspotServiceService)
    serv.client = ss.client
    return serv
}

func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId (ctx context.Context, zoneId string) (*common.CreateResult, error) {
}

func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId (ctx context.Context, zoneId string) (*common.CreateResult, error) {
}

func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId (ctx context.Context, zoneId string) (*common.CreateResult, error) {
}

func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById (ctx context.Context, id string, zoneId string) (*portalservice.Hotspot, error) {
}

func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId (ctx context.Context, zoneId string) (*portalservice.PortalServiceList, error) {
}

