package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPProfileService struct {
	apiClient *APIClient
}

func NewWSGDPDHCPProfileService(c *APIClient) *WSGDPDHCPProfileService {
	s := new(WSGDPDHCPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPDHCPProfileService() *WSGDPDHCPProfileService {
	serv := new(WSGDPDHCPProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddDpDhcpProfiles
//
// Use this API command to create basic DP DHCP profile - basic.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfileBasicBO
func (s *WSGDPDHCPProfileService) AddDpDhcpProfiles(ctx context.Context, body *dpprofile.DpDhcpProfileBasicBO) (*dpprofile.DpDhcpProfileBasicBO, error) {
}

// AddDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to create DP DHCP profile - host.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfileHostBO
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *dpprofile.DpDhcpProfileHostBO, pId string) (*dpprofile.DpDhcpProfileHostBO, error) {
}

// AddDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to create DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfileOptionSpaceBO
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *dpprofile.DpDhcpProfileOptionSpaceBO, pId string) error {
}

// AddDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to create DP DHCP profile - pool.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfilePoolBO
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *dpprofile.DpDhcpProfilePoolBO, pId string) (*dpprofile.DpDhcpProfilePoolBO, error) {
}

// DeleteDpDhcpProfiles
//
// Use this API command to delete DP DHCP profiles.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfiles(ctx context.Context, body *dpprofile.BulkDelete) error {
}

// DeleteDpDhcpProfilesById
//
// Use this API command to delete DP DHCP profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesById(ctx context.Context, pId string) error {
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to delete DP DHCP profile - host.
//
// Path Parameters:
// - pHostId string
//		- required
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, pHostId string, pId string) error {
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to delete DP DHCP profile - hosts.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *dpprofile.BulkDelete, pId string) error {
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to delete DP DHCP profile - option43 spaces.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *dpprofile.BulkDelete, pId string) error {
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to delete DP DHCP profile - option43 space.
//
// Path Parameters:
// - pId string
//		- required
// - pSpaceId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, pId string, pSpaceId string) error {
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to delete DP DHCP profile - pools.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *dpprofile.BulkDelete, pId string) error {
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to delete DP DHCP profile - pool.
//
// Path Parameters:
// - pId string
//		- required
// - pPoolId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, pId string, pPoolId string) error {
}

// FindDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context) (*dpprofile.DpDhcpProfileBasicBOList, error) {
}

// FindDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileBasicBO, error) {
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
//
// Path Parameters:
// - pHostId string
//		- required
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, pHostId string, pId string) (*dpprofile.DpDhcpProfileHostBO, error) {
}

// FindDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileHostBOList, error) {
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBOList, error) {
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
//
// Path Parameters:
// - pId string
//		- required
// - pSpaceId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, pId string, pSpaceId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBO, error) {
}

// FindDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfilePoolBOList, error) {
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Path Parameters:
// - pId string
//		- required
// - pPoolId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, pId string, pPoolId string) (*dpprofile.DpDhcpProfilePoolBO, error) {
}

// UpdateDpDhcpProfilesById
//
// Use this API command to modify DP DHCP profile - basic.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfileBasicBO
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesById(ctx context.Context, body *dpprofile.DpDhcpProfileBasicBO, pId string) (*dpprofile.DpDhcpProfileBasicBO, error) {
}

// UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to modify DP DHCP profile - host.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfileHostBO
//
// Path Parameters:
// - pHostId string
//		- required
// - pId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, body *dpprofile.DpDhcpProfileHostBO, pHostId string, pId string) (*dpprofile.DpDhcpProfileHostBO, error) {
}

// UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to update DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfileOptionSpaceBO
//
// Path Parameters:
// - pId string
//		- required
// - pSpaceId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, body *dpprofile.DpDhcpProfileOptionSpaceBO, pId string, pSpaceId string) error {
}

// UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to modify DP DHCP profile - pool.
//
// Request Body:
//	 - body *dpprofile.DpDhcpProfilePoolBO
//
// Path Parameters:
// - pId string
//		- required
// - pPoolId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, body *dpprofile.DpDhcpProfilePoolBO, pId string, pPoolId string) (*dpprofile.DpDhcpProfilePoolBO, error) {
}
