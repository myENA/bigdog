package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGAPGroupService struct {
	apiClient *VSZClient
}

func NewWSGAPGroupService(c *VSZClient) *WSGAPGroupService {
	s := new(WSGAPGroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAPGroupService() *WSGAPGroupService {
	return NewWSGAPGroupService(ss.apiClient)
}

// WSGAPGroupAddMembers
//
// Definition: apgroup_addMembers
type WSGAPGroupAddMembers struct {
	// MemberList
	// List of apMac
	MemberList []*WSGAPGroupMember `json:"memberList,omitempty"`
}

func NewWSGAPGroupAddMembers() *WSGAPGroupAddMembers {
	m := new(WSGAPGroupAddMembers)
	return m
}

// WSGAPGroupConfiguration
//
// Definition: apgroup_apGroupConfiguration
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

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonOverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

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

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Wifi24 *WSGCommonRadio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonRadio50SuperSet `json:"wifi50,omitempty"`

	WlanGroup24 *WSGCommonGenericRef `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WSGCommonGenericRef `json:"wlanGroup50,omitempty"`

	// ZoneId
	// Identifier of the zone to which the AP group belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGAPGroupConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPGroupConfiguration
}

func newWSGAPGroupConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPGroupConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPGroupConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGAPGroupConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPGroupConfiguration() *WSGAPGroupConfiguration {
	m := new(WSGAPGroupConfiguration)
	return m
}

// WSGAPGroupList
//
// Definition: apgroup_apGroupList
type WSGAPGroupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPGroupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAPGroupListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPGroupList
}

func newWSGAPGroupListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPGroupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPGroupListAPIResponse) Hydrate() error {
	r.Data = new(WSGAPGroupList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPGroupList() *WSGAPGroupList {
	m := new(WSGAPGroupList)
	return m
}

// WSGAPGroupMember
//
// Definition: apgroup_apGroupMember
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

// WSGAPGroupSummary
//
// Definition: apgroup_apGroupSummary
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

// WSGAPGroupCreateAPGroup
//
// Definition: apgroup_createAPGroup
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

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonOverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

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

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

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

// WSGAPGroupModifyAPGroup
//
// Definition: apgroup_modifyAPGroup
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

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *WSGCommonOverrideGenericRef `json:"locationBasedService,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

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

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

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
// Operation ID: addRkszonesApgroupsByZoneId
//
// Use this API command to create new AP group within a zone.
//
// Request Body:
//	 - body *WSGAPGroupCreateAPGroup
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGAPGroupService) AddRkszonesApgroupsByZoneId(ctx context.Context, body *WSGAPGroupCreateAPGroup, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesApgroupsByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesApgroupsMembersByApMac
//
// Operation ID: addRkszonesApgroupsMembersByApMac
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
func (s *WSGAPGroupService) AddRkszonesApgroupsMembersByApMac(ctx context.Context, apMac string, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesApgroupsMembersByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddRkszonesApgroupsMembersById
//
// Operation ID: addRkszonesApgroupsMembersById
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
func (s *WSGAPGroupService) AddRkszonesApgroupsMembersById(ctx context.Context, body *WSGAPGroupAddMembers, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesApgroupsMembersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsAltitudeById
//
// Operation ID: deleteRkszonesApgroupsAltitudeById
//
// Use this API command to clear the altitude of AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAltitudeById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAltitudeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsApMgmtVlanById
//
// Operation ID: deleteRkszonesApgroupsApMgmtVlanById
//
// Disable AP Management Vlan Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsApMgmtVlanById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsApMgmtVlanById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsApmodelByModel
//
// Operation ID: deleteRkszonesApgroupsApmodelByModel
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
func (s *WSGAPGroupService) DeleteRkszonesApgroupsApmodelByModel(ctx context.Context, id string, model string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsApmodelByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("model", model)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsAutoChannelSelection24ById
//
// Operation ID: deleteRkszonesApgroupsAutoChannelSelection24ById
//
// Disable Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection24ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAutoChannelSelection24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsAutoChannelSelection50ById
//
// Operation ID: deleteRkszonesApgroupsAutoChannelSelection50ById
//
// Disable Radio 5G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAutoChannelSelection50ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAutoChannelSelection50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsAwsVenueById
//
// Operation ID: deleteRkszonesApgroupsAwsVenueById
//
// Use this API command to disable AWS venue override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsAwsVenueById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsAwsVenueById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsById
//
// Operation ID: deleteRkszonesApgroupsById
//
// Use this API command to delete an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsChannelEvaluationIntervalById
//
// Operation ID: deleteRkszonesApgroupsChannelEvaluationIntervalById
//
// Disable Channel Evaluation Interval Override of an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsChannelEvaluationIntervalById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsChannelEvaluationIntervalById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsClientAdmissionControl24ById
//
// Operation ID: deleteRkszonesApgroupsClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl24ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsClientAdmissionControl24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsClientAdmissionControl50ById
//
// Operation ID: deleteRkszonesApgroupsClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsClientAdmissionControl50ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsClientAdmissionControl50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById
//
// Operation ID: deleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById
//
// Use this API command to disable Directed Multicast from Network to wired/wireless client configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsDirectedMulticastFromNetworkEnabledById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById
//
// Operation ID: deleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById
//
// Use this API command to disable Directed Multicast from wired client to Network configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsDirectedMulticastFromWiredClientEnabledById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById
//
// Operation ID: deleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById
//
// Use this API command to disable Directed Multicast from wireless client to Network configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsDirectedMulticastFromWirelessClientEnabledById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsLocationAdditionalInfoById
//
// Operation ID: deleteRkszonesApgroupsLocationAdditionalInfoById
//
// Use this API command to disable location additionalInfo override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationAdditionalInfoById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLocationAdditionalInfoById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsLocationBasedServiceById
//
// Operation ID: deleteRkszonesApgroupsLocationBasedServiceById
//
// Use this API command to disable location based service override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationBasedServiceById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLocationBasedServiceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsLocationById
//
// Operation ID: deleteRkszonesApgroupsLocationById
//
// Use this API command to disable location override for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLocationById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLocationById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsLteBandLockChannelsById
//
// Operation ID: deleteRkszonesApgroupsLteBandLockChannelsById
//
// Use this API command to disable LTE band lock channel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsLteBandLockChannelsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsLteBandLockChannelsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsMembersByApMac
//
// Operation ID: deleteRkszonesApgroupsMembersByApMac
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
func (s *WSGAPGroupService) DeleteRkszonesApgroupsMembersByApMac(ctx context.Context, apMac string, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsMembersByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsProtectionMode24ById
//
// Operation ID: deleteRkszonesApgroupsProtectionMode24ById
//
// Use this API command to disable 2.4GHz radio protection mode configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsProtectionMode24ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsProtectionMode24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsRecoverySsidById
//
// Operation ID: deleteRkszonesApgroupsRecoverySsidById
//
// Use this API command to disable Recovery SSID configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRecoverySsidById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRecoverySsidById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsRksGreForwardBroadcastById
//
// Operation ID: deleteRkszonesApgroupsRksGreForwardBroadcastById
//
// Use this API command to disable Ruckus GRE Broadcast packet forwarding override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRksGreForwardBroadcastById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRksGreForwardBroadcastById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsRogueApAggressivenessModeById
//
// Operation ID: deleteRkszonesApgroupsRogueApAggressivenessModeById
//
// Use this API command to disable rogue AP aggressiveness mode override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApAggressivenessModeById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRogueApAggressivenessModeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsRogueApJammingThresholdById
//
// Operation ID: deleteRkszonesApgroupsRogueApJammingThresholdById
//
// Use this API command to disable rogue AP jamming threshold override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApJammingThresholdById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRogueApJammingThresholdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsRogueApReportThresholdById
//
// Operation ID: deleteRkszonesApgroupsRogueApReportThresholdById
//
// Use this API command to disable rogue AP report threshold override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsRogueApReportThresholdById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsRogueApReportThresholdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsVenueProfileById
//
// Operation ID: deleteRkszonesApgroupsVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsVenueProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsVenueProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi24ById
//
// Operation ID: deleteRkszonesApgroupsWifi24ById
//
// Use this API command to disable 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi24ChannelById
//
// Operation ID: deleteRkszonesApgroupsWifi24ChannelById
//
// Use this API command to disable 2.4GHz radio channel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ChannelById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi24ChannelRangeById
//
// Operation ID: deleteRkszonesApgroupsWifi24ChannelRangeById
//
// Use this API command to disable 2.4GHz radio channelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelRangeById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ChannelRangeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi24ChannelWidthById
//
// Operation ID: deleteRkszonesApgroupsWifi24ChannelWidthById
//
// Use this API command to disable 2.4GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24ChannelWidthById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24ChannelWidthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi24TxPowerById
//
// Operation ID: deleteRkszonesApgroupsWifi24TxPowerById
//
// Use this API command to disable 2.4GHz radio txPower override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi24TxPowerById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi24TxPowerById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50ById
//
// Operation ID: deleteRkszonesApgroupsWifi50ById
//
// Use this API command to disable 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50ChannelWidthById
//
// Operation ID: deleteRkszonesApgroupsWifi50ChannelWidthById
//
// Use this API command to disable 5GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50ChannelWidthById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50ChannelWidthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50IndoorChannelById
//
// Operation ID: deleteRkszonesApgroupsWifi50IndoorChannelById
//
// Use this API command to disable 5GHz radio indoorChannel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50IndoorChannelById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50IndoorChannelRangeById
//
// Operation ID: deleteRkszonesApgroupsWifi50IndoorChannelRangeById
//
// Use this API command to disable 5GHz radio indoorChannelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50IndoorChannelRangeById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50IndoorChannelRangeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50OutdoorChannelById
//
// Operation ID: deleteRkszonesApgroupsWifi50OutdoorChannelById
//
// Use this API command to disable 5GHz radio outdoorChannel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50OutdoorChannelById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById
//
// Operation ID: deleteRkszonesApgroupsWifi50OutdoorChannelRangeById
//
// Use this API command to disable 5GHz radio outdoorChannelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50OutdoorChannelRangeById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50OutdoorChannelRangeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWifi50TxPowerById
//
// Operation ID: deleteRkszonesApgroupsWifi50TxPowerById
//
// Use this API command to disable 5GHz radio txPower override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWifi50TxPowerById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWifi50TxPowerById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWlanGroup24ById
//
// Operation ID: deleteRkszonesApgroupsWlanGroup24ById
//
// Use this API command to disable WLAN group on 2.4GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup24ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWlanGroup24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesApgroupsWlanGroup50ById
//
// Operation ID: deleteRkszonesApgroupsWlanGroup50ById
//
// Use this API command to disable WLAN group on 5GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) DeleteRkszonesApgroupsWlanGroup50ById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesApgroupsWlanGroup50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindRkszonesApgroupsApmodelByModel
//
// Operation ID: findRkszonesApgroupsApmodelByModel
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
func (s *WSGAPGroupService) FindRkszonesApgroupsApmodelByModel(ctx context.Context, id string, model string, zoneId string, mutators ...RequestMutator) (*WSGZoneAPModelApModelAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneAPModelApModelAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGZoneAPModelApModelAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesApgroupsApmodelByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("model", model)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneAPModelApModelAPIResponse), err
}

// FindRkszonesApgroupsById
//
// Operation ID: findRkszonesApgroupsById
//
// Use this API command to retrieve information about an AP group.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAPGroupConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPGroupConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPGroupConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesApgroupsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPGroupConfigurationAPIResponse), err
}

// FindRkszonesApgroupsByZoneId
//
// Operation ID: findRkszonesApgroupsByZoneId
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
func (s *WSGAPGroupService) FindRkszonesApgroupsByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAPGroupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPGroupListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPGroupListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesApgroupsByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPGroupListAPIResponse), err
}

// FindRkszonesApgroupsDefaultByZoneId
//
// Operation ID: findRkszonesApgroupsDefaultByZoneId
//
// Use this API command to retrieve information about default AP group of zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGAPGroupService) FindRkszonesApgroupsDefaultByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGAPGroupConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPGroupConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPGroupConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesApgroupsDefaultByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPGroupConfigurationAPIResponse), err
}

// PartialUpdateRkszonesApgroupsById
//
// Operation ID: partialUpdateRkszonesApgroupsById
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
func (s *WSGAPGroupService) PartialUpdateRkszonesApgroupsById(ctx context.Context, body *WSGAPGroupModifyAPGroup, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesApgroupsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateRkszonesApgroupsApmodelByModel
//
// Operation ID: updateRkszonesApgroupsApmodelByModel
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
func (s *WSGAPGroupService) UpdateRkszonesApgroupsApmodelByModel(ctx context.Context, body *WSGZoneAPModelApModel, id string, model string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesApgroupsApmodelByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("model", model)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateRkszonesApgroupsById
//
// Operation ID: updateRkszonesApgroupsById
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
func (s *WSGAPGroupService) UpdateRkszonesApgroupsById(ctx context.Context, body *WSGAPGroupModifyAPGroup, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesApgroupsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
