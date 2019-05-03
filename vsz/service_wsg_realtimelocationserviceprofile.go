package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRealTimeLocationServiceProfileService struct {
    client *Client
}

func NewWSGRealTimeLocationServiceProfileService (client *Client) *WSGRealTimeLocationServiceProfileService {
    s := new(WSGRealTimeLocationServiceProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService () *WSGRealTimeLocationServiceProfileService {
    serv := new(WSGRealTimeLocationServiceProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceById (ctx context.Context, id string, zoneId string) (*profile.CreateRtlsProfile, error) {
}

func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceByZoneId (ctx context.Context, zoneId string) (*profile.RtlsProfileList, error) {
}

