package ruckus

// API Version: v9_0

type SwitchMIPConfigCreate struct {
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

type SwitchMIPConfigCreateResult struct {
	// Id
	// The ID of Setting
	Id *string `json:"id,omitempty"`
}

func NewSwitchMIPConfigCreateResult() *SwitchMIPConfigCreateResult {
	m := new(SwitchMIPConfigCreateResult)
	return m
}

type SwitchMIPConfig struct {
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

func NewSwitchMIPConfig() *SwitchMIPConfig {
	m := new(SwitchMIPConfig)
	return m
}

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

func NewSwitchMIPConfigList() *SwitchMIPConfigList {
	m := new(SwitchMIPConfigList)
	return m
}

type SwitchMIPConfigModify struct {
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
