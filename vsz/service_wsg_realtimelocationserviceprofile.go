package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRealTimeLocationServiceProfileService struct {
	client *Client
}

func NewWSGRealTimeLocationServiceProfileService(client *Client) *WSGRealTimeLocationServiceProfileService {
	s := new(WSGRealTimeLocationServiceProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService() *WSGRealTimeLocationServiceProfileService {
	serv := new(WSGRealTimeLocationServiceProfileService)
	serv.client = ss.client
	return serv
}

// FindRkszonesRealTimeLocationServiceById
//
// Use this API command to Get RTLS Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceById(ctx context.Context, pId string, pZoneId string) (*profile.CreateRtlsProfile, error) {
}

// FindRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to Get RTLS Profile by zone ID.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, pZoneId string) (*profile.RtlsProfileList, error) {
}
