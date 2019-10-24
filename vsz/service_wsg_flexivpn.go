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

func NewWSGFlexiVPNService(client *Client) *WSGFlexiVPNService {
	s := new(WSGFlexiVPNService)
	s.client = client
	return s
}

func (ss *WSGService) WSGFlexiVPNService() *WSGFlexiVPNService {
	serv := new(WSGFlexiVPNService)
	serv.client = ss.client
	return serv
}

// DeleteRkszonesWlansFlexiVpnProfileById
//
// Use this API command to delete Flexi-VPN on WLAN
func (s *WSGFlexiVPNService) DeleteRkszonesWlansFlexiVpnProfileById(ctx context.Context, pId string, pZoneId string) error {
}

// FindServicesFlexiVpnProfileByQueryCriteria
//
// Use this API command to query Flexi-VPN profiles.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGFlexiVPNService) FindServicesFlexiVpnProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.FlexiVpnProfileList, error) {
}
