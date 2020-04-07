package vsz

// API Version: v9_0

type SwitchMVeConfigCreate struct {
	// DhcpRelayAgent
	// DHCP Replay IP Address
	// Constraints:
	//    - nullable
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group ID
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	// Constraints:
	//    - nullable
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	// name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// OspfArea
	// OSPF IP Address
	// Constraints:
	//    - nullable
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	// Constraints:
	//    - nullable
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// SubnetMask
	// Subnet Mask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch ID
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// VeId
	// VE Id
	// Constraints:
	//    - nullable
	VeId *int `json:"veId,omitempty"`

	// VlanId
	// VLAN ID
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVeConfigCreate() *SwitchMVeConfigCreate {
	m := new(SwitchMVeConfigCreate)
	return m
}

type SwitchMVeConfigCreateResult interface{}

func MakeSwitchMVeConfigCreateResult() SwitchMVeConfigCreateResult {
	m := new(SwitchMVeConfigCreateResult)
	return m
}

type SwitchMVeConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	// Constraints:
	//    - nullable
	List []*SwitchMVeConfig `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total configs count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMVeConfigList() *SwitchMVeConfigList {
	m := new(SwitchMVeConfigList)
	return m
}

type SwitchMVeConfigModify struct {
	// DhcpRelayAgent
	// DHCP Replay IP Address
	// Constraints:
	//    - nullable
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	// Constraints:
	//    - nullable
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	// name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// OspfArea
	// OSPF IP Address
	// Constraints:
	//    - nullable
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	// Constraints:
	//    - nullable
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// SubnetMask
	// Subnet Mask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch ID
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// VlanId
	// VLAN ID
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVeConfigModify() *SwitchMVeConfigModify {
	m := new(SwitchMVeConfigModify)
	return m
}

type SwitchMVeConfig struct {
	// CreatedTime
	// Created Time
	// Constraints:
	//    - nullable
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpRelayAgent
	// DHCP Replay IP Address
	// Constraints:
	//    - nullable
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group ID
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// ID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// InAclConfigName
	// Ingress ACL Config Name
	// Constraints:
	//    - nullable
	InAclConfigName *string `json:"inAclConfigName,omitempty"`

	// InAclConfigUUID
	// Ingress ACL Config UUID
	// Constraints:
	//    - nullable
	InAclConfigUUID *string `json:"inAclConfigUUID,omitempty"`

	// IpAddress
	// IP Address
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	// name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// OspfArea
	// OSPF IP Address
	// Constraints:
	//    - nullable
	OspfArea *string `json:"ospfArea,omitempty"`

	// OutAclConfigName
	// Egress ACL Config Name
	// Constraints:
	//    - nullable
	OutAclConfigName *string `json:"outAclConfigName,omitempty"`

	// OutAclConfigUUID
	// Egress ACL Config UUID
	// Constraints:
	//    - nullable
	OutAclConfigUUID *string `json:"outAclConfigUUID,omitempty"`

	// SubnetMask
	// Subnet Mask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch ID
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch Name
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// SwitchStatus
	// Switch Status
	// Constraints:
	//    - nullable
	SwitchStatus *string `json:"switchStatus,omitempty"`

	// UpdatedTime
	// Updated Time
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// VeId
	// VE Id
	// Constraints:
	//    - nullable
	VeId *int `json:"veId,omitempty"`

	// VlanId
	// VLAN ID
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVeConfig() *SwitchMVeConfig {
	m := new(SwitchMVeConfig)
	return m
}
