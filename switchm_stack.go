package bigdog

// API Version: v9_1

import (
	"context"
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

// SwitchMSwitchStackConfigAuditIdList
//
// Definition: stack_auditIdList
type SwitchMSwitchStackConfigAuditIdList struct {
	Extra interface{} `json:"extra,omitempty"`

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

// SwitchMSwitchStackConfigList
//
// Definition: stack_list
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

// SwitchMSwitchStackConfigMember
//
// Definition: stack_member
type SwitchMSwitchStackConfigMember struct {
	// ActiveMode
	// Role of stack
	ActiveMode *string `json:"activeMode,omitempty"`

	// Model
	// Switch model of stack
	Model *string `json:"model,omitempty"`

	// Poe
	// Information of PoE
	Poe *SwitchMSwitchStackConfigMemberPoeType `json:"poe,omitempty"`

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
	SwitchPorts []*SwitchMSwitchPortDetails `json:"switchPorts,omitempty"`

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

// SwitchMSwitchStackConfigMemberPoeType
//
// Definition: stack_memberPoeType
//
// Information of PoE
type SwitchMSwitchStackConfigMemberPoeType struct {
	// Free
	// Free power capacity of a switch unit in stack
	Free *int `json:"free,omitempty"`

	// Percent
	// Percentage of power usage for a switch unit in stack
	Percent *float64 `json:"percent,omitempty"`

	// Total
	// Total power capacity of a switch unit in stack
	Total *int `json:"total,omitempty"`
}

func NewSwitchMSwitchStackConfigMemberPoeType() *SwitchMSwitchStackConfigMemberPoeType {
	m := new(SwitchMSwitchStackConfigMemberPoeType)
	return m
}

// SwitchMSwitchStackConfigMemberPortStatusType
//
// Definition: stack_memberPortStatusType
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

// SwitchMSwitchStackConfigStackConfig
//
// Definition: stack_stackConfig
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

// SwitchMSwitchStackConfigStackConfigList
//
// Definition: stack_stackConfigList
type SwitchMSwitchStackConfigStackConfigList []*SwitchMSwitchStackConfigStackConfig

func MakeSwitchMSwitchStackConfigStackConfigList() SwitchMSwitchStackConfigStackConfigList {
	m := make(SwitchMSwitchStackConfigStackConfigList, 0)
	return m
}

// AddStack
//
// Operation ID: addStack
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchStackConfigAuditIdList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindStackBySwitchId
//
// Operation ID: findStackBySwitchId
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchStackConfigStackConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindStackMemberBySerialNumber
//
// Operation ID: findStackMemberBySerialNumber
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("serialNumber", serialNumber)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSwitchStackConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
