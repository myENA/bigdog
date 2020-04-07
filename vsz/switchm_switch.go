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
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchAuditId() *SwitchMSwitchAuditId {
	m := new(SwitchMSwitchAuditId)
	return m
}

type SwitchMSwitchBarChart struct {
	// Id
	// Identifier of the barchart
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Key
	// Key of the barchart
	// Constraints:
	//    - nullable
	Key *string `json:"key,omitempty"`

	// Value
	// Metrics of the barchart
	// Constraints:
	//    - nullable
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMSwitchBarChart() *SwitchMSwitchBarChart {
	m := new(SwitchMSwitchBarChart)
	return m
}

type SwitchMSwitchConnectedAPsQueryList struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchConnectedAPsQueryListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch connected AP returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch connected AP after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchConnectedDevice `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Connected AP list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total connected AP list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchConnectedAPsQueryList() *SwitchMSwitchConnectedAPsQueryList {
	m := new(SwitchMSwitchConnectedAPsQueryList)
	return m
}

// SwitchMSwitchConnectedAPsQueryListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of connected device
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IsRuckusAP
	// Remote connected device is Ruckus AP, True or False.
	// Constraints:
	//    - nullable
	IsRuckusAP *string `json:"isRuckusAP,omitempty"`

	// LocalPort
	// Local port which connect to remote device
	// Constraints:
	//    - nullable
	LocalPort *string `json:"localPort,omitempty"`

	// LocalPortIfaceName
	// Local port interface name
	// Constraints:
	//    - nullable
	LocalPortIfaceName *string `json:"localPortIfaceName,omitempty"`

	// LocalPortMac
	// Local port mac address
	// Constraints:
	//    - nullable
	LocalPortMac *string `json:"localPortMac,omitempty"`

	// RemoteDeviceMac
	// Constraints:
	//    - nullable
	RemoteDeviceMac *string `json:"remoteDeviceMac,omitempty"`

	// RemoteDeviceName
	// Remote connected device name
	// Constraints:
	//    - nullable
	RemoteDeviceName *string `json:"remoteDeviceName,omitempty"`

	// RemotePort
	// Remote device port which connected to local device
	// Constraints:
	//    - nullable
	RemotePort *string `json:"remotePort,omitempty"`

	// RemotePortDesc
	// Remote connected device port description
	// Constraints:
	//    - nullable
	RemotePortDesc *string `json:"remotePortDesc,omitempty"`

	// RemotePortMac
	// Remote connected device port mac address
	// Constraints:
	//    - nullable
	RemotePortMac *string `json:"remotePortMac,omitempty"`

	// RemotePortType
	// Remote connected device port type
	// Constraints:
	//    - nullable
	RemotePortType *string `json:"remotePortType,omitempty"`

	// SampledInstant
	// Sampled instant
	// Constraints:
	//    - nullable
	SampledInstant *SwitchMSwitchConnectedDeviceSampledInstantType `json:"sampledInstant,omitempty"`

	// SwitchGroup
	// Switch group
	// Constraints:
	//    - nullable
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchGroupLevelOneId
	// Switch group level one Id
	// Constraints:
	//    - nullable
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Switch group level two Id
	// Constraints:
	//    - nullable
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch name
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// UnitId
	// Unit Id
	// Constraints:
	//    - nullable
	UnitId *string `json:"unitId,omitempty"`
}

func NewSwitchMSwitchConnectedDevice() *SwitchMSwitchConnectedDevice {
	m := new(SwitchMSwitchConnectedDevice)
	return m
}

// SwitchMSwitchConnectedDeviceSampledInstantType
//
// Sampled instant
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchConnectedDevicesQueryListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch connected devices returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch connected devices after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchConnectedDevice `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Connected devices list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total connected devices list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchConnectedDevicesQueryList() *SwitchMSwitchConnectedDevicesQueryList {
	m := new(SwitchMSwitchConnectedDevicesQueryList)
	return m
}

// SwitchMSwitchConnectedDevicesQueryListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchDeleteSwitchesResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first delete switches returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more delete switches after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchAuditId `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Delete switches list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total delete switches list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchDeleteSwitchesResultList() *SwitchMSwitchDeleteSwitchesResultList {
	m := new(SwitchMSwitchDeleteSwitchesResultList)
	return m
}

