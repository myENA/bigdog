package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apgroup"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
)

type WSGAPGroupService struct {
    client *Client
}

func NewWSGAPGroupService (client *Client) *WSGAPGroupService {
    s := new(WSGAPGroupService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAPGroupService () *WSGAPGroupService {
    serv := new(WSGAPGroupService)
    serv.client = ss.client
    return serv
}

func (s *WSGAPGroupService) AddRkszonesApgroupsMembersByApMac (ctx context.Context, apMac string, id string, zoneId string) error {
}

func (s *WSGAPGroupService) AddRkszonesApgroupsMembersById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsAltitudeById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsApMgmtVlanById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection24ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection50ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsChannelEvaluationIntervalById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl24ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl50ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationAdditionalInfoById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationBasedServiceById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsProtectionMode24ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRecoverySsidById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApAggressivenessModeById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApJammingThresholdById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApReportThresholdById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsVenueProfileById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelRangeById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelWidthById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24TxPowerById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ChannelWidthById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelRangeById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50TxPowerById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup24ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup50ById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsApmodelByModel (ctx context.Context, id string, model string, zoneId string) (zoneapmodel.ApModel, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsById (ctx context.Context, id string, zoneId string) (apgroup.ApGroupConfiguration, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsByZoneId (ctx context.Context, zoneId string) (apgroup.ApGroupList, error) {
}

func (s *WSGAPGroupService) FindRkszonesApgroupsDefaultByZoneId (ctx context.Context, zoneId string) (apgroup.ApGroupConfiguration, error) {
}

