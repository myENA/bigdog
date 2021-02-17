package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

func newSwitchMStaticRouteSettingStaticRouteAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMStaticRouteSettingStaticRouteAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMStaticRouteSettingStaticRouteAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMStaticRouteSettingStaticRoute)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
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

func newSwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMStaticRouteSettingStaticRoutesQueryResult)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
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
// Use this API command to Create Static Route.
//
// Operation ID: addStaticRoutes
// Operation path: /staticRoutes
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMStaticRouteSettingCreateStaticRoute
func (s *SwitchMStaticRouteSettingService) AddStaticRoutes(ctx context.Context, body *SwitchMStaticRouteSettingCreateStaticRoute, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddStaticRoutes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteStaticRoutes
//
// Use this API command to Delete Static Route by Id list.
//
// Operation ID: deleteStaticRoutes
// Operation path: /staticRoutes
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMStaticRouteSettingService) DeleteStaticRoutes(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteStaticRoutes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteStaticRoutesById
//
// Use this API command to Delete Static Route.
//
// Operation ID: deleteStaticRoutesById
// Operation path: /staticRoutes/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteSettingService) DeleteStaticRoutesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteStaticRoutesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindStaticRoutesById
//
// Use this API command to Retrieve Static Route.
//
// Operation ID: findStaticRoutesById
// Operation path: /staticRoutes/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteSettingService) FindStaticRoutesById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMStaticRouteSettingStaticRouteAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMStaticRouteSettingStaticRouteAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindStaticRoutesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMStaticRouteSettingStaticRouteAPIResponse), err
}

// FindStaticRoutesByQueryCriteria
//
// Use this API command to Retrieve Static Route list.
//
// Operation ID: findStaticRoutesByQueryCriteria
// Operation path: /staticRoutes/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMStaticRouteSettingService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindStaticRoutesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMStaticRouteSettingStaticRoutesQueryResultAPIResponse), err
}

// UpdateStaticRoutesById
//
// Use this API command to Update Static Route.
//
// Operation ID: updateStaticRoutesById
// Operation path: /staticRoutes/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMStaticRouteSettingUpdateStaticRoute
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteSettingService) UpdateStaticRoutesById(ctx context.Context, body *SwitchMStaticRouteSettingUpdateStaticRoute, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateStaticRoutesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
