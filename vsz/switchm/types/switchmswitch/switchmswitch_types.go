package switchmswitch

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/switchm/types/common"
)

type AuditID struct {
	// ID
	// Audit Id
	ID *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

type BarChart struct {
	// ID
	// Identifier of the barchart
	ID *string `json:"id,omitempty"`

	// Key
	// Key of the barchart
	Key *string `json:"key,omitempty"`

	// Value
	// Metrics of the barchart
	Value *float64 `json:"value,omitempty"`
}

type ConnectedAPsQueryList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch connected AP returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch connected AP after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ConnectedDevice `json:"list,omitempty"`

	// RawDataTotalCount
	// Connected AP list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total connected AP list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type ConnectedDevice struct {
	// DomainID
	// Identifier of the management domain to which the connected device belong
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of connected device
	ID *string `json:"id,omitempty"`

	// IsRuckusAP
	// Remote connected device is Ruckus AP, True or False.
	IsRuckusAP *string `json:"isRuckusAP,omitempty"`

	// LocalPort
	// Local port which connect to remote device
	LocalPort *string `json:"localPort,omitempty"`

	// LocalPortIfaceName
	// Local port interface name
	LocalPortIfaceName *string `json:"localPortIfaceName,omitempty"`

	// LocalPortMac
	// Local port mac address
	LocalPortMac *string `json:"localPortMac,omitempty"`

	// RemoteDeviceName
	// Remote connected device name
	RemoteDeviceName *string `json:"remoteDeviceName,omitempty"`

	// RemotePort
	// Remote device port which connected to local device
	RemotePort *string `json:"remotePort,omitempty"`

	// RemotePortDesc
	// Remote connected device port description
	RemotePortDesc *string `json:"remotePortDesc,omitempty"`

	// RemotePortMac
	// Remote connected device port mac address
	RemotePortMac *string `json:"remotePortMac,omitempty"`

	// RemotePortType
	// Remote connected device port type
	RemotePortType *string `json:"remotePortType,omitempty"`

	// SampledInstant
	// Sampled instant
	SampledInstant map[string]interface{} `json:"sampledInstant,omitempty"`

	// SwitchGroup
	// Switch group
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchGroupLevelOneID
	// Switch group level one Id
	SwitchGroupLevelOneID *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoID
	// Switch group level two Id
	SwitchGroupLevelTwoID *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch name
	SwitchName *string `json:"switchName,omitempty"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`

	// UnitID
	// Unit Id
	UnitID *string `json:"unitId,omitempty"`
}

type ConnectedDevicesQueryList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch connected devices returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch connected devices after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ConnectedDevice `json:"list,omitempty"`

	// RawDataTotalCount
	// Connected devices list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total connected devices list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type DeleteSwitchesResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first delete switches returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more delete switches after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*AuditID `json:"list,omitempty"`

	// RawDataTotalCount
	// Delete switches list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total delete switches list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type Firmware struct {
	// FromVersion
	// Original fireware version before firmware update
	FromVersion *string `json:"fromVersion,omitempty"`

	// Timestamp
	// Timestamp of fireware update
	Timestamp *float64 `json:"timestamp,omitempty"`

	// ToVersion
	// Firmware version after firmware update
	ToVersion *string `json:"toVersion,omitempty"`
}

type FirmwareHistoryQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware history returned out of the complete query list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more firmware history after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Firmware `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware history list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total firmware history list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type NetworkSwitch struct {
	// Alarm
	// Count of switch alarm
	Alarm *int `json:"alarm,omitempty"`

	// DefaultGateway
	// Default gateway of switch
	DefaultGateway *string `json:"defaultGateway,omitempty"`

	// DomainID
	// Identifier of the management domain to which the switch belong
	DomainID *string `json:"domainId,omitempty"`

	// FirmwareUpdate
	// Information of firmware update
	FirmwareUpdate *NetworkSwitchFirmwareUpdateType `json:"firmwareUpdate,omitempty"`

	// FirmwareVersion
	// Switch firmware version
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// GroupID
	// Identifier of switch group
	GroupID *string `json:"groupId,omitempty"`

	// GroupName
	// Name of switch group
	GroupName *string `json:"groupName,omitempty"`

	// ID
	// Identifier of switch
	ID *string `json:"id,omitempty"`

	// IPAddress
	// switch IP address
	IPAddress *string `json:"ipAddress,omitempty"`

	// IPAddressType
	// IP address type
	IPAddressType *string `json:"ipAddressType,omitempty"`

	// LastBackupStatus
	// Last config backup status of switch
	LastBackupStatus *string `json:"lastBackupStatus,omitempty"`

	// LastBackupTime
	// Last config backup time of switch
	LastBackupTime *int `json:"lastBackupTime,omitempty"`

	// LastRestoreStatus
	// Last config restore status of switch
	LastRestoreStatus *string `json:"lastRestoreStatus,omitempty"`

	// LastRestoreTime
	// Last config restore time of switch
	LastRestoreTime *int `json:"lastRestoreTime,omitempty"`

	// MacAddress
	// Switch mac address
	MacAddress *string `json:"macAddress,omitempty"`

	// Model
	// Switch model
	Model *string `json:"model,omitempty"`

	// Modules
	// Stack or Switch
	Modules *string `json:"modules,omitempty"`

	// NumOfUnits
	// Count of switch unit
	NumOfUnits *int `json:"numOfUnits,omitempty"`

	// Poe
	// Information of PoE
	Poe *NetworkSwitchPoeType `json:"poe,omitempty"`

	// Ports
	// Total port count
	Ports *int `json:"ports,omitempty"`

	// PortStatus
	// Information of port status
	PortStatus *NetworkSwitchPortStatusType `json:"portStatus,omitempty"`

	// RegistrationStatus
	// Status for switch registater to ICX-M
	RegistrationStatus *string `json:"registrationStatus,omitempty"`

	// SerialNumber
	// SWitch serial number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// StackID
	// Stack Id
	StackID *string `json:"stackId,omitempty"`

	// Status
	// Status of switch, Ex: ONLINE, OFFLINE
	Status *string `json:"status,omitempty"`

	// SwitchName
	// Switch name
	SwitchName *string `json:"switchName,omitempty"`

	// UpTime
	// SWitch uptime
	UpTime *string `json:"upTime,omitempty"`
}

// NetworkSwitchFirmwareUpdateType
//
// Information of firmware update
type NetworkSwitchFirmwareUpdateType struct {
	// ModifiedTime
	// Modified time of the firmware update scheduled
	ModifiedTime *string `json:"modifiedTime,omitempty"`

	// ScheduledTime
	// Scheduled time of firmware update
	ScheduledTime *string `json:"scheduledTime,omitempty"`

	// ScheduleID
	// Schedule Id of firmware update
	ScheduleID *string `json:"scheduleId,omitempty"`

	// Status
	// Status of firmware update
	Status *string `json:"status,omitempty"`

	// ToVersion
	// Update to which firmware version
	ToVersion *string `json:"toVersion,omitempty"`
}

// NetworkSwitchPoeType
//
// Information of PoE
type NetworkSwitchPoeType struct {
	// Free
	// Free power capacity of a switch
	Free *int `json:"free,omitempty"`

	// Percent
	// Percentage of power usage for a switch
	Percent *float64 `json:"percent,omitempty"`

	// Total
	// Total power capacity of a switch
	Total *int `json:"total,omitempty"`
}

// NetworkSwitchPortStatusType
//
// Information of port status
type NetworkSwitchPortStatusType struct {
	// AdminDown
	// Count for port status is admin down of switch
	AdminDown *int `json:"adminDown,omitempty"`

	// Down
	// Count for port status is down of switch
	Down *int `json:"down,omitempty"`

	// Speed
	// Port speed of switch
	Speed *string `json:"speed,omitempty"`

	// SpeedInt
	// Switch port fully speed
	SpeedInt *int `json:"speedInt,omitempty"`

	// Total
	// Total count for port status of switch
	Total *int `json:"total,omitempty"`

	// Up
	// Count for port status is up of switch
	Up *int `json:"up,omitempty"`

	// Warning
	// Count for port status is warring of switch
	Warning *int `json:"warning,omitempty"`
}

type PortDetails struct {
	// AdminStatus
	// Admin status of switch port, UP or DOWN
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	ConnectedDevice *PortDetailsConnectedDeviceType `json:"connectedDevice,omitempty"`

	// ID
	// Identifier of switch port
	ID *string `json:"id,omitempty"`

	// InUtilization
	// Switch port traffic in utilization
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
	// Switch port connected neighbor name
	NeighborName *string `json:"neighborName,omitempty"`

	// OpticsType
	// Switch port optics type
	OpticsType *string `json:"opticsType,omitempty"`

	// OutUtilization
	// Switch port traffic out utilization
	OutUtilization *float64 `json:"outUtilization,omitempty"`

	// Packets
	// Port packet transmit information
	Packets *PortDetailsPacketsType `json:"packets,omitempty"`

	// Poe
	// POE information of switch port
	Poe *PortDetailsPoeType `json:"poe,omitempty"`

	// PoeEnabled
	// POE Enabled, True or False
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoeType
	// POE type
	PoeType *string `json:"poeType,omitempty"`

	// PortError
	// Port error Information
	PortError *PortDetailsPortErrorType `json:"portError,omitempty"`

	// PortIdentifier
	// Port Identifier of switch port
	PortIdentifier *string `json:"portIdentifier,omitempty"`

	// PortSpeed
	// Switch port speed
	PortSpeed *string `json:"portSpeed,omitempty"`

	// SampledInstant
	// Sampled instant of switch port
	SampledInstant *string `json:"sampledInstant,omitempty"`

	// Status
	// Status of switch port, UP or DOWN
	Status *string `json:"status,omitempty"`

	// StpState
	// Switch port STP state
	StpState *int `json:"stpState,omitempty"`

	// SwitchGroup
	// Switch group of switch port
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchName
	// Switch Name of switch port
	SwitchName *string `json:"switchName,omitempty"`

	// TrafficUsage
	// Traffic usage information
	TrafficUsage *PortDetailsTrafficUsageType `json:"trafficUsage,omitempty"`

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

// PortDetailsConnectedDeviceType
//
// Connected device information
type PortDetailsConnectedDeviceType struct {
	// DomainID
	// Identifier of the management domain to which the connected device belong
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of connected device
	ID *string `json:"id,omitempty"`

	// IsRuckusAP
	// Connected devices is RuckusAP,True or False
	IsRuckusAP *string `json:"isRuckusAP,omitempty"`

	// LocalPort
	// Local port description to connected device
	LocalPort *string `json:"localPort,omitempty"`

	// LocalPortIfaceName
	// Local port interface name
	LocalPortIfaceName *string `json:"localPortIfaceName,omitempty"`

	// LocalPortMac
	// Local port mac address to connected device
	LocalPortMac *string `json:"localPortMac,omitempty"`

	// RemoteDeviceName
	// Remote connected device name
	RemoteDeviceName *string `json:"remoteDeviceName,omitempty"`

	// RemotePort
	// Remote port number of connected device
	RemotePort *string `json:"remotePort,omitempty"`

	// RemotePortDesc
	// Remote port description of connected device
	RemotePortDesc *string `json:"remotePortDesc,omitempty"`

	// RemotePortMac
	// Remote port mac address of local device
	RemotePortMac *string `json:"remotePortMac,omitempty"`

	// RemotePortType
	// Remote port type of connected device
	RemotePortType *string `json:"remotePortType,omitempty"`

	// SwitchGroup
	// Switch group
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchGroupLevelOneID
	// Switch group level one Id
	SwitchGroupLevelOneID *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoID
	// Switch group level two Id
	SwitchGroupLevelTwoID *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch name
	SwitchName *string `json:"switchName,omitempty"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`

	// UnitID
	// Unit Id
	UnitID *string `json:"unitId,omitempty"`
}

// PortDetailsPacketsType
//
// Port packet transmit information
type PortDetailsPacketsType struct {
	// BroadcastIn
	// Switch port broadcast in packet count
	BroadcastIn *int `json:"broadcastIn,omitempty"`

	// BroadcastOut
	// Switch port broadcast out packet count
	BroadcastOut *int `json:"broadcastOut,omitempty"`

	// MulticastIn
	// Switch port multicast in packet count
	MulticastIn *int `json:"multicastIn,omitempty"`

	// MulticastOut
	// Switch port multicast out packet count
	MulticastOut *int `json:"multicastOut,omitempty"`
}

// PortDetailsPoeType
//
// POE information of switch port
type PortDetailsPoeType struct {
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

// PortDetailsPortErrorType
//
// Port error Information
type PortDetailsPortErrorType struct {
	// CrcError
	// Switch port CRC error count
	CrcError *int `json:"crcError,omitempty"`

	// InDiscard
	// Switch port traffic in discard count
	InDiscard *int `json:"inDiscard,omitempty"`

	// InError
	// Switch port traffic in error count
	InError *int `json:"inError,omitempty"`

	// OutError
	// Switch port traffic out error count
	OutError *int `json:"outError,omitempty"`
}

type PortDetailsQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch port detail returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch port detail after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortDetails `json:"list,omitempty"`

	// RawDataTotalCount
	// Switch port detail list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch port detail list count
	TotalCount *int `json:"totalCount,omitempty"`
}

// PortDetailsTrafficUsageType
//
// Traffic usage information
type PortDetailsTrafficUsageType struct {
	// Rx
	// Rx traffic usage of switch port
	Rx *int `json:"rx,omitempty"`

	// Tx
	// Tx traffic usage of switch port
	Tx *int `json:"tx,omitempty"`
}

type PortStatus struct {
	// AdminDown
	// Admin down port count
	AdminDown *int `json:"adminDown,omitempty"`

	// Down
	// Down port count
	Down *int `json:"down,omitempty"`

	// Speed
	// Switch port speed
	Speed *string `json:"speed,omitempty"`

	// SpeedInt
	// Switch port fully speed
	SpeedInt *int `json:"speedInt,omitempty"`

	// Total
	// Total port count
	Total *int `json:"total,omitempty"`

	// Up
	// Up port count
	Up *int `json:"up,omitempty"`

	// Warning
	// Warring port count
	Warning *int `json:"warning,omitempty"`
}

type StackMember struct {
	// ActiveMode
	// Role of stack
	ActiveMode *string `json:"activeMode,omitempty"`

	// Model
	// Switch model of stack
	Model *string `json:"model,omitempty"`

	// Ports
	// Port count of stack
	Ports *int `json:"ports,omitempty"`

	PortStatus *PortStatus `json:"portStatus,omitempty"`

	// SerialNumber
	// Serial number of stack
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SwitchModule
	// Switch module of stack
	SwitchModule *string `json:"switchModule,omitempty"`

	// SwitchName
	// Switch name of stack
	SwitchName *string `json:"switchName,omitempty"`

	SwitchPorts *PortDetails `json:"switchPorts,omitempty"`

	// SwitchUnit
	// Switch unit of stack
	SwitchUnit *string `json:"switchUnit,omitempty"`
}

type StackMemberQueryResult struct {
	// Extra
	// Extra information for stack member list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first stack member returned out of the complete stack member list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more stack member after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*StackMember `json:"list,omitempty"`

	// RawDataTotalCount
	// Total stack member count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current stack member count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchIDList []string

type SwitchPortsSummaryQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch ports summary returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch ports summary after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortStatus `json:"list,omitempty"`

	// RawDataTotalCount
	// Switch ports summary list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch ports summary list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchQueryResultList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more  after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*NetworkSwitch `json:"list,omitempty"`

	// RawDataTotalCount
	// Switch query result list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch query result list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type TopSwitchesByFirmwareQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top switches by firmware returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more top switches by firmware after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*BarChart `json:"list,omitempty"`

	// RawDataTotalCount
	// Top switches by firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total top switches by firmware list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type TopSwitchesByModelQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top switches by model returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are top switches by model after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*BarChart `json:"list,omitempty"`

	// RawDataTotalCount
	// Top switches by model list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total top switches by model list count
	TotalCount *int `json:"totalCount,omitempty"`
}
