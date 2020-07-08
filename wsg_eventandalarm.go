package bigdog

// API Version: v9_0

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
// Operation ID: addAlertAlarmList
//
// Query Alarms with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertAlarmList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAlarmListAlarmQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAlarmListAlarmQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAlertAlarmList, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAlarmListAlarmQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAlertAlarmSummary
//
// Operation ID: addAlertAlarmSummary
//
// Use this API command to retrieve the alarm summary with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertAlarmSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAlertSummaryAlarmSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAlertSummaryAlarmSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAlertAlarmSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAlertSummaryAlarmSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAlertEventList
//
// Operation ID: addAlertEventList
//
// Query Events with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertEventList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGEventListEventQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGEventListEventQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAlertEventList, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGEventListEventQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAlertEventSummary
//
// Operation ID: addAlertEventSummary
//
// Use this API command to retrieve the event summary with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventAndAlarmService) AddAlertEventSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAlertSummaryEventSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAlertSummaryEventSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAlertEventSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAlertSummaryEventSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateAlertAlarmAck
//
// Operation ID: updateAlertAlarmAck
//
// Acknowledge multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *WSGAlertAckBulkAlarms
func (s *WSGEventAndAlarmService) UpdateAlertAlarmAck(ctx context.Context, body *WSGAlertAckBulkAlarms, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAlertAlarmAck, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateAlertAlarmAckByAlarmID
//
// Operation ID: updateAlertAlarmAckByAlarmID
//
// Acknowledge a single Alarm with provided alarmID.
//
// Required Parameters:
// - alarmID string
//		- required
func (s *WSGEventAndAlarmService) UpdateAlertAlarmAckByAlarmID(ctx context.Context, alarmID string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAlertAlarmAckByAlarmID, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("alarmID", alarmID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateAlertAlarmClear
//
// Operation ID: updateAlertAlarmClear
//
// Clear multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *WSGAlertClearBulkAlarms
func (s *WSGEventAndAlarmService) UpdateAlertAlarmClear(ctx context.Context, body *WSGAlertClearBulkAlarms, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAlertAlarmClear, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateAlertAlarmClearByAlarmID
//
// Operation ID: updateAlertAlarmClearByAlarmID
//
// Clear a single Alarm with provided alarmID.
//
// Required Parameters:
// - alarmID string
//		- required
func (s *WSGEventAndAlarmService) UpdateAlertAlarmClearByAlarmID(ctx context.Context, alarmID string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAlertAlarmClearByAlarmID, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("alarmID", alarmID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
