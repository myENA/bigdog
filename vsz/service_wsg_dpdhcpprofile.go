package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPProfileService struct {
    client *Client
}

func NewWSGDPDHCPProfileService (client *Client) *WSGDPDHCPProfileService {
    s := new(WSGDPDHCPProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDPDHCPProfileService () *WSGDPDHCPProfileService {
    serv := new(WSGDPDHCPProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles (ctx context.Context) (*dpprofile.DpDhcpProfileBasicBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById (ctx context.Context, id string) (*dpprofile.DpDhcpProfileBasicBO, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId (ctx context.Context, hostId string, id string) (*dpprofile.DpDhcpProfileHostBO, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById (ctx context.Context, id string) (*dpprofile.DpDhcpProfileHostBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById (ctx context.Context, id string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId (ctx context.Context, id string, spaceId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBO, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById (ctx context.Context, id string) (*dpprofile.DpDhcpProfilePoolBOList, error) {
}

func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId (ctx context.Context, id string, poolId string) (*dpprofile.DpDhcpProfilePoolBO, error) {
}