// SwitchMSwitchDeleteSwitchesResultListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	FromVersion *string `json:"fromVersion,omitempty"`

	// Timestamp
	// Timestamp of fireware update
	// Constraints:
	//    - nullable
	Timestamp *float64 `json:"timestamp,omitempty"`

	// ToVersion
	// Firmware version after firmware update
	// Constraints:
	//    - nullable
	ToVersion *string `json:"toVersion,omitempty"`
}

func NewSwitchMSwitchFirmware() *SwitchMSwitchFirmware {
	m := new(SwitchMSwitchFirmware)
	return m
}

type SwitchMSwitchFirmwareHistoryQueryResultList struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchFirmwareHistoryQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware history returned out of the complete query list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more firmware history after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchFirmware `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Firmware history list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total firmware history list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchFirmwareHistoryQueryResultList() *SwitchMSwitchFirmwareHistoryQueryResultList {
	m := new(SwitchMSwitchFirmwareHistoryQueryResultList)
	return m
}

// SwitchMSwitchFirmwareHistoryQueryResultListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	Alarm *int `json:"alarm,omitempty"`

	// DefaultGateway
	// Default gateway of switch
	// Constraints:
	//    - nullable
	DefaultGateway *string `json:"defaultGateway,omitempty"`

	// DomainId
	// Identifier of the management domain to which the switch belong
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Family
	// Switch Model Family
	// Constraints:
	//    - nullable
	Family *string `json:"family,omitempty"`

	// FirmwareUpdate
	// Information of firmware update
	// Constraints:
	//    - nullable
	FirmwareUpdate *SwitchMSwitchNetworkSwitchFirmwareUpdateType `json:"firmwareUpdate,omitempty"`

	// FirmwareVersion
	// Switch firmware version
	// Constraints:
	//    - nullable
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// GroupFirmware
	// Firmware of switch group
	// Constraints:
	//    - nullable
	GroupFirmware *string `json:"groupFirmware,omitempty"`

	// GroupId
	// Identifier of switch group
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// GroupName
	// Name of switch group
	// Constraints:
	//    - nullable
	GroupName *string `json:"groupName,omitempty"`

	// Id
	// Identifier of switch
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpAddress
	// switch IP address
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpAddressType
	// IP address type
	// Constraints:
	//    - nullable
	IpAddressType *string `json:"ipAddressType,omitempty"`

	// LastBackupStatus
	// Last config backup status of switch
	// Constraints:
	//    - nullable
	//    - oneof:[STARTED,SUCCESS,FAILED]
	LastBackupStatus *string `json:"lastBackupStatus,omitempty" validate:"omitempty,oneof=STARTED SUCCESS FAILED"`

	// LastBackupTime
	// Last config backup time of switch
	// Constraints:
	//    - nullable
	LastBackupTime *int `json:"lastBackupTime,omitempty"`

	// LastRestoreStatus
	// Last config restore status of switch
	// Constraints:
	//    - nullable
	//    - oneof:[STARTED,SUCCESS,FAILED]
	LastRestoreStatus *string `json:"lastRestoreStatus,omitempty" validate:"omitempty,oneof=STARTED SUCCESS FAILED"`

	// LastRestoreTime
	// Last config restore time of switch
	// Constraints:
	//    - nullable
	LastRestoreTime *int `json:"lastRestoreTime,omitempty"`

	// MacAddress
	// Switch mac address
	// Constraints:
	//    - nullable
	MacAddress *string `json:"macAddress,omitempty"`

	// Model
	// Switch model
	// Constraints:
	//    - nullable
	Model *string `json:"model,omitempty"`

	// Modules
	// Stack or Switch
	// Constraints:
	//    - nullable
	Modules *string `json:"modules,omitempty"`

	// NumOfUnits
	// Count of switch unit
	// Constraints:
	//    - nullable
	NumOfUnits *int `json:"numOfUnits,omitempty"`

	// ParentGroupId
	// Identifier of parent switch group
	// Constraints:
	//    - nullable
	ParentGroupId *string `json:"parentGroupId,omitempty"`

	// Poe
	// Information of PoE
	// Constraints:
	//    - nullable
	Poe *SwitchMSwitchNetworkSwitchPoeType `json:"poe,omitempty"`

	// Ports
	// Total port count
	// Constraints:
	//    - nullable
	Ports *int `json:"ports,omitempty"`

	// PortStatus
	// Information of port status
	// Constraints:
	//    - nullable
	PortStatus *SwitchMSwitchNetworkSwitchPortStatusType `json:"portStatus,omitempty"`

	// RegistrationStatus
	// Status for switch registater to ICX-M
	// Constraints:
	//    - nullable
	RegistrationStatus *string `json:"registrationStatus,omitempty"`

	// SerialNumber
	// SWitch serial number
	// Constraints:
	//    - nullable
	SerialNumber *string `json:"serialNumber,omitempty"`

	// StackId
	// Stack Id
	// Constraints:
	//    - nullable
	StackId *string `json:"stackId,omitempty"`

	// Status
	// Status of switch, Ex: ONLINE, OFFLINE
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// SwitchName
	// Switch name
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// UpTime
	// SWitch uptime
	// Constraints:
	//    - nullable
	UpTime *string `json:"upTime,omitempty"`
}

func NewSwitchMSwitchNetworkSwitch() *SwitchMSwitchNetworkSwitch {
	m := new(SwitchMSwitchNetworkSwitch)
	return m
}

// SwitchMSwitchNetworkSwitchFirmwareUpdateType
//
// Information of firmware update
// Constraints:
//    - nullable
type SwitchMSwitchNetworkSwitchFirmwareUpdateType struct {
	// ModifiedTime
	// Modified time of the firmware update scheduled
	// Constraints:
	//    - nullable
	ModifiedTime *string `json:"modifiedTime,omitempty"`

	// ScheduledTime
	// Scheduled time of firmware update
	// Constraints:
	//    - nullable
	ScheduledTime *string `json:"scheduledTime,omitempty"`

	// ScheduleId
	// Schedule Id of firmware update
	// Constraints:
	//    - nullable
	ScheduleId *string `json:"scheduleId,omitempty"`

	// Status
	// Status of firmware update
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// ToVersion
	// Update to which firmware version
	// Constraints:
	//    - nullable
	ToVersion *string `json:"toVersion,omitempty"`
}

func NewSwitchMSwitchNetworkSwitchFirmwareUpdateType() *SwitchMSwitchNetworkSwitchFirmwareUpdateType {
	m := new(SwitchMSwitchNetworkSwitchFirmwareUpdateType)
	return m
}

// SwitchMSwitchNetworkSwitchPoeType
//
// Information of PoE
// Constraints:
//    - nullable
type SwitchMSwitchNetworkSwitchPoeType struct {
	// Free
	// Free power capacity of a switch
	// Constraints:
	//    - nullable
	Free *int `json:"free,omitempty"`

	// Percent
	// Percentage of power usage for a switch
	// Constraints:
	//    - nullable
	Percent *float64 `json:"percent,omitempty"`

	// Total
	// Total power capacity of a switch
	// Constraints:
	//    - nullable
	Total *int `json:"total,omitempty"`
}

func NewSwitchMSwitchNetworkSwitchPoeType() *SwitchMSwitchNetworkSwitchPoeType {
	m := new(SwitchMSwitchNetworkSwitchPoeType)
	return m
}

// SwitchMSwitchNetworkSwitchPortStatusType
//
// Information of port status
// Constraints:
//    - nullable
type SwitchMSwitchNetworkSwitchPortStatusType struct {
	// AdminDown
	// Count for port status is admin down of switch
	// Constraints:
	//    - nullable
	AdminDown *int `json:"adminDown,omitempty"`

	// Down
	// Count for port status is down of switch
	// Constraints:
	//    - nullable
	Down *int `json:"down,omitempty"`

	// Speed
	// Port speed of switch
	// Constraints:
	//    - nullable
	Speed *string `json:"speed,omitempty"`

	// SpeedInt
	// Switch port fully speed
	// Constraints:
	//    - nullable
	SpeedInt *int `json:"speedInt,omitempty"`

	// Total
	// Total count for port status of switch
	// Constraints:
	//    - nullable
	Total *int `json:"total,omitempty"`

	// Up
	// Count for port status is up of switch
	// Constraints:
	//    - nullable
	Up *int `json:"up,omitempty"`

	// Warning
	// Count for port status is warring of switch
	// Constraints:
	//    - nullable
	Warning *int `json:"warning,omitempty"`
}

func NewSwitchMSwitchNetworkSwitchPortStatusType() *SwitchMSwitchNetworkSwitchPortStatusType {
	m := new(SwitchMSwitchNetworkSwitchPortStatusType)
	return m
}

type SwitchMSwitchPortDetails struct {
	// AdminStatus
	// Admin status of switch port, UP or DOWN
	// Constraints:
	//    - nullable
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	// Constraints:
	//    - nullable
	ConnectedDevice *SwitchMSwitchPortDetailsConnectedDeviceType `json:"connectedDevice,omitempty"`

	// Id
	// Identifier of switch port
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// InUtilization
	// Switch port traffic in utilization
	// Constraints:
	//    - nullable
	InUtilization *float64 `json:"inUtilization,omitempty"`

	// LagName
	// LAG name of switch port
	// Constraints:
	//    - nullable
	LagName *string `json:"lagName,omitempty"`

	// Mac
	// Mac address of switch port
	// Constraints:
	//    - nullable
	Mac *string `json:"mac,omitempty"`

	// Name
	// Name of switch port
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// NeighborName
	// Switch port connected neighbor name
	// Constraints:
	//    - nullable
	NeighborName *string `json:"neighborName,omitempty"`

	// OpticsType
	// Switch port optics type
	// Constraints:
	//    - nullable
	OpticsType *string `json:"opticsType,omitempty"`

	// OutUtilization
	// Switch port traffic out utilization
	// Constraints:
	//    - nullable
	OutUtilization *float64 `json:"outUtilization,omitempty"`

	// Packets
	// Port packet transmit information
	// Constraints:
	//    - nullable
	Packets *SwitchMSwitchPortDetailsPacketsType `json:"packets,omitempty"`

	// Poe
	// PoE information of switch port
	// Constraints:
	//    - nullable
	Poe *SwitchMSwitchPortDetailsPoeType `json:"poe,omitempty"`

	// PoeEnabled
	// PoE Enabled, True or False
	// Constraints:
	//    - nullable
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PoeType
	// PoE type
	// Constraints:
	//    - nullable
	PoeType *string `json:"poeType,omitempty"`

	// PortError
	// Port error Information
	// Constraints:
	//    - nullable
	PortError *SwitchMSwitchPortDetailsPortErrorType `json:"portError,omitempty"`

	// PortIdentifier
	// Port Identifier of switch port
	// Constraints:
	//    - nullable
	PortIdentifier *string `json:"portIdentifier,omitempty"`

	// PortSpeed
	// Switch port speed
	// Constraints:
	//    - nullable
	PortSpeed *string `json:"portSpeed,omitempty"`

	// SampledInstant
	// Sampled instant of switch port
	// Constraints:
	//    - nullable
	SampledInstant *string `json:"sampledInstant,omitempty"`

	// Status
	// Status of switch port, UP or DOWN
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// StpState
	// Switch port STP state
	// Constraints:
	//    - nullable
	StpState *int `json:"stpState,omitempty"`

	// SwitchGroup
	// Switch group of switch port
	// Constraints:
	//    - nullable
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchName
	// Switch Name of switch port
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// TrafficUsage
	// Traffic usage information
	// Constraints:
	//    - nullable
	TrafficUsage *SwitchMSwitchPortDetailsTrafficUsageType `json:"trafficUsage,omitempty"`

	// Type
	// Type of switch port
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`

	// UnTaggedVlan
	// Untagged vlan of switch port
	// Constraints:
	//    - nullable
	UnTaggedVlan *string `json:"unTaggedVlan,omitempty"`

	// UsedInFormingStack
	// Used in forming stack, True or False
	// Constraints:
	//    - nullable
	UsedInFormingStack *bool `json:"usedInFormingStack,omitempty"`

	// Vlans
	// Switch port include vlans
	// Constraints:
	//    - nullable
	Vlans *string `json:"vlans,omitempty"`
}

