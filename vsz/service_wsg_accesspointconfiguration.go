package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshnodeinfo"
)

type WSGAccessPointConfigurationService struct {
    client *Client
}

func NewWSGAccessPointConfigurationService (client *Client) *WSGAccessPointConfigurationService {
    s := new(WSGAccessPointConfigurationService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccessPointConfigurationService () *WSGAccessPointConfigurationService {
    serv := new(WSGAccessPointConfigurationService)
    serv.client = ss.client
    return serv
}

func (s *WSGAccessPointConfigurationService) DeleteApsAltitudeByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsApMgmtVlanByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsAutoChannelSelection24ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsAutoChannelSelection50ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsBonjourGatewayByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsChannelEvaluationIntervalByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsClientAdmissionControl24ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsClientAdmissionControl50ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromNetworkEnabledByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromWiredClientEnabledByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromWirelessClientEnabledByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsGpsCoordinatesByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsLocationAdditionalInfoByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsLocationByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsLoginByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsMeshOptionsByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsProtectionMode24ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsRecoverySsidByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsRogueApAggressivenessModeByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsRogueApJammingThresholdByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsRogueApReportThresholdByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsSmartMonitorByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsSyslogByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsVenueProfileByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelRangeByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelWidthByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi24TxPowerByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelRangeByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelWidthByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWifi50TxPowerByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWlanGroup24ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) DeleteApsWlanGroup50ByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) FindAps (ctx context.Context) (ap.ApListEntry, error) {
}

func (s *WSGAccessPointConfigurationService) FindApsByApMac (ctx context.Context, apMac string) (ap.ApConfiguration, error) {
}

func (s *WSGAccessPointConfigurationService) FindApsPictureByApMac (ctx context.Context, apMac string) (json.RawMessage, error) {
}

func (s *WSGAccessPointConfigurationService) FindApsSupportLogByApMac (ctx context.Context, apMac string) (json.RawMessage, error) {
}

func (s *WSGAccessPointConfigurationService) FindMeshZeroTouch (ctx context.Context) (meshnodeinfo.MeshNodeInfoList, error) {
}

func (s *WSGAccessPointConfigurationService) UpdateApsRebootByApMac (ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointConfigurationService) UpdateApsSpecificByApMac (ctx context.Context, apMac string) error {
}

