package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGFlexiVPNService struct {
    client *Client
}

func NewWSGFlexiVPNService (client *Client) *WSGFlexiVPNService {
    s := new(WSGFlexiVPNService)
    s.client = client
    return s
}

func (ss *WSGService) WSGFlexiVPNService () *WSGFlexiVPNService {
    serv := new(WSGFlexiVPNService)
    serv.client = ss.client
    return serv
}

func (s *WSGFlexiVPNService) DeleteRkszonesWlansFlexiVpnProfileById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGFlexiVPNService) FindServicesFlexiVpnProfileByQueryCriteria (ctx context.Context) (*profile.FlexiVpnProfileList, error) {
}

