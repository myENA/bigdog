package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGEventandalarmService struct {
	apiClient *VSZClient
}

func NewWSGEventandalarmService(c *VSZClient) *WSGEventandalarmService {
	s := new(WSGEventandalarmService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEventandalarmService() *WSGEventandalarmService {
	return NewWSGEventandalarmService(ss.apiClient)
}

// AddAlertAlarmList
//
// Query Alarms with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandalarmService) AddAlertAlarmList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAlarmListAlarmQueryResultList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAlarmListAlarmQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAlertAlarmSummary
//
// Use this API command to retrieve the alarm summary with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandalarmService) AddAlertAlarmSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAlertsummaryAlarmSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAlertsummaryAlarmSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAlertAlarmSummary, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAlertsummaryAlarmSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAlertEventList
//
// Query Events with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandalarmService) AddAlertEventList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGEventListEventQueryResultList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEventListEventQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAlertEventSummary
//
// Use this API command to retrieve the event summary with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandalarmService) AddAlertEventSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAlertsummaryEventSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAlertsummaryEventSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAlertEventSummary, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAlertsummaryEventSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateAlertAlarmAck
//
// Acknowledge multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *WSGAlertAckBulkAlarms
func (s *WSGEventandalarmService) UpdateAlertAlarmAck(ctx context.Context, body *WSGAlertAckBulkAlarms) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateAlertAlarmAckByAlarmID
//
// Acknowledge a single Alarm with provided alarmID.
//
// Required Parameters:
// - alarmID string
//		- required
func (s *WSGEventandalarmService) UpdateAlertAlarmAckByAlarmID(ctx context.Context, alarmID string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAlertAlarmAckByAlarmID, true)
	req.SetPathParameter("alarmID", alarmID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateAlertAlarmClear
//
// Clear multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *WSGAlertClearBulkAlarms
func (s *WSGEventandalarmService) UpdateAlertAlarmClear(ctx context.Context, body *WSGAlertClearBulkAlarms) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateAlertAlarmClearByAlarmID
//
// Clear a single Alarm with provided alarmID.
//
// Required Parameters:
// - alarmID string
//		- required
func (s *WSGEventandalarmService) UpdateAlertAlarmClearByAlarmID(ctx context.Context, alarmID string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAlertAlarmClearByAlarmID, true)
	req.SetPathParameter("alarmID", alarmID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
