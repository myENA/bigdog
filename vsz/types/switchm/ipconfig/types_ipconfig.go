package ipconfig

// API Version: v8_1

import (
	"encoding/json"
)

type Create struct {
	// DHCPRelayAgent
	// DHCP Relay IP Address
	DHCPRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupID
	// Switch Group Id
	GroupID *string `json:"groupId,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IPAddress
	// IP Address
	IPAddress *string `json:"ipAddress,omitempty"`

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

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`
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

type IPConfig struct {
	// CreatedTime
	// Config Created Time
	CreatedTime *int `json:"createdTime,omitempty"`

	// DHCPRelayAgent
	// DHCP Replay IP Address
	DHCPRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupID
	// Switch Group ID
	GroupID *string `json:"groupId,omitempty"`

	// ID
	// Config ID
	ID *string `json:"id,omitempty"`

	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IPAddress
	// IP Address
	IPAddress *string `json:"ipAddress,omitempty"`

	// OspfArea
	// OSPF IP Address
	OspfArea *string `json:"ospfArea,omitempty"`

	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL ConfigUUID
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
	// Config Updated Time
	UpdatedTime *int `json:"updatedTime,omitempty"`
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
	List []*IPConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type Modify struct {
	// DHCPRelayAgent
	// DHCP Relay IP Address
	DHCPRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IPAddress
	// IP Address
	IPAddress *string `json:"ipAddress,omitempty"`

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
