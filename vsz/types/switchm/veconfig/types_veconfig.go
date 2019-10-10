package veconfig

// API Version: v8_1

import (
	"encoding/json"
)

type Create struct {
	// DHCPRelayAgent
	// DHCP Replay IP Address
	DHCPRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupID
	// Switch Group ID
	GroupID *string `json:"groupId,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IPAddress
	// IP Address
	IPAddress *string `json:"ipAddress,omitempty"`

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

	// SwitchID
	// Switch ID
	SwitchID *string `json:"switchId,omitempty"`

	// VeID
	// VE Id
	VeID *int `json:"veId,omitempty"`

	// VlanID
	// VLAN ID
	VlanID *int `json:"vlanId,omitempty"`
}

type CreateResult interface{}

type EmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *EmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = EmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *EmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type List struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*VeConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type Modify struct {
	// DHCPRelayAgent
	// DHCP Replay IP Address
	DHCPRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IPAddress
	// IP Address
	IPAddress *string `json:"ipAddress,omitempty"`

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

	// SwitchID
	// Switch ID
	SwitchID *string `json:"switchId,omitempty"`

	// VlanID
	// VLAN ID
	VlanID *int `json:"vlanId,omitempty"`
}

type VeConfig struct {
	// CreatedTime
	// Created Time
	CreatedTime *int `json:"createdTime,omitempty"`

	// DHCPRelayAgent
	// DHCP Replay IP Address
	DHCPRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupID
	// Switch Group ID
	GroupID *string `json:"groupId,omitempty"`

	// ID
	// ID
	ID *string `json:"id,omitempty"`

	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IPAddress
	// IP Address
	IPAddress *string `json:"ipAddress,omitempty"`

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

	// SwitchID
	// Switch ID
	SwitchID *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch Name
	SwitchName *string `json:"switchName,omitempty"`

	// SwitchStatus
	// Switch Status
	SwitchStatus *string `json:"switchStatus,omitempty"`

	// UpdatedTime
	// Updated Time
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// VeID
	// VE Id
	VeID *int `json:"veId,omitempty"`

	// VlanID
	// VLAN ID
	VlanID *int `json:"vlanId,omitempty"`
}
