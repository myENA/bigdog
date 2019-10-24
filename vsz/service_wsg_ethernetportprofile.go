package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ethernetport"
)

type WSGEthernetPortProfileService struct {
	client *Client
}

func NewWSGEthernetPortProfileService(client *Client) *WSGEthernetPortProfileService {
	s := new(WSGEthernetPortProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGEthernetPortProfileService() *WSGEthernetPortProfileService {
	serv := new(WSGEthernetPortProfileService)
	serv.client = ss.client
	return serv
}

// FindRkszonesProfileEthernetPortById
//
// Retrieve a Ethernet Port Porfile.
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById(ctx context.Context, pId string, pZoneId string) (*ethernetport.EthernetPortProfile, error) {
}

// FindRkszonesProfileEthernetPortByZoneId
//
// Retrieve a list of Ethernet Port Porfiles within a zone.
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*ethernetport.ProfileList, error) {
}
