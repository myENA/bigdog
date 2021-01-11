package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type SwitchMPortSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMPortSettingsService(c *VSZClient) *SwitchMPortSettingsService {
	s := new(SwitchMPortSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortSettingsService() *SwitchMPortSettingsService {
	return NewSwitchMPortSettingsService(ss.apiClient)
}

// SwitchMPortSettingsCreateBulk
//
// Definition: portSettings_createBulk
type SwitchMPortSettingsCreateBulk struct {
	AltoId *string `json:"altoId,omitempty"`

	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// GroupId
	// Group Id
	GroupId *string `json:"groupId,omitempty"`

	IdList SwitchMCommonIdList `json:"idList,omitempty"`

	// IgnoreList
	// attributes not to overwrite
	IgnoreList []string `json:"ignoreList,omitempty"`

	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	LldpQosList []*SwitchMPortSettingsLldpQos `json:"lldpQosList,omitempty"`

	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	PoeBudget *int `json:"poeBudget,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - oneof:[0,1,2,3,4,5,6,7,8]
	PoeClass *string `json:"poeClass,omitempty"`

	// PoeEnabled
	// PoE Enabled
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoePriority
	// PoE Priority
	PoePriority *int `json:"poePriority,omitempty"`

	// Port
	// Port
	Port *string `json:"port,omitempty"`

	// PortEnabled
	// Port Enabled
	PortEnabled *bool `json:"portEnabled,omitempty"`

	// PortName
	// Port Name
	PortName *string `json:"portName,omitempty"`

	PortProtected *bool `json:"portProtected,omitempty"`

	// PortSpeed
	// Port Speed
	// Constraints:
	//    - oneof:[NONE,1G,10G,AUTO,10-HALF,10-FULL,100-HALF,100-FULL,1000-FULL,1000-FULL-MASTER,1000-FULL-SLAVE,2500-FULL,2500-FULL-MASTER,2500-FULL-SLAVE,5G-FULL,5G-FULL-MASTER,5G-FULL-SLAVE,10G-FULL,10G-FULL-MASTER,10G-FULL-SLAVE,25G-FULL,40G-FULL,100G-FULL]
	PortSpeed *string `json:"portSpeed,omitempty"`

	// RstpAdminEdgePortEnabled
	// RSTP Admin Edge Port Enabled
	RstpAdminEdgePortEnabled *bool `json:"rstpAdminEdgePortEnabled,omitempty"`

	// StpBpduGuardEnabled
	// STP BPDU GUARD Enabled
	StpBpduGuardEnabled *bool `json:"stpBpduGuardEnabled,omitempty"`

	// StpRootGuardEnabled
	// STP Root Guard  Enabled
	StpRootGuardEnabled *bool `json:"stpRootGuardEnabled,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// TaggedVlans
	// Tagged Vlans
	TaggedVlans *string `json:"taggedVlans,omitempty"`

	// UntaggedVlans
	// Untagged vlans
	UntaggedVlans *string `json:"untaggedVlans,omitempty"`

	VoiceVlanId *int `json:"voiceVlanId,omitempty"`
}

func NewSwitchMPortSettingsCreateBulk() *SwitchMPortSettingsCreateBulk {
	m := new(SwitchMPortSettingsCreateBulk)
	return m
}

// SwitchMPortSettingsLldpQos
//
// Definition: portSettings_LldpQos
type SwitchMPortSettingsLldpQos struct {
	// ApplicationType
	// Constraints:
	//    - oneof:[GUEST_VOICE,GUEST_VOICE_SIGNALING,SOFTPHONE_VOICE,STREAMING_VIDEO,VIDEO_CONFERENCING,VIDEO_SIGNALING,VOICE,VOICE_SIGNALING]
	ApplicationType *string `json:"applicationType,omitempty"`

	Dscp *int `json:"dscp,omitempty"`

	Priority *int `json:"priority,omitempty"`

	// QosVlanType
	// Constraints:
	//    - oneof:[PRIORITY_TAGGED,TAGGED,UNTAGGED]
	QosVlanType *string `json:"qosVlanType,omitempty"`

	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMPortSettingsLldpQos() *SwitchMPortSettingsLldpQos {
	m := new(SwitchMPortSettingsLldpQos)
	return m
}

// SwitchMPortSettings
//
// Definition: portSettings_PortSettings
type SwitchMPortSettings struct {
	AltoId *string `json:"altoId,omitempty"`

	// CreatedTime
	// The create time of the Port Settings
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// GroupId
	// Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`

	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	LldpQosList []*SwitchMPortSettingsLldpQos `json:"lldpQosList,omitempty"`

	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	PoeBudget *int `json:"poeBudget,omitempty"`

	// PoeCapability
	// PoE Capability
	PoeCapability *bool `json:"poeCapability,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - oneof:[0,1,2,3,4,5,6,7,8]
	PoeClass *string `json:"poeClass,omitempty"`

	// PoeEnabled
	// PoE Enabled
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoePriority
	// PoE Priority
	PoePriority *int `json:"poePriority,omitempty"`

	// Port
	// Port
	Port *string `json:"port,omitempty"`

	// PortEnabled
	// Port Enabled
	PortEnabled *bool `json:"portEnabled,omitempty"`

	// PortName
	// Port Name
	PortName *string `json:"portName,omitempty"`

	PortProtected *bool `json:"portProtected,omitempty"`

	// PortSpeed
	// Port Speed
	// Constraints:
	//    - oneof:[NONE,1G,10G,AUTO,10-HALF,10-FULL,100-HALF,100-FULL,1000-FULL,1000-FULL-MASTER,1000-FULL-SLAVE,2500-FULL,2500-FULL-MASTER,2500-FULL-SLAVE,5G-FULL,5G-FULL-MASTER,5G-FULL-SLAVE,10G-FULL,10G-FULL-MASTER,10G-FULL-SLAVE,25G-FULL,40G-FULL,100G-FULL]
	PortSpeed *string `json:"portSpeed,omitempty"`

	// RstpAdminEdgePortEnabled
	// RSTP Admin Edge Port Enabled
	RstpAdminEdgePortEnabled *bool `json:"rstpAdminEdgePortEnabled,omitempty"`

	// StpBpduGuardEnabled
	// STP BPDU GUARD Enabled
	StpBpduGuardEnabled *bool `json:"stpBpduGuardEnabled,omitempty"`

	// StpRootGuardEnabled
	// STP Root Guard  Enabled
	StpRootGuardEnabled *bool `json:"stpRootGuardEnabled,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// TaggedVlans
	// Tagged Vlans
	TaggedVlans *string `json:"taggedVlans,omitempty"`

	// UntaggedVlans
	// Untagged vlans
	UntaggedVlans *string `json:"untaggedVlans,omitempty"`

	// UpdatedTime
	// The modify time of the Port Settings
	UpdatedTime *int `json:"updatedTime,omitempty"`

	VoiceVlanId *int `json:"voiceVlanId,omitempty"`
}

type SwitchMPortSettingsAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMPortSettings
}

func newSwitchMPortSettingsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMPortSettingsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMPortSettingsAPIResponse) Hydrate() error {
	r.Data = new(SwitchMPortSettings)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMPortSettings() *SwitchMPortSettings {
	m := new(SwitchMPortSettings)
	return m
}

// SwitchMPortSettingsQueryResult
//
// Definition: portSettings_PortSettingsQueryResult
type SwitchMPortSettingsQueryResult struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Port Settings returned out of the complete Port Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Port Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMPortSettings `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Port Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Port Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMPortSettingsQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMPortSettingsQueryResult
}

func newSwitchMPortSettingsQueryResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMPortSettingsQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMPortSettingsQueryResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMPortSettingsQueryResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMPortSettingsQueryResult() *SwitchMPortSettingsQueryResult {
	m := new(SwitchMPortSettingsQueryResult)
	return m
}

// SwitchMPortSettingsUpdatePortSettings
//
// Definition: portSettings_UpdatePortSettings
type SwitchMPortSettingsUpdatePortSettings struct {
	AltoId *string `json:"altoId,omitempty"`

	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	LldpQosList []*SwitchMPortSettingsLldpQos `json:"lldpQosList,omitempty"`

	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	PoeBudget *int `json:"poeBudget,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - oneof:[0,1,2,3,4,5,6,7,8]
	PoeClass *string `json:"poeClass,omitempty"`

	// PoeEnabled
	// PoE Enabled
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoePriority
	// PoE Priority
	PoePriority *int `json:"poePriority,omitempty"`

	// PortEnabled
	// Port Enabled
	PortEnabled *bool `json:"portEnabled,omitempty"`

	// PortName
	// Port Name
	PortName *string `json:"portName,omitempty"`

	PortProtected *bool `json:"portProtected,omitempty"`

	// PortSpeed
	// Port Speed
	// Constraints:
	//    - oneof:[NONE,1G,10G,AUTO,10-HALF,10-FULL,100-HALF,100-FULL,1000-FULL,1000-FULL-MASTER,1000-FULL-SLAVE,2500-FULL,2500-FULL-MASTER,2500-FULL-SLAVE,5G-FULL,5G-FULL-MASTER,5G-FULL-SLAVE,10G-FULL,10G-FULL-MASTER,10G-FULL-SLAVE,25G-FULL,40G-FULL,100G-FULL]
	PortSpeed *string `json:"portSpeed,omitempty"`

	// RstpAdminEdgePortEnabled
	// RSTP Admin Edge Port Enabled
	RstpAdminEdgePortEnabled *bool `json:"rstpAdminEdgePortEnabled,omitempty"`

	// StpBpduGuardEnabled
	// STP BPDU GUARD Enabled
	StpBpduGuardEnabled *bool `json:"stpBpduGuardEnabled,omitempty"`

	// StpRootGuardEnabled
	// STP Root Guard  Enabled
	StpRootGuardEnabled *bool `json:"stpRootGuardEnabled,omitempty"`

	// TaggedVlans
	// Tagged Vlans
	TaggedVlans *string `json:"taggedVlans,omitempty"`

	// UntaggedVlans
	// Untagged vlans
	UntaggedVlans *string `json:"untaggedVlans,omitempty"`

	VoiceVlanId *int `json:"voiceVlanId,omitempty"`
}

func NewSwitchMPortSettingsUpdatePortSettings() *SwitchMPortSettingsUpdatePortSettings {
	m := new(SwitchMPortSettingsUpdatePortSettings)
	return m
}

// AddPortSettingsBulk
//
// Use this API command to Bulk update the port setting
//
// Operation ID: addPortSettingsBulk
// Operation path: /portSettings/bulk
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMPortSettingsCreateBulk
func (s *SwitchMPortSettingsService) AddPortSettingsBulk(ctx context.Context, body *SwitchMPortSettingsCreateBulk, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddPortSettingsBulk, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindPortSettings
//
// Use this API command to Retrieve all Port Settings list.
//
// Operation ID: findPortSettings
// Operation path: /portSettings
// Success code: 200 (OK)
func (s *SwitchMPortSettingsService) FindPortSettings(ctx context.Context, mutators ...RequestMutator) (*SwitchMPortSettingsQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMPortSettingsQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMPortSettingsQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindPortSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMPortSettingsQueryResultAPIResponse), err
}

// FindPortSettingsById
//
// Use this API command to Retrieve Port Settings.
//
// Operation ID: findPortSettingsById
// Operation path: /portSettings/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMPortSettingsService) FindPortSettingsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMPortSettingsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMPortSettingsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMPortSettingsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindPortSettingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMPortSettingsAPIResponse), err
}

// FindPortSettingsByQueryCriteria
//
// Use this API command to Retrieve Port Settings list.
//
// Operation ID: findPortSettingsByQueryCriteria
// Operation path: /portSettings/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortSettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMPortSettingsQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMPortSettingsQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMPortSettingsQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindPortSettingsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMPortSettingsQueryResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMPortSettingsQueryResultAPIResponse), err
}

// UpdatePortSettingsById
//
// Use this API command to Update Port Settings.
//
// Operation ID: updatePortSettingsById
// Operation path: /portSettings/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMPortSettingsUpdatePortSettings
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMPortSettingsService) UpdatePortSettingsById(ctx context.Context, body *SwitchMPortSettingsUpdatePortSettings, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdatePortSettingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
