package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGZoneAffinityProfileService struct {
	client *Client
}

func NewWSGZoneAffinityProfileService(client *Client) *WSGZoneAffinityProfileService {
	s := new(WSGZoneAffinityProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGZoneAffinityProfileService() *WSGZoneAffinityProfileService {
	serv := new(WSGZoneAffinityProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, qVdpId string) (*profile.ZoneAffinityProfileList, error) {
}

func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, pId string) (*profile.ReturnZoneAffinityProfile, error) {
}
