package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGWLANSchedulerService struct {
	apiClient *VSZClient
}

func NewWSGWLANSchedulerService(c *VSZClient) *WSGWLANSchedulerService {
	s := new(WSGWLANSchedulerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWLANSchedulerService() *WSGWLANSchedulerService {
	return NewWSGWLANSchedulerService(ss.apiClient)
}

// WSGWLANSchedulerCreateWlanScheduler
//
// Definition: wlanscheduler_createWlanScheduler
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
	Name *WSGCommonNormalName `json:"name"`

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

// WSGWLANSchedulerModifyWlanScheduler
//
// Definition: wlanscheduler_modifyWlanScheduler
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

// WSGWLANSchedulerWlanSchedule
//
// Definition: wlanscheduler_wlanSchedule
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

type WSGWLANSchedulerWlanScheduleAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANSchedulerWlanSchedule
}

func newWSGWLANSchedulerWlanScheduleAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANSchedulerWlanScheduleAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANSchedulerWlanScheduleAPIResponse) Hydrate() error {
	r.Data = new(WSGWLANSchedulerWlanSchedule)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGWLANSchedulerWlanSchedule() *WSGWLANSchedulerWlanSchedule {
	m := new(WSGWLANSchedulerWlanSchedule)
	return m
}

// WSGWLANSchedulerWlanScheduleList
//
// Definition: wlanscheduler_wlanScheduleList
type WSGWLANSchedulerWlanScheduleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANSchedulerWlanSchedule `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGWLANSchedulerWlanScheduleListAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANSchedulerWlanScheduleList
}

func newWSGWLANSchedulerWlanScheduleListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANSchedulerWlanScheduleListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANSchedulerWlanScheduleListAPIResponse) Hydrate() error {
	r.Data = new(WSGWLANSchedulerWlanScheduleList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGWLANSchedulerWlanScheduleList() *WSGWLANSchedulerWlanScheduleList {
	m := new(WSGWLANSchedulerWlanScheduleList)
	return m
}

// WSGWLANSchedulerWlanScheduleQueryResultEntry
//
// Definition: wlanscheduler_wlanScheduleQueryResultEntry
type WSGWLANSchedulerWlanScheduleQueryResultEntry struct {
	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	// Id
	// This is seemingly always empty.  Leaving it here in case it suddenly becomes used...
	Id interface{} `json:"id,omitempty"`

	// Key
	// This is the actual ID of the WLAN Scheduler.  I don't know why.
	Key *string `json:"key,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	OptimisticLockVersion *int `json:"optimisticLockVersion,omitempty"`

	// ScheduleList
	// Comma-delimited scheduler time list
	ScheduleList *string `json:"scheduleList,omitempty"`

	// TenantId
	// This seems to always be the same value as tenantUUID.  I don't know why.
	TenantId *string `json:"tenantId,omitempty"`

	// TenantUUID
	// This seems to always be the same value as tenantId.  I don't know why.
	TenantUUID *string `json:"tenantUUID,omitempty"`

	// WlanIds
	// This seems to always be empty
	WlanIds []interface{} `json:"wlanIds,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`

	ZoneUUID *string `json:"zoneUUID,omitempty"`
}

func NewWSGWLANSchedulerWlanScheduleQueryResultEntry() *WSGWLANSchedulerWlanScheduleQueryResultEntry {
	m := new(WSGWLANSchedulerWlanScheduleQueryResultEntry)
	return m
}

// WSGWLANSchedulerWlanScheduleQueryResultList
//
// Definition: wlanscheduler_wlanScheduleQueryResultList
type WSGWLANSchedulerWlanScheduleQueryResultList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANSchedulerWlanScheduleQueryResultEntry `json:"list,omitempty"`

	// RawDataTotalCount
	// Seemingly always empty
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANSchedulerWlanScheduleQueryResultList
}

func newWSGWLANSchedulerWlanScheduleQueryResultListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse) Hydrate() error {
	r.Data = new(WSGWLANSchedulerWlanScheduleQueryResultList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGWLANSchedulerWlanScheduleQueryResultList() *WSGWLANSchedulerWlanScheduleQueryResultList {
	m := new(WSGWLANSchedulerWlanScheduleQueryResultList)
	return m
}

// AddRkszonesWlanSchedulersByZoneId
//
// Use this API command to create a new WLAN schedule.
//
// Operation ID: addRkszonesWlanSchedulersByZoneId
// Operation path: /rkszones/{zoneId}/wlanSchedulers
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANSchedulerCreateWlanScheduler
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) AddRkszonesWlanSchedulersByZoneId(ctx context.Context, body *WSGWLANSchedulerCreateWlanScheduler, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlanSchedulersByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesWlanSchedulersById
//
// Use this API command to delete a WLAN schedule.
//
// Operation ID: deleteRkszonesWlanSchedulersById
// Operation path: /rkszones/{zoneId}/wlanSchedulers/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) DeleteRkszonesWlanSchedulersById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlanSchedulersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesWlanSchedulersById
//
// Use this API command to retrieve a WLAN schedule.
//
// Operation ID: findRkszonesWlanSchedulersById
// Operation path: /rkszones/{zoneId}/wlanSchedulers/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGWLANSchedulerWlanScheduleAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANSchedulerWlanScheduleAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANSchedulerWlanScheduleAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesWlanSchedulersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANSchedulerWlanScheduleAPIResponse), err
}

// FindRkszonesWlanSchedulersByZoneId
//
// Use this API command to retrieve the list of WLAN schedule from a zone.
//
// Operation ID: findRkszonesWlanSchedulersByZoneId
// Operation path: /rkszones/{zoneId}/wlanSchedulers
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGWLANSchedulerWlanScheduleListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANSchedulerWlanScheduleListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANSchedulerWlanScheduleListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesWlanSchedulersByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANSchedulerWlanScheduleListAPIResponse), err
}

// FindServicesWlanSchedulerByQueryCriteria
//
// Query Wlan Schedulers with specified filters.
//
// Operation ID: findServicesWlanSchedulerByQueryCriteria
// Operation path: /query/services/wlanScheduler
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWLANSchedulerService) FindServicesWlanSchedulerByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANSchedulerWlanScheduleQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesWlanSchedulerByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANSchedulerWlanScheduleQueryResultListAPIResponse), err
}

// PartialUpdateRkszonesWlanSchedulersById
//
// Use this API command to modify the configuration of a WLAN schedule.
//
// Operation ID: partialUpdateRkszonesWlanSchedulersById
// Operation path: /rkszones/{zoneId}/wlanSchedulers/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGWLANSchedulerModifyWlanScheduler
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANSchedulerService) PartialUpdateRkszonesWlanSchedulersById(ctx context.Context, body *WSGWLANSchedulerModifyWlanScheduler, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlanSchedulersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
