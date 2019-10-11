package stack

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type AuditIdList []*switchmswitch.AuditId

func NewAuditIdList() *AuditIdList {
	auditIdListType := make(AuditIdList, 0)
	return &auditIdListType
}

func NewAuditIdListWithDefaults() *AuditIdList {
	auditIdListType := make(AuditIdList, 0)
	return &auditIdListType
}

type List struct {
	// FirstIndex
	// Index of the first stack returned out of the complete stack list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more stack after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// List of stack
	List []*Member `json:"list,omitempty"`

	// RawDataTotalCount
	// Stack count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Stack count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewList() *List {
	listType := new(List)
	return listType
}

func NewListWithDefaults() *List {
	listType := new(List)
	return listType
}

type Member struct {
	// ActiveMode
	// Role of stack
	ActiveMode *string `json:"activeMode,omitempty"`

	// Model
	// Switch model of stack
	Model *string `json:"model,omitempty"`

	// Ports
	// Port count  of stack
	Ports *int `json:"ports,omitempty"`

	// PortStatus
	// Port status Information
	PortStatus *MemberPortStatusType `json:"portStatus,omitempty"`

	// SerialNumber
	// Serial number of stack
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SwitchModule
	// Switch module of stack
	SwitchModule *string `json:"switchModule,omitempty"`

	// SwitchName
	// Switch name of stack
	SwitchName *string `json:"switchName,omitempty"`

	// SwitchPorts
	// Switch port information of stack
	SwitchPorts []*MemberSwitchPortsType `json:"switchPorts,omitempty"`

	// SwitchUnit
	// Switch unit of stack
	SwitchUnit *string `json:"switchUnit,omitempty"`
}

func NewMember() *Member {
	memberType := new(Member)
	return memberType
}

func NewMemberWithDefaults() *Member {
	memberType := new(Member)
	return memberType
}

// MemberPortStatusType
//
// Port status Information
type MemberPortStatusType struct {
	// AdminDown
	// Count for port status is admin down of stack
	AdminDown *int `json:"adminDown,omitempty"`

	// Down
	// Count for port status is down of stack
	Down *int `json:"down,omitempty"`

	// Speed
	// Port speed of stack
	Speed *string `json:"speed,omitempty"`

	// Total
	// Total port count of stack
	Total *int `json:"total,omitempty"`

	// Up
	// Count for port status is up of stack
	Up *int `json:"up,omitempty"`

	// Warning
	// Count for port status is warring of stack
	Warning *int `json:"warning,omitempty"`
}

func NewMemberPortStatusType() *MemberPortStatusType {
	memberPortStatusTypeType := new(MemberPortStatusType)
	return memberPortStatusTypeType
}

func NewMemberPortStatusTypeWithDefaults() *MemberPortStatusType {
	memberPortStatusTypeType := new(MemberPortStatusType)
	return memberPortStatusTypeType
}

type MemberSwitchPortsType struct {
	// AdminStatus
	// Admin Status of switch port
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	ConnectedDevice *MemberSwitchPortsTypeConnectedDeviceType `json:"connectedDevice,omitempty"`

	// Id
	// Identifier of switch port
	Id *string `json:"id,omitempty"`

	// InUtilization
	// In utilization of switch port
	InUtilization *float64 `json:"inUtilization,omitempty"`

	// LagName
	// LAG name of switch port
	LagName *string `json:"lagName,omitempty"`

	// Mac
	// Mac address of switch port
	Mac *string `json:"mac,omitempty"`

	// Name
	// Name of switch port
	Name *string `json:"name,omitempty"`

	// NeighborName
	// Neighbor name of switch port
	NeighborName *string `json:"neighborName,omitempty"`

	// OpticsType
	// Optics type of switch port
	OpticsType *string `json:"opticsType,omitempty"`

	// OutUtilization
	// Out utilization of switch port
	OutUtilization *float64 `json:"outUtilization,omitempty"`

	// Poe
	// POE information of switch port
	Poe *MemberSwitchPortsTypePoeType `json:"poe,omitempty"`

	// PoeEnabled
	// POE Enabled, True or False
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PortIdentifier
	// Port Identifier of switch port
	PortIdentifier *string `json:"portIdentifier,omitempty"`

	// PortSpeed
	// Port speed of switch port
	PortSpeed *string `json:"portSpeed,omitempty"`

	// SampledInstant
	// Sampled instant of switch port
	SampledInstant *string `json:"sampledInstant,omitempty"`

	// Status
	// Status of switch port
	Status *string `json:"status,omitempty"`

	// StpState
	// STP state of switch port
	StpState *int `json:"stpState,omitempty"`

	// SwitchGroup
	// Switch group of switch port
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchName
	// Switch name of stack
	SwitchName *string `json:"switchName,omitempty"`

	// TrafficUsage
	// Traffic usage information
	TrafficUsage *MemberSwitchPortsTypeTrafficUsageType `json:"trafficUsage,omitempty"`

	// Type
	// Type of switch port
	Type *string `json:"type,omitempty"`

	// UnTaggedVlan
	// Untagged vlan of switch port
	UnTaggedVlan *string `json:"unTaggedVlan,omitempty"`

	// UsedInFormingStack
	// Used in forming stack, True or False
	UsedInFormingStack *bool `json:"usedInFormingStack,omitempty"`

	// Vlans
	// Switch port include vlans
	Vlans *string `json:"vlans,omitempty"`
}

func NewMemberSwitchPortsType() *MemberSwitchPortsType {
	memberSwitchPortsTypeType := new(MemberSwitchPortsType)
	return memberSwitchPortsTypeType
}

