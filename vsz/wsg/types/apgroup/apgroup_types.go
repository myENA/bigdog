package apgroup

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type AddMembers struct {
	// MemberList
	// List of apMac
	MemberList []*ApGroupMember `json:"memberList,omitempty"`
}

type ApGroupConfiguration struct {
	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the ap group
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// Description
	// Description of the AP group
	Description *string `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// ID
	// Identifier of the AP group
	ID *string `json:"id,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.OverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// Members
	// Members of the AP group
	Members []*ApGroupMember `json:"members,omitempty"`

	// Name
	// Name of the AP group
	Name *string `json:"name,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Wifi24 *common.Radio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50SuperSet `json:"wifi50,omitempty"`

	WLANGroup24 *common.GenericRef `json:"wlanGroup24,omitempty"`

	WLANGroup50 *common.GenericRef `json:"wlanGroup50,omitempty"`

	// ZoneID
	// Identifier of the zone to which the AP group belongs
	ZoneID *string `json:"zoneId,omitempty"`
}

type ApGroupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApGroupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApGroupMember struct {
	ApMac *common.Mac `json:"apMac,omitempty"`
}

type ApGroupSummary struct {
	// ID
	// Identifier of the AP group
	ID *string `json:"id,omitempty"`

	// Name
	// Description of the AP group
	Name *string `json:"name,omitempty"`
}

type CreateAPGroup struct {
}

type ModifyAPGroup struct {
	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the ap group
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network to wired / wireless client.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client to network.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client to network.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.OverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Wifi24 *common.Radio24 `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50 `json:"wifi50,omitempty"`

	WLANGroup24 *common.GenericRef `json:"wlanGroup24,omitempty"`

	WLANGroup50 *common.GenericRef `json:"wlanGroup50,omitempty"`
}
