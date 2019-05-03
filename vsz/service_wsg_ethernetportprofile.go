package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ethernetport"
)

type WSGEthernetPortProfileService struct {
    client *Client
}

func NewWSGEthernetPortProfileService (client *Client) *WSGEthernetPortProfileService {
    s := new(WSGEthernetPortProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGEthernetPortProfileService () *WSGEthernetPortProfileService {
    serv := new(WSGEthernetPortProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById (ctx context.Context, id string, zoneId string) (*ethernetport.EthernetPortProfile, error) {
}

func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId (ctx context.Context, zoneId string) (*ethernetport.ProfileList, error) {
}

