package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCISessionsSummaryReportService struct {
	apiClient *SCIClient
}

func NewSCISessionsSummaryReportService(c *SCIClient) *SCISessionsSummaryReportService {
	s := new(SCISessionsSummaryReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCISessionsSummaryReportService() *SCISessionsSummaryReportService {
	return NewSCISessionsSummaryReportService(ss.apiClient)
}

// ReportSessionsSummaryReport33TopTable
//
// Operation ID: report_SessionsSummaryReport_33_topTable
//
// Sessions Summary Report - Top Sessions Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport33TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport33topTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSessionsSummaryReport33topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSessionsSummaryReport33TopTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSessionsSummaryReport33topTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSessionsSummaryReport34Overview
//
// Operation ID: report_SessionsSummaryReport_34_overview
//
// Sessions Summary Report - Average Durations
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport34Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport34overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSessionsSummaryReport34overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSessionsSummaryReport34Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSessionsSummaryReport34overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSessionsSummaryReport42DurationPercentile
//
// Operation ID: report_SessionsSummaryReport_42_durationPercentile
//
// Sessions Summary Report - Session Duration Percentiles
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport42DurationPercentile(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport42durationPercentile200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSessionsSummaryReport42durationPercentile200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSessionsSummaryReport42DurationPercentile, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSessionsSummaryReport42durationPercentile200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
