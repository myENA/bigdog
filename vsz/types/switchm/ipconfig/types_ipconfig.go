package ipconfig

// API Version: v8_1

import (
	"encoding/json"
)

type Create struct {
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

func NewCreate() *Create {
	createType := new(Create)
	return createType
}

func NewDefaultCreate() *Create {
	createType := new(Create)
	return createType
}

type CreateResult interface{}

func NewCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

func NewDefaultCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

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

func NewEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

func NewDefaultEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

type IpConfig struct {
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

func NewIpConfig() *IpConfig {
	ipConfigType := new(IpConfig)
	return ipConfigType
}

func NewDefaultIpConfig() *IpConfig {
	ipConfigType := new(IpConfig)
	return ipConfigType
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
	List []*IpConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewList() *List {
	listType := new(List)
	return listType
}

func NewDefaultList() *List {
	listType := new(List)
	return listType
}

type Modify struct {
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

func NewModify() *Modify {
	modifyType := new(Modify)
	return modifyType
}

func NewDefaultModify() *Modify {
	modifyType := new(Modify)
	return modifyType
}
