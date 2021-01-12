package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type SwitchMSwitchService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchService(c *VSZClient) *SwitchMSwitchService {
	s := new(SwitchMSwitchService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchService() *SwitchMSwitchService {
	return NewSwitchMSwitchService(ss.apiClient)
}

// SwitchMSwitchAuditId
//
// Definition: switch_auditId
type SwitchMSwitchAuditId struct {
	// Id
	// Audit Id
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

type SwitchMSwitchAuditIdAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchAuditId
}

func newSwitchMSwitchAuditIdAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchAuditIdAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchAuditIdAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchAuditId)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchAuditId() *SwitchMSwitchAuditId {
	m := new(SwitchMSwitchAuditId)
	return m
}

// SwitchMSwitchBarChart
//
// Definition: switch_barChart
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

// SwitchMSwitchConnectedAPsQueryList
//
// Definition: switch_connectedAPsQueryList
type SwitchMSwitchConnectedAPsQueryList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchConnectedAPsQueryListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchConnectedAPsQueryList
}

func newSwitchMSwitchConnectedAPsQueryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchConnectedAPsQueryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchConnectedAPsQueryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchConnectedAPsQueryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchConnectedAPsQueryList() *SwitchMSwitchConnectedAPsQueryList {
	m := new(SwitchMSwitchConnectedAPsQueryList)
	return m
}

// SwitchMSwitchConnectedDevice
//
// Definition: switch_connectedDevice
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
	SampledInstant interface{} `json:"sampledInstant,omitempty"`

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

// SwitchMSwitchConnectedDevicesQueryList
//
// Definition: switch_connectedDevicesQueryList
type SwitchMSwitchConnectedDevicesQueryList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchConnectedDevicesQueryListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchConnectedDevicesQueryList
}

func newSwitchMSwitchConnectedDevicesQueryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchConnectedDevicesQueryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchConnectedDevicesQueryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchConnectedDevicesQueryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchConnectedDevicesQueryList() *SwitchMSwitchConnectedDevicesQueryList {
	m := new(SwitchMSwitchConnectedDevicesQueryList)
	return m
}

// SwitchMSwitchDeleteSwitchesResultList
//
// Definition: switch_deleteSwitchesResultList
type SwitchMSwitchDeleteSwitchesResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchDeleteSwitchesResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchDeleteSwitchesResultList
}

func newSwitchMSwitchDeleteSwitchesResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchDeleteSwitchesResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchDeleteSwitchesResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchDeleteSwitchesResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchDeleteSwitchesResultList() *SwitchMSwitchDeleteSwitchesResultList {
	m := new(SwitchMSwitchDeleteSwitchesResultList)
	return m
}

// SwitchMSwitchFirmware
//
// Definition: switch_firmware
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

// SwitchMSwitchFirmwareHistoryQueryResultList
//
// Definition: switch_firmwareHistoryQueryResultList
type SwitchMSwitchFirmwareHistoryQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchFirmwareHistoryQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchFirmwareHistoryQueryResultList
}

func newSwitchMSwitchFirmwareHistoryQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchFirmwareHistoryQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchFirmwareHistoryQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchFirmwareHistoryQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchFirmwareHistoryQueryResultList() *SwitchMSwitchFirmwareHistoryQueryResultList {
	m := new(SwitchMSwitchFirmwareHistoryQueryResultList)
	return m
}

// SwitchMSwitchNetworkSwitch
//
// Definition: switch_networkSwitch
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
	LastBackupStatus *string `json:"lastBackupStatus,omitempty"`

	// LastBackupTime
	// Last config backup time of switch
	LastBackupTime *int `json:"lastBackupTime,omitempty"`

	// LastRestoreStatus
	// Last config restore status of switch
	// Constraints:
	//    - oneof:[STARTED,SUCCESS,FAILED]
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

	SupportedCsl *int `json:"supportedCsl,omitempty"`

	// SwitchName
	// Switch name
	SwitchName *string `json:"switchName,omitempty"`

	// UpTime
	// SWitch uptime
	UpTime *string `json:"upTime,omitempty"`
}

type SwitchMSwitchNetworkSwitchAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchNetworkSwitch
}

func newSwitchMSwitchNetworkSwitchAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchNetworkSwitchAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchNetworkSwitchAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchNetworkSwitch)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchNetworkSwitch() *SwitchMSwitchNetworkSwitch {
	m := new(SwitchMSwitchNetworkSwitch)
	return m
}

// SwitchMSwitchNetworkSwitchFirmwareUpdateType
//
// Definition: switch_networkSwitchFirmwareUpdateType
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
// Definition: switch_networkSwitchPoeType
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
// Definition: switch_networkSwitchPortStatusType
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

