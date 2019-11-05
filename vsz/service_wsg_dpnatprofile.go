package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPNATProfileService struct {
	apiClient *APIClient
}

func NewWSGDPNATProfileService(c *APIClient) *WSGDPNATProfileService {
	s := new(WSGDPNATProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPNATProfileService() *WSGDPNATProfileService {
	serv := new(WSGDPNATProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddDpNatProfiles
//
// Use this API command to create DHCP NAT profile - basic.
//
// Request Body:
//	 - body *dpprofile.DpNatProfileBasicBO
func (s *WSGDPNATProfileService) AddDpNatProfiles(ctx context.Context, body *dpprofile.DpNatProfileBasicBO) (*dpprofile.DpNatProfileBasicBO, error) {
}

// AddDpNatProfilesDpNatPoolsById
//
// Use this API command to create DHCP NAT profile - pool.
//
// Request Body:
//	 - body *dpprofile.DpNatProfilePoolBO
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPNATProfileService) AddDpNatProfilesDpNatPoolsById(ctx context.Context, body *dpprofile.DpNatProfilePoolBO, pId string) (*dpprofile.DpNatProfilePoolBO, error) {
}

// DeleteDpNatProfiles
//
// Use this API command to delete DHCP NAT profiles.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
func (s *WSGDPNATProfileService) DeleteDpNatProfiles(ctx context.Context, body *dpprofile.BulkDelete) error {
}

// DeleteDpNatProfilesById
//
// Use this API command to delete DHCP NAT profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesById(ctx context.Context, pId string) error {
}

// DeleteDpNatProfilesDpNatPoolsById
//
// Use this API command to delete DP NAT profile - pools.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsById(ctx context.Context, body *dpprofile.BulkDelete, pId string) error {
}

// DeleteDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to delete DP NAT profile - pool.
//
// Path Parameters:
// - pId string
//		- required
// - pPoolId string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, pId string, pPoolId string) error {
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

// UpdateDpNatProfilesById
//
// Use this API command to modify DHCP NAT profile - basic.
//
// Request Body:
//	 - body *dpprofile.DpNatProfileBasicBO
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesById(ctx context.Context, body *dpprofile.DpNatProfileBasicBO, pId string) (*dpprofile.DpNatProfileBasicBO, error) {
}

// UpdateDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to modify DHCP NAT profile - pool.
//
// Request Body:
//	 - body *dpprofile.DpNatProfilePoolBO
//
// Path Parameters:
// - pId string
//		- required
// - pPoolId string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, body *dpprofile.DpNatProfilePoolBO, pId string, pPoolId string) (*dpprofile.DpNatProfilePoolBO, error) {
}
