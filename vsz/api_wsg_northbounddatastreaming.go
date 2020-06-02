package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGNorthbounddatastreamingService struct {
	apiClient *APIClient
}

func NewWSGNorthbounddatastreamingService(c *APIClient) *WSGNorthbounddatastreamingService {
	s := new(WSGNorthbounddatastreamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGNorthbounddatastreamingService() *WSGNorthbounddatastreamingService {
	return NewWSGNorthbounddatastreamingService(ss.apiClient)
}

type WSGNorthbounddatastreamingCreateNorthboundDataStreamingProfile struct {
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

func NewWSGNorthbounddatastreamingCreateNorthboundDataStreamingProfile() *WSGNorthbounddatastreamingCreateNorthboundDataStreamingProfile {
	m := new(WSGNorthbounddatastreamingCreateNorthboundDataStreamingProfile)
	return m
}

type WSGNorthbounddatastreamingModifyNorthboundDataStreamingEventCodes struct {
	// NorthboundDataStreamingAcceptedEventCodes
	// Constraints:
	//    - required
	NorthboundDataStreamingAcceptedEventCodes []int `json:"northboundDataStreamingAcceptedEventCodes"`
}

func NewWSGNorthbounddatastreamingModifyNorthboundDataStreamingEventCodes() *WSGNorthbounddatastreamingModifyNorthboundDataStreamingEventCodes {
	m := new(WSGNorthbounddatastreamingModifyNorthboundDataStreamingEventCodes)
	return m
}

type WSGNorthbounddatastreamingModifyNorthboundDataStreamingProfile struct {
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

func NewWSGNorthbounddatastreamingModifyNorthboundDataStreamingProfile() *WSGNorthbounddatastreamingModifyNorthboundDataStreamingProfile {
	m := new(WSGNorthbounddatastreamingModifyNorthboundDataStreamingProfile)
	return m
}

type WSGNorthbounddatastreamingEventCodes struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more Northbound Data Streaming accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGNorthbounddatastreamingEventCodesListType `json:"list,omitempty"`

	// TotalCount
	// Total Northbound Data Streaming accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGNorthbounddatastreamingEventCodes() *WSGNorthbounddatastreamingEventCodes {
	m := new(WSGNorthbounddatastreamingEventCodes)
	return m
}

type WSGNorthbounddatastreamingEventCodesListType struct {
	// Code
	// Northbound Data Streaming accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// Northbound Data Streaming accepted event type
	Type *string `json:"type,omitempty"`
}

func NewWSGNorthbounddatastreamingEventCodesListType() *WSGNorthbounddatastreamingEventCodesListType {
	m := new(WSGNorthbounddatastreamingEventCodesListType)
	return m
}

type WSGNorthbounddatastreamingProfile struct {
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

func NewWSGNorthbounddatastreamingProfile() *WSGNorthbounddatastreamingProfile {
	m := new(WSGNorthbounddatastreamingProfile)
	return m
}

type WSGNorthbounddatastreamingProfileList struct {
	Extra *WSGNorthbounddatastreamingProfileListExtraType `json:"extra,omitempty"`

	List []*WSGNorthbounddatastreamingProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGNorthbounddatastreamingProfileList) UnmarshalJSON(b []byte) error {
	tmpt := new(WSGNorthbounddatastreamingProfileList)
	if err := json.Unmarshal(b, tmpt); err != nil {
		return err
	}
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	delete(tmp, "extra")
	delete(tmp, "list")
	tmpt.XAdditionalProperties = tmp
	*t = *tmpt
	return nil
}

func (t *WSGNorthbounddatastreamingProfileList) MarshalJSON() ([]byte, error) {
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

func NewWSGNorthbounddatastreamingProfileList() *WSGNorthbounddatastreamingProfileList {
	m := new(WSGNorthbounddatastreamingProfileList)
	return m
}

type WSGNorthbounddatastreamingProfileListExtraType struct {
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

func NewWSGNorthbounddatastreamingProfileListExtraType() *WSGNorthbounddatastreamingProfileListExtraType {
	m := new(WSGNorthbounddatastreamingProfileListExtraType)
	return m
}

type WSGNorthbounddatastreamingSettings struct {
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

func NewWSGNorthbounddatastreamingSettings() *WSGNorthbounddatastreamingSettings {
	m := new(WSGNorthbounddatastreamingSettings)
	return m
}

// AddNorthboundDataStreamingProfile
//
// Use this API command to create northbound Data Streaming Profile
//
// Request Body:
//	 - body *WSGNorthbounddatastreamingCreateNorthboundDataStreamingProfile
func (s *WSGNorthbounddatastreamingService) AddNorthboundDataStreamingProfile(ctx context.Context, body *WSGNorthbounddatastreamingCreateNorthboundDataStreamingProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddNorthboundDataStreamingProfile, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteNorthboundDataStreamingProfileById
//
// Use this API command to delete northbound Data Streaming Profile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthbounddatastreamingService) DeleteNorthboundDataStreamingProfileById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteNorthboundDataStreamingProfileById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// FindNorthboundDataStreamingEventCodes
//
// Use this API command to retrieve NorthboundDataStreamingEventCodes.
func (s *WSGNorthbounddatastreamingService) FindNorthboundDataStreamingEventCodes(ctx context.Context) (*WSGNorthbounddatastreamingEventCodes, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGNorthbounddatastreamingEventCodes
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindNorthboundDataStreamingEventCodes, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGNorthbounddatastreamingEventCodes()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindNorthboundDataStreamingProfileById
//
// Use this API command to retrieve northbound Data Streaming Profile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthbounddatastreamingService) FindNorthboundDataStreamingProfileById(ctx context.Context, id string) (*WSGNorthbounddatastreamingProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGNorthbounddatastreamingProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGNorthbounddatastreamingProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindNorthboundDataStreamingProfileList
//
// Use this API command to retrieve northbound Data Streaming Profile List
func (s *WSGNorthbounddatastreamingService) FindNorthboundDataStreamingProfileList(ctx context.Context) (*WSGNorthbounddatastreamingProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGNorthbounddatastreamingProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileList, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGNorthbounddatastreamingProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateNorthboundDataStreamingEventCodes
//
// Use this API command to modify NorthboundDataStreamingEventCodes.
//
// Request Body:
//	 - body *WSGNorthbounddatastreamingModifyNorthboundDataStreamingEventCodes
func (s *WSGNorthbounddatastreamingService) UpdateNorthboundDataStreamingEventCodes(ctx context.Context, body *WSGNorthbounddatastreamingModifyNorthboundDataStreamingEventCodes) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingEventCodes, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateNorthboundDataStreamingProfileById
//
// Use this API command to update northbound Data Streaming Profile
//
// Request Body:
//	 - body *WSGNorthbounddatastreamingModifyNorthboundDataStreamingProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthbounddatastreamingService) UpdateNorthboundDataStreamingProfileById(ctx context.Context, body *WSGNorthbounddatastreamingModifyNorthboundDataStreamingProfile, id string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingProfileById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateNorthboundDataStreamingSettings
//
// Use this API command to modify Northbound Data Streaming Settings.
//
// Request Body:
//	 - body *WSGNorthbounddatastreamingSettings
func (s *WSGNorthbounddatastreamingService) UpdateNorthboundDataStreamingSettings(ctx context.Context, body *WSGNorthbounddatastreamingSettings) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingSettings, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
