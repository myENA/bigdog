package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIComparisonReportService struct {
	apiClient *SCIClient
}

func NewSCIComparisonReportService(c *SCIClient) *SCIComparisonReportService {
	s := new(SCIComparisonReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIComparisonReportService() *SCIComparisonReportService {
	return NewSCIComparisonReportService(ss.apiClient)
}

// ReportComparisonReport140ComparisionOverview
//
// Operation ID: report.ComparisonReport.140.comparisionOverview
//
// Comparison Report - Overview
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport140ComparisionOverview(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport140comparisionOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport140comparisionOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportComparisonReport140ComparisionOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport140comparisionOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportComparisonReport145ComparisionMetric1
//
// Operation ID: report.ComparisonReport.145.comparisionMetric1
//
// Comparison Report - Metric 1 Over Time
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport145ComparisionMetric1(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport145comparisionMetric1200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport145comparisionMetric1200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportComparisonReport145ComparisionMetric1, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport145comparisionMetric1200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportComparisonReport146ComparisionMetric2
//
// Operation ID: report.ComparisonReport.146.comparisionMetric2
//
// Comparison Report - Metric 2 Over Time
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport146ComparisionMetric2(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport146comparisionMetric2200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport146comparisionMetric2200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportComparisonReport146ComparisionMetric2, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport146comparisionMetric2200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportComparisonReport147ComparisionTable
//
// Operation ID: report.ComparisonReport.147.comparisionTable
//
// Comparison Report - Table
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport147ComparisionTable(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport147comparisionTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport147comparisionTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportComparisonReport147ComparisionTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport147comparisionTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
