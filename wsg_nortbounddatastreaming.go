package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGNorthboundDataStreamingService struct {
	apiClient *VSZClient
}

func NewWSGNorthboundDataStreamingService(c *VSZClient) *WSGNorthboundDataStreamingService {
	s := new(WSGNorthboundDataStreamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGNorthboundDataStreamingService() *WSGNorthboundDataStreamingService {
	return NewWSGNorthboundDataStreamingService(ss.apiClient)
}

// WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile
//
// Definition: northboundDataStreaming_createNorthboundDataStreamingProfile
type WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile struct {
	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerHost *string `json:"serverHost"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerPort *string `json:"serverPort"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	SystemId *string `json:"systemId"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	User *string `json:"user"`
}

func NewWSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile)
	return m
}

// WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes
//
// Definition: northboundDataStreaming_modifyNorthboundDataStreamingEventCodes
type WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes struct {
	// NorthboundDataStreamingAcceptedEventCodes
	// Constraints:
	//    - required
	NorthboundDataStreamingAcceptedEventCodes []int `json:"northboundDataStreamingAcceptedEventCodes"`
}

func NewWSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes() *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes {
	m := new(WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes)
	return m
}

// WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile
//
// Definition: northboundDataStreaming_modifyNorthboundDataStreamingProfile
type WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile struct {
	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerHost *string `json:"serverHost"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerPort *string `json:"serverPort"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	SystemId *string `json:"systemId"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	User *string `json:"user"`
}

func NewWSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile)
	return m
}

// WSGNorthboundDataStreamingEventCodes
//
// Definition: northboundDataStreaming_northboundDataStreamingEventCodes
type WSGNorthboundDataStreamingEventCodes struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more Northbound Data Streaming accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGNorthboundDataStreamingEventCodesListType `json:"list,omitempty"`

	// TotalCount
	// Total Northbound Data Streaming accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGNorthboundDataStreamingEventCodesAPIResponse struct {
	*RawAPIResponse
	Data *WSGNorthboundDataStreamingEventCodes
}

func newWSGNorthboundDataStreamingEventCodesAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGNorthboundDataStreamingEventCodesAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGNorthboundDataStreamingEventCodesAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGNorthboundDataStreamingEventCodes)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGNorthboundDataStreamingEventCodes() *WSGNorthboundDataStreamingEventCodes {
	m := new(WSGNorthboundDataStreamingEventCodes)
	return m
}

// WSGNorthboundDataStreamingEventCodesListType
//
// Definition: northboundDataStreaming_northboundDataStreamingEventCodesListType
type WSGNorthboundDataStreamingEventCodesListType struct {
	// Code
	// Northbound Data Streaming accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// Northbound Data Streaming accepted event type
	Type *string `json:"type,omitempty"`
}

func NewWSGNorthboundDataStreamingEventCodesListType() *WSGNorthboundDataStreamingEventCodesListType {
	m := new(WSGNorthboundDataStreamingEventCodesListType)
	return m
}

// WSGNorthboundDataStreamingProfile
//
// Definition: northboundDataStreaming_northboundDataStreamingProfile
type WSGNorthboundDataStreamingProfile struct {
	// ConnectionStatus
	// Connection status of the Northbound Data Streaming Profile
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty"`

	// Id
	// UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Id *string `json:"id,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Name *string `json:"name,omitempty"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Password *string `json:"password,omitempty"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerHost *string `json:"serverHost,omitempty"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerPort *string `json:"serverPort,omitempty"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	SystemId *string `json:"systemId,omitempty"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	User *string `json:"user,omitempty"`
}

type WSGNorthboundDataStreamingProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGNorthboundDataStreamingProfile
}

func newWSGNorthboundDataStreamingProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGNorthboundDataStreamingProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGNorthboundDataStreamingProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGNorthboundDataStreamingProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingProfile)
	return m
}

// WSGNorthboundDataStreamingProfileList
//
// Definition: northboundDataStreaming_northboundDataStreamingProfileList
type WSGNorthboundDataStreamingProfileList struct {
	Extra *WSGNorthboundDataStreamingProfileListExtraType `json:"extra,omitempty"`

	List []*WSGNorthboundDataStreamingProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGNorthboundDataStreamingProfileList) UnmarshalJSON(b []byte) error {
	type _WSGNorthboundDataStreamingProfileList WSGNorthboundDataStreamingProfileList
	tmpType := new(_WSGNorthboundDataStreamingProfileList)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "list")
	*t = WSGNorthboundDataStreamingProfileList(*tmpType)
	return nil
}

func (t *WSGNorthboundDataStreamingProfileList) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	return json.Marshal(tmp)
}

type WSGNorthboundDataStreamingProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGNorthboundDataStreamingProfileList
}

func newWSGNorthboundDataStreamingProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGNorthboundDataStreamingProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGNorthboundDataStreamingProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGNorthboundDataStreamingProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGNorthboundDataStreamingProfileList() *WSGNorthboundDataStreamingProfileList {
	m := new(WSGNorthboundDataStreamingProfileList)
	return m
}

