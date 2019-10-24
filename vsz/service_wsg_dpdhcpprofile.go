package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPProfileService struct {
	client *Client
}

func NewWSGDPDHCPProfileService(client *Client) *WSGDPDHCPProfileService {
	s := new(WSGDPDHCPProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDPDHCPProfileService() *WSGDPDHCPProfileService {
	serv := new(WSGDPDHCPProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context) (*dpprofile.DpDhcpProfileBasicBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileBasicBO, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, pHostId string, pId string) (*dpprofile.DpDhcpProfileHostBO, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileHostBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, pId string, pSpaceId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBO, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfilePoolBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, pId string, pPoolId string) (*dpprofile.DpDhcpProfilePoolBO, error) {
}
