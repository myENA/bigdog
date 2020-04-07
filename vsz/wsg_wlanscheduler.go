package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWLANSchedulerService struct {
	apiClient *APIClient
}

func NewWSGWLANSchedulerService(c *APIClient) *WSGWLANSchedulerService {
	s := new(WSGWLANSchedulerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWLANSchedulerService() *WSGWLANSchedulerService {
	return NewWSGWLANSchedulerService(ss.apiClient)
}

type WSGWLANSchedulerCreateWlanScheduler struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	// Constraints:
	//    - nullable
	Fri []string `json:"fri,omitempty" validate:"omitempty,dive"`

	// Mon
	// Schedules on Monday
	// Constraints:
	//    - nullable
	Mon []string `json:"mon,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Sat
	// Schedules on Saturday
	// Constraints:
	//    - nullable
	Sat []string `json:"sat,omitempty" validate:"omitempty,dive"`

	// Sun
	// Schedules on Sunday
	// Constraints:
	//    - nullable
	Sun []string `json:"sun,omitempty" validate:"omitempty,dive"`

	// Thu
	// Schedules on Thursday
	// Constraints:
	//    - nullable
	Thu []string `json:"thu,omitempty" validate:"omitempty,dive"`

	// Tue
	// Schedules on Tuesday
	// Constraints:
	//    - nullable
	Tue []string `json:"tue,omitempty" validate:"omitempty,dive"`

	// Wed
	// Schedules on Wednesday
	// Constraints:
	//    - nullable
	Wed []string `json:"wed,omitempty" validate:"omitempty,dive"`
}

func NewWSGWLANSchedulerCreateWlanScheduler() *WSGWLANSchedulerCreateWlanScheduler {
	m := new(WSGWLANSchedulerCreateWlanScheduler)
	return m
}

type WSGWLANSchedulerModifyWlanScheduler struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	// Constraints:
	//    - nullable
	Fri []string `json:"fri,omitempty" validate:"omitempty,dive"`

	// Mon
	// Schedules on Monday
	// Constraints:
	//    - nullable
	Mon []string `json:"mon,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	// Constraints:
	//    - nullable
	Sat []string `json:"sat,omitempty" validate:"omitempty,dive"`

	// Sun
	// Schedules on Sunday
	// Constraints:
	//    - nullable
	Sun []string `json:"sun,omitempty" validate:"omitempty,dive"`

	// Thu
	// Schedules on Thursday
	// Constraints:
	//    - nullable
	Thu []string `json:"thu,omitempty" validate:"omitempty,dive"`

	// Tue
	// Schedules on Tuesday
	// Constraints:
	//    - nullable
	Tue []string `json:"tue,omitempty" validate:"omitempty,dive"`

	// Wed
	// schedules on Wednesday
	// Constraints:
	//    - nullable
	Wed []string `json:"wed,omitempty" validate:"omitempty,dive"`
}

func NewWSGWLANSchedulerModifyWlanScheduler() *WSGWLANSchedulerModifyWlanScheduler {
	m := new(WSGWLANSchedulerModifyWlanScheduler)
	return m
}

type WSGWLANSchedulerWlanSchedule struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	// Constraints:
	//    - nullable
	Fri []string `json:"fri,omitempty" validate:"omitempty,dive"`

	// Id
	// Identifier of the WLAN schedule
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Mon
	// Schedules on Monday
	// Constraints:
	//    - nullable
	Mon []string `json:"mon,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	// Constraints:
	//    - nullable
	Sat []string `json:"sat,omitempty" validate:"omitempty,dive"`

	// Sun
	// Schedules on Sunday
	// Constraints:
	//    - nullable
	Sun []string `json:"sun,omitempty" validate:"omitempty,dive"`

	// Thu
	// Schedules on Thursday
	// Constraints:
	//    - nullable
	Thu []string `json:"thu,omitempty" validate:"omitempty,dive"`

	// Tue
	// Schedules on Tuesday
	// Constraints:
	//    - nullable
	Tue []string `json:"tue,omitempty" validate:"omitempty,dive"`

	// Wed
	// Schedules on Wednesday
	// Constraints:
	//    - nullable
	Wed []string `json:"wed,omitempty" validate:"omitempty,dive"`

	// ZoneId
	// Identifier of the zone to which the WLAN schedule belongs
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGWLANSchedulerWlanSchedule() *WSGWLANSchedulerWlanSchedule {
	m := new(WSGWLANSchedulerWlanSchedule)
	return m
}

type WSGWLANSchedulerWlanScheduleList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGWLANSchedulerWlanSchedule `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANSchedulerWlanScheduleList() *WSGWLANSchedulerWlanScheduleList {
	m := new(WSGWLANSchedulerWlanScheduleList)
	return m
}

// AddRkszonesWlanSchedulersByZoneId
//
// Use this API command to create a new WLAN schedule.
//
// Request Body:
//	 - body *WSGWLANSchedulerCreateWlanScheduler
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) AddRkszonesWlanSchedulersByZoneId(ctx context.Context, body *WSGWLANSchedulerCreateWlanScheduler, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesWlanSchedulersByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesWlanSchedulersById
//
// Use this API command to delete a WLAN schedule.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) DeleteRkszonesWlanSchedulersById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlanSchedulersById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesWlanSchedulersById
//
// Use this API command to retrieve a WLAN schedule.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersById(ctx context.Context, id string, zoneId string) (*WSGWLANSchedulerWlanSchedule, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANSchedulerWlanSchedule
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlanSchedulersById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANSchedulerWlanSchedule()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesWlanSchedulersByZoneId
//
// Use this API command to retrieve the list of WLAN schedule from a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGWLANSchedulerWlanScheduleList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANSchedulerWlanScheduleList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesWlanSchedulersByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANSchedulerWlanScheduleList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesWlanSchedulerByQueryCriteria
//
// Query Wlan Schedulers with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWLANSchedulerService) FindServicesWlanSchedulerByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesWlanSchedulerByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesWlanSchedulersById
//
// Use this API command to modify the configuration of a WLAN schedule.
//
// Request Body:
//	 - body *WSGWLANSchedulerModifyWlanScheduler
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) PartialUpdateRkszonesWlanSchedulersById(ctx context.Context, body *WSGWLANSchedulerModifyWlanScheduler, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlanSchedulersById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
