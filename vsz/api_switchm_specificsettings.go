package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSpecificsettingsService struct {
	apiClient *APIClient
}

func NewSwitchMSpecificsettingsService(c *APIClient) *SwitchMSpecificsettingsService {
	s := new(SwitchMSpecificsettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSpecificsettingsService() *SwitchMSpecificsettingsService {
	return NewSwitchMSpecificsettingsService(ss.apiClient)
}

type SwitchMSpecificsettingsDHCPOption struct {
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

func NewSwitchMSpecificsettingsDHCPOption() *SwitchMSpecificsettingsDHCPOption {
	m := new(SwitchMSpecificsettingsDHCPOption)
	return m
}

type SwitchMSpecificsettingsDHCPServer struct {
	// DefaultRouterIp
	// Default Router Ip
	DefaultRouterIp *string `json:"defaultRouterIp,omitempty"`

	DhcpOptions []*SwitchMSpecificsettingsDHCPOption `json:"dhcpOptions,omitempty"`

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

func NewSwitchMSpecificsettingsDHCPServer() *SwitchMSpecificsettingsDHCPServer {
	m := new(SwitchMSpecificsettingsDHCPServer)
	return m
}

type SwitchMSpecificsettingsIdList struct {
	// Hostname
	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`
}

func NewSwitchMSpecificsettingsIdList() *SwitchMSpecificsettingsIdList {
	m := new(SwitchMSpecificsettingsIdList)
	return m
}

type SwitchMSpecificsettings struct {
	// CreatedTime
	// The create time of the Specific Settings
	CreatedTime *int `json:"createdTime,omitempty"`

	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*SwitchMSpecificsettingsDHCPServer `json:"dhcpServers,omitempty"`

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

func NewSwitchMSpecificsettings() *SwitchMSpecificsettings {
	m := new(SwitchMSpecificsettings)
	return m
}

type SwitchMSpecificsettingsAllResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSpecificsettingsAllResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Specific Settings returned out of the complete Specific Settings list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Specific Settings after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSpecificsettingsIdList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Specific Settings count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Specific Settings count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSpecificsettingsAllResult() *SwitchMSpecificsettingsAllResult {
	m := new(SwitchMSpecificsettingsAllResult)
	return m
}

// SwitchMSpecificsettingsAllResultExtraType
//
// Any additional response data
type SwitchMSpecificsettingsAllResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSpecificsettingsAllResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSpecificsettingsAllResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSpecificsettingsAllResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSpecificsettingsAllResultExtraType() *SwitchMSpecificsettingsAllResultExtraType {
	m := new(SwitchMSpecificsettingsAllResultExtraType)
	return m
}

type SwitchMSpecificsettingsUpdateSpecificSettings struct {
	// DhcpServerEnabled
	// DHCP server enabled
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`

	DhcpServers []*SwitchMSpecificsettingsDHCPServer `json:"dhcpServers,omitempty"`

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

func NewSwitchMSpecificsettingsUpdateSpecificSettings() *SwitchMSpecificsettingsUpdateSpecificSettings {
	m := new(SwitchMSpecificsettingsUpdateSpecificSettings)
	return m
}

// DeleteSpecificSettingsById
//
// Use this API command to Delete Specific Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificsettingsService) DeleteSpecificSettingsById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
func (s *SwitchMSpecificsettingsService) FindSpecificSettings(ctx context.Context) (*SwitchMSpecificsettingsAllResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSpecificsettingsAllResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSpecificSettings, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSpecificsettingsAllResult()
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
func (s *SwitchMSpecificsettingsService) FindSpecificSettingsById(ctx context.Context, id string) (*SwitchMSpecificsettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSpecificsettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSpecificSettingsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSpecificsettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSpecificSettingsById
//
// Use this API command to Update Specific Settings.
//
// Request Body:
//	 - body *SwitchMSpecificsettingsUpdateSpecificSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificsettingsService) UpdateSpecificSettingsById(ctx context.Context, body *SwitchMSpecificsettingsUpdateSpecificSettings, id string) (*APIResponseMeta, error) {
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