func NewMemberSwitchPortsTypeWithDefaults() *MemberSwitchPortsType {
	memberSwitchPortsTypeType := new(MemberSwitchPortsType)
	return memberSwitchPortsTypeType
}

// MemberSwitchPortsTypeConnectedDeviceType
//
// Connected device information
type MemberSwitchPortsTypeConnectedDeviceType struct {
	// DomainId
	// Identifier of the management domain to which the connected device belong
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of switch port connected device
	Id *string `json:"id,omitempty"`

	// IsRuckusAP
	// Connected devices is RuckusAP,True or False
	IsRuckusAP *string `json:"isRuckusAP,omitempty"`

	// LocalPort
	// Local port description of connected device
	LocalPort *string `json:"localPort,omitempty"`

	// LocalPortIfaceName
	// Local port interface name
	LocalPortIfaceName *string `json:"localPortIfaceName,omitempty"`

	// LocalPortMac
	// Local port mac address of connected device
	LocalPortMac *string `json:"localPortMac,omitempty"`

	// RemoteDeviceName
	// Remote device name of connected device
	RemoteDeviceName *string `json:"remoteDeviceName,omitempty"`

	// RemotePort
	// Remote port number of connected device
	RemotePort *string `json:"remotePort,omitempty"`

	// RemotePortDesc
	// Remote port description of connected device
	RemotePortDesc *string `json:"remotePortDesc,omitempty"`

	// RemotePortMac
	// Remote port mac address of connected device
	RemotePortMac *string `json:"remotePortMac,omitempty"`

	// RemotePortType
	// Remote port type of connected device
	RemotePortType *string `json:"remotePortType,omitempty"`

	// SwitchGroup
	// Switch group
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchGroupLevelOneId
	// Switch group level one Id
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Switch group level two Id
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch name
	SwitchName *string `json:"switchName,omitempty"`

	// TenantId
	// Tenant Id of stack
	TenantId *string `json:"tenantId,omitempty"`

	// UnitId
	// Unit Id
	UnitId *string `json:"unitId,omitempty"`
}

func NewMemberSwitchPortsTypeConnectedDeviceType() *MemberSwitchPortsTypeConnectedDeviceType {
	memberSwitchPortsTypeConnectedDeviceTypeType := new(MemberSwitchPortsTypeConnectedDeviceType)
	return memberSwitchPortsTypeConnectedDeviceTypeType
}

func NewMemberSwitchPortsTypeConnectedDeviceTypeWithDefaults() *MemberSwitchPortsTypeConnectedDeviceType {
	memberSwitchPortsTypeConnectedDeviceTypeType := new(MemberSwitchPortsTypeConnectedDeviceType)
	return memberSwitchPortsTypeConnectedDeviceTypeType
}

// MemberSwitchPortsTypePoeType
//
// POE information of switch port
type MemberSwitchPortsTypePoeType struct {
	// Free
	// Free power capacity of switch port
	Free *int `json:"free,omitempty"`

	// Percent
	// Power used percentage of switch port
	Percent *float64 `json:"percent,omitempty"`

	// Total
	// Total power capacity of switch port
	Total *int `json:"total,omitempty"`
}

func NewMemberSwitchPortsTypePoeType() *MemberSwitchPortsTypePoeType {
	memberSwitchPortsTypePoeTypeType := new(MemberSwitchPortsTypePoeType)
	return memberSwitchPortsTypePoeTypeType
}

func NewMemberSwitchPortsTypePoeTypeWithDefaults() *MemberSwitchPortsTypePoeType {
	memberSwitchPortsTypePoeTypeType := new(MemberSwitchPortsTypePoeType)
	return memberSwitchPortsTypePoeTypeType
}

// MemberSwitchPortsTypeTrafficUsageType
//
// Traffic usage information
type MemberSwitchPortsTypeTrafficUsageType struct {
	// Rx
	// Rx traffic usage of switch port
	Rx *int `json:"rx,omitempty"`

	// Tx
	// Tx traffic usage of switch port
	Tx *int `json:"tx,omitempty"`
}

func NewMemberSwitchPortsTypeTrafficUsageType() *MemberSwitchPortsTypeTrafficUsageType {
	memberSwitchPortsTypeTrafficUsageTypeType := new(MemberSwitchPortsTypeTrafficUsageType)
	return memberSwitchPortsTypeTrafficUsageTypeType
}

func NewMemberSwitchPortsTypeTrafficUsageTypeWithDefaults() *MemberSwitchPortsTypeTrafficUsageType {
	memberSwitchPortsTypeTrafficUsageTypeType := new(MemberSwitchPortsTypeTrafficUsageType)
	return memberSwitchPortsTypeTrafficUsageTypeType
}

type StackConfig struct {
	// ActiveSwitchId
	// Switch Id of Active Unit
	ActiveSwitchId *string `json:"activeSwitchId,omitempty"`

	// IsActiveRole
	// Switch role is Active, True (Active) or False (Standby or Member)
	IsActiveRole *bool `json:"isActiveRole,omitempty"`

	// SuggestedId
	// Suggested switch unit Id in stack, 1 ~ 12
	SuggestedId *int `json:"suggestedId,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

func NewStackConfig() *StackConfig {
	stackConfigType := new(StackConfig)
	return stackConfigType
}

func NewStackConfigWithDefaults() *StackConfig {
	stackConfigType := new(StackConfig)
	return stackConfigType
}

type StackConfigList []*StackConfig

func NewStackConfigList() *StackConfigList {
	stackConfigListType := make(StackConfigList, 0)
	return &stackConfigListType
}

func NewStackConfigListWithDefaults() *StackConfigList {
	stackConfigListType := make(StackConfigList, 0)
	return &stackConfigListType
}
