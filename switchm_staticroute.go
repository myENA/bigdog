package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchStaticRouteSettingService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchStaticRouteSettingService(c *VSZClient) *SwitchMSwitchStaticRouteSettingService {
	s := new(SwitchMSwitchStaticRouteSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchStaticRouteSettingService() *SwitchMSwitchStaticRouteSettingService {
	return NewSwitchMSwitchStaticRouteSettingService(ss.apiClient)
}

type SwitchMSwitchStaticRouteSettingCreateStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *int `json:"adminDistance,omitempty"`

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

func NewSwitchMSwitchStaticRouteSettingCreateStaticRoute() *SwitchMSwitchStaticRouteSettingCreateStaticRoute {
	m := new(SwitchMSwitchStaticRouteSettingCreateStaticRoute)
	return m
}

type SwitchMSwitchStaticRouteSettingStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *string `json:"adminDistance,omitempty"`

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

func NewSwitchMSwitchStaticRouteSettingStaticRoute() *SwitchMSwitchStaticRouteSettingStaticRoute {
	m := new(SwitchMSwitchStaticRouteSettingStaticRoute)
	return m
}

type SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult struct {
	// Extra
	// Any additional response data
	Extra **SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Static Route returned out of the complete Static Route list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Static Routes after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Static Route count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Static Route Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchStaticRouteSettingStaticRoutesQueryResult() *SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult {
	m := new(SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult)
	return m
}

// SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType
//
// Any additional response data
type SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType() *SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType {
	m := new(SwitchMSwitchStaticRouteSettingStaticRoutesQueryResultExtraType)
	return m
}

type SwitchMSwitchStaticRouteSettingUpdateStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *int `json:"adminDistance,omitempty"`

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

func NewSwitchMSwitchStaticRouteSettingUpdateStaticRoute() *SwitchMSwitchStaticRouteSettingUpdateStaticRoute {
	m := new(SwitchMSwitchStaticRouteSettingUpdateStaticRoute)
	return m
}

// AddStaticRoutes
//
// Use this API command to Create Static Route.
//
// Request Body:
//	 - body *SwitchMSwitchStaticRouteSettingCreateStaticRoute
func (s *SwitchMSwitchStaticRouteSettingService) AddStaticRoutes(ctx context.Context, body *SwitchMSwitchStaticRouteSettingCreateStaticRoute) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddStaticRoutes, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteStaticRoutes
//
// Use this API command to Delete Static Route by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMSwitchStaticRouteSettingService) DeleteStaticRoutes(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteStaticRoutes, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteStaticRoutesById
//
// Use this API command to Delete Static Route.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchStaticRouteSettingService) DeleteStaticRoutesById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteStaticRoutesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindStaticRoutesById
//
// Use this API command to Retrieve Static Route.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchStaticRouteSettingService) FindStaticRoutesById(ctx context.Context, id string) (*SwitchMSwitchStaticRouteSettingStaticRoute, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchStaticRouteSettingStaticRoute
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindStaticRoutesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchStaticRouteSettingStaticRoute()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindStaticRoutesByQueryCriteria
//
// Use this API command to Retrieve Static Route list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchStaticRouteSettingService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchStaticRouteSettingStaticRoutesQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindStaticRoutesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchStaticRouteSettingStaticRoutesQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateStaticRoutesById
//
// Use this API command to Update Static Route.
//
// Request Body:
//	 - body *SwitchMSwitchStaticRouteSettingUpdateStaticRoute
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMSwitchStaticRouteSettingService) UpdateStaticRoutesById(ctx context.Context, body *SwitchMSwitchStaticRouteSettingUpdateStaticRoute, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateStaticRoutesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
