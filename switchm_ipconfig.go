package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// SwitchMIPConfigCreate
//
// Definition: ipConfig_create
type SwitchMIPConfigCreate struct {
	AltoId *string `json:"altoId,omitempty"`

	// DhcpRelayAgent
	// DHCP Relay IP Address
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// Port
	// Switch Port
	Port *string `json:"port,omitempty"`

	// PortName
	// Switch Port Name
	PortName *string `json:"portName,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMIPConfigCreate() *SwitchMIPConfigCreate {
	m := new(SwitchMIPConfigCreate)
	return m
}

// SwitchMIPConfigCreateResult
//
// Definition: ipConfig_createResult
type SwitchMIPConfigCreateResult struct {
	// Id
	// The ID of Setting
	Id *string `json:"id,omitempty"`
}

type SwitchMIPConfigCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMIPConfigCreateResult
}

func newSwitchMIPConfigCreateResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMIPConfigCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMIPConfigCreateResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMIPConfigCreateResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMIPConfigCreateResult() *SwitchMIPConfigCreateResult {
	m := new(SwitchMIPConfigCreateResult)
	return m
}

// SwitchMIPConfig
//
// Definition: ipConfig_ipConfig
type SwitchMIPConfig struct {
	AltoId *string `json:"altoId,omitempty"`

	// CreatedTime
	// Config Created Time
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpRelayAgent
	// DHCP Replay IP Address
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group ID
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Config ID
	Id *string `json:"id,omitempty"`

	// InAclConfigName
	// Ingress ACL Config Name
	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigName
	// Egress ACL Config Name
	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// Port
	// Switch Port
	Port *string `json:"port,omitempty"`

	// PortName
	// Switch Port Name
	PortName *string `json:"portName,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch ID
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch Name
	SwitchName *string `json:"switchName,omitempty"`

	// SwitchStatus
	// Switch Status
	SwitchStatus *string `json:"switchStatus,omitempty"`

	// UpdatedTime
	// Config Updated Time
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type SwitchMIPConfigAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMIPConfig
}

func newSwitchMIPConfigAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMIPConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMIPConfigAPIResponse) Hydrate() error {
	r.Data = new(SwitchMIPConfig)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMIPConfig() *SwitchMIPConfig {
	m := new(SwitchMIPConfig)
	return m
}

// SwitchMIPConfigList
//
// Definition: ipConfig_list
type SwitchMIPConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMIPConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMIPConfigListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMIPConfigList
}

func newSwitchMIPConfigListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMIPConfigListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMIPConfigListAPIResponse) Hydrate() error {
	r.Data = new(SwitchMIPConfigList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMIPConfigList() *SwitchMIPConfigList {
	m := new(SwitchMIPConfigList)
	return m
}

// SwitchMIPConfigModify
//
// Definition: ipConfig_modify
type SwitchMIPConfigModify struct {
	AltoId *string `json:"altoId,omitempty"`

	// DhcpRelayAgent
	// DHCP Relay IP Address
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// Port
	// Switch Port
	Port *string `json:"port,omitempty"`

	// PortName
	// Switch Port Name
	PortName *string `json:"portName,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewSwitchMIPConfigModify() *SwitchMIPConfigModify {
	m := new(SwitchMIPConfigModify)
	return m
}
