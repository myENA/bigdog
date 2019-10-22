package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGVDPProfileService struct {
    client *Client
}

func NewWSGVDPProfileService (client *Client) *WSGVDPProfileService {
    s := new(WSGVDPProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGVDPProfileService () *WSGVDPProfileService {
    serv := new(WSGVDPProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGVDPProfileService) FindProfilesVdp (ctx context.Context) (*profile.ProfileList, error) {
}

func (s *WSGVDPProfileService) FindProfilesVdpById (ctx context.Context, id string) (*profile.VdpProfile, error) {
}

func (s *WSGVDPProfileService) UpdateProfilesVdpApproveById (ctx context.Context, id string) error {
}

