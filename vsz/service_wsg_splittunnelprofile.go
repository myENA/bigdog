package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/splittunnel"
)

type WSGSplitTunnelProfileService struct {
    client *Client
}

func NewWSGSplitTunnelProfileService (client *Client) *WSGSplitTunnelProfileService {
    s := new(WSGSplitTunnelProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSplitTunnelProfileService () *WSGSplitTunnelProfileService {
    serv := new(WSGSplitTunnelProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById (ctx context.Context, id string, zoneId string) (*splittunnel.SplitTunnelProfile, error) {
}

func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria (ctx context.Context) (*splittunnel.SplitTunnelProfileQuery, error) {
}

func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId (ctx context.Context, zoneId string) (*splittunnel.SplitTunnelProfileList, error) {
}

