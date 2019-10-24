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

func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *common.DoAssignIp, pZoneId string) (*common.DhcpSiteConfigListRef, error) {
}

func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDual(ctx context.Context, body *zone.CreateZone) (*common.CreateResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) AddRkszonesIpv6(ctx context.Context, body *zone.CreateZone) (*common.CreateResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesAltitudeById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning24ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning50ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBandBalancingById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl24ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl50ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing24ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing50ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesIpsecProfilesById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLocationBasedServiceById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRecoverySsidById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRogueById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSmartMonitorById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSnmpAgentById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSoftGreTunnelProfliesById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSyslogById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesVenueProfileById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszones(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*zone.ZoneList, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApFirmwareByZoneId(ctx context.Context, pZoneId string) (*zone.ApFirmwareList, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelByModel(ctx context.Context, pModel string, pZoneId string) (*zoneapmodel.ApModel, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelCommonAttributeByModel(ctx context.Context, pModel string, pZoneId string) (*apmodel.CommonAttribute, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableIpsecProfilesByZoneId(ctx context.Context, pZoneId string) (*zone.AvailableTunnelProfileList, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableTunnelProfilesByZoneId(ctx context.Context, pZoneId string) (*zone.AvailableTunnelProfileList, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesById(ctx context.Context, pId string) (*zone.ZoneConfiguration, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesMeshById(ctx context.Context, pId string) (*zone.MeshConfiguration, error) {
}

func (s *WSGRuckusWirelessAPZoneService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *zone.QueryCriteria) (*zone.DhcpSiteConfigList, error) {
}
