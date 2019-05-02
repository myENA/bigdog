package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGIPSECProfileService struct {
    client *Client
}

func NewWSGIPSECProfileService (client *Client) *WSGIPSECProfileService {
    s := new(WSGIPSECProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIPSECProfileService () *WSGIPSECProfileService {
    serv := new(WSGIPSECProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGIPSECProfileService) FindProfilesTunnelIpsec (ctx context.Context) (profile.ProfileList, error) {
}

func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecById (ctx context.Context, id string) (profile.IpsecProfile, error) {
}

func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecByQueryCriteria (ctx context.Context) (profile.IpsecProfileList, error) {
}
