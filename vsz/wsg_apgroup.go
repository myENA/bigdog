package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAPGroupService struct {
	apiClient *APIClient
}

func NewWSGAPGroupService(c *APIClient) *WSGAPGroupService {
	s := new(WSGAPGroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAPGroupService() *WSGAPGroupService {
	return NewWSGAPGroupService(ss.apiClient)
}

type WSGAPGroupAddMembers struct {
	// MemberList
	// List of apMac
	MemberList []*WSGAPGroupMember `json:"memberList,omitempty"`
}

func NewWSGAPGroupAddMembers() *WSGAPGroupAddMembers {
	m := new(WSGAPGroupAddMembers)
	return m
}

type WSGAPGroupConfiguration struct {
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApGroupRoguePolicy
	// Override Rogue AP policy ID (only for monitoring group).
	ApGroupRoguePolicy *string `json:"apGroupRoguePolicy,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the ap group
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// Description
	// Description of the AP group
	Description *string `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// Id
	// Identifier of the AP group
	Id *string `json:"id,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonOverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	// Members
	// Members of the AP group
	Members []*WSGAPGroupMember `json:"members,omitempty"`

	// MonitoringEnabled
	// Set AP group as monitoring group.
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty"`

	// Name
	// Name of the AP group
	// Constraints:
	//    - required
	Name *string `json:"name"`

	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	PaloAltoFirewallProfileOverrided *bool `json:"paloAltoFirewallProfileOverrided,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// RogueScanFreq
	// Set rogue scan frequency (only for monitoring group).
	// Constraints:
	//    - oneof:[LOW,MEDIUM,HIGH]
	RogueScanFreq *string `json:"rogueScanFreq,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Wifi24 *WSGCommonRadio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50SuperSet `json:"wifi50,omitempty"`

	WlanGroup24 *WSGCommonGenericRef `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WSGCommonGenericRef `json:"wlanGroup50,omitempty"`

	// ZoneId
	// Identifier of the zone to which the AP group belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAPGroupConfiguration() *WSGAPGroupConfiguration {
	m := new(WSGAPGroupConfiguration)
	return m
}

type WSGAPGroupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPGroupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPGroupList() *WSGAPGroupList {
	m := new(WSGAPGroupList)
	return m
}

type WSGAPGroupMember struct {
	ApMac *WSGCommonMac `json:"apMac,omitempty"`

	// ApSerial
	// Serial number of member AP
	ApSerial *string `json:"apSerial,omitempty"`
}

func NewWSGAPGroupMember() *WSGAPGroupMember {
	m := new(WSGAPGroupMember)
	return m
}

type WSGAPGroupSummary struct {
	// Id
	// Identifier of the AP group
	Id *string `json:"id,omitempty"`

	// Name
	// Description of the AP group
	Name *string `json:"name,omitempty"`
}

func NewWSGAPGroupSummary() *WSGAPGroupSummary {
	m := new(WSGAPGroupSummary)
	return m
}

type WSGAPGroupCreateAPGroup struct {
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApGroupRoguePolicy
	// Override Rogue AP policy ID (only for monitoring group).
	ApGroupRoguePolicy *string `json:"apGroupRoguePolicy,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the ap group
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network to wired / wireless client.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client to network.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client to network.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonOverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	// MonitoringEnabled
	// Set AP group as monitoring group.
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	PaloAltoFirewallProfileOverrided *bool `json:"paloAltoFirewallProfileOverrided,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// RogueScanFreq
	// Set rogue scan frequency (only for monitoring group).
	// Constraints:
	//    - oneof:[LOW,MEDIUM,HIGH]
	RogueScanFreq *string `json:"rogueScanFreq,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50 `json:"wifi50,omitempty"`

	WlanGroup24 *WSGCommonGenericRef `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WSGCommonGenericRef `json:"wlanGroup50,omitempty"`
}

func NewWSGAPGroupCreateAPGroup() *WSGAPGroupCreateAPGroup {
	m := new(WSGAPGroupCreateAPGroup)
	return m
}

type WSGAPGroupModifyAPGroup struct {
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApGroupRoguePolicy
	// Override Rogue AP policy ID (only for monitoring group).
	ApGroupRoguePolicy *string `json:"apGroupRoguePolicy,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the ap group
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network to wired / wireless client.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client to network.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client to network.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonOverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	// MonitoringEnabled
	// Set AP group as monitoring group.
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	PaloAltoFirewallProfileOverrided *bool `json:"paloAltoFirewallProfileOverrided,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// RogueScanFreq
	// Set rogue scan frequency (only for monitoring group).
	// Constraints:
	//    - oneof:[LOW,MEDIUM,HIGH]
	RogueScanFreq *string `json:"rogueScanFreq,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50 `json:"wifi50,omitempty"`

	WlanGroup24 *WSGCommonGenericRef `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WSGCommonGenericRef `json:"wlanGroup50,omitempty"`
}

func NewWSGAPGroupModifyAPGroup() *WSGAPGroupModifyAPGroup {
	m := new(WSGAPGroupModifyAPGroup)
	return m
}

// AddRkszonesApgroupsByZoneId
//
// Use this API command to create new AP group within a zone.
//
// Request Body:
//	 - body *WSGAPGroupCreateAPGroup
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsByZoneId(ctx context.Context, body *WSGAPGroupCreateAPGroup, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesApgroupsByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesApgroupsMembersByApMac
//
// Use this API command to add a member AP to an AP group.
//
// Required Parameters:
// - apMac string
//		- required
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsMembersByApMac(ctx context.Context, apMac string, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesApgroupsMembersByApMac, true)
	req.SetPathParameter("apMac", apMac)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// AddRkszonesApgroupsMembersById
//
// Add multiple members to an AP group.
//
// Request Body:
//	 - body *WSGAPGroupAddMembers
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsMembersById(ctx context.Context, body *WSGAPGroupAddMembers, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesApgroupsMembersById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsAltitudeById
//
// Use this API command to clear the altitude of AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAltitudeById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAltitudeById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsApMgmtVlanById
//
// Disable AP Management Vlan Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsApMgmtVlanById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsApMgmtVlanById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsApmodelByModel
//
// Use this API command to disable AP model specific configuration override zone that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsApmodelByModel(ctx context.Context, id string, model string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, model, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsApmodelByModel, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("model", model)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsAutoChannelSelection24ById
//
// Disable Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection24ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAutoChannelSelection24ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsAutoChannelSelection50ById
//
// Disable Radio 5G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection50ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAutoChannelSelection50ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsAwsVenueById
//
// Use this API command to disable AWS venue override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAwsVenueById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAwsVenueById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsById
//
// Use this API command to delete an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsChannelEvaluationIntervalById
//
// Disable Channel Evaluation Interval Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsChannelEvaluationIntervalById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsChannelEvaluationIntervalById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl24ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsClientAdmissionControl24ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl50ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsClientAdmissionControl50ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById
//
// Use this API command to disable Directed Multicast from Network to wired/wireless client configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById
//
// Use this API command to disable Directed Multicast from wired client to Network configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById
//
// Use this API command to disable Directed Multicast from wireless client to Network configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsLocationAdditionalInfoById
//
// Use this API command to disable location additionalInfo override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationAdditionalInfoById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLocationAdditionalInfoById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsLocationBasedServiceById
//
// Use this API command to disable location based service override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationBasedServiceById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLocationBasedServiceById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsLocationById
//
// Use this API command to disable location override for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLocationById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsLteBandLockChannelsById
//
// Use this API command to disable LTE band lock channel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLteBandLockChannelsById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLteBandLockChannelsById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsMembersByApMac
//
// Use this API command to remove a member AP from an AP group.
//
// Required Parameters:
// - apMac string
//		- required
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsMembersByApMac(ctx context.Context, apMac string, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsMembersByApMac, true)
	req.SetPathParameter("apMac", apMac)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsProtectionMode24ById
//
// Use this API command to disable 2.4GHz radio protection mode configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsProtectionMode24ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsProtectionMode24ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsRecoverySsidById
//
// Use this API command to disable Recovery SSID configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRecoverySsidById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRecoverySsidById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsRksGreForwardBroadcastById
//
// Use this API command to disable Ruckus GRE Broadcast packet forwarding override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRksGreForwardBroadcastById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRksGreForwardBroadcastById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsRogueApAggressivenessModeById
//
// Use this API command to disable rogue AP aggressiveness mode override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApAggressivenessModeById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRogueApAggressivenessModeById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsRogueApJammingThresholdById
//
// Use this API command to disable rogue AP jamming threshold override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApJammingThresholdById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRogueApJammingThresholdById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsRogueApReportThresholdById
//
// Use this API command to disable rogue AP report threshold override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApReportThresholdById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRogueApReportThresholdById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsVenueProfileById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsVenueProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi24ById
//
// Use this API command to disable 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi24ChannelById
//
// Use this API command to disable 2.4GHz radio channel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ChannelById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi24ChannelRangeById
//
// Use this API command to disable 2.4GHz radio channelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelRangeById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ChannelRangeById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi24ChannelWidthById
//
// Use this API command to disable 2.4GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelWidthById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ChannelWidthById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi24TxPowerById
//
// Use this API command to disable 2.4GHz radio txPower override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24TxPowerById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24TxPowerById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50ById
//
// Use this API command to disable 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50ChannelWidthById
//
// Use this API command to disable 5GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ChannelWidthById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50ChannelWidthById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50IndoorChannelById
//
// Use this API command to disable 5GHz radio indoorChannel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50IndoorChannelById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50IndoorChannelRangeById
//
// Use this API command to disable 5GHz radio indoorChannelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelRangeById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50IndoorChannelRangeById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50OutdoorChannelById
//
// Use this API command to disable 5GHz radio outdoorChannel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50OutdoorChannelById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById
//
// Use this API command to disable 5GHz radio outdoorChannelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50OutdoorChannelRangeById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWifi50TxPowerById
//
// Use this API command to disable 5GHz radio txPower override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50TxPowerById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50TxPowerById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWlanGroup24ById
//
// Use this API command to disable WLAN group on 2.4GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup24ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWlanGroup24ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesApgroupsWlanGroup50ById
//
// Use this API command to disable WLAN group on 5GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup50ById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWlanGroup50ById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesApgroupsApmodelByModel
//
// Use this API command to retrieve AP model specific configuration override zone that belong to an AP group, NULL means the override setting is not checked inside an AP group.
//
// Required Parameters:
// - id string
//		- required
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsApmodelByModel(ctx context.Context, id string, model string, zoneId string) (*WSGZoneAPModelApModel, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGZoneAPModelApModel
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, model, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApgroupsApmodelByModel, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("model", model)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGZoneAPModelApModel()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesApgroupsById
//
// Use this API command to retrieve information about an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsById(ctx context.Context, id string, zoneId string) (*WSGAPGroupConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPGroupConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApgroupsById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPGroupConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesApgroupsByZoneId
//
// Use this API command to retrieve the list of AP groups that belong to a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGAPGroupService) FindRkszonesApgroupsByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGAPGroupList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPGroupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApgroupsByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPGroupList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesApgroupsDefaultByZoneId
//
// Use this API command to retrieve information about default AP group of zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsDefaultByZoneId(ctx context.Context, zoneId string) (*WSGAPGroupConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPGroupConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApgroupsDefaultByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPGroupConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesApgroupsById
//
// Use this API command to modify the configuration of an AP group.
//
// Request Body:
//	 - body *WSGAPGroupModifyAPGroup
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) PartialUpdateRkszonesApgroupsById(ctx context.Context, body *WSGAPGroupModifyAPGroup, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesApgroupsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesApgroupsApmodelByModel
//
// Use this API command to modify AP model specific configuration override zone that belong to an AP group.
//
// Request Body:
//	 - body *WSGZoneAPModelApModel
//
// Required Parameters:
// - id string
//		- required
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) UpdateRkszonesApgroupsApmodelByModel(ctx context.Context, body *WSGZoneAPModelApModel, id string, model string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, model, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesApgroupsApmodelByModel, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("model", model)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesApgroupsById
//
// Use this API command to modify the entire information of an AP group.
//
// Request Body:
//	 - body *WSGAPGroupModifyAPGroup
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) UpdateRkszonesApgroupsById(ctx context.Context, body *WSGAPGroupModifyAPGroup, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesApgroupsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
