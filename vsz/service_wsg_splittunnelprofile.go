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

func NewWSGSplitTunnelProfileService(client *Client) *WSGSplitTunnelProfileService {
	s := new(WSGSplitTunnelProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSplitTunnelProfileService() *WSGSplitTunnelProfileService {
	serv := new(WSGSplitTunnelProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById(ctx context.Context, pId string, pZoneId string) (*splittunnel.SplitTunnelProfile, error) {
}

func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*splittunnel.SplitTunnelProfileQuery, error) {
}

func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, pZoneId string) (*splittunnel.SplitTunnelProfileList, error) {
}
