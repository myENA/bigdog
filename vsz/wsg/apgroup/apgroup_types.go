package apgroup

// API Version: v8_0

type AddMembers struct {
	MemberList []*ApGroupMember `json:"memberList,omitempty"`
}

type ApGroupConfiguration struct {
	Altitude                                   *common.Altitude                       `json:"altitude,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan               `json:"apMgmtVlan,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection           `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection           `json:"autoChannelSelection50,omitempty"`
	AwsVenue                                   *string                                `json:"awsVenue,omitempty"`
	ChannelEvaluationInterval                  *int                                   `json:"channelEvaluationInterval,omitempty"`
	ClientAdmissionControl24                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	Description                                *string                                `json:"description,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                                  `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                                  `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                                  `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	ID                                         *string                                `json:"id,omitempty"`
	Latitude                                   *float64                               `json:"latitude,omitempty"`
	Location                                   *string                                `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                                `json:"locationAdditionalInfo,omitempty"`
	LocationBasedService                       *common.OverrideGenericRef             `json:"locationBasedService,omitempty"`
	Longitude                                  *float64                               `json:"longitude,omitempty"`
	Members                                    []*ApGroupMember                       `json:"members,omitempty"`
	Name                                       *string                                `json:"name,omitempty"`
	ProtectionMode24                           *string                                `json:"protectionMode24,omitempty"`
	RecoverySsid                               *common.RecoverySsid                   `json:"recoverySsid,omitempty"`
	RogueApAggressivenessMode                  *int                                   `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingThreshold                    *int                                   `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                                   `json:"rogueApReportThreshold,omitempty"`
	VenueProfile                               *common.GenericRef                     `json:"venueProfile,omitempty"`
	Wifi24                                     *common.Radio24SuperSet                `json:"wifi24,omitempty"`
	Wifi50                                     *common.Radio50SuperSet                `json:"wifi50,omitempty"`
	WLANGroup24                                *common.GenericRef                     `json:"wlanGroup24,omitempty"`
	WLANGroup50                                *common.GenericRef                     `json:"wlanGroup50,omitempty"`
	ZoneID                                     *string                                `json:"zoneId,omitempty"`
}

type ApGroupList struct {
	FirstIndex *int              `json:"firstIndex,omitempty"`
	HasMore    *bool             `json:"hasMore,omitempty"`
	List       []*ApGroupSummary `json:"list,omitempty"`
	TotalCount *int              `json:"totalCount,omitempty"`
}

type ApGroupMember struct {
	ApMac *string `json:"apMac,omitempty"`
}

type ApGroupSummary struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateAPGroup struct {
}

type ModifyAPGroup struct {
	Altitude                                   *common.Altitude                       `json:"altitude,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan               `json:"apMgmtVlan,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection           `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection           `json:"autoChannelSelection50,omitempty"`
	AwsVenue                                   *string                                `json:"awsVenue,omitempty"`
	ChannelEvaluationInterval                  *int                                   `json:"channelEvaluationInterval,omitempty"`
	ClientAdmissionControl24                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	Description                                *string                                `json:"description,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                                  `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                                  `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                                  `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	Latitude                                   *float64                               `json:"latitude,omitempty"`
	Location                                   *string                                `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                                `json:"locationAdditionalInfo,omitempty"`
	LocationBasedService                       *common.OverrideGenericRef             `json:"locationBasedService,omitempty"`
	Longitude                                  *float64                               `json:"longitude,omitempty"`
	Name                                       *string                                `json:"name,omitempty"`
	ProtectionMode24                           *string                                `json:"protectionMode24,omitempty"`
	RecoverySsid                               *common.RecoverySsid                   `json:"recoverySsid,omitempty"`
	RogueApAggressivenessMode                  *int                                   `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingThreshold                    *int                                   `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                                   `json:"rogueApReportThreshold,omitempty"`
	VenueProfile                               *common.GenericRef                     `json:"venueProfile,omitempty"`
	Wifi24                                     *common.Radio24                        `json:"wifi24,omitempty"`
	Wifi50                                     *common.Radio50                        `json:"wifi50,omitempty"`
	WLANGroup24                                *common.GenericRef                     `json:"wlanGroup24,omitempty"`
	WLANGroup50                                *common.GenericRef                     `json:"wlanGroup50,omitempty"`
}
