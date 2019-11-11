package vsz

// API Version: v8_1

import (
	"encoding/json"
)

type SwitchMVeConfigCreate struct {
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

type SwitchMVeConfigCreateResult interface{}

type SwitchMVeConfigEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMVeConfigEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMVeConfigEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMVeConfigEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMVeConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMVeConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMVeConfigModify struct {
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

type SwitchMVeConfig struct {
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
