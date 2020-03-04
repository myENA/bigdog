package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Sat
	// Schedules on Saturday
	Sat []string `json:"sat,omitempty"`

	// Sun
	// Schedules on Sunday
	Sun []string `json:"sun,omitempty"`

	// Thu
	// Schedules on Thursday
	Thu []string `json:"thu,omitempty"`

	// Tue
	// Schedules on Tuesday
	Tue []string `json:"tue,omitempty"`

	// Wed
	// Schedules on Wednesday
	Wed []string `json:"wed,omitempty"`
}

func NewWSGWLANSchedulerCreateWlanScheduler() *WSGWLANSchedulerCreateWlanScheduler {
	m := new(WSGWLANSchedulerCreateWlanScheduler)
	return m
}

type WSGWLANSchedulerModifyWlanScheduler struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	Sat []string `json:"sat,omitempty"`

	// Sun
	// Schedules on Sunday
	Sun []string `json:"sun,omitempty"`

	// Thu
	// Schedules on Thursday
	Thu []string `json:"thu,omitempty"`

	// Tue
	// Schedules on Tuesday
	Tue []string `json:"tue,omitempty"`

	// Wed
	// schedules on Wednesday
	Wed []string `json:"wed,omitempty"`
}

func NewWSGWLANSchedulerModifyWlanScheduler() *WSGWLANSchedulerModifyWlanScheduler {
	m := new(WSGWLANSchedulerModifyWlanScheduler)
	return m
}

type WSGWLANSchedulerWlanSchedule struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Id
	// Identifier of the WLAN schedule
	Id *string `json:"id,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	Sat []string `json:"sat,omitempty"`

	// Sun
	// Schedules on Sunday
	Sun []string `json:"sun,omitempty"`

	// Thu
	// Schedules on Thursday
	Thu []string `json:"thu,omitempty"`

	// Tue
	// Schedules on Tuesday
	Tue []string `json:"tue,omitempty"`

	// Wed
	// Schedules on Wednesday
	Wed []string `json:"wed,omitempty"`

	// ZoneId
	// Identifier of the zone to which the WLAN schedule belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGWLANSchedulerWlanSchedule() *WSGWLANSchedulerWlanSchedule {
	m := new(WSGWLANSchedulerWlanSchedule)
	return m
}

type WSGWLANSchedulerWlanScheduleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANSchedulerWlanSchedule `json:"list,omitempty"`

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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddRkszonesWlanSchedulersByZoneId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
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
func (s *WSGWLANSchedulerService) DeleteRkszonesWlanSchedulersById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteRkszonesWlanSchedulersById), true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesWlanSchedulersById), true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANSchedulerWlanSchedule()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesWlanSchedulersByZoneId), true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANSchedulerWlanScheduleList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindServicesWlanSchedulerByQueryCriteria), true)
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
func (s *WSGWLANSchedulerService) PartialUpdateRkszonesWlanSchedulersById(ctx context.Context, body *WSGWLANSchedulerModifyWlanScheduler, id string, zoneId string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateRkszonesWlanSchedulersById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
