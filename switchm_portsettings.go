package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchPortSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchPortSettingsService(c *VSZClient) *SwitchMSwitchPortSettingsService {
	s := new(SwitchMSwitchPortSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortSettingsService() *SwitchMSwitchPortSettingsService {
	return NewSwitchMSwitchPortSettingsService(ss.apiClient)
}

type SwitchMSwitchPortSettingsCreateBulk struct {
	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// GroupId
	// Group Id
	GroupId *string `json:"groupId,omitempty"`

	IdList *SwitchMSwitchPortSettingsCreateBulk `json:"idList,omitempty"`

	// IgnoreList
	// attributes not to overwrite
	IgnoreList []string `json:"ignoreList,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - oneof:[0,1,2,3,4]
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
}

func NewSwitchMSwitchPortSettingsCreateBulk() *SwitchMSwitchPortSettingsCreateBulk {
	m := new(SwitchMSwitchPortSettingsCreateBulk)
	return m
}

type SwitchMSwitchPortSettingsPortSettings struct {
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

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// PoeCapability
	// PoE Capability
	PoeCapability *bool `json:"poeCapability,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - oneof:[0,1,2,3,4]
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
}

func NewSwitchMSwitchPortSettingsPortSettings() *SwitchMSwitchPortSettingsPortSettings {
	m := new(SwitchMSwitchPortSettingsPortSettings)
	return m
}

type SwitchMSwitchPortSettingsPortSettingsQueryResult struct {
	// Extra
	// Any additional response data
	Extra **SwitchMSwitchPortSettingsPortSettingsQueryResult `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Port Settings returned out of the complete Port Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Port Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchPortSettingsPortSettingsQueryResult `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Port Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Port Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchPortSettingsPortSettingsQueryResult() *SwitchMSwitchPortSettingsPortSettingsQueryResult {
	m := new(SwitchMSwitchPortSettingsPortSettingsQueryResult)
	return m
}

// SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType
//
// Any additional response data
type SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchPortSettingsPortSettingsQueryResultExtraType() *SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType {
	m := new(SwitchMSwitchPortSettingsPortSettingsQueryResultExtraType)
	return m
}

type SwitchMSwitchPortSettingsUpdatePortSettings struct {
	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - oneof:[0,1,2,3,4]
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
}

func NewSwitchMSwitchPortSettingsUpdatePortSettings() *SwitchMSwitchPortSettingsUpdatePortSettings {
	m := new(SwitchMSwitchPortSettingsUpdatePortSettings)
	return m
}

// AddPortSettingsBulk
//
// Use this API command to Bulk update the port setting
//
// Request Body:
//	 - body *SwitchMSwitchPortSettingsCreateBulk
func (s *SwitchMSwitchPortSettingsService) AddPortSettingsBulk(ctx context.Context, body *SwitchMSwitchPortSettingsCreateBulk) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddPortSettingsBulk, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindPortSettings
//
// Use this API command to Retrieve all Port Settings list.
func (s *SwitchMSwitchPortSettingsService) FindPortSettings(ctx context.Context) (*SwitchMSwitchPortSettingsPortSettingsQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchPortSettingsPortSettingsQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortSettings, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchPortSettingsPortSettingsQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPortSettingsById
//
// Use this API command to Retrieve Port Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchPortSettingsService) FindPortSettingsById(ctx context.Context, id string) (*SwitchMSwitchPortSettingsPortSettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchPortSettingsPortSettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortSettingsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchPortSettingsPortSettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPortSettingsByQueryCriteria
//
// Use this API command to Retrieve Port Settings list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchPortSettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchPortSettingsPortSettingsQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchPortSettingsPortSettingsQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindPortSettingsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchPortSettingsPortSettingsQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdatePortSettingsById
//
// Use this API command to Update Port Settings.
//
// Request Body:
//	 - body *SwitchMSwitchPortSettingsUpdatePortSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchPortSettingsService) UpdatePortSettingsById(ctx context.Context, body *SwitchMSwitchPortSettingsUpdatePortSettings, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdatePortSettingsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
