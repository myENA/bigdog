package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGWebAuthenticationService struct {
    client *Client
}

func NewWSGWebAuthenticationService (client *Client) *WSGWebAuthenticationService {
    s := new(WSGWebAuthenticationService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWebAuthenticationService () *WSGWebAuthenticationService {
    serv := new(WSGWebAuthenticationService)
    serv.client = ss.client
    return serv
}

func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthRedirectById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthById (ctx context.Context, id string, zoneId string) (portalservice.WebAuthentication, error) {
}

func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthByZoneId (ctx context.Context, zoneId string) (portalservice.PortalServiceList, error) {
}

