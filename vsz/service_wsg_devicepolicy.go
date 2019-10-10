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

func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyById(ctx context.Context, id string, zoneId string) (*devicepolicy.DevicePolicyPorfile, error) {
}

func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyByZoneId(ctx context.Context, zoneId string) (*devicepolicy.PorfileList, error) {
}
