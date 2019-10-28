package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPNATProfileService struct {
	client *Client
}

func NewWSGDPDHCPNATProfileService(client *Client) *WSGDPDHCPNATProfileService {
	s := new(WSGDPDHCPNATProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDPDHCPNATProfileService() *WSGDPDHCPNATProfileService {
	serv := new(WSGDPDHCPNATProfileService)
	serv.client = ss.client
	return serv
}

// FindDpProfileSettings
//
// Use this API command to retrieve DP DHCP & NAT profile setting list.
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettings(ctx context.Context) (*dpprofile.DpProfileSettingBOList, error) {
}

// FindDpProfileSettingsByDpKey
//
// Use this API command to retrieve DP DHCP & NAT profile setting.
//
// Path Parameters:
// - pDpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettingsByDpKey(ctx context.Context, pDpKey string) (*dpprofile.DpProfileSettingBO, error) {
}
