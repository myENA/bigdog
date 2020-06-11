package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSpecificSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMSpecificSettingsService(c *VSZClient) *SwitchMSpecificSettingsService {
	s := new(SwitchMSpecificSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSpecificSettingsService() *SwitchMSpecificSettingsService {
	return NewSwitchMSpecificSettingsService(ss.apiClient)
}

type SwitchMSpecificSettingsDHCPOption struct {
	// Seq
	// Seq of Option
	Seq *int `json:"seq,omitempty"`

	// Type
	// Type of Option
	Type *string `json:"type,omitempty"`

	// Value
	// Value of Option
	Value *string `json:"value,omitempty"`
}

func NewSwitchMSpecificSettingsDHCPOption() *SwitchMSpecificSettingsDHCPOption {
	m := new(SwitchMSpecificSettingsDHCPOption)
	return m
}

type SwitchMSpecificSettingsDHCPServer struct {
	// DefaultRouterIp
	// Default Router Ip
	DefaultRouterIp *string `json:"defaultRouterIp,omitempty"`

	DhcpOptions []*SwitchMSpecificSettingsDHCPOption `json:"dhcpOptions,omitempty"`

	// ExcludedEnd
	// Excluded range end
	ExcludedEnd *string `json:"excludedEnd,omitempty"`

	// ExcludedStart
	// Excluded range start
	ExcludedStart *string `json:"excludedStart,omitempty"`

	// LeaseDays
	// Lease Days
	LeaseDays *int `json:"leaseDays,omitempty"`

	// LeaseHrs
	// Lease Hours
	LeaseHrs *int `json:"leaseHrs,omitempty"`

	// LeaseMins
	// Lease Mins
	LeaseMins *int `json:"leaseMins,omitempty"`

	// Network
	// Network/Mask
	Network *string `json:"network,omitempty"`

	// PoolName
	// Pool Name
	PoolName *string `json:"poolName,omitempty"`
}

func NewSwitchMSpecificSettingsDHCPServer() *SwitchMSpecificSettingsDHCPServer {
	m := new(SwitchMSpecificSettingsDHCPServer)
	return m
}

type SwitchMSpecificSettingsIdList struct {
	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMSpecificSettingsIdList() *SwitchMSpecificSettingsIdList {
	m := new(SwitchMSpecificSettingsIdList)
	return m
}

type SwitchMSpecificSettings struct {
	// CreatedTime
	// The create time of the Specific Settings
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*SwitchMSpecificSettingsDHCPServer `json:"dhcpServers,omitempty"`

	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`

	// IgmpSnooping
	// IGMP snopping
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// JumboMode
	// JumboMode
	JumboMode *bool `json:"jumboMode,omitempty"`

	// UpdatedTime
	// The modify time of the Specific Settings
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMSpecificSettings() *SwitchMSpecificSettings {
	m := new(SwitchMSpecificSettings)
	return m
}

type SwitchMSpecificSettingsAllResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSpecificSettingsAllResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Specific Settings returned out of the complete Specific Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Specific Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSpecificSettingsIdList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Specific Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Specific Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSpecificSettingsAllResult() *SwitchMSpecificSettingsAllResult {
	m := new(SwitchMSpecificSettingsAllResult)
	return m
}

// SwitchMSpecificSettingsAllResultExtraType
//
// Any additional response data
type SwitchMSpecificSettingsAllResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSpecificSettingsAllResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSpecificSettingsAllResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSpecificSettingsAllResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSpecificSettingsAllResultExtraType() *SwitchMSpecificSettingsAllResultExtraType {
	m := new(SwitchMSpecificSettingsAllResultExtraType)
	return m
}

type SwitchMSpecificSettingsUpdateSpecificSettings struct {
	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*SwitchMSpecificSettingsDHCPServer `json:"dhcpServers,omitempty"`

	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// IgmpSnooping
	// IGMP snopping
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// JumboMode
	// JumboMode
	JumboMode *bool `json:"jumboMode,omitempty"`
}

func NewSwitchMSpecificSettingsUpdateSpecificSettings() *SwitchMSpecificSettingsUpdateSpecificSettings {
	m := new(SwitchMSpecificSettingsUpdateSpecificSettings)
	return m
}

// DeleteSpecificSettingsById
//
// Use this API command to Delete Specific Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificSettingsService) DeleteSpecificSettingsById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSpecificSettingsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindSpecificSettings
//
// Use this API command to Retrieve all Specific Setting list.
func (s *SwitchMSpecificSettingsService) FindSpecificSettings(ctx context.Context, mutators ...RequestMutator) (*SwitchMSpecificSettingsAllResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSpecificSettingsAllResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSpecificSettings, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSpecificSettingsAllResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSpecificSettingsById
//
// Use this API command to Retrieve Specific Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMSpecificSettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSpecificSettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSpecificSettingsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSpecificSettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSpecificSettingsById
//
// Use this API command to Update Specific Settings.
//
// Request Body:
//	 - body *SwitchMSpecificSettingsUpdateSpecificSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificSettingsService) UpdateSpecificSettingsById(ctx context.Context, body *SwitchMSpecificSettingsUpdateSpecificSettings, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSpecificSettingsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
