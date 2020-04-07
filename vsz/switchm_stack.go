package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMStackService struct {
	apiClient *APIClient
}

func NewSwitchMStackService(c *APIClient) *SwitchMStackService {
	s := new(SwitchMStackService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMStackService() *SwitchMStackService {
	return NewSwitchMStackService(ss.apiClient)
}

type SwitchMStackAuditIdList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *SwitchMStackAuditIdListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMSwitchAuditId `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMStackAuditIdList() *SwitchMStackAuditIdList {
	m := new(SwitchMStackAuditIdList)
	return m
}

type SwitchMStackAuditIdListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMStackAuditIdListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMStackAuditIdListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMStackAuditIdListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMStackAuditIdListExtraType() *SwitchMStackAuditIdListExtraType {
	m := new(SwitchMStackAuditIdListExtraType)
	return m
}

type SwitchMStackList struct {
	// FirstIndex
	// Index of the first stack returned out of the complete stack list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more stack after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// List of stack
	// Constraints:
	//    - nullable
	List []*SwitchMStackMember `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Stack count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Stack count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMStackList() *SwitchMStackList {
	m := new(SwitchMStackList)
	return m
}

type SwitchMStackMember struct {
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
	// Port count  of stack
	// Constraints:
	//    - nullable
	Ports *int `json:"ports,omitempty"`

	// PortStatus
	// Port status Information
	// Constraints:
	//    - nullable
	PortStatus *SwitchMStackMemberPortStatusType `json:"portStatus,omitempty"`

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
	// Switch port information of stack
	// Constraints:
	//    - nullable
	SwitchPorts []*SwitchMStackMemberSwitchPortsType `json:"switchPorts,omitempty" validate:"omitempty,dive"`

	// SwitchUnit
	// Switch unit of stack
	// Constraints:
	//    - nullable
	SwitchUnit *string `json:"switchUnit,omitempty"`

	// SwitchUnitState
	// Switch unit state of stack
	// Constraints:
	//    - nullable
	SwitchUnitState *string `json:"switchUnitState,omitempty"`
}

func NewSwitchMStackMember() *SwitchMStackMember {
	m := new(SwitchMStackMember)
	return m
}

// SwitchMStackMemberPortStatusType
//
// Port status Information
// Constraints:
//    - nullable
type SwitchMStackMemberPortStatusType struct {
	// AdminDown
	// Count for port status is admin down of stack
	// Constraints:
	//    - nullable
	AdminDown *int `json:"adminDown,omitempty"`

	// Down
	// Count for port status is down of stack
	// Constraints:
	//    - nullable
	Down *int `json:"down,omitempty"`

	// Speed
	// Port speed of stack
	// Constraints:
	//    - nullable
	Speed *string `json:"speed,omitempty"`

	// Total
	// Total port count of stack
	// Constraints:
	//    - nullable
	Total *int `json:"total,omitempty"`

	// Up
	// Count for port status is up of stack
	// Constraints:
	//    - nullable
	Up *int `json:"up,omitempty"`

	// Warning
	// Count for port status is warring of stack
	// Constraints:
	//    - nullable
	Warning *int `json:"warning,omitempty"`
}

func NewSwitchMStackMemberPortStatusType() *SwitchMStackMemberPortStatusType {
	m := new(SwitchMStackMemberPortStatusType)
	return m
}

type SwitchMStackMemberSwitchPortsType struct {
	// AdminStatus
	// Admin Status of switch port
	// Constraints:
	//    - nullable
	AdminStatus *string `json:"adminStatus,omitempty"`

	// ConnectedDevice
	// Connected device information
	// Constraints:
	//    - nullable
	ConnectedDevice *SwitchMStackMemberSwitchPortsTypeConnectedDeviceType `json:"connectedDevice,omitempty"`

	// Id
	// Identifier of switch port
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// InUtilization
	// In utilization of switch port
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
	// Neighbor name of switch port
	// Constraints:
	//    - nullable
	NeighborName *string `json:"neighborName,omitempty"`

	// OpticsType
	// Optics type of switch port
	// Constraints:
	//    - nullable
	OpticsType *string `json:"opticsType,omitempty"`

	// OutUtilization
	// Out utilization of switch port
	// Constraints:
	//    - nullable
	OutUtilization *float64 `json:"outUtilization,omitempty"`

	// Poe
	// PoE information of switch port
	// Constraints:
	//    - nullable
	Poe *SwitchMStackMemberSwitchPortsTypePoeType `json:"poe,omitempty"`

	// PoeEnabled
	// PoE Enabled, True or False
	// Constraints:
	//    - nullable
	PoeEnabled *bool `json:"poeEnabled,omitempty"`

	// PortIdentifier
	// Port Identifier of switch port
	// Constraints:
	//    - nullable
	PortIdentifier *string `json:"portIdentifier,omitempty"`

	// PortSpeed
	// Port speed of switch port
	// Constraints:
	//    - nullable
	PortSpeed *string `json:"portSpeed,omitempty"`

	// SampledInstant
	// Sampled instant of switch port
	// Constraints:
	//    - nullable
	SampledInstant *string `json:"sampledInstant,omitempty"`

	// Status
	// Status of switch port
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// StpState
	// STP state of switch port
	// Constraints:
	//    - nullable
	StpState *int `json:"stpState,omitempty"`

	// SwitchGroup
	// Switch group of switch port
	// Constraints:
	//    - nullable
	SwitchGroup *string `json:"switchGroup,omitempty"`

	// SwitchName
	// Switch name of stack
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// TrafficUsage
	// Traffic usage information
	// Constraints:
	//    - nullable
	TrafficUsage *SwitchMStackMemberSwitchPortsTypeTrafficUsageType `json:"trafficUsage,omitempty"`

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

func NewSwitchMStackMemberSwitchPortsType() *SwitchMStackMemberSwitchPortsType {
	m := new(SwitchMStackMemberSwitchPortsType)
	return m
}

// SwitchMStackMemberSwitchPortsTypeConnectedDeviceType
//
// Connected device information
// Constraints:
//    - nullable
type SwitchMStackMemberSwitchPortsTypeConnectedDeviceType struct {
	// DomainId
	// Identifier of the management domain to which the connected device belong
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of switch port connected device
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IsRuckusAP
	// Connected devices is RuckusAP,True or False
	// Constraints:
	//    - nullable
	IsRuckusAP *string `json:"isRuckusAP,omitempty"`

	// LocalPort
	// Local port description of connected device
	// Constraints:
	//    - nullable
	LocalPort *string `json:"localPort,omitempty"`

	// LocalPortIfaceName
	// Local port interface name
	// Constraints:
	//    - nullable
	LocalPortIfaceName *string `json:"localPortIfaceName,omitempty"`

	// LocalPortMac
	// Local port mac address of connected device
	// Constraints:
	//    - nullable
	LocalPortMac *string `json:"localPortMac,omitempty"`

	// RemoteDeviceName
	// Remote device name of connected device
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
	// Remote port mac address of connected device
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
	// Tenant Id of stack
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// UnitId
	// Unit Id
	// Constraints:
	//    - nullable
	UnitId *string `json:"unitId,omitempty"`
}

func NewSwitchMStackMemberSwitchPortsTypeConnectedDeviceType() *SwitchMStackMemberSwitchPortsTypeConnectedDeviceType {
	m := new(SwitchMStackMemberSwitchPortsTypeConnectedDeviceType)
	return m
}

// SwitchMStackMemberSwitchPortsTypePoeType
//
// PoE information of switch port
// Constraints:
//    - nullable
type SwitchMStackMemberSwitchPortsTypePoeType struct {
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

func NewSwitchMStackMemberSwitchPortsTypePoeType() *SwitchMStackMemberSwitchPortsTypePoeType {
	m := new(SwitchMStackMemberSwitchPortsTypePoeType)
	return m
}

// SwitchMStackMemberSwitchPortsTypeTrafficUsageType
//
// Traffic usage information
// Constraints:
//    - nullable
type SwitchMStackMemberSwitchPortsTypeTrafficUsageType struct {
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

func NewSwitchMStackMemberSwitchPortsTypeTrafficUsageType() *SwitchMStackMemberSwitchPortsTypeTrafficUsageType {
	m := new(SwitchMStackMemberSwitchPortsTypeTrafficUsageType)
	return m
}

type SwitchMStackConfig struct {
	// ActiveSwitchId
	// Switch Id of Active Unit
	// Constraints:
	//    - nullable
	ActiveSwitchId *string `json:"activeSwitchId,omitempty"`

	// IsActiveRole
	// Switch role is Active, True (Active) or False (Standby or Member)
	// Constraints:
	//    - nullable
	IsActiveRole *bool `json:"isActiveRole,omitempty"`

	// SuggestedId
	// Suggested switch unit Id in stack, 1 ~ 12
	// Constraints:
	//    - nullable
	SuggestedId *int `json:"suggestedId,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMStackConfig() *SwitchMStackConfig {
	m := new(SwitchMStackConfig)
	return m
}

type SwitchMStackConfigList []*SwitchMStackConfig

func MakeSwitchMStackConfigList() SwitchMStackConfigList {
	m := make(SwitchMStackConfigList, 0)
	return m
}

// AddStack
//
// Use this API command to create a stack configuration.
//
// Request Body:
//	 - body SwitchMStackConfigList
func (s *SwitchMStackService) AddStack(ctx context.Context, body SwitchMStackConfigList) (*SwitchMStackAuditIdList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMStackAuditIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddStack, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMStackAuditIdList()
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
func (s *SwitchMStackService) FindStackBySwitchId(ctx context.Context, switchId string) (*SwitchMStackConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMStackConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, switchId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindStackBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMStackConfig()
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
func (s *SwitchMStackService) FindStackMemberBySerialNumber(ctx context.Context, serialNumber string) (*SwitchMStackList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMStackList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, serialNumber, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindStackMemberBySerialNumber, true)
	req.SetPathParameter("serialNumber", serialNumber)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMStackList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
