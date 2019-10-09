package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dhcpmsgstats"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dhcppools"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGDHCPService struct {
    client *Client
}

func NewWSGDHCPService (client *Client) *WSGDHCPService {
    s := new(WSGDHCPService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDHCPService () *WSGDHCPService {
    serv := new(WSGDHCPService)
    serv.client = ss.client
    return serv
}

func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac (ctx context.Context, apMac string) (*dhcpmsgstats.DhcpMsgStats, error) {
}

func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac (ctx context.Context, apMac string) (*dhcppools.DhcpPools, error) {
}

func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex (ctx context.Context, apMac string, poolIndex string) (*dhcppools.DhcpPoolInfo, error) {
}

func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById (ctx context.Context, id string, zoneId string) (*common.DhcpProfileRef, error) {
}

func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId (ctx context.Context, zoneId string) (*profile.DhcpProfileList, error) {
}

func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId (ctx context.Context, zoneId string) (*common.DhcpSiteConfigListRef, error) {
}

