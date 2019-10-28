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

func NewWSGGuestAccessService(client *Client) *WSGGuestAccessService {
	s := new(WSGGuestAccessService)
	s.client = client
	return s
}

func (ss *WSGService) WSGGuestAccessService() *WSGGuestAccessService {
	serv := new(WSGGuestAccessService)
	serv.client = ss.client
	return serv
}

// DeleteRkszonesPortalsGuestRedirectById
//
// Use this API command to set redirect to the URL that user intends to visit on guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestRedirectById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesPortalsGuestSmsGatewayById
//
// Use this API command to disable SMS gateway on guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestSmsGatewayById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesPortalsGuestById
//
// Use this API command to retrieve guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) FindRkszonesPortalsGuestById(ctx context.Context, pId string, pZoneId string) (*portalservice.GuestAccess, error) {
}

// FindRkszonesPortalsGuestByZoneId
//
// Use this API command to retrieve a list of guest access of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) FindRkszonesPortalsGuestByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
}