func NewSwitchMSwitchPortDetails() *SwitchMSwitchPortDetails {
	m := new(SwitchMSwitchPortDetails)
	return m
}

// SwitchMSwitchPortDetailsConnectedDeviceType
//
// Connected device information
// Constraints:
//    - nullable
type SwitchMSwitchPortDetailsConnectedDeviceType struct {
	// DomainId
	// Identifier of the management domain to which the connected device belong
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of connected device
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IsRuckusAP
	// Connected devices is RuckusAP,True or False
	// Constraints:
	//    - nullable
	IsRuckusAP *string `json:"isRuckusAP,omitempty"`

	// LocalPort
	// Local port description to connected device
	// Constraints:
	//    - nullable
	LocalPort *string `json:"localPort,omitempty"`

	// LocalPortIfaceName
	// Local port interface name
	// Constraints:
	//    - nullable
	LocalPortIfaceName *string `json:"localPortIfaceName,omitempty"`

	// LocalPortMac
	// Local port mac address to connected device
	// Constraints:
	//    - nullable
	LocalPortMac *string `json:"localPortMac,omitempty"`

	// RemoteDeviceMac
	// Constraints:
	//    - nullable
	RemoteDeviceMac *string `json:"remoteDeviceMac,omitempty"`

	// RemoteDeviceName
	// Remote connected device name
	// Constraints:
	//    - nullable
	RemoteDeviceName *string `json:"remoteDeviceName,omitempty"`

	// RemotePort
	// Remote port number of connected device
	// Constraints:
	//    - nullable
	RemotePort *string `json:"remotePort,omitempty"`

	// RemotePortDesc
	// Remote port description of connected device
	// Constraints:
	//    - nullable
	RemotePortDesc *string `json:"remotePortDesc,omitempty"`

	// RemotePortMac
	// Remote port mac address of local device
	// Constraints:
	//    - nullable
	RemotePortMac *string `json:"remotePortMac,omitempty"`

	// RemotePortType
	// Remote port type of connected device
	// Constraints:
	//    - nullable
	RemotePortType *string `json:"remotePortType,omitempty"`

	// SwitchGroup
	// Switch group
	// Constraints:
	//    - nullable
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchGroupLevelOneId
	// Switch group level one Id
	// Constraints:
	//    - nullable
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Switch group level two Id
	// Constraints:
	//    - nullable
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch name
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// UnitId
	// Unit Id
	// Constraints:
	//    - nullable
	UnitId *string `json:"unitId,omitempty"`
}

