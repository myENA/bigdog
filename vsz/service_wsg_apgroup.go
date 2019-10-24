package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apgroup"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
)

type WSGAPGroupService struct {
	client *Client
}

func NewWSGAPGroupService(client *Client) *WSGAPGroupService {
	s := new(WSGAPGroupService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAPGroupService() *WSGAPGroupService {
	serv := new(WSGAPGroupService)
	serv.client = ss.client
	return serv
}

func (s *WSGAPGroupService) AddRkszonesApgroupsMembersByApMac(ctx context.Context, pApMac string, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) AddRkszonesApgroupsMembersById(ctx context.Context, body *apgroup.AddMembers, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsAltitudeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsApMgmtVlanById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsChannelEvaluationIntervalById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationAdditionalInfoById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationBasedServiceById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLteBandLockChannelsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsProtectionMode24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRecoverySsidById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApAggressivenessModeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApJammingThresholdById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApReportThresholdById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsVenueProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelRangeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelWidthById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24TxPowerById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ChannelWidthById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelRangeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50TxPowerById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsApmodelByModel(ctx context.Context, pId string, pModel string, pZoneId string) (*zoneapmodel.ApModel, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsById(ctx context.Context, pId string, pZoneId string) (*apgroup.ApGroupConfiguration, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*apgroup.ApGroupList, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsDefaultByZoneId(ctx context.Context, pZoneId string) (*apgroup.ApGroupConfiguration, error) {
}
