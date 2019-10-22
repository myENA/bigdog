package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGGuestAccessService struct {
    client *Client
}

func NewWSGGuestAccessService (client *Client) *WSGGuestAccessService {
    s := new(WSGGuestAccessService)
    s.client = client
    return s
}

func (ss *WSGService) WSGGuestAccessService () *WSGGuestAccessService {
    serv := new(WSGGuestAccessService)
    serv.client = ss.client
    return serv
}

func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestRedirectById (ctx context.Context, id string, zoneId string) (*common.EmptyResult, error) {
}

func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestSmsGatewayById (ctx context.Context, id string, zoneId string) (*common.EmptyResult, error) {
}

func (s *WSGGuestAccessService) FindRkszonesPortalsGuestById (ctx context.Context, id string, zoneId string) (*portalservice.GuestAccess, error) {
}

func (s *WSGGuestAccessService) FindRkszonesPortalsGuestByZoneId (ctx context.Context, zoneId string) (*portalservice.PortalServiceList, error) {
}