func NewSwitchMSwitchPortDetailsConnectedDeviceType() *SwitchMSwitchPortDetailsConnectedDeviceType {
	m := new(SwitchMSwitchPortDetailsConnectedDeviceType)
	return m
}

// SwitchMSwitchPortDetailsPacketsType
//
// Port packet transmit information
// Constraints:
//    - nullable
type SwitchMSwitchPortDetailsPacketsType struct {
	// BroadcastIn
	// Switch port broadcast in packet count
	// Constraints:
	//    - nullable
	BroadcastIn *int `json:"broadcastIn,omitempty"`

	// BroadcastOut
	// Switch port broadcast out packet count
	// Constraints:
	//    - nullable
	BroadcastOut *int `json:"broadcastOut,omitempty"`

	// MulticastIn
	// Switch port multicast in packet count
	// Constraints:
	//    - nullable
	MulticastIn *int `json:"multicastIn,omitempty"`

	// MulticastOut
	// Switch port multicast out packet count
	// Constraints:
	//    - nullable
	MulticastOut *int `json:"multicastOut,omitempty"`
}

func NewSwitchMSwitchPortDetailsPacketsType() *SwitchMSwitchPortDetailsPacketsType {
	m := new(SwitchMSwitchPortDetailsPacketsType)
	return m
}

