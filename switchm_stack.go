package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchStackConfigService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchStackConfigService(c *VSZClient) *SwitchMSwitchStackConfigService {
	s := new(SwitchMSwitchStackConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchStackConfigService() *SwitchMSwitchStackConfigService {
	return NewSwitchMSwitchStackConfigService(ss.apiClient)
}

type SwitchMSwitchStackConfigAuditIdList struct {
	Extra *SwitchMSwitchStackConfigAuditIdListExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchAuditId `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchStackConfigAuditIdList() *SwitchMSwitchStackConfigAuditIdList {
	m := new(SwitchMSwitchStackConfigAuditIdList)
	return m
}

type SwitchMSwitchStackConfigAuditIdListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchStackConfigAuditIdListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchStackConfigAuditIdListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchStackConfigAuditIdListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchStackConfigAuditIdListExtraType() *SwitchMSwitchStackConfigAuditIdListExtraType {
	m := new(SwitchMSwitchStackConfigAuditIdListExtraType)
	return m
}

type SwitchMSwitchStackConfigList struct {
	// FirstIndex
	// Index of the first stack returned out of the complete stack list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more stack after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// List of stack
	List []*SwitchMSwitchStackConfigMember `json:"list,omitempty"`

	// RawDataTotalCount
	// Stack count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Stack count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchStackConfigList() *SwitchMSwitchStackConfigList {
	m := new(SwitchMSwitchStackConfigList)
	return m
}

type SwitchMSwitchStackConfigMember struct {
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
	PortStatus *SwitchMSwitchStackConfigMemberPortStatusType `json:"portStatus,omitempty"`

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
	SwitchPorts []*SwitchMSwitchStackConfigMemberSwitchPortsType `json:"switchPorts,omitempty"`

	// SwitchUnit
	// Switch unit of stack
	SwitchUnit *string `json:"switchUnit,omitempty"`

	// SwitchUnitState
	// Switch unit state of stack
	SwitchUnitState *string `json:"switchUnitState,omitempty"`
}

func NewSwitchMSwitchStackConfigMember() *SwitchMSwitchStackConfigMember {
	m := new(SwitchMSwitchStackConfigMember)
	return m
}

// SwitchMSwitchStackConfigMemberPortStatusType
//
// Port status Information
type SwitchMSwitchStackConfigMemberPortStatusType struct {
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

func NewSwitchMSwitchStackConfigMemberPortStatusType() *SwitchMSwitchStackConfigMemberPortStatusType {
	m := new(SwitchMSwitchStackConfigMemberPortStatusType)
	return m
}

type SwitchMSwitchStackConfigMemberSwitchPortsType struct {
	// AdminStatus
	// Admin Status of switch port
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	ConnectedDevice *SwitchMSwitchStackConfigMemberSwitchPortsTypeConnectedDeviceType `json:"connectedDevice,omitempty"`

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
	// PoE information of switch port
	Poe *SwitchMSwitchStackConfigMemberSwitchPortsTypePoeType `json:"poe,omitempty"`

	// PoeEnabled
	// PoE Enabled, True or False
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
	TrafficUsage *SwitchMSwitchStackConfigMemberSwitchPortsTypeTrafficUsageType `json:"trafficUsage,omitempty"`

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

func NewSwitchMSwitchStackConfigMemberSwitchPortsType() *SwitchMSwitchStackConfigMemberSwitchPortsType {
	m := new(SwitchMSwitchStackConfigMemberSwitchPortsType)
	return m
}

// SwitchMSwitchStackConfigMemberSwitchPortsTypeConnectedDeviceType
//
// Connected device information
type SwitchMSwitchStackConfigMemberSwitchPortsTypeConnectedDeviceType struct {
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

func NewSwitchMSwitchStackConfigMemberSwitchPortsTypeConnectedDeviceType() *SwitchMSwitchStackConfigMemberSwitchPortsTypeConnectedDeviceType {
	m := new(SwitchMSwitchStackConfigMemberSwitchPortsTypeConnectedDeviceType)
	return m
}

// SwitchMSwitchStackConfigMemberSwitchPortsTypePoeType
//
// PoE information of switch port
type SwitchMSwitchStackConfigMemberSwitchPortsTypePoeType struct {
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

func NewSwitchMSwitchStackConfigMemberSwitchPortsTypePoeType() *SwitchMSwitchStackConfigMemberSwitchPortsTypePoeType {
	m := new(SwitchMSwitchStackConfigMemberSwitchPortsTypePoeType)
	return m
}

// SwitchMSwitchStackConfigMemberSwitchPortsTypeTrafficUsageType
//
// Traffic usage information
type SwitchMSwitchStackConfigMemberSwitchPortsTypeTrafficUsageType struct {
	// Rx
	// Rx traffic usage of switch port
	Rx *int `json:"rx,omitempty"`

	// Tx
	// Tx traffic usage of switch port
	Tx *int `json:"tx,omitempty"`
}

func NewSwitchMSwitchStackConfigMemberSwitchPortsTypeTrafficUsageType() *SwitchMSwitchStackConfigMemberSwitchPortsTypeTrafficUsageType {
	m := new(SwitchMSwitchStackConfigMemberSwitchPortsTypeTrafficUsageType)
	return m
}

type SwitchMSwitchStackConfigStackConfig struct {
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

func NewSwitchMSwitchStackConfigStackConfig() *SwitchMSwitchStackConfigStackConfig {
	m := new(SwitchMSwitchStackConfigStackConfig)
	return m
}

type SwitchMSwitchStackConfigStackConfigList []*SwitchMSwitchStackConfigStackConfig

func MakeSwitchMSwitchStackConfigStackConfigList() SwitchMSwitchStackConfigStackConfigList {
	m := make(SwitchMSwitchStackConfigStackConfigList, 0)
	return m
}

// AddStack
//
// Use this API command to create a stack configuration.
//
// Request Body:
//	 - body SwitchMSwitchStackConfigStackConfigList
func (s *SwitchMSwitchStackConfigService) AddStack(ctx context.Context, body SwitchMSwitchStackConfigStackConfigList, mutators ...RequestMutator) (*SwitchMSwitchStackConfigAuditIdList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchStackConfigAuditIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddStack, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchStackConfigAuditIdList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindStackBySwitchId
//
// Use this API command to retrieve a stack configuration configured via SZ.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMSwitchStackConfigService) FindStackBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMSwitchStackConfigStackConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchStackConfigStackConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindStackBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchStackConfigStackConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindStackMemberBySerialNumber
//
// Use this API command to retrieve the member of switches in a stack.
//
// Required Parameters:
// - serialNumber string
//		- required
func (s *SwitchMSwitchStackConfigService) FindStackMemberBySerialNumber(ctx context.Context, serialNumber string, mutators ...RequestMutator) (*SwitchMSwitchStackConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchStackConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindStackMemberBySerialNumber, true)
	req.SetPathParameter("serialNumber", serialNumber)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchStackConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
