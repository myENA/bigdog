package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGZoneAffinityProfileService struct {
    client *Client
}

func NewWSGZoneAffinityProfileService (client *Client) *WSGZoneAffinityProfileService {
    s := new(WSGZoneAffinityProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGZoneAffinityProfileService () *WSGZoneAffinityProfileService {
    serv := new(WSGZoneAffinityProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity (ctx context.Context) (profile.ZoneAffinityProfileList, error) {
}

func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById (ctx context.Context, id string) (profile.ReturnZoneAffinityProfile, error) {
}