// SwitchMSwitchPortDetailsPoeType
//
// PoE information of switch port
// Constraints:
//    - nullable
type SwitchMSwitchPortDetailsPoeType struct {
	// Free
	// Free power capacity of switch port
	// Constraints:
	//    - nullable
	Free *int `json:"free,omitempty"`

	// Percent
	// Power used percentage of switch port
	// Constraints:
	//    - nullable
	Percent *float64 `json:"percent,omitempty"`

	// Total
	// Total power capacity of switch port
	// Constraints:
	//    - nullable
	Total *int `json:"total,omitempty"`
}

func NewSwitchMSwitchPortDetailsPoeType() *SwitchMSwitchPortDetailsPoeType {
	m := new(SwitchMSwitchPortDetailsPoeType)
	return m
}

// SwitchMSwitchPortDetailsPortErrorType
//
// Port error Information
// Constraints:
//    - nullable
type SwitchMSwitchPortDetailsPortErrorType struct {
	// CrcError
	// Switch port CRC error count
	// Constraints:
	//    - nullable
	CrcError *int `json:"crcError,omitempty"`

	// InDiscard
	// Switch port traffic in discard count
	// Constraints:
	//    - nullable
	InDiscard *int `json:"inDiscard,omitempty"`

	// InError
	// Switch port traffic in error count
	// Constraints:
	//    - nullable
	InError *int `json:"inError,omitempty"`

	// OutError
	// Switch port traffic out error count
	// Constraints:
	//    - nullable
	OutError *int `json:"outError,omitempty"`
}

