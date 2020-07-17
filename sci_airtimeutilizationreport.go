package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIAirtimeUtilizationReportService struct {
	apiClient *SCIClient
}

func NewSCIAirtimeUtilizationReportService(c *SCIClient) *SCIAirtimeUtilizationReportService {
	s := new(SCIAirtimeUtilizationReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAirtimeUtilizationReportService() *SCIAirtimeUtilizationReportService {
	return NewSCIAirtimeUtilizationReportService(ss.apiClient)
}

// ReportAirtimeUtilizationReport1Overview
//
// Operation ID: report_AirtimeUtilizationReport_1_overview
//
// Airtime Utilization Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport1Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport1overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport1overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport1Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport1overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport2TopChart
//
// Operation ID: report_AirtimeUtilizationReport_2_topChart
//
// Airtime Utilization Report - Top 10 APs by Airtime Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport2TopChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport2topChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport2topChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport2TopChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport2topChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport3TopAPsByAirtime24Table
//
// Operation ID: report_AirtimeUtilizationReport_3_topAPsByAirtime24Table
//
// Airtime Utilization Report - Top APs by Airtime Utilization for 2.4 GHz
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport3TopAPsByAirtime24Table(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport3TopAPsByAirtime24Table, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport4TopAPsByAirtime5Table
//
// Operation ID: report_AirtimeUtilizationReport_4_topAPsByAirtime5Table
//
// Airtime Utilization Report - Top APs by Airtime Utilization for 5 GHz
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport4TopAPsByAirtime5Table(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport4TopAPsByAirtime5Table, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport5TrendChart
//
// Operation ID: report_AirtimeUtilizationReport_5_trendChart
//
// Airtime Utilization Report - Airtime Utilization Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport5TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport5trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport5trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport5TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport5trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport6TrendTable
//
// Operation ID: report_AirtimeUtilizationReport_6_trendTable
//
// Airtime Utilization Report - Airtime Utilization Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport6TrendTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport6trendTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport6trendTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport6TrendTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport6trendTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
