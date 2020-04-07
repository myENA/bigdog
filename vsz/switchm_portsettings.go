package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMPortSettingsService struct {
	apiClient *APIClient
}

func NewSwitchMPortSettingsService(c *APIClient) *SwitchMPortSettingsService {
	s := new(SwitchMPortSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortSettingsService() *SwitchMPortSettingsService {
	return NewSwitchMPortSettingsService(ss.apiClient)
}

type SwitchMPortSettingsCreateBulk struct {
	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	// Constraints:
	//    - nullable
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// GroupId
	// Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// IdList
	// Constraints:
	//    - nullable
	IdList SwitchMCommonIdList `json:"idList,omitempty"`

	// IgnoreList
	// attributes not to overwrite
	// Constraints:
	//    - nullable
	IgnoreList []string `json:"ignoreList,omitempty" validate:"omitempty,dive"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	// Constraints:
	//    - nullable
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	// Constraints:
	//    - nullable
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	// Constraints:
	//    - nullable
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	// Constraints:
	//    - nullable
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - nullable
	//    - oneof:[0,1,2,3,4]
	PoeClass *string `json:"poeClass,omitempty" validate:"omitempty,oneof=0 1 2 3 4"`

	// PoeEnabled
	// PoE Enabled
	// Constraints:
	//    - nullable
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoePriority
	// PoE Priority
	// Constraints:
	//    - nullable
	PoePriority *int `json:"poePriority,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`

	// PortEnabled
	// Port Enabled
	// Constraints:
	//    - nullable
	PortEnabled *bool `json:"portEnabled,omitempty"`

	// PortSpeed
	// Port Speed
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,1G,10G,AUTO,10-HALF,10-FULL,100-HALF,100-FULL,1000-FULL,1000-FULL-MASTER,1000-FULL-SLAVE,2500-FULL,2500-FULL-MASTER,2500-FULL-SLAVE,5G-FULL,5G-FULL-MASTER,5G-FULL-SLAVE,10G-FULL,10G-FULL-MASTER,10G-FULL-SLAVE,25G-FULL,40G-FULL,100G-FULL]
	PortSpeed *string `json:"portSpeed,omitempty" validate:"omitempty,oneof=NONE 1G 10G AUTO 10-HALF 10-FULL 100-HALF 100-FULL 1000-FULL 1000-FULL-MASTER 1000-FULL-SLAVE 2500-FULL 2500-FULL-MASTER 2500-FULL-SLAVE 5G-FULL 5G-FULL-MASTER 5G-FULL-SLAVE 10G-FULL 10G-FULL-MASTER 10G-FULL-SLAVE 25G-FULL 40G-FULL 100G-FULL"`

	// RstpAdminEdgePortEnabled
	// RSTP Admin Edge Port Enabled
	// Constraints:
	//    - nullable
	RstpAdminEdgePortEnabled *bool `json:"rstpAdminEdgePortEnabled,omitempty"`

	// StpBpduGuardEnabled
	// STP BPDU GUARD Enabled
	// Constraints:
	//    - nullable
	StpBpduGuardEnabled *bool `json:"stpBpduGuardEnabled,omitempty"`

	// StpRootGuardEnabled
	// STP Root Guard  Enabled
	// Constraints:
	//    - nullable
	StpRootGuardEnabled *bool `json:"stpRootGuardEnabled,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// TaggedVlans
	// Tagged Vlans
	// Constraints:
	//    - nullable
	TaggedVlans *string `json:"taggedVlans,omitempty"`

	// UntaggedVlans
	// Untagged vlans
	// Constraints:
	//    - nullable
	UntaggedVlans *string `json:"untaggedVlans,omitempty"`
}

func NewSwitchMPortSettingsCreateBulk() *SwitchMPortSettingsCreateBulk {
	m := new(SwitchMPortSettingsCreateBulk)
	return m
}

type SwitchMPortSettings struct {
	// CreatedTime
	// The create time of the Port Settings
	// Constraints:
	//    - nullable
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	// Constraints:
	//    - nullable
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// GroupId
	// Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	// Constraints:
	//    - nullable
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	// Constraints:
	//    - nullable
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	// Constraints:
	//    - nullable
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	// Constraints:
	//    - nullable
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// PoeCapability
	// PoE Capability
	// Constraints:
	//    - nullable
	PoeCapability *bool `json:"poeCapability,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - nullable
	//    - oneof:[0,1,2,3,4]
	PoeClass *string `json:"poeClass,omitempty" validate:"omitempty,oneof=0 1 2 3 4"`

	// PoeEnabled
	// PoE Enabled
	// Constraints:
	//    - nullable
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoePriority
	// PoE Priority
	// Constraints:
	//    - nullable
	PoePriority *int `json:"poePriority,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`

	// PortEnabled
	// Port Enabled
	// Constraints:
	//    - nullable
	PortEnabled *bool `json:"portEnabled,omitempty"`

	// PortName
	// Port Name
	// Constraints:
	//    - nullable
	PortName *string `json:"portName,omitempty"`

	// PortSpeed
	// Port Speed
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,1G,10G,AUTO,10-HALF,10-FULL,100-HALF,100-FULL,1000-FULL,1000-FULL-MASTER,1000-FULL-SLAVE,2500-FULL,2500-FULL-MASTER,2500-FULL-SLAVE,5G-FULL,5G-FULL-MASTER,5G-FULL-SLAVE,10G-FULL,10G-FULL-MASTER,10G-FULL-SLAVE,25G-FULL,40G-FULL,100G-FULL]
	PortSpeed *string `json:"portSpeed,omitempty" validate:"omitempty,oneof=NONE 1G 10G AUTO 10-HALF 10-FULL 100-HALF 100-FULL 1000-FULL 1000-FULL-MASTER 1000-FULL-SLAVE 2500-FULL 2500-FULL-MASTER 2500-FULL-SLAVE 5G-FULL 5G-FULL-MASTER 5G-FULL-SLAVE 10G-FULL 10G-FULL-MASTER 10G-FULL-SLAVE 25G-FULL 40G-FULL 100G-FULL"`

	// RstpAdminEdgePortEnabled
	// RSTP Admin Edge Port Enabled
	// Constraints:
	//    - nullable
	RstpAdminEdgePortEnabled *bool `json:"rstpAdminEdgePortEnabled,omitempty"`

	// StpBpduGuardEnabled
	// STP BPDU GUARD Enabled
	// Constraints:
	//    - nullable
	StpBpduGuardEnabled *bool `json:"stpBpduGuardEnabled,omitempty"`

	// StpRootGuardEnabled
	// STP Root Guard  Enabled
	// Constraints:
	//    - nullable
	StpRootGuardEnabled *bool `json:"stpRootGuardEnabled,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// TaggedVlans
	// Tagged Vlans
	// Constraints:
	//    - nullable
	TaggedVlans *string `json:"taggedVlans,omitempty"`

	// UntaggedVlans
	// Untagged vlans
	// Constraints:
	//    - nullable
	UntaggedVlans *string `json:"untaggedVlans,omitempty"`

	// UpdatedTime
	// The modify time of the Port Settings
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMPortSettings() *SwitchMPortSettings {
	m := new(SwitchMPortSettings)
	return m
}

type SwitchMPortSettingsQueryResult struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMPortSettingsQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Port Settings returned out of the complete Port Settings list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Port Settings after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMPortSettings `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total Port Settings count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Port Settings count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMPortSettingsQueryResult() *SwitchMPortSettingsQueryResult {
	m := new(SwitchMPortSettingsQueryResult)
	return m
}

// SwitchMPortSettingsQueryResultExtraType
//
// Any additional response data
// Constraints:
//    - nullable
type SwitchMPortSettingsQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMPortSettingsQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMPortSettingsQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMPortSettingsQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMPortSettingsQueryResultExtraType() *SwitchMPortSettingsQueryResultExtraType {
	m := new(SwitchMPortSettingsQueryResultExtraType)
	return m
}

type SwitchMPortSettingsUpdatePortSettings struct {
	// DhcpSnoopingTrustPortEnabled
	// DHCP Snooping Trust Port Enabled
	// Constraints:
	//    - nullable
	DhcpSnoopingTrustPortEnabled *bool `json:"dhcpSnoopingTrustPortEnabled,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	// Constraints:
	//    - nullable
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpsgEnabled
	// IPSG Enabled
	// Constraints:
	//    - nullable
	IpsgEnabled *bool `json:"ipsgEnabled,omitempty"`

	// LldpEnabled
	// LLDP Enabled
	// Constraints:
	//    - nullable
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	// Constraints:
	//    - nullable
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// PoeClass
	// PoE Class
	// Constraints:
	//    - nullable
	//    - oneof:[0,1,2,3,4]
	PoeClass *string `json:"poeClass,omitempty" validate:"omitempty,oneof=0 1 2 3 4"`

	// PoeEnabled
	// PoE Enabled
	// Constraints:
	//    - nullable
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoePriority
	// PoE Priority
	// Constraints:
	//    - nullable
	PoePriority *int `json:"poePriority,omitempty"`

	// PortEnabled
	// Port Enabled
	// Constraints:
	//    - nullable
	PortEnabled *bool `json:"portEnabled,omitempty"`

	// PortName
	// Port Name
	// Constraints:
	//    - nullable
	PortName *string `json:"portName,omitempty"`

	// PortSpeed
	// Port Speed
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,1G,10G,AUTO,10-HALF,10-FULL,100-HALF,100-FULL,1000-FULL,1000-FULL-MASTER,1000-FULL-SLAVE,2500-FULL,2500-FULL-MASTER,2500-FULL-SLAVE,5G-FULL,5G-FULL-MASTER,5G-FULL-SLAVE,10G-FULL,10G-FULL-MASTER,10G-FULL-SLAVE,25G-FULL,40G-FULL,100G-FULL]
	PortSpeed *string `json:"portSpeed,omitempty" validate:"omitempty,oneof=NONE 1G 10G AUTO 10-HALF 10-FULL 100-HALF 100-FULL 1000-FULL 1000-FULL-MASTER 1000-FULL-SLAVE 2500-FULL 2500-FULL-MASTER 2500-FULL-SLAVE 5G-FULL 5G-FULL-MASTER 5G-FULL-SLAVE 10G-FULL 10G-FULL-MASTER 10G-FULL-SLAVE 25G-FULL 40G-FULL 100G-FULL"`

	// RstpAdminEdgePortEnabled
	// RSTP Admin Edge Port Enabled
	// Constraints:
	//    - nullable
	RstpAdminEdgePortEnabled *bool `json:"rstpAdminEdgePortEnabled,omitempty"`

	// StpBpduGuardEnabled
	// STP BPDU GUARD Enabled
	// Constraints:
	//    - nullable
	StpBpduGuardEnabled *bool `json:"stpBpduGuardEnabled,omitempty"`

	// StpRootGuardEnabled
	// STP Root Guard  Enabled
	// Constraints:
	//    - nullable
	StpRootGuardEnabled *bool `json:"stpRootGuardEnabled,omitempty"`

	// TaggedVlans
	// Tagged Vlans
	// Constraints:
	//    - nullable
	TaggedVlans *string `json:"taggedVlans,omitempty"`

	// UntaggedVlans
	// Untagged vlans
	// Constraints:
	//    - nullable
	UntaggedVlans *string `json:"untaggedVlans,omitempty"`
}

func NewSwitchMPortSettingsUpdatePortSettings() *SwitchMPortSettingsUpdatePortSettings {
	m := new(SwitchMPortSettingsUpdatePortSettings)
	return m
}

// AddPortSettingsBulk
//
// Use this API command to Bulk update the port setting
//
// Request Body:
//	 - body *SwitchMPortSettingsCreateBulk
func (s *SwitchMPortSettingsService) AddPortSettingsBulk(ctx context.Context, body *SwitchMPortSettingsCreateBulk) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddPortSettingsBulk, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindPortSettings
//
// Use this API command to Retrieve all Port Settings list.
func (s *SwitchMPortSettingsService) FindPortSettings(ctx context.Context) (*SwitchMPortSettingsQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortSettingsQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortSettings, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortSettingsQueryResult()
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
func (s *SwitchMPortSettingsService) FindPortSettingsById(ctx context.Context, id string) (*SwitchMPortSettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortSettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortSettingsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortSettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPortSettingsByQueryCriteria
//
// Use this API command to Retrieve Port Settings list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortSettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMPortSettingsQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortSettingsQueryResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindPortSettingsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortSettingsQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdatePortSettingsById
//
// Use this API command to Update Port Settings.
//
// Request Body:
//	 - body *SwitchMPortSettingsUpdatePortSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMPortSettingsService) UpdatePortSettingsById(ctx context.Context, body *SwitchMPortSettingsUpdatePortSettings, id string) (*APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdatePortSettingsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
