package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/devicepolicy"
)

type WSGDevicePolicyService struct {
	client *Client
}

func NewWSGDevicePolicyService(client *Client) *WSGDevicePolicyService {
	s := new(WSGDevicePolicyService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDevicePolicyService() *WSGDevicePolicyService {
	serv := new(WSGDevicePolicyService)
	serv.client = ss.client
	return serv
}

// FindRkszonesDevicePolicyById
//
// Retrieve a Device Policy Porfile.
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*devicepolicy.DevicePolicyPorfile, error) {
}

// FindRkszonesDevicePolicyByZoneId
//
// Retrieve a list of Device Policy Porfiles within a zone.
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*devicepolicy.PorfileList, error) {
}
