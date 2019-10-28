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

// FindDpNatProfiles
//
// Use this API command to retrieve DHCP NAT profile - basic list.
func (s *WSGDPNATProfileService) FindDpNatProfiles(ctx context.Context) (*dpprofile.DpNatProfileBasicBOList, error) {
}

// FindDpNatProfilesById
//
// Use this API command to retrieve DHCP NAT profile - basic.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesById(ctx context.Context, pId string) (*dpprofile.DpNatProfileBasicBO, error) {
}

// FindDpNatProfilesDpNatPoolsById
//
// Use this API command to retrieve DP NAT profile - pool list.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsById(ctx context.Context, pId string) (*dpprofile.DpNatProfilePoolBOList, error) {
}

// FindDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Path Parameters:
// - pId string
//		- required
// - pPoolId string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, pId string, pPoolId string) (*dpprofile.DpNatProfilePoolBO, error) {
}
