package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGVDPProfileService struct {
	client *Client
}

func NewWSGVDPProfileService(client *Client) *WSGVDPProfileService {
	s := new(WSGVDPProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGVDPProfileService() *WSGVDPProfileService {
	serv := new(WSGVDPProfileService)
	serv.client = ss.client
	return serv
}

// FindProfilesVdp
//
// Use this API command to retrieve a list of vdp.
func (s *WSGVDPProfileService) FindProfilesVdp(ctx context.Context, qIndex string, qListSize string) (*profile.ProfileList, error) {
}

// FindProfilesVdpById
//
// Use this API command to retrieve an vdp.
func (s *WSGVDPProfileService) FindProfilesVdpById(ctx context.Context, pId string) (*profile.VdpProfile, error) {
}

// UpdateProfilesVdpApproveById
//
// Use this API command to approve vdp.
func (s *WSGVDPProfileService) UpdateProfilesVdpApproveById(ctx context.Context, pId string) error {
}
