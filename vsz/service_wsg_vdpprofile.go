package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGVDPProfileService struct {
	apiClient *APIClient
}

func NewWSGVDPProfileService(c *APIClient) *WSGVDPProfileService {
	s := new(WSGVDPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVDPProfileService() *WSGVDPProfileService {
	serv := new(WSGVDPProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// DeleteProfilesVdpById
//
// Use this API command to delete an vdp.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) DeleteProfilesVdpById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesVdp
//
// Use this API command to retrieve a list of vdp.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGVDPProfileService) FindProfilesVdp(ctx context.Context, qIndex string, qListSize string) (*profile.ProfileList, error) {
}

// FindProfilesVdpById
//
// Use this API command to retrieve an vdp.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) FindProfilesVdpById(ctx context.Context, pId string) (*profile.VdpProfile, error) {
}

// UpdateProfilesVdpApproveById
//
// Use this API command to approve vdp.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) UpdateProfilesVdpApproveById(ctx context.Context, pId string) error {
}
