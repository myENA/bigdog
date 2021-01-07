package bigdog

// API Version: v9_1

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

// SwitchMSpecificSettingsDHCPOption
//
// Definition: specificSettings_DHCPOption
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

// SwitchMSpecificSettingsDHCPServer
//
// Definition: specificSettings_DHCPServer
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

// SwitchMSpecificSettingsIdList
//
// Definition: specificSettings_idList
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

// SwitchMSpecificSettings
//
// Definition: specificSettings_SpecificSettings
type SwitchMSpecificSettings struct {
	AltoId *string `json:"altoId,omitempty"`

	// CreatedTime
	// The create time of the Specific Settings
	CreatedTime *int `json:"createdTime,omitempty"`

	// DefaultGateway
	// DefaultGateway
	DefaultGateway *string `json:"defaultGateway,omitempty"`

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

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// JumboMode
	// JumboMode
	JumboMode *bool `json:"jumboMode,omitempty"`

	// SubnetMask
	// SubnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// UpdatedTime
	// The modify time of the Specific Settings
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type SwitchMSpecificSettingsAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSpecificSettings
}

func newSwitchMSpecificSettingsAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMSpecificSettingsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMSpecificSettingsAPIResponse) Hydrate() error {
	r.Data = new(SwitchMSpecificSettings)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMSpecificSettings() *SwitchMSpecificSettings {
	m := new(SwitchMSpecificSettings)
	return m
}

// SwitchMSpecificSettingsAllResult
//
// Definition: specificSettings_specificSettingsAllResult
type SwitchMSpecificSettingsAllResult struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMSpecificSettingsAllResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSpecificSettingsAllResult
}

func newSwitchMSpecificSettingsAllResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMSpecificSettingsAllResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMSpecificSettingsAllResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMSpecificSettingsAllResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMSpecificSettingsAllResult() *SwitchMSpecificSettingsAllResult {
	m := new(SwitchMSpecificSettingsAllResult)
	return m
}

// SwitchMSpecificSettingsUpdateSpecificSettings
//
// Definition: specificSettings_UpdateSpecificSettings
type SwitchMSpecificSettingsUpdateSpecificSettings struct {
	AltoId *string `json:"altoId,omitempty"`

	// DefaultGateway
	// DefaultGateway
	DefaultGateway *string `json:"defaultGateway,omitempty"`

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

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// JumboMode
	// JumboMode
	JumboMode *bool `json:"jumboMode,omitempty"`

	// SubnetMask
	// SubnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewSwitchMSpecificSettingsUpdateSpecificSettings() *SwitchMSpecificSettingsUpdateSpecificSettings {
	m := new(SwitchMSpecificSettingsUpdateSpecificSettings)
	return m
}

// DeleteSpecificSettingsById
//
// Operation ID: deleteSpecificSettingsById
//
// Use this API command to Delete Specific Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificSettingsService) DeleteSpecificSettingsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteSpecificSettingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindSpecificSettings
//
// Operation ID: findSpecificSettings
//
// Use this API command to Retrieve all Specific Setting list.
func (s *SwitchMSpecificSettingsService) FindSpecificSettings(ctx context.Context, mutators ...RequestMutator) (*SwitchMSpecificSettingsAllResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSpecificSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSpecificSettingsAllResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSpecificSettingsById
//
// Operation ID: findSpecificSettingsById
//
// Use this API command to Retrieve Specific Settings.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMSpecificSettingsAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSpecificSettingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSpecificSettings()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSpecificSettingsById
//
// Operation ID: updateSpecificSettingsById
//
// Use this API command to Update Specific Settings.
//
// Request Body:
//	 - body *SwitchMSpecificSettingsUpdateSpecificSettings
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSpecificSettingsService) UpdateSpecificSettingsById(ctx context.Context, body *SwitchMSpecificSettingsUpdateSpecificSettings, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateSpecificSettingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
