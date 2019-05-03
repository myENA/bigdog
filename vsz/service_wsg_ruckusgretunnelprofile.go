package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRuckusGRETunnelProfileService struct {
    client *Client
}

func NewWSGRuckusGRETunnelProfileService (client *Client) *WSGRuckusGRETunnelProfileService {
    s := new(WSGRuckusGRETunnelProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGRuckusGRETunnelProfileService () *WSGRuckusGRETunnelProfileService {
    serv := new(WSGRuckusGRETunnelProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgre (ctx context.Context) (*profile.ProfileList, error) {
}

func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreById (ctx context.Context, id string) (*profile.RuckusGREProfile, error) {
}

func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreByQueryCriteria (ctx context.Context) (*profile.RuckusGREProfileList, error) {
}

