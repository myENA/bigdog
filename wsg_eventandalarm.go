package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGEventAndAlarmService struct {
	apiClient *VSZClient
}

func NewWSGEventAndAlarmService(c *VSZClient) *WSGEventAndAlarmService {
	s := new(WSGEventAndAlarmService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEventAndAlarmService() *WSGEventAndAlarmService {
	return NewWSGEventAndAlarmService(ss.apiClient)
}

// AddAlertAlarmList
//
// Query Alarms with specified filters.
//
// Operation ID: addAlertAlarmList
// Operation path: /alert/alarm/list
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertAlarmList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAlarmListAlarmQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAlarmListAlarmQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAlarmListAlarmQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAlertAlarmList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAlarmListAlarmQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAlarmListAlarmQueryResultListAPIResponse), err
}

// AddAlertAlarmSummary
//
// Use this API command to retrieve the alarm summary with specified filters.
//
// Operation ID: addAlertAlarmSummary
// Operation path: /alert/alarmSummary
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertAlarmSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAlertSummaryAlarmSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAlertSummaryAlarmSummaryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAlertSummaryAlarmSummaryAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAlertAlarmSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAlertSummaryAlarmSummaryAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAlertSummaryAlarmSummaryAPIResponse), err
}

// AddAlertEventList
//
// Query Events with specified filters.
//
// Operation ID: addAlertEventList
// Operation path: /alert/event/list
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertEventList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGEventListEventQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGEventListEventQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGEventListEventQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAlertEventList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGEventListEventQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGEventListEventQueryResultListAPIResponse), err
}

// AddAlertEventSummary
//
// Use this API command to retrieve the event summary with specified filters.
//
// Operation ID: addAlertEventSummary
// Operation path: /alert/eventSummary
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertEventSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAlertSummaryEventSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAlertSummaryEventSummaryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAlertSummaryEventSummaryAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAlertEventSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAlertSummaryEventSummaryAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAlertSummaryEventSummaryAPIResponse), err
}

// UpdateAlertAlarmAck
//
// Acknowledge multiple Alarms with provided alarmIDs.
//
// Operation ID: updateAlertAlarmAck
// Operation path: /alert/alarm/ack
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAlertAckBulkAlarms
func (s *WSGEventAndAlarmService) UpdateAlertAlarmAck(ctx context.Context, body *WSGAlertAckBulkAlarms, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateAlertAlarmAck, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateAlertAlarmAckByAlarmID
//
// Acknowledge a single Alarm with provided alarmID.
//
// Operation ID: updateAlertAlarmAckByAlarmID
// Operation path: /alert/alarm/{alarmID}/ack
// Success code: 200 (OK)
//
// Required parameters:
// - alarmID string
//		- required
func (s *WSGEventAndAlarmService) UpdateAlertAlarmAckByAlarmID(ctx context.Context, alarmID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateAlertAlarmAckByAlarmID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("alarmID", alarmID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateAlertAlarmClear
//
// Clear multiple Alarms with provided alarmIDs.
//
// Operation ID: updateAlertAlarmClear
// Operation path: /alert/alarm/clear
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAlertClearBulkAlarms
func (s *WSGEventAndAlarmService) UpdateAlertAlarmClear(ctx context.Context, body *WSGAlertClearBulkAlarms, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateAlertAlarmClear, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateAlertAlarmClearByAlarmID
//
// Clear a single Alarm with provided alarmID.
//
// Operation ID: updateAlertAlarmClearByAlarmID
// Operation path: /alert/alarm/{alarmID}/clear
// Success code: 200 (OK)
//
// Required parameters:
// - alarmID string
//		- required
func (s *WSGEventAndAlarmService) UpdateAlertAlarmClearByAlarmID(ctx context.Context, alarmID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateAlertAlarmClearByAlarmID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("alarmID", alarmID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