func NewSwitchMSwitchPortDetailsPortErrorType() *SwitchMSwitchPortDetailsPortErrorType {
	m := new(SwitchMSwitchPortDetailsPortErrorType)
	return m
}

type SwitchMSwitchPortDetailsQueryResultList struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchPortDetailsQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch port detail returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch port detail after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchPortDetails `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Switch port detail list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch port detail list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchPortDetailsQueryResultList() *SwitchMSwitchPortDetailsQueryResultList {
	m := new(SwitchMSwitchPortDetailsQueryResultList)
	return m
}

// SwitchMSwitchPortDetailsQueryResultListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
type SwitchMSwitchPortDetailsTrafficUsageType struct {
	// Rx
	// Rx traffic usage of switch port
	// Constraints:
	//    - nullable
	Rx *int `json:"rx,omitempty"`

	// Tx
	// Tx traffic usage of switch port
	// Constraints:
	//    - nullable
	Tx *int `json:"tx,omitempty"`
}

func NewSwitchMSwitchPortDetailsTrafficUsageType() *SwitchMSwitchPortDetailsTrafficUsageType {
	m := new(SwitchMSwitchPortDetailsTrafficUsageType)
	return m
}

// SwitchMSwitchPortStatus
//
// $
// Constraints:
//    - nullable
type SwitchMSwitchPortStatus struct {
	// AdminDown
	// Admin down port count
	// Constraints:
	//    - nullable
	AdminDown *int `json:"adminDown,omitempty"`

	// Down
	// Down port count
	// Constraints:
	//    - nullable
	Down *int `json:"down,omitempty"`

	// Speed
	// Switch port speed
	// Constraints:
	//    - nullable
	Speed *string `json:"speed,omitempty"`

	// SpeedInt
	// Switch port fully speed
	// Constraints:
	//    - nullable
	SpeedInt *int `json:"speedInt,omitempty"`

	// Total
	// Total port count
	// Constraints:
	//    - nullable
	Total *int `json:"total,omitempty"`

	// Up
	// Up port count
	// Constraints:
	//    - nullable
	Up *int `json:"up,omitempty"`

	// Warning
	// Warring port count
	// Constraints:
	//    - nullable
	Warning *int `json:"warning,omitempty"`
}

func NewSwitchMSwitchPortStatus() *SwitchMSwitchPortStatus {
	m := new(SwitchMSwitchPortStatus)
	return m
}

type SwitchMSwitchStackMember struct {
	// ActiveMode
	// Role of stack
	// Constraints:
	//    - nullable
	ActiveMode *string `json:"activeMode,omitempty"`

	// Model
	// Switch model of stack
	// Constraints:
	//    - nullable
	Model *string `json:"model,omitempty"`

	// Ports
	// Port count of stack
	// Constraints:
	//    - nullable
	Ports *int `json:"ports,omitempty"`

	// PortStatus
	// Constraints:
	//    - nullable
	PortStatus *SwitchMSwitchPortStatus `json:"portStatus,omitempty"`

	// SerialNumber
	// Serial number of stack
	// Constraints:
	//    - nullable
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SwitchModule
	// Switch module of stack
	// Constraints:
	//    - nullable
	SwitchModule *string `json:"switchModule,omitempty"`

	// SwitchName
	// Switch name of stack
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// SwitchPorts
	// Constraints:
	//    - nullable
	SwitchPorts *SwitchMSwitchPortDetails `json:"switchPorts,omitempty"`

	// SwitchUnit
	// Switch unit of stack
	// Constraints:
	//    - nullable
	SwitchUnit *string `json:"switchUnit,omitempty"`
}

func NewSwitchMSwitchStackMember() *SwitchMSwitchStackMember {
	m := new(SwitchMSwitchStackMember)
	return m
}

type SwitchMSwitchStackMemberQueryResult struct {
	// Extra
	// Extra information for stack member list
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchStackMemberQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first stack member returned out of the complete stack member list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more stack member after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchStackMember `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total stack member count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current stack member count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchStackMemberQueryResult() *SwitchMSwitchStackMemberQueryResult {
	m := new(SwitchMSwitchStackMemberQueryResult)
	return m
}

