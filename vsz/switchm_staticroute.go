package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMStaticRouteService struct {
	apiClient *APIClient
}

func NewSwitchMStaticRouteService(c *APIClient) *SwitchMStaticRouteService {
	s := new(SwitchMStaticRouteService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMStaticRouteService() *SwitchMStaticRouteService {
	return NewSwitchMStaticRouteService(ss.apiClient)
}

type SwitchMStaticRouteCreateStaticRoute struct {
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
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMStaticRouteCreateStaticRoute() *SwitchMStaticRouteCreateStaticRoute {
	m := new(SwitchMStaticRouteCreateStaticRoute)
	return m
}

type SwitchMStaticRouteEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMStaticRouteEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMStaticRouteEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMStaticRouteEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMStaticRouteEmptyResult() *SwitchMStaticRouteEmptyResult {
	m := new(SwitchMStaticRouteEmptyResult)
	return m
}

type SwitchMStaticRoute struct {
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
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the Static Route
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMStaticRoute() *SwitchMStaticRoute {
	m := new(SwitchMStaticRoute)
	return m
}

type SwitchMStaticRoutesQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMStaticRoutesQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Static Route returned out of the complete Static Route list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Static Routes after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMStaticRoute `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Static Route count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Static Route Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMStaticRoutesQueryResult() *SwitchMStaticRoutesQueryResult {
	m := new(SwitchMStaticRoutesQueryResult)
	return m
}

// SwitchMStaticRoutesQueryResultExtraType
//
// Any additional response data
type SwitchMStaticRoutesQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMStaticRoutesQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMStaticRoutesQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMStaticRoutesQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMStaticRoutesQueryResultExtraType() *SwitchMStaticRoutesQueryResultExtraType {
	m := new(SwitchMStaticRoutesQueryResultExtraType)
	return m
}

type SwitchMStaticRouteUpdateStaticRoute struct {
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
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`
}

func NewSwitchMStaticRouteUpdateStaticRoute() *SwitchMStaticRouteUpdateStaticRoute {
	m := new(SwitchMStaticRouteUpdateStaticRoute)
	return m
}

// AddStaticRoutes
//
// Use this API command to Create Static Route.
//
// Request Body:
//	 - body *SwitchMStaticRouteCreateStaticRoute
func (s *SwitchMStaticRouteService) AddStaticRoutes(ctx context.Context, body *SwitchMStaticRouteCreateStaticRoute) (*SwitchMCommonCreateResult, error) {
	var (
		resp *SwitchMCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddStaticRoutes, true)
}

// DeleteStaticRoutes
//
// Use this API command to Delete Static Route by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMStaticRouteService) DeleteStaticRoutes(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteStaticRoutes, true)
}

// DeleteStaticRoutesById
//
// Use this API command to Delete Static Route.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteService) DeleteStaticRoutesById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteStaticRoutesById, true)
}

// FindStaticRoutesById
//
// Use this API command to Retrieve Static Route.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteService) FindStaticRoutesById(ctx context.Context, id string) (*SwitchMStaticRoute, error) {
	var (
		resp *SwitchMStaticRoute
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindStaticRoutesById, true)
}

// FindStaticRoutesByQueryCriteria
//
// Use this API command to Retrieve Static Route list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMStaticRouteService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMStaticRoutesQueryResult, error) {
	var (
		resp *SwitchMStaticRoutesQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMFindStaticRoutesByQueryCriteria, true)
}

// UpdateStaticRoutesById
//
// Use this API command to Update Static Route.
//
// Request Body:
//	 - body *SwitchMStaticRouteUpdateStaticRoute
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMStaticRouteService) UpdateStaticRoutesById(ctx context.Context, body *SwitchMStaticRouteUpdateStaticRoute, id string) (*SwitchMStaticRouteEmptyResult, error) {
	var (
		resp *SwitchMStaticRouteEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateStaticRoutesById, true)
}
