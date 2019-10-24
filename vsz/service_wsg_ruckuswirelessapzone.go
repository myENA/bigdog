package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
)

type WSGRuckusWirelessAPZoneService struct {
	client *Client
}

func NewWSGRuckusWirelessAPZoneService(client *Client) *WSGRuckusWirelessAPZoneService {
	s := new(WSGRuckusWirelessAPZoneService)
	s.client = client
	return s
}

func (ss *WSGService) WSGRuckusWirelessAPZoneService() *WSGRuckusWirelessAPZoneService {
	serv := new(WSGRuckusWirelessAPZoneService)
	serv.client = ss.client
	return serv
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *common.DoAssignIp
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *common.DoAssignIp, pZoneId string) (*common.DhcpSiteConfigListRef, error) {
}

// AddRkszonesDual
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv4/IPv6.
//
// Request Body:
//	 - body *zone.CreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDual(ctx context.Context, body *zone.CreateZone) (*common.CreateResult, error) {
}

// AddRkszonesIpv6
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv6.
//
// Request Body:
//	 - body *zone.CreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesIpv6(ctx context.Context, body *zone.CreateZone) (*common.CreateResult, error) {
}

// DeleteRkszonesAltitudeById
//
// Use this API command to disable altitude configuration of zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesAltitudeById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesBackgroundScanning24ById
//
// Use this API command to disable background scanning 2.4GHz radio configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning24ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesBackgroundScanning50ById
//
// Use this API command to disable background scanning 5GHz radio configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning50ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesBandBalancingById
//
// Use this API command to disable band balancing for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBandBalancingById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl24ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl50ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesClientLoadBalancing24ById
//
// Use this API command to disable client load balancing 2.4GHz radio configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing24ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesClientLoadBalancing50ById
//
// Use this API command to disable client load balancing 5GHz radio configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing50ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesIpsecProfilesById
//
// Use this API command to Delete IPsec profiles.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesIpsecProfilesById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesLocationBasedServiceById
//
// Use this API command to disable location based service for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLocationBasedServiceById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesRecoverySsidById
//
// Use this API command to clear recovery ssid setting of a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRecoverySsidById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesRogueById
//
// Use this API command to disable rogue AP detection for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRogueById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesSmartMonitorById
//
// Use this API command to disable smart monitor for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSmartMonitorById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesSnmpAgentById
//
// Use this API command to clear SNMPv2 and SNMPv3 agent that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSnmpAgentById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesSoftGreTunnelProfliesById
//
// Use this API command to Delete IPsec profiles.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSoftGreTunnelProfliesById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesSyslogById
//
// Use this API command to disable syslog configuration for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSyslogById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesVenueProfileById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindRkszones
//
// Use this API command to retrieve the list of Ruckus Wireless AP zones that belong to a domain.
func (s *WSGRuckusWirelessAPZoneService) FindRkszones(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*zone.ZoneList, error) {
}

// FindRkszonesApFirmwareByZoneId
//
// Use this API command to retrieve AP Firmware the list that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApFirmwareByZoneId(ctx context.Context, pZoneId string) (*zone.ApFirmwareList, error) {
}

// FindRkszonesApmodelByModel
//
// Use this API command to retrieve AP model specific configuration that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelByModel(ctx context.Context, pModel string, pZoneId string) (*zoneapmodel.ApModel, error) {
}

// FindRkszonesApmodelCommonAttributeByModel
//
// Use this API command to retrieve AP model common attribute that belong to a zone.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelCommonAttributeByModel(ctx context.Context, pModel string, pZoneId string) (*apmodel.CommonAttribute, error) {
}

// FindRkszonesAvailableIpsecProfilesByZoneId
//
// Get available IPSec tunnel profiles of this Zone.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableIpsecProfilesByZoneId(ctx context.Context, pZoneId string) (*zone.AvailableTunnelProfileList, error) {
}

// FindRkszonesAvailableTunnelProfilesByZoneId
//
// Get available GRE tunnel profiles of this Zone.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableTunnelProfilesByZoneId(ctx context.Context, pZoneId string) (*zone.AvailableTunnelProfileList, error) {
}

// FindRkszonesById
//
// Use this API command to retrieve Ruckus Wireless AP zones configuration.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesById(ctx context.Context, pId string) (*zone.ZoneConfiguration, error) {
}

// FindRkszonesMeshById
//
// Use this API command to retrieve the mesh configuration of a zone.
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesMeshById(ctx context.Context, pId string) (*zone.MeshConfiguration, error) {
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *zone.QueryCriteria
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *zone.QueryCriteria) (*zone.DhcpSiteConfigList, error) {
}
