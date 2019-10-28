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

func NewWSGDHCPService(client *Client) *WSGDHCPService {
	s := new(WSGDHCPService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDHCPService() *WSGDHCPService {
	serv := new(WSGDHCPService)
	serv.client = ss.client
	return serv
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, pApMac string) (*dhcpmsgstats.DhcpMsgStats, error) {
}

// FindDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, pApMac string) (*dhcppools.DhcpPools, error) {
}

// FindDhcpDataDhcpPoolsByPoolIndex
//
// Use this API command to get AP DHCP Pool Usage by pool's index.
//
// Path Parameters:
// - pApMac string
//		- required
// - pPoolIndex string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, pApMac string, pPoolIndex string) (*dhcppools.DhcpPoolInfo, error) {
}

// FindRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to get DHCP Pool by pool's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, pId string, pZoneId string) (*common.DhcpProfileRef, error) {
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, pZoneId string) (*profile.DhcpProfileList, error) {
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, pZoneId string) (*common.DhcpSiteConfigListRef, error) {
}
