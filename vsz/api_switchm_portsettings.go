package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMPortsettingsService struct {
	apiClient *APIClient
}

func NewSwitchMPortsettingsService(c *APIClient) *SwitchMPortsettingsService {
	s := new(SwitchMPortsettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortsettingsService() *SwitchMPortsettingsService {
	return NewSwitchMPortsettingsService(ss.apiClient)
}

type SwitchMPortsettingsCreateBulk struct {
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

func NewSwitchMPortsettingsCreateBulk() *SwitchMPortsettingsCreateBulk {
	m := new(SwitchMPortsettingsCreateBulk)
	return m
}

type SwitchMPortsettings struct {
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

func NewSwitchMPortsettings() *SwitchMPortsettings {
	m := new(SwitchMPortsettings)
	return m
}

type SwitchMPortsettingsQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMPortsettingsQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Port Settings returned out of the complete Port Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Port Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMPortsettings `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Port Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Port Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMPortsettingsQueryResult() *SwitchMPortsettingsQueryResult {
	m := new(SwitchMPortsettingsQueryResult)
	return m
}

// SwitchMPortsettingsQueryResultExtraType
//
// Any additional response data
type SwitchMPortsettingsQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMPortsettingsQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMPortsettingsQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMPortsettingsQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMPortsettingsQueryResultExtraType() *SwitchMPortsettingsQueryResultExtraType {
	m := new(SwitchMPortsettingsQueryResultExtraType)
	return m
}

type SwitchMPortsettingsUpdatePortSettings struct {
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

func NewSwitchMPortsettingsUpdatePortSettings() *SwitchMPortsettingsUpdatePortSettings {
	m := new(SwitchMPortsettingsUpdatePortSettings)
	return m
}

// AddPortSettingsBulk
//
// Use this API command to Bulk update the port setting
//
// Request Body:
//	 - body *SwitchMPortsettingsCreateBulk
func (s *SwitchMPortsettingsService) AddPortSettingsBulk(ctx context.Context, body *SwitchMPortsettingsCreateBulk) (*APIResponseMeta, error) {
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
func (s *SwitchMPortsettingsService) FindPortSettings(ctx context.Context) (*SwitchMPortsettingsQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortsettingsQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortSettings, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortsettingsQueryResult()
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
func (s *SwitchMPortsettingsService) FindPortSettingsById(ctx context.Context, id string) (*SwitchMPortsettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortsettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortSettingsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortsettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPortSettingsByQueryCriteria
//
// Use this API command to Retrieve Port Settings list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMPortsettingsQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortsettingsQueryResult
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
	resp = NewSwitchMPortsettingsQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdatePortSettingsById
//
// Use this API command to Update Port Settings.
//
// Request Body:
//	 - body *SwitchMPortsettingsUpdatePortSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMPortsettingsService) UpdatePortSettingsById(ctx context.Context, body *SwitchMPortsettingsUpdatePortSettings, id string) (*APIResponseMeta, error) {
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
