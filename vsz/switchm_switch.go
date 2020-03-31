package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchService(c *APIClient) *SwitchMSwitchService {
	s := new(SwitchMSwitchService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchService() *SwitchMSwitchService {
	return NewSwitchMSwitchService(ss.apiClient)
}

type SwitchMSwitchAuditId struct {
	// Id
	// Audit Id
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchAuditId() *SwitchMSwitchAuditId {
	m := new(SwitchMSwitchAuditId)
	return m
}

type SwitchMSwitchBarChart struct {
	// Id
	// Identifier of the barchart
	Id *string `json:"id,omitempty"`

	// Key
	// Key of the barchart
	Key *string `json:"key,omitempty"`

	// Value
	// Metrics of the barchart
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMSwitchBarChart() *SwitchMSwitchBarChart {
	m := new(SwitchMSwitchBarChart)
	return m
}

type SwitchMSwitchConnectedAPsQueryList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchConnectedAPsQueryListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch connected AP returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch connected AP after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchConnectedDevice `json:"list,omitempty"`

	// RawDataTotalCount
	// Connected AP list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total connected AP list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchConnectedAPsQueryList() *SwitchMSwitchConnectedAPsQueryList {
	m := new(SwitchMSwitchConnectedAPsQueryList)
	return m
}

// SwitchMSwitchConnectedAPsQueryListExtraType
//
// Any additional response data
type SwitchMSwitchConnectedAPsQueryListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchConnectedAPsQueryListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchConnectedAPsQueryListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchConnectedAPsQueryListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchConnectedAPsQueryListExtraType() *SwitchMSwitchConnectedAPsQueryListExtraType {
	m := new(SwitchMSwitchConnectedAPsQueryListExtraType)
	return m
}

type SwitchMSwitchConnectedDevice struct {
	// DomainId
	// Identifier of the management domain to which the connected device belong
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of connected device
	Id *string `json:"id,omitempty"`

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

	RemoteDeviceMac *string `json:"remoteDeviceMac,omitempty"`

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
	SampledInstant *SwitchMSwitchConnectedDeviceSampledInstantType `json:"sampledInstant,omitempty"`

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
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// UnitId
	// Unit Id
	UnitId *string `json:"unitId,omitempty"`
}

func NewSwitchMSwitchConnectedDevice() *SwitchMSwitchConnectedDevice {
	m := new(SwitchMSwitchConnectedDevice)
	return m
}

// SwitchMSwitchConnectedDeviceSampledInstantType
//
// Sampled instant
type SwitchMSwitchConnectedDeviceSampledInstantType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchConnectedDeviceSampledInstantType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchConnectedDeviceSampledInstantType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchConnectedDeviceSampledInstantType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchConnectedDeviceSampledInstantType() *SwitchMSwitchConnectedDeviceSampledInstantType {
	m := new(SwitchMSwitchConnectedDeviceSampledInstantType)
	return m
}

type SwitchMSwitchConnectedDevicesQueryList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchConnectedDevicesQueryListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch connected devices returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch connected devices after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchConnectedDevice `json:"list,omitempty"`

	// RawDataTotalCount
	// Connected devices list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total connected devices list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchConnectedDevicesQueryList() *SwitchMSwitchConnectedDevicesQueryList {
	m := new(SwitchMSwitchConnectedDevicesQueryList)
	return m
}

// SwitchMSwitchConnectedDevicesQueryListExtraType
//
// Any additional response data
type SwitchMSwitchConnectedDevicesQueryListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchConnectedDevicesQueryListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchConnectedDevicesQueryListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchConnectedDevicesQueryListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchConnectedDevicesQueryListExtraType() *SwitchMSwitchConnectedDevicesQueryListExtraType {
	m := new(SwitchMSwitchConnectedDevicesQueryListExtraType)
	return m
}

type SwitchMSwitchDeleteSwitchesResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchDeleteSwitchesResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first delete switches returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more delete switches after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchAuditId `json:"list,omitempty"`

	// RawDataTotalCount
	// Delete switches list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total delete switches list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchDeleteSwitchesResultList() *SwitchMSwitchDeleteSwitchesResultList {
	m := new(SwitchMSwitchDeleteSwitchesResultList)
	return m
}

// SwitchMSwitchDeleteSwitchesResultListExtraType
//
// Any additional response data
type SwitchMSwitchDeleteSwitchesResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchDeleteSwitchesResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchDeleteSwitchesResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchDeleteSwitchesResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchDeleteSwitchesResultListExtraType() *SwitchMSwitchDeleteSwitchesResultListExtraType {
	m := new(SwitchMSwitchDeleteSwitchesResultListExtraType)
	return m
}

type SwitchMSwitchFirmware struct {
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

func NewSwitchMSwitchFirmware() *SwitchMSwitchFirmware {
	m := new(SwitchMSwitchFirmware)
	return m
}

type SwitchMSwitchFirmwareHistoryQueryResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchFirmwareHistoryQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware history returned out of the complete query list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more firmware history after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchFirmware `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware history list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total firmware history list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchFirmwareHistoryQueryResultList() *SwitchMSwitchFirmwareHistoryQueryResultList {
	m := new(SwitchMSwitchFirmwareHistoryQueryResultList)
	return m
}

// SwitchMSwitchFirmwareHistoryQueryResultListExtraType
//
// Any additional response data
type SwitchMSwitchFirmwareHistoryQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchFirmwareHistoryQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchFirmwareHistoryQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchFirmwareHistoryQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchFirmwareHistoryQueryResultListExtraType() *SwitchMSwitchFirmwareHistoryQueryResultListExtraType {
	m := new(SwitchMSwitchFirmwareHistoryQueryResultListExtraType)
	return m
}

type SwitchMSwitchNetworkSwitch struct {
	// Alarm
	// Count of switch alarm
	Alarm *int `json:"alarm,omitempty"`

	// DefaultGateway
	// Default gateway of switch
	DefaultGateway *string `json:"defaultGateway,omitempty"`

	// DomainId
	// Identifier of the management domain to which the switch belong
	DomainId *string `json:"domainId,omitempty"`

	// Family
	// Switch Model Family
	Family *string `json:"family,omitempty"`

	// FirmwareUpdate
	// Information of firmware update
	FirmwareUpdate *SwitchMSwitchNetworkSwitchFirmwareUpdateType `json:"firmwareUpdate,omitempty"`

	// FirmwareVersion
	// Switch firmware version
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// GroupFirmware
	// Firmware of switch group
	GroupFirmware *string `json:"groupFirmware,omitempty"`

	// GroupId
	// Identifier of switch group
	GroupId *string `json:"groupId,omitempty"`

	// GroupName
	// Name of switch group
	GroupName *string `json:"groupName,omitempty"`

	// Id
	// Identifier of switch
	Id *string `json:"id,omitempty"`

	// IpAddress
	// switch IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpAddressType
	// IP address type
	IpAddressType *string `json:"ipAddressType,omitempty"`

	// LastBackupStatus
	// Last config backup status of switch
	// Constraints:
	//    - oneof:[STARTED,SUCCESS,FAILED]
	LastBackupStatus *string `json:"lastBackupStatus,omitempty" validate:"oneof=STARTED SUCCESS FAILED"`

	// LastBackupTime
	// Last config backup time of switch
	LastBackupTime *int `json:"lastBackupTime,omitempty"`

	// LastRestoreStatus
	// Last config restore status of switch
	// Constraints:
	//    - oneof:[STARTED,SUCCESS,FAILED]
	LastRestoreStatus *string `json:"lastRestoreStatus,omitempty" validate:"oneof=STARTED SUCCESS FAILED"`

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

	// ParentGroupId
	// Identifier of parent switch group
	ParentGroupId *string `json:"parentGroupId,omitempty"`

	// Poe
	// Information of PoE
	Poe *SwitchMSwitchNetworkSwitchPoeType `json:"poe,omitempty"`

	// Ports
	// Total port count
	Ports *int `json:"ports,omitempty"`

	// PortStatus
	// Information of port status
	PortStatus *SwitchMSwitchNetworkSwitchPortStatusType `json:"portStatus,omitempty"`

	// RegistrationStatus
	// Status for switch registater to ICX-M
	RegistrationStatus *string `json:"registrationStatus,omitempty"`

	// SerialNumber
	// SWitch serial number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// StackId
	// Stack Id
	StackId *string `json:"stackId,omitempty"`

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

func NewSwitchMSwitchNetworkSwitch() *SwitchMSwitchNetworkSwitch {
	m := new(SwitchMSwitchNetworkSwitch)
	return m
}

// SwitchMSwitchNetworkSwitchFirmwareUpdateType
//
// Information of firmware update
type SwitchMSwitchNetworkSwitchFirmwareUpdateType struct {
	// ModifiedTime
	// Modified time of the firmware update scheduled
	ModifiedTime *string `json:"modifiedTime,omitempty"`

	// ScheduledTime
	// Scheduled time of firmware update
	ScheduledTime *string `json:"scheduledTime,omitempty"`

	// ScheduleId
	// Schedule Id of firmware update
	ScheduleId *string `json:"scheduleId,omitempty"`

	// Status
	// Status of firmware update
	Status *string `json:"status,omitempty"`

	// ToVersion
	// Update to which firmware version
	ToVersion *string `json:"toVersion,omitempty"`
}

func NewSwitchMSwitchNetworkSwitchFirmwareUpdateType() *SwitchMSwitchNetworkSwitchFirmwareUpdateType {
	m := new(SwitchMSwitchNetworkSwitchFirmwareUpdateType)
	return m
}

// SwitchMSwitchNetworkSwitchPoeType
//
// Information of PoE
type SwitchMSwitchNetworkSwitchPoeType struct {
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

func NewSwitchMSwitchNetworkSwitchPoeType() *SwitchMSwitchNetworkSwitchPoeType {
	m := new(SwitchMSwitchNetworkSwitchPoeType)
	return m
}

// SwitchMSwitchNetworkSwitchPortStatusType
//
// Information of port status
type SwitchMSwitchNetworkSwitchPortStatusType struct {
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

func NewSwitchMSwitchNetworkSwitchPortStatusType() *SwitchMSwitchNetworkSwitchPortStatusType {
	m := new(SwitchMSwitchNetworkSwitchPortStatusType)
	return m
}

type SwitchMSwitchPortDetails struct {
	// AdminStatus
	// Admin status of switch port, UP or DOWN
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	ConnectedDevice *SwitchMSwitchPortDetailsConnectedDeviceType `json:"connectedDevice,omitempty"`

	// Id
	// Identifier of switch port
	Id *string `json:"id,omitempty"`

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
	Packets *SwitchMSwitchPortDetailsPacketsType `json:"packets,omitempty"`

	// Poe
	// PoE information of switch port
	Poe *SwitchMSwitchPortDetailsPoeType `json:"poe,omitempty"`

	// PoeEnabled
	// PoE Enabled, True or False
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoeType
	// PoE type
	PoeType *string `json:"poeType,omitempty"`

	// PortError
	// Port error Information
	PortError *SwitchMSwitchPortDetailsPortErrorType `json:"portError,omitempty"`

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
	TrafficUsage *SwitchMSwitchPortDetailsTrafficUsageType `json:"trafficUsage,omitempty"`

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

func NewSwitchMSwitchPortDetails() *SwitchMSwitchPortDetails {
	m := new(SwitchMSwitchPortDetails)
	return m
}

// SwitchMSwitchPortDetailsConnectedDeviceType
//
// Connected device information
type SwitchMSwitchPortDetailsConnectedDeviceType struct {
	// DomainId
	// Identifier of the management domain to which the connected device belong
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of connected device
	Id *string `json:"id,omitempty"`

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

	RemoteDeviceMac *string `json:"remoteDeviceMac,omitempty"`

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
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// UnitId
	// Unit Id
	UnitId *string `json:"unitId,omitempty"`
}

func NewSwitchMSwitchPortDetailsConnectedDeviceType() *SwitchMSwitchPortDetailsConnectedDeviceType {
	m := new(SwitchMSwitchPortDetailsConnectedDeviceType)
	return m
}

// SwitchMSwitchPortDetailsPacketsType
//
// Port packet transmit information
type SwitchMSwitchPortDetailsPacketsType struct {
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

func NewSwitchMSwitchPortDetailsPacketsType() *SwitchMSwitchPortDetailsPacketsType {
	m := new(SwitchMSwitchPortDetailsPacketsType)
	return m
}

// SwitchMSwitchPortDetailsPoeType
//
// PoE information of switch port
type SwitchMSwitchPortDetailsPoeType struct {
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

func NewSwitchMSwitchPortDetailsPoeType() *SwitchMSwitchPortDetailsPoeType {
	m := new(SwitchMSwitchPortDetailsPoeType)
	return m
}

// SwitchMSwitchPortDetailsPortErrorType
//
// Port error Information
type SwitchMSwitchPortDetailsPortErrorType struct {
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

func NewSwitchMSwitchPortDetailsPortErrorType() *SwitchMSwitchPortDetailsPortErrorType {
	m := new(SwitchMSwitchPortDetailsPortErrorType)
	return m
}

type SwitchMSwitchPortDetailsQueryResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchPortDetailsQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch port detail returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch port detail after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchPortDetails `json:"list,omitempty"`

	// RawDataTotalCount
	// Switch port detail list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch port detail list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchPortDetailsQueryResultList() *SwitchMSwitchPortDetailsQueryResultList {
	m := new(SwitchMSwitchPortDetailsQueryResultList)
	return m
}

// SwitchMSwitchPortDetailsQueryResultListExtraType
//
// Any additional response data
type SwitchMSwitchPortDetailsQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchPortDetailsQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchPortDetailsQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchPortDetailsQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchPortDetailsQueryResultListExtraType() *SwitchMSwitchPortDetailsQueryResultListExtraType {
	m := new(SwitchMSwitchPortDetailsQueryResultListExtraType)
	return m
}

// SwitchMSwitchPortDetailsTrafficUsageType
//
// Traffic usage information
type SwitchMSwitchPortDetailsTrafficUsageType struct {
	// Rx
	// Rx traffic usage of switch port
	Rx *int `json:"rx,omitempty"`

	// Tx
	// Tx traffic usage of switch port
	Tx *int `json:"tx,omitempty"`
}

func NewSwitchMSwitchPortDetailsTrafficUsageType() *SwitchMSwitchPortDetailsTrafficUsageType {
	m := new(SwitchMSwitchPortDetailsTrafficUsageType)
	return m
}

// SwitchMSwitchPortStatus
//
// $
type SwitchMSwitchPortStatus struct {
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

func NewSwitchMSwitchPortStatus() *SwitchMSwitchPortStatus {
	m := new(SwitchMSwitchPortStatus)
	return m
}

type SwitchMSwitchStackMember struct {
	// ActiveMode
	// Role of stack
	ActiveMode *string `json:"activeMode,omitempty"`

	// Model
	// Switch model of stack
	Model *string `json:"model,omitempty"`

	// Ports
	// Port count of stack
	Ports *int `json:"ports,omitempty"`

	PortStatus *SwitchMSwitchPortStatus `json:"portStatus,omitempty"`

	// SerialNumber
	// Serial number of stack
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SwitchModule
	// Switch module of stack
	SwitchModule *string `json:"switchModule,omitempty"`

	// SwitchName
	// Switch name of stack
	SwitchName *string `json:"switchName,omitempty"`

	SwitchPorts *SwitchMSwitchPortDetails `json:"switchPorts,omitempty"`

	// SwitchUnit
	// Switch unit of stack
	SwitchUnit *string `json:"switchUnit,omitempty"`
}

func NewSwitchMSwitchStackMember() *SwitchMSwitchStackMember {
	m := new(SwitchMSwitchStackMember)
	return m
}

type SwitchMSwitchStackMemberQueryResult struct {
	// Extra
	// Extra information for stack member list
	Extra *SwitchMSwitchStackMemberQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first stack member returned out of the complete stack member list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more stack member after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchStackMember `json:"list,omitempty"`

	// RawDataTotalCount
	// Total stack member count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current stack member count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchStackMemberQueryResult() *SwitchMSwitchStackMemberQueryResult {
	m := new(SwitchMSwitchStackMemberQueryResult)
	return m
}

// SwitchMSwitchStackMemberQueryResultExtraType
//
// Extra information for stack member list
type SwitchMSwitchStackMemberQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchStackMemberQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchStackMemberQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchStackMemberQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchStackMemberQueryResultExtraType() *SwitchMSwitchStackMemberQueryResultExtraType {
	m := new(SwitchMSwitchStackMemberQueryResultExtraType)
	return m
}

// SwitchMSwitchIdList
//
// $
type SwitchMSwitchIdList []string

func MakeSwitchMSwitchIdList() SwitchMSwitchIdList {
	m := make(SwitchMSwitchIdList, 0)
	return m
}

type SwitchMSwitchPortsSummaryQueryResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchPortsSummaryQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch ports summary returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch ports summary after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchPortStatus `json:"list,omitempty"`

	// RawDataTotalCount
	// Switch ports summary list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch ports summary list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchPortsSummaryQueryResultList() *SwitchMSwitchPortsSummaryQueryResultList {
	m := new(SwitchMSwitchPortsSummaryQueryResultList)
	return m
}

// SwitchMSwitchPortsSummaryQueryResultListExtraType
//
// Any additional response data
type SwitchMSwitchPortsSummaryQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchPortsSummaryQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchPortsSummaryQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchPortsSummaryQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchPortsSummaryQueryResultListExtraType() *SwitchMSwitchPortsSummaryQueryResultListExtraType {
	m := new(SwitchMSwitchPortsSummaryQueryResultListExtraType)
	return m
}

type SwitchMSwitchQueryResultList struct {
	Extra *SwitchMCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more  after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchNetworkSwitch `json:"list,omitempty"`

	// RawDataTotalCount
	// Switch query result list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch query result list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchQueryResultList() *SwitchMSwitchQueryResultList {
	m := new(SwitchMSwitchQueryResultList)
	return m
}

type SwitchMSwitchTopSwitchesByFirmwareQueryResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top switches by firmware returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more top switches by firmware after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchBarChart `json:"list,omitempty"`

	// RawDataTotalCount
	// Top switches by firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total top switches by firmware list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTopSwitchesByFirmwareQueryResultList() *SwitchMSwitchTopSwitchesByFirmwareQueryResultList {
	m := new(SwitchMSwitchTopSwitchesByFirmwareQueryResultList)
	return m
}

// SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType
//
// Any additional response data
type SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType() *SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType {
	m := new(SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType)
	return m
}

type SwitchMSwitchTopSwitchesByModelQueryResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchTopSwitchesByModelQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top switches by model returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are top switches by model after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchBarChart `json:"list,omitempty"`

	// RawDataTotalCount
	// Top switches by model list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total top switches by model list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTopSwitchesByModelQueryResultList() *SwitchMSwitchTopSwitchesByModelQueryResultList {
	m := new(SwitchMSwitchTopSwitchesByModelQueryResultList)
	return m
}

// SwitchMSwitchTopSwitchesByModelQueryResultListExtraType
//
// Any additional response data
type SwitchMSwitchTopSwitchesByModelQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTopSwitchesByModelQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTopSwitchesByModelQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTopSwitchesByModelQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTopSwitchesByModelQueryResultListExtraType() *SwitchMSwitchTopSwitchesByModelQueryResultListExtraType {
	m := new(SwitchMSwitchTopSwitchesByModelQueryResultListExtraType)
	return m
}

// AddSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitch(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitch, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchSnmpSyncedSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone and SNMP synced.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchSnmpSyncedSwitch(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchSnmpSyncedSwitch, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchViewDetails
//
// Use this API command to retrieve switch and port details for the selected Switch/SwitchGroup/Domain.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchViewDetails(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchStackMemberQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchStackMemberQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchViewDetails, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchStackMemberQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSwitch
//
// Use this API command to delete multiple switches managed by SmartZone
//
// Request Body:
//	 - body SwitchMSwitchIdList
func (s *SwitchMSwitchService) DeleteSwitch(ctx context.Context, body SwitchMSwitchIdList) (*SwitchMSwitchDeleteSwitchesResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchDeleteSwitchesResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSwitch, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchDeleteSwitchesResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSwitchById
//
// Use this API command to delete a switch managed by SmartZone.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchService) DeleteSwitchById(ctx context.Context, id string) (*SwitchMSwitchAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSwitchById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSwitchById
//
// Use this API command to retrieve a switch status.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchService) FindSwitchById(ctx context.Context, id string) (*SwitchMSwitchNetworkSwitch, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchNetworkSwitch
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSwitchById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchNetworkSwitch()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSwitchFirmwareBySwitchId
//
// Use this API command to get a list of firmware update history.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMSwitchService) FindSwitchFirmwareBySwitchId(ctx context.Context, switchId string) (*SwitchMSwitchFirmwareHistoryQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchFirmwareHistoryQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, switchId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSwitchFirmwareBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchFirmwareHistoryQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSwitchMoveByDestinationSwitchGroupId
//
// Use this API command to move a list of switches to a switch group.
//
// Request Body:
//	 - body SwitchMSwitchIdList
//
// Required Parameters:
// - destinationSwitchGroupId string
//		- required
func (s *SwitchMSwitchService) UpdateSwitchMoveByDestinationSwitchGroupId(ctx context.Context, body SwitchMSwitchIdList, destinationSwitchGroupId string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, destinationSwitchGroupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSwitchMoveByDestinationSwitchGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("destinationSwitchGroupId", destinationSwitchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
