package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchSpecificSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchSpecificSettingsService(c *VSZClient) *SwitchMSwitchSpecificSettingsService {
	s := new(SwitchMSwitchSpecificSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchSpecificSettingsService() *SwitchMSwitchSpecificSettingsService {
	return NewSwitchMSwitchSpecificSettingsService(ss.apiClient)
}

type SwitchMSwitchSpecificSettingsDHCPOption struct {
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

func NewSwitchMSwitchSpecificSettingsDHCPOption() *SwitchMSwitchSpecificSettingsDHCPOption {
	m := new(SwitchMSwitchSpecificSettingsDHCPOption)
	return m
}

type SwitchMSwitchSpecificSettingsDHCPServer struct {
	// DefaultRouterIp
	// Default Router Ip
	DefaultRouterIp *string `json:"defaultRouterIp,omitempty"`

	DhcpOptions []*SwitchMSwitchSpecificSettingsDHCPServer `json:"dhcpOptions,omitempty"`

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

func NewSwitchMSwitchSpecificSettingsDHCPServer() *SwitchMSwitchSpecificSettingsDHCPServer {
	m := new(SwitchMSwitchSpecificSettingsDHCPServer)
	return m
}

type SwitchMSwitchSpecificSettingsIdList struct {
	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMSwitchSpecificSettingsIdList() *SwitchMSwitchSpecificSettingsIdList {
	m := new(SwitchMSwitchSpecificSettingsIdList)
	return m
}

type SwitchMSwitchSpecificSettingsSpecificSettings struct {
	// CreatedTime
	// The create time of the Specific Settings
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*SwitchMSwitchSpecificSettingsSpecificSettings `json:"dhcpServers,omitempty"`

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

func NewSwitchMSwitchSpecificSettingsSpecificSettings() *SwitchMSwitchSpecificSettingsSpecificSettings {
	m := new(SwitchMSwitchSpecificSettingsSpecificSettings)
	return m
}

type SwitchMSwitchSpecificSettingsSpecificSettingsAllResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchSpecificSettingsSpecificSettingsAllResult `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Specific Settings returned out of the complete Specific Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Specific Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchSpecificSettingsSpecificSettingsAllResult `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Specific Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Specific Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchSpecificSettingsSpecificSettingsAllResult() *SwitchMSwitchSpecificSettingsSpecificSettingsAllResult {
	m := new(SwitchMSwitchSpecificSettingsSpecificSettingsAllResult)
	return m
}

// SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType
//
// Any additional response data
type SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType() *SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType {
	m := new(SwitchMSwitchSpecificSettingsSpecificSettingsAllResultExtraType)
	return m
}

type SwitchMSwitchSpecificSettingsUpdateSpecificSettings struct {
	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*SwitchMSwitchSpecificSettingsUpdateSpecificSettings `json:"dhcpServers,omitempty"`

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

func NewSwitchMSwitchSpecificSettingsUpdateSpecificSettings() *SwitchMSwitchSpecificSettingsUpdateSpecificSettings {
	m := new(SwitchMSwitchSpecificSettingsUpdateSpecificSettings)
	return m
}

// DeleteSpecificSettingsById
//
// Use this API command to Delete Specific Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) DeleteSpecificSettingsById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindSpecificSettings
//
// Use this API command to Retrieve all Specific Setting list.
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettings(ctx context.Context) (*SwitchMSwitchSpecificSettingsSpecificSettingsAllResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchSpecificSettingsSpecificSettingsAllResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSpecificSettings, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchSpecificSettingsSpecificSettingsAllResult()
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
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, id string) (*SwitchMSwitchSpecificSettingsSpecificSettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchSpecificSettingsSpecificSettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSpecificSettingsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchSpecificSettingsSpecificSettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSpecificSettingsById
//
// Use this API command to Update Specific Settings.
//
// Request Body:
//	 - body *SwitchMSwitchSpecificSettingsUpdateSpecificSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) UpdateSpecificSettingsById(ctx context.Context, body *SwitchMSwitchSpecificSettingsUpdateSpecificSettings, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