// SwitchMSwitchStackMemberQueryResultExtraType
//
// Extra information for stack member list
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
type SwitchMSwitchIdList []string

func MakeSwitchMSwitchIdList() SwitchMSwitchIdList {
	m := make(SwitchMSwitchIdList, 0)
	return m
}

type SwitchMSwitchPortsSummaryQueryResultList struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchPortsSummaryQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first switch ports summary returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more switch ports summary after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchPortStatus `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Switch ports summary list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch ports summary list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchPortsSummaryQueryResultList() *SwitchMSwitchPortsSummaryQueryResultList {
	m := new(SwitchMSwitchPortsSummaryQueryResultList)
	return m
}

// SwitchMSwitchPortsSummaryQueryResultListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Extra
	// Constraints:
	//    - nullable
	Extra *SwitchMCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more  after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchNetworkSwitch `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Switch query result list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total switch query result list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchQueryResultList() *SwitchMSwitchQueryResultList {
	m := new(SwitchMSwitchQueryResultList)
	return m
}

type SwitchMSwitchTopSwitchesByFirmwareQueryResultList struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top switches by firmware returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more top switches by firmware after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchBarChart `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Top switches by firmware list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total top switches by firmware list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTopSwitchesByFirmwareQueryResultList() *SwitchMSwitchTopSwitchesByFirmwareQueryResultList {
	m := new(SwitchMSwitchTopSwitchesByFirmwareQueryResultList)
	return m
}

// SwitchMSwitchTopSwitchesByFirmwareQueryResultListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	Extra *SwitchMSwitchTopSwitchesByModelQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top switches by model returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are top switches by model after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchBarChart `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Top switches by model list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total top switches by model list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchTopSwitchesByModelQueryResultList() *SwitchMSwitchTopSwitchesByModelQueryResultList {
	m := new(SwitchMSwitchTopSwitchesByModelQueryResultList)
	return m
}

// SwitchMSwitchTopSwitchesByModelQueryResultListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
