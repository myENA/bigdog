package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMStaticRouteSettingService struct {
	apiClient *VSZClient
}

func NewSwitchMStaticRouteSettingService(c *VSZClient) *SwitchMStaticRouteSettingService {
	s := new(SwitchMStaticRouteSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMStaticRouteSettingService() *SwitchMStaticRouteSettingService {
	return NewSwitchMStaticRouteSettingService(ss.apiClient)
}

// SwitchMStaticRouteSettingCreateStaticRoute
//
// Definition: staticRoute_createStaticRoute
type SwitchMStaticRouteSettingCreateStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *int `json:"adminDistance,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// DestinationIp
	// Destination IP
	DestinationIp *string `json:"destinationIp,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// NextHop
	// Next Hop
	NextHop *string `json:"nextHop,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMStaticRouteSettingCreateStaticRoute() *SwitchMStaticRouteSettingCreateStaticRoute {
	m := new(SwitchMStaticRouteSettingCreateStaticRoute)
	return m
}

// SwitchMStaticRouteSettingStaticRoute
//
// Definition: staticRoute_staticRoute
type SwitchMStaticRouteSettingStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance interface{} `json:"adminDistance,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// CreatedTime
	// The create time of the Static Route
	CreatedTime *int `json:"createdTime,omitempty"`

	// DestinationIp
	// Destination IP
	DestinationIp *string `json:"destinationIp,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Static Route Id
	Id *string `json:"id,omitempty"`

	// NextHop
	// Next Hop
	NextHop *string `json:"nextHop,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the Static Route
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type SwitchMStaticRouteSettingStaticRouteAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMStaticRouteSettingStaticRoute
}

func newSwitchMStaticRouteSettingStaticRouteAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMStaticRouteSettingStaticRouteAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMStaticRouteSettingStaticRouteAPIResponse) Hydrate() error {
	r.Data = new(SwitchMStaticRouteSettingStaticRoute)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMStaticRouteSettingStaticRoute() *SwitchMStaticRouteSettingStaticRoute {
	m := new(SwitchMStaticRouteSettingStaticRoute)
	return m
}

// SwitchMStaticRouteSettingStaticRoutesQueryResult
//
// Definition: staticRoute_staticRoutesQueryResult
type SwitchMStaticRouteSettingStaticRoutesQueryResult struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Static Route returned out of the complete Static Route list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Static Routes after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMStaticRouteSettingStaticRoute `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Static Route count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Static Route Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMStaticRouteSettingStaticRoutesQueryResult
}

func newSwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMStaticRouteSettingStaticRoutesQueryResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMStaticRouteSettingStaticRoutesQueryResult() *SwitchMStaticRouteSettingStaticRoutesQueryResult {
	m := new(SwitchMStaticRouteSettingStaticRoutesQueryResult)
	return m
}

// SwitchMStaticRouteSettingUpdateStaticRoute
//
// Definition: staticRoute_updateStaticRoute
type SwitchMStaticRouteSettingUpdateStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *int `json:"adminDistance,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// DestinationIp
	// Destination IP
	DestinationIp *string `json:"destinationIp,omitempty"`

	// NextHop
	// Next Hop
	NextHop *string `json:"nextHop,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`
}

func NewSwitchMStaticRouteSettingUpdateStaticRoute() *SwitchMStaticRouteSettingUpdateStaticRoute {
	m := new(SwitchMStaticRouteSettingUpdateStaticRoute)
	return m
}

// AddStaticRoutes
//
// Operation ID: addStaticRoutes
//
// Use this API command to Create Static Route.
//
// Request Body:
//	 - body *SwitchMStaticRouteSettingCreateStaticRoute
func (s *SwitchMStaticRouteSettingService) AddStaticRoutes(ctx context.Context, body *SwitchMStaticRouteSettingCreateStaticRoute, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddStaticRoutes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteStaticRoutes
//
// Operation ID: deleteStaticRoutes
//
// Use this API command to Delete Static Route by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMStaticRouteSettingService) DeleteStaticRoutes(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteStaticRoutes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteStaticRoutesById
//
// Operation ID: deleteStaticRoutesById
//
// Use this API command to Delete Static Route.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteSettingService) DeleteStaticRoutesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteStaticRoutesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindStaticRoutesById
//
// Operation ID: findStaticRoutesById
//
// Use this API command to Retrieve Static Route.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteSettingService) FindStaticRoutesById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMStaticRouteSettingStaticRouteAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMStaticRouteSettingStaticRouteAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMStaticRouteSettingStaticRouteAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindStaticRoutesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMStaticRouteSettingStaticRouteAPIResponse), err
}

// FindStaticRoutesByQueryCriteria
//
// Operation ID: findStaticRoutesByQueryCriteria
//
// Use this API command to Retrieve Static Route list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMStaticRouteSettingService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindStaticRoutesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse), err
}

// UpdateStaticRoutesById
//
// Operation ID: updateStaticRoutesById
//
// Use this API command to Update Static Route.
//
// Request Body:
//	 - body *SwitchMStaticRouteSettingUpdateStaticRoute
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteSettingService) UpdateStaticRoutesById(ctx context.Context, body *SwitchMStaticRouteSettingUpdateStaticRoute, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateStaticRoutesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