// WSGNorthboundDataStreamingProfileListExtraType
//
// Definition: northboundDataStreaming_northboundDataStreamingProfileListExtraType
type WSGNorthboundDataStreamingProfileListExtraType struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled,omitempty"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled,omitempty"`

	// StreamingDomainIds
	// Domain Ids for 'streamingByDomainZoneEnabled' settings
	StreamingDomainIds []string `json:"streamingDomainIds,omitempty"`

	// StreamingZoneIds
	// Zone Ids for 'streamingByDomainZoneEnabled' settings
	StreamingZoneIds []string `json:"streamingZoneIds,omitempty"`
}

func NewWSGNorthboundDataStreamingProfileListExtraType() *WSGNorthboundDataStreamingProfileListExtraType {
	m := new(WSGNorthboundDataStreamingProfileListExtraType)
	return m
}

// WSGNorthboundDataStreamingSettings
//
// Definition: northboundDataStreaming_northboundDataStreamingSettings
type WSGNorthboundDataStreamingSettings struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	// Constraints:
	//    - required
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	// Constraints:
	//    - required
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled"`

	// StreamingDomainIds
	// Domain Ids for 'streamingByDomainZoneEnabled' settings
	StreamingDomainIds []string `json:"streamingDomainIds,omitempty"`

	// StreamingZoneIds
	// Zone Ids for 'streamingByDomainZoneEnabled' settings
	StreamingZoneIds []string `json:"streamingZoneIds,omitempty"`
}

func NewWSGNorthboundDataStreamingSettings() *WSGNorthboundDataStreamingSettings {
	m := new(WSGNorthboundDataStreamingSettings)
	return m
}

// AddNorthboundDataStreamingProfile
//
// Use this API command to create northbound Data Streaming Profile
//
// Operation ID: addNorthboundDataStreamingProfile
// Operation path: /northboundDataStreamingProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile
func (s *WSGNorthboundDataStreamingService) AddNorthboundDataStreamingProfile(ctx context.Context, body *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddNorthboundDataStreamingProfile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteNorthboundDataStreamingProfileById
//
// Use this API command to delete northbound Data Streaming Profile
//
// Operation ID: deleteNorthboundDataStreamingProfileById
// Operation path: /northboundDataStreamingProfile/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) DeleteNorthboundDataStreamingProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteNorthboundDataStreamingProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindNorthboundDataStreamingEventCodes
//
// Use this API command to retrieve NorthboundDataStreamingEventCodes.
//
// Operation ID: findNorthboundDataStreamingEventCodes
// Operation path: /northboundDataStreamingEventCodes
// Success code: 200 (OK)
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingEventCodes(ctx context.Context, mutators ...RequestMutator) (*WSGNorthboundDataStreamingEventCodesAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGNorthboundDataStreamingEventCodesAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindNorthboundDataStreamingEventCodes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGNorthboundDataStreamingEventCodesAPIResponse), err
}

// FindNorthboundDataStreamingProfileById
//
// Use this API command to retrieve northbound Data Streaming Profile
//
// Operation ID: findNorthboundDataStreamingProfileById
// Operation path: /northboundDataStreamingProfile/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGNorthboundDataStreamingProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGNorthboundDataStreamingProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGNorthboundDataStreamingProfileAPIResponse), err
}

// FindNorthboundDataStreamingProfileList
//
// Use this API command to retrieve northbound Data Streaming Profile List
//
// Operation ID: findNorthboundDataStreamingProfileList
// Operation path: /northboundDataStreamingProfileList
// Success code: 200 (OK)
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileList(ctx context.Context, mutators ...RequestMutator) (*WSGNorthboundDataStreamingProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGNorthboundDataStreamingProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGNorthboundDataStreamingProfileListAPIResponse), err
}

// UpdateNorthboundDataStreamingEventCodes
//
// Use this API command to modify NorthboundDataStreamingEventCodes.
//
// Operation ID: updateNorthboundDataStreamingEventCodes
// Operation path: /northboundDataStreamingEventCodes
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingEventCodes(ctx context.Context, body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateNorthboundDataStreamingEventCodes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// UpdateNorthboundDataStreamingProfileById
//
// Use this API command to update northbound Data Streaming Profile
//
// Operation ID: updateNorthboundDataStreamingProfileById
// Operation path: /northboundDataStreamingProfile/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingProfileById(ctx context.Context, body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateNorthboundDataStreamingProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// UpdateNorthboundDataStreamingSettings
//
// Use this API command to modify Northbound Data Streaming Settings.
//
// Operation ID: updateNorthboundDataStreamingSettings
// Operation path: /northboundDataStreamingSettings
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGNorthboundDataStreamingSettings
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingSettings(ctx context.Context, body *WSGNorthboundDataStreamingSettings, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateNorthboundDataStreamingSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
