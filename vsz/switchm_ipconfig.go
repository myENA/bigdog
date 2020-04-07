package vsz

// API Version: v9_0

type SwitchMIpConfigCreate struct {
	// DhcpRelayAgent
	// DHCP Relay IP Address
	// Constraints:
	//    - nullable
	DhcpRelayAgent *string `json:"dhcpRelayAgent,omitempty"`

	// GroupId
	// Switch Group Id
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

	// Port
	// Switch Port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`

	// PortName
	// Switch Port Name
	// Constraints:
	//    - nullable
	PortName *string `json:"portName,omitempty"`

	// SubnetMask
	// Subnet mask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMIpConfigCreate() *SwitchMIpConfigCreate {
	m := new(SwitchMIpConfigCreate)
	return m
}

type SwitchMIpConfigCreateResult interface{}

func MakeSwitchMIpConfigCreateResult() SwitchMIpConfigCreateResult {
	m := new(SwitchMIpConfigCreateResult)
	return m
}

type SwitchMIpConfig struct {
	// CreatedTime
	// Config Created Time
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
	// Config ID
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

	// Port
	// Switch Port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`

	// PortName
	// Switch Port Name
	// Constraints:
	//    - nullable
	PortName *string `json:"portName,omitempty"`

	// SubnetMask
	// Subnet mask
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
	// Config Updated Time
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMIpConfig() *SwitchMIpConfig {
	m := new(SwitchMIpConfig)
	return m
}

type SwitchMIpConfigList struct {
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
	List []*SwitchMIpConfig `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total configs count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMIpConfigList() *SwitchMIpConfigList {
	m := new(SwitchMIpConfigList)
	return m
}

type SwitchMIpConfigModify struct {
	// DhcpRelayAgent
	// DHCP Relay IP Address
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

	// Port
	// Switch Port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`

	// PortName
	// Switch Port Name
	// Constraints:
	//    - nullable
	PortName *string `json:"portName,omitempty"`

	// SubnetMask
	// Subnet mask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewSwitchMIpConfigModify() *SwitchMIpConfigModify {
	m := new(SwitchMIpConfigModify)
	return m
}
