package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apgroup"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAPGroupService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGAPGroupService(c *APIClient) *WSGAPGroupService {
	s := new(WSGAPGroupService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGAPGroupService() *WSGAPGroupService {
	return NewWSGAPGroupService(ss.apiClient)
}

// AddRkszonesApgroupsByZoneId
//
// Use this API command to create new AP group within a zone.
//
// Request Body:
//	 - body *apgroup.CreateAPGroup
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsByZoneId(ctx context.Context, body *apgroup.CreateAPGroup, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// AddRkszonesApgroupsMembersByApMac
//
// Use this API command to add a member AP to an AP group.
//
// Path Parameters:
// - pApMac string
//		- required
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsMembersByApMac(ctx context.Context, pApMac string, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesApgroupsMembersById
//
// Add multiple members to an AP group.
//
// Request Body:
//	 - body *apgroup.AddMembers
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsMembersById(ctx context.Context, body *apgroup.AddMembers, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteRkszonesApgroupsAltitudeById
//
// Use this API command to clear the altitude of AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAltitudeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsApMgmtVlanById
//
// Disable AP Management Vlan Override of an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsApMgmtVlanById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsApmodelByModel
//
// Use this API command to disable AP model specific configuration override zone that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pModel string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsApmodelByModel(ctx context.Context, pId string, pModel string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsAutoChannelSelection24ById
//
// Disable Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsAutoChannelSelection50ById
//
// Disable Radio 5G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsById
//
// Use this API command to delete an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsChannelEvaluationIntervalById
//
// Disable Channel Evaluation Interval Override of an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsChannelEvaluationIntervalById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById
//
// Use this API command to disable Directed Multicast from Network to wired/wireless client configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById
//
// Use this API command to disable Directed Multicast from wired client to Network configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById
//
// Use this API command to disable Directed Multicast from wireless client to Network configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsLocationAdditionalInfoById
//
// Use this API command to disable location additionalInfo override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationAdditionalInfoById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsLocationBasedServiceById
//
// Use this API command to disable location based service override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationBasedServiceById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsLocationById
//
// Use this API command to disable location override for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsLteBandLockChannelsById
//
// Use this API command to disable LTE band lock channel override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLteBandLockChannelsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsMembersByApMac
//
// Use this API command to remove a member AP from an AP group.
//
// Path Parameters:
// - pApMac string
//		- required
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsMembersByApMac(ctx context.Context, pApMac string, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsProtectionMode24ById
//
// Use this API command to disable 2.4GHz radio protection mode configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsProtectionMode24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsRecoverySsidById
//
// Use this API command to disable Recovery SSID configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRecoverySsidById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsRogueApAggressivenessModeById
//
// Use this API command to disable rogue AP aggressiveness mode override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApAggressivenessModeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsRogueApJammingThresholdById
//
// Use this API command to disable rogue AP jamming threshold override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApJammingThresholdById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsRogueApReportThresholdById
//
// Use this API command to disable rogue AP report threshold override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApReportThresholdById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsVenueProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi24ById
//
// Use this API command to disable 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi24ChannelById
//
// Use this API command to disable 2.4GHz radio channel override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi24ChannelRangeById
//
// Use this API command to disable 2.4GHz radio channelRange override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelRangeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi24ChannelWidthById
//
// Use this API command to disable 2.4GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelWidthById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi24TxPowerById
//
// Use this API command to disable 2.4GHz radio txPower override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24TxPowerById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50ById
//
// Use this API command to disable 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50ChannelWidthById
//
// Use this API command to disable 5GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ChannelWidthById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50IndoorChannelById
//
// Use this API command to disable 5GHz radio indoorChannel override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50IndoorChannelRangeById
//
// Use this API command to disable 5GHz radio indoorChannelRange override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelRangeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50OutdoorChannelById
//
// Use this API command to disable 5GHz radio outdoorChannel override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById
//
// Use this API command to disable 5GHz radio outdoorChannelRange override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWifi50TxPowerById
//
// Use this API command to disable 5GHz radio txPower override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50TxPowerById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWlanGroup24ById
//
// Use this API command to disable WLAN group on 2.4GHz radio override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup24ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesApgroupsWlanGroup50ById
//
// Use this API command to disable WLAN group on 5GHz radio override zone for APs that belong to an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup50ById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApgroupsApmodelByModel
//
// Use this API command to retrieve AP model specific configuration override zone that belong to an AP group, NULL means the override setting is not checked inside an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pModel string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsApmodelByModel(ctx context.Context, pId string, pModel string, pZoneId string) (*zoneapmodel.ApModel, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApgroupsById
//
// Use this API command to retrieve information about an AP group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsById(ctx context.Context, pId string, pZoneId string) (*apgroup.ApGroupConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApgroupsByZoneId
//
// Use this API command to retrieve the list of AP groups that belong to a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGAPGroupService) FindRkszonesApgroupsByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*apgroup.ApGroupList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApgroupsDefaultByZoneId
//
// Use this API command to retrieve information about default AP group of zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsDefaultByZoneId(ctx context.Context, pZoneId string) (*apgroup.ApGroupConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesApgroupsById
//
// Use this API command to modify the basic information of an AP group.
//
// Request Body:
//	 - body *apgroup.ModifyAPGroup
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) PartialUpdateRkszonesApgroupsById(ctx context.Context, body *apgroup.ModifyAPGroup, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// UpdateRkszonesApgroupsApmodelByModel
//
// Use this API command to modify AP model specific configuration override zone that belong to an AP group.
//
// Request Body:
//	 - body *zoneapmodel.ApModel
//
// Path Parameters:
// - pId string
//		- required
// - pModel string
//		- required
// - pZoneId string
//		- required
func (s *WSGAPGroupService) UpdateRkszonesApgroupsApmodelByModel(ctx context.Context, body *zoneapmodel.ApModel, pId string, pModel string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
