package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

type SwitchMSwitchStackConfigAuditIdListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchStackConfigAuditIdList
}

func newSwitchMSwitchStackConfigAuditIdListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchStackConfigAuditIdListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchStackConfigAuditIdListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchStackConfigAuditIdList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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

type SwitchMSwitchStackConfigListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchStackConfigList
}

func newSwitchMSwitchStackConfigListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchStackConfigListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchStackConfigListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchStackConfigList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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

type SwitchMSwitchStackConfigStackConfigAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchStackConfigStackConfig
}

func newSwitchMSwitchStackConfigStackConfigAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchStackConfigStackConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchStackConfigStackConfigAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSwitchStackConfigStackConfig)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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
// Use this API command to create a stack configuration.
//
// Operation ID: addStack
// Operation path: /stack
// Success code: 200 (OK)
//
// Request body:
//	 - body SwitchMSwitchStackConfigStackConfigList
func (s *SwitchMSwitchStackConfigService) AddStack(ctx context.Context, body SwitchMSwitchStackConfigStackConfigList, mutators ...RequestMutator) (*SwitchMSwitchStackConfigAuditIdListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchStackConfigAuditIdListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddStack, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchStackConfigAuditIdListAPIResponse), err
}

// FindStackBySwitchId
//
// Use this API command to retrieve a stack configuration configured via SZ.
//
// Operation ID: findStackBySwitchId
// Operation path: /stack/{switchId}
// Success code: 200 (OK)
//
// Required parameters:
// - switchId string
//		- required
func (s *SwitchMSwitchStackConfigService) FindStackBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMSwitchStackConfigStackConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchStackConfigStackConfigAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindStackBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchStackConfigStackConfigAPIResponse), err
}

// FindStackMemberBySerialNumber
//
// Use this API command to retrieve the member of switches in a stack.
//
// Operation ID: findStackMemberBySerialNumber
// Operation path: /stack/member/{serialNumber}
// Success code: 200 (OK)
//
// Required parameters:
// - serialNumber string
//		- required
func (s *SwitchMSwitchStackConfigService) FindStackMemberBySerialNumber(ctx context.Context, serialNumber string, mutators ...RequestMutator) (*SwitchMSwitchStackConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchStackConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindStackMemberBySerialNumber, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("serialNumber", serialNumber)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchStackConfigListAPIResponse), err
}
