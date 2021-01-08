package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// SwitchMVEConfigCreate
//
// Definition: veConfig_create
type SwitchMVEConfigCreate struct {
	AltoId *string `json:"altoId,omitempty"`

	// DhcpRelayAgent
	// DHCP Replay IP Address
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group ID
	GroupId *string `json:"groupId,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// SubnetMask
	// Subnet Mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch ID
	SwitchId *string `json:"switchId,omitempty"`

	// VeId
	// VE Id
	VeId *int `json:"veId,omitempty"`

	// VlanId
	// VLAN ID
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVEConfigCreate() *SwitchMVEConfigCreate {
	m := new(SwitchMVEConfigCreate)
	return m
}

// SwitchMVEConfigCreateResult
//
// Definition: veConfig_createResult
type SwitchMVEConfigCreateResult struct {
	// Id
	// The ID of Setting
	Id *string `json:"id,omitempty"`
}

type SwitchMVEConfigCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMVEConfigCreateResult
}

func newSwitchMVEConfigCreateResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMVEConfigCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMVEConfigCreateResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMVEConfigCreateResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMVEConfigCreateResult() *SwitchMVEConfigCreateResult {
	m := new(SwitchMVEConfigCreateResult)
	return m
}

// SwitchMVEConfigList
//
// Definition: veConfig_list
type SwitchMVEConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMVEConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMVEConfigListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMVEConfigList
}

func newSwitchMVEConfigListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMVEConfigListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMVEConfigListAPIResponse) Hydrate() error {
	r.Data = new(SwitchMVEConfigList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMVEConfigList() *SwitchMVEConfigList {
	m := new(SwitchMVEConfigList)
	return m
}

// SwitchMVEConfigModify
//
// Definition: veConfig_modify
type SwitchMVEConfigModify struct {
	AltoId *string `json:"altoId,omitempty"`

	// DhcpRelayAgent
	// DHCP Replay IP Address
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// SubnetMask
	// Subnet Mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch ID
	SwitchId *string `json:"switchId,omitempty"`

	// VlanId
	// VLAN ID
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVEConfigModify() *SwitchMVEConfigModify {
	m := new(SwitchMVEConfigModify)
	return m
}

// SwitchMVEConfig
//
// Definition: veConfig_veConfig
type SwitchMVEConfig struct {
	AltoId *string `json:"altoId,omitempty"`

	// CreatedTime
	// Created Time
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpRelayAgent
	// DHCP Replay IP Address
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group ID
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// ID
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

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigName
	// Egress ACL Config Name
	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// SubnetMask
	// Subnet Mask
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
	// Updated Time
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// VeId
	// VE Id
	VeId *int `json:"veId,omitempty"`

	// VlanId
	// VLAN ID
	VlanId *int `json:"vlanId,omitempty"`
}

type SwitchMVEConfigAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMVEConfig
}

func newSwitchMVEConfigAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMVEConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMVEConfigAPIResponse) Hydrate() error {
	r.Data = new(SwitchMVEConfig)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMVEConfig() *SwitchMVEConfig {
	m := new(SwitchMVEConfig)
	return m
}
