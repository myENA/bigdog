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
// Operation ID: report_ComparisonReport_140_comparisionOverview
//
// Comparison Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport140ComparisionOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport140ComparisionOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportComparisonReport140comparisionOverview200ResponseTypeAPIResponse), err
}

// ReportComparisonReport145ComparisionMetric1
//
// Operation ID: report_ComparisonReport_145_comparisionMetric1
//
// Comparison Report - Metric 1 Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport145ComparisionMetric1(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport145ComparisionMetric1, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportComparisonReport145comparisionMetric1200ResponseTypeAPIResponse), err
}

// ReportComparisonReport146ComparisionMetric2
//
// Operation ID: report_ComparisonReport_146_comparisionMetric2
//
// Comparison Report - Metric 2 Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport146ComparisionMetric2(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport146ComparisionMetric2, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportComparisonReport146comparisionMetric2200ResponseTypeAPIResponse), err
}

// ReportComparisonReport147ComparisionTable
//
// Operation ID: report_ComparisonReport_147_comparisionTable
//
// Comparison Report - Table
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIComparisonReportService) ReportComparisonReport147ComparisionTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport147ComparisionTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportComparisonReport147comparisionTable200ResponseTypeAPIResponse), err
}
