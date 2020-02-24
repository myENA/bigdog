package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGNorthboundDataStreamingService struct {
	apiClient *APIClient
}

func NewWSGNorthboundDataStreamingService(c *APIClient) *WSGNorthboundDataStreamingService {
	s := new(WSGNorthboundDataStreamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGNorthboundDataStreamingService() *WSGNorthboundDataStreamingService {
	return NewWSGNorthboundDataStreamingService(ss.apiClient)
}

type WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile struct {
	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty" validate:"omitempty,dive"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Name *string `json:"name" validate:"required"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerHost *string `json:"serverHost" validate:"required"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerPort *string `json:"serverPort" validate:"required"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	SystemId *string `json:"systemId" validate:"required"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	User *string `json:"user" validate:"required"`
}

func NewWSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile)
	return m
}

type WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes struct {
	// NorthboundDataStreamingAcceptedEventCodes
	// Constraints:
	//    - required
	NorthboundDataStreamingAcceptedEventCodes []int `json:"northboundDataStreamingAcceptedEventCodes" validate:"required,dive"`
}

func NewWSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes() *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes {
	m := new(WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes)
	return m
}

type WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile struct {
	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty" validate:"omitempty,dive"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Name *string `json:"name" validate:"required"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerHost *string `json:"serverHost" validate:"required"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerPort *string `json:"serverPort" validate:"required"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	SystemId *string `json:"systemId" validate:"required"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	User *string `json:"user" validate:"required"`
}

func NewWSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile)
	return m
}

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

func NewWSGNorthboundDataStreamingEventCodes() *WSGNorthboundDataStreamingEventCodes {
	m := new(WSGNorthboundDataStreamingEventCodes)
	return m
}

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

type WSGNorthboundDataStreamingProfile struct {
	// ConnectionStatus
	// Connection status of the Northbound Data Streaming Profile
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty" validate:"omitempty,dive"`

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

func NewWSGNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingProfile)
	return m
}

type WSGNorthboundDataStreamingProfileList struct {
	Extra *WSGNorthboundDataStreamingProfileListExtraType `json:"extra,omitempty"`

	List []*WSGNorthboundDataStreamingProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGNorthboundDataStreamingProfileList) UnmarshalJSON(b []byte) error {
	tmpt := new(WSGNorthboundDataStreamingProfileList)
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

func NewWSGNorthboundDataStreamingProfileList() *WSGNorthboundDataStreamingProfileList {
	m := new(WSGNorthboundDataStreamingProfileList)
	return m
}

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

type WSGNorthboundDataStreamingSettings struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	// Constraints:
	//    - required
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled" validate:"required"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	// Constraints:
	//    - required
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled" validate:"required"`

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
// Request Body:
//	 - body *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile
func (s *WSGNorthboundDataStreamingService) AddNorthboundDataStreamingProfile(ctx context.Context, body *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddNorthboundDataStreamingProfile, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteNorthboundDataStreamingProfileById
//
// Use this API command to delete northbound Data Streaming Profile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) DeleteNorthboundDataStreamingProfileById(ctx context.Context, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteNorthboundDataStreamingProfileById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindNorthboundDataStreamingEventCodes
//
// Use this API command to retrieve NorthboundDataStreamingEventCodes.
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingEventCodes(ctx context.Context) (*WSGNorthboundDataStreamingEventCodes, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGNorthboundDataStreamingEventCodes
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindNorthboundDataStreamingEventCodes, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGNorthboundDataStreamingEventCodes()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindNorthboundDataStreamingProfileById
//
// Use this API command to retrieve northbound Data Streaming Profile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileById(ctx context.Context, id string) (*WSGNorthboundDataStreamingProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGNorthboundDataStreamingProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGNorthboundDataStreamingProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindNorthboundDataStreamingProfileList
//
// Use this API command to retrieve northbound Data Streaming Profile List
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileList(ctx context.Context) (*WSGNorthboundDataStreamingProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGNorthboundDataStreamingProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileList, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGNorthboundDataStreamingProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateNorthboundDataStreamingEventCodes
//
// Use this API command to modify NorthboundDataStreamingEventCodes.
//
// Request Body:
//	 - body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingEventCodes(ctx context.Context, body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingEventCodes, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
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
//	 - body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingProfileById(ctx context.Context, body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile, id string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingProfileById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
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
//	 - body *WSGNorthboundDataStreamingSettings
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingSettings(ctx context.Context, body *WSGNorthboundDataStreamingSettings) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingSettings, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
