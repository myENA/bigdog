package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPNATProfileService struct {
	client *Client
}

func NewWSGDPNATProfileService(client *Client) *WSGDPNATProfileService {
	s := new(WSGDPNATProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDPNATProfileService() *WSGDPNATProfileService {
	serv := new(WSGDPNATProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGDPNATProfileService) FindDpNatProfiles(ctx context.Context) (*dpprofile.DpNatProfileBasicBOList, error) {
}

func (s *WSGDPNATProfileService) FindDpNatProfilesById(ctx context.Context, id string) (*dpprofile.DpNatProfileBasicBO, error) {
}

func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsById(ctx context.Context, id string) (*dpprofile.DpNatProfilePoolBOList, error) {
}

func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string) (*dpprofile.DpNatProfilePoolBO, error) {
}
