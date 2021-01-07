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
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport140comparisionOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport140ComparisionOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport140comparisionOverview200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport145comparisionMetric1200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport145ComparisionMetric1, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport145comparisionMetric1200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport146comparisionMetric2200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport146ComparisionMetric2, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport146comparisionMetric2200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportComparisonReport147comparisionTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportComparisonReport147ComparisionTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportComparisonReport147comparisionTable200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
