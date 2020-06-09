package bigdog

// API Version: v9_0

type SwitchMVEConfigCreate struct {
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

type SwitchMVEConfigCreateResult struct {
	// Id
	// The ID of Setting
	Id *string `json:"id,omitempty"`
}

func NewSwitchMVEConfigCreateResult() *SwitchMVEConfigCreateResult {
	m := new(SwitchMVEConfigCreateResult)
	return m
}

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

func NewSwitchMVEConfigList() *SwitchMVEConfigList {
	m := new(SwitchMVEConfigList)
	return m
}

type SwitchMVEConfigModify struct {
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

type SwitchMVEConfig struct {
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

func NewSwitchMVEConfig() *SwitchMVEConfig {
	m := new(SwitchMVEConfig)
	return m
}
