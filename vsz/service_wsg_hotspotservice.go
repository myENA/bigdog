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

func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *portalservice.CreateHotspotExternal, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *portalservice.CreateHotspotInternal, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *portalservice.CreateHotspotSmartClientOnly, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot, error) {
}

func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
}