// SwitchMSwitchPortDetails
//
// Definition: switch_portDetails
type SwitchMSwitchPortDetails struct {
	// AdminStatus
	// Admin status of switch port, UP or DOWN
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	ConnectedDevice []*SwitchMSwitchConnectedDevice `json:"connectedDevice,omitempty"`

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

	SwitchFamily *string `json:"switchFamily,omitempty"`

	// SwitchGroup
	// Switch group of switch port
	SwitchGroup *string `json:"switchGroup,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchModel *string `json:"switchModel,omitempty"`

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

// SwitchMSwitchPortDetailsPacketsType
//
// Definition: switch_portDetailsPacketsType
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
// Definition: switch_portDetailsPoeType
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
// Definition: switch_portDetailsPortErrorType
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

// SwitchMSwitchPortDetailsQueryResultList
//
// Definition: switch_portDetailsQueryResultList
type SwitchMSwitchPortDetailsQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchPortDetailsQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchPortDetailsQueryResultList
}

func newSwitchMSwitchPortDetailsQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchPortDetailsQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchPortDetailsQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchPortDetailsQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchPortDetailsQueryResultList() *SwitchMSwitchPortDetailsQueryResultList {
	m := new(SwitchMSwitchPortDetailsQueryResultList)
	return m
}

// SwitchMSwitchPortDetailsTrafficUsageType
//
// Definition: switch_portDetailsTrafficUsageType
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
// Definition: switch_portStatus
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

// SwitchMSwitchStackMember
//
// Definition: switch_stackMember
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

	SwitchPorts []*SwitchMSwitchPortDetails `json:"switchPorts,omitempty"`

	// SwitchUnit
	// Switch unit of stack
	SwitchUnit *string `json:"switchUnit,omitempty"`
}

func NewSwitchMSwitchStackMember() *SwitchMSwitchStackMember {
	m := new(SwitchMSwitchStackMember)
	return m
}

// SwitchMSwitchStackMemberQueryResult
//
// Definition: switch_stackMemberQueryResult
type SwitchMSwitchStackMemberQueryResult struct {
	// Extra
	// Extra information for stack member list
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchStackMemberQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchStackMemberQueryResult
}

func newSwitchMSwitchStackMemberQueryResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchStackMemberQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchStackMemberQueryResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchStackMemberQueryResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchStackMemberQueryResult() *SwitchMSwitchStackMemberQueryResult {
	m := new(SwitchMSwitchStackMemberQueryResult)
	return m
}

// SwitchMSwitchIdList
//
// Definition: switch_switchIdList
//
// $
type SwitchMSwitchIdList []string

func MakeSwitchMSwitchIdList() SwitchMSwitchIdList {
	m := make(SwitchMSwitchIdList, 0)
	return m
}

// SwitchMSwitchPortsSummaryQueryResultList
//
// Definition: switch_switchPortsSummaryQueryResultList
type SwitchMSwitchPortsSummaryQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchPortsSummaryQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchPortsSummaryQueryResultList
}

func newSwitchMSwitchPortsSummaryQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchPortsSummaryQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchPortsSummaryQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchPortsSummaryQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchPortsSummaryQueryResultList() *SwitchMSwitchPortsSummaryQueryResultList {
	m := new(SwitchMSwitchPortsSummaryQueryResultList)
	return m
}

// SwitchMSwitchQueryResultList
//
// Definition: switch_switchQueryResultList
type SwitchMSwitchQueryResultList struct {
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchQueryResultList
}

func newSwitchMSwitchQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchQueryResultList() *SwitchMSwitchQueryResultList {
	m := new(SwitchMSwitchQueryResultList)
	return m
}

// SwitchMSwitchRebootResponse
//
// Definition: switch_switchRebootResponse
type SwitchMSwitchRebootResponse struct {
	Id *string `json:"id,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRebootResponse) UnmarshalJSON(b []byte) error {
	type _SwitchMSwitchRebootResponse SwitchMSwitchRebootResponse
	tmpType := new(_SwitchMSwitchRebootResponse)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "id")
	*t = SwitchMSwitchRebootResponse(*tmpType)
	return nil
}

func (t *SwitchMSwitchRebootResponse) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	return json.Marshal(tmp)
}

type SwitchMSwitchRebootResponseAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchRebootResponse
}

func newSwitchMSwitchRebootResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchRebootResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchRebootResponseAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchRebootResponse)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchRebootResponse() *SwitchMSwitchRebootResponse {
	m := new(SwitchMSwitchRebootResponse)
	return m
}

// SwitchMSwitchTopSwitchesByFirmwareQueryResultList
//
// Definition: switch_topSwitchesByFirmwareQueryResultList
type SwitchMSwitchTopSwitchesByFirmwareQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchTopSwitchesByFirmwareQueryResultList
}

func newSwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchTopSwitchesByFirmwareQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchTopSwitchesByFirmwareQueryResultList() *SwitchMSwitchTopSwitchesByFirmwareQueryResultList {
	m := new(SwitchMSwitchTopSwitchesByFirmwareQueryResultList)
	return m
}

// SwitchMSwitchTopSwitchesByModelQueryResultList
//
// Definition: switch_topSwitchesByModelQueryResultList
type SwitchMSwitchTopSwitchesByModelQueryResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchTopSwitchesByModelQueryResultList
}

func newSwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchTopSwitchesByModelQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSwitchTopSwitchesByModelQueryResultList() *SwitchMSwitchTopSwitchesByModelQueryResultList {
	m := new(SwitchMSwitchTopSwitchesByModelQueryResultList)
	return m
}

// AddSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone.
//
// Operation ID: addSwitch
// Operation path: /switch
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitch(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchQueryResultListAPIResponse), err
}

// AddSwitchReboot
//
// Reboot switch by MAC address
//
// Operation ID: addSwitchReboot
// Operation path: /switch/{switchId}/reboot
// Success code: 200 (OK)
//
// Required parameters:
// - switchId string
//		- required
func (s *SwitchMSwitchService) AddSwitchReboot(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMSwitchRebootResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchRebootResponseAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMAddSwitchReboot, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchRebootResponseAPIResponse), err
}

// AddSwitchSnmpSyncedSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone and SNMP synced.
//
// Operation ID: addSwitchSnmpSyncedSwitch
// Operation path: /switch/snmpSyncedSwitch
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchSnmpSyncedSwitch(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchSnmpSyncedSwitch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchQueryResultListAPIResponse), err
}

// AddSwitchViewDetails
//
// Use this API command to retrieve switch and port details for the selected Switch/SwitchGroup/Domain.
//
// Operation ID: addSwitchViewDetails
// Operation path: /switch/view/details
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchViewDetails(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchStackMemberQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchStackMemberQueryResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchViewDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchStackMemberQueryResultAPIResponse), err
}

// DeleteSwitch
//
// Use this API command to delete multiple switches managed by SmartZone
//
// Operation ID: deleteSwitch
// Operation path: /switch
// Success code: 200 (OK)
//
// Request body:
//	 - body SwitchMSwitchIdList
func (s *SwitchMSwitchService) DeleteSwitch(ctx context.Context, body SwitchMSwitchIdList, mutators ...RequestMutator) (*SwitchMSwitchDeleteSwitchesResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchDeleteSwitchesResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteSwitch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchDeleteSwitchesResultListAPIResponse), err
}

// DeleteSwitchById
//
// Use this API command to delete a switch managed by SmartZone.
//
// Operation ID: deleteSwitchById
// Operation path: /switch/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMSwitchService) DeleteSwitchById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMSwitchAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteSwitchById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchAuditIdAPIResponse), err
}

// FindSwitchById
//
// Use this API command to retrieve a switch status.
//
// Operation ID: findSwitchById
// Operation path: /switch/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMSwitchService) FindSwitchById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMSwitchNetworkSwitchAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchNetworkSwitchAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSwitchById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchNetworkSwitchAPIResponse), err
}

// FindSwitchFirmwareBySwitchId
//
// Use this API command to get a list of firmware update history.
//
// Operation ID: findSwitchFirmwareBySwitchId
// Operation path: /switch/{switchId}/firmware
// Success code: 200 (OK)
//
// Required parameters:
// - switchId string
//		- required
func (s *SwitchMSwitchService) FindSwitchFirmwareBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMSwitchFirmwareHistoryQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchFirmwareHistoryQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSwitchFirmwareBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchFirmwareHistoryQueryResultListAPIResponse), err
}

// FindSwitchModelList
//
// Use this API command to Retrieve Switch Model List.
//
// Operation ID: findSwitchModelList
// Operation path: /switchModel/list
// Success code: 200 (OK)
func (s *SwitchMSwitchService) FindSwitchModelList(ctx context.Context, mutators ...RequestMutator) (*SwitchMSwitchModelResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchModelResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSwitchModelList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchModelResultAPIResponse), err
}

// UpdateSwitchMoveByDestinationSwitchGroupId
//
// Use this API command to move a list of switches to a switch group.
//
// Operation ID: updateSwitchMoveByDestinationSwitchGroupId
// Operation path: /switch/move/{destinationSwitchGroupId}
// Success code: 200 (OK)
//
// Request body:
//	 - body SwitchMSwitchIdList
//
// Required parameters:
// - destinationSwitchGroupId string
//		- required
func (s *SwitchMSwitchService) UpdateSwitchMoveByDestinationSwitchGroupId(ctx context.Context, body SwitchMSwitchIdList, destinationSwitchGroupId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateSwitchMoveByDestinationSwitchGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("destinationSwitchGroupId", destinationSwitchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
