package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIClientHealthReportService struct {
	apiClient *SCIClient
}

func NewSCIClientHealthReportService(c *SCIClient) *SCIClientHealthReportService {
	s := new(SCIClientHealthReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIClientHealthReportService() *SCIClientHealthReportService {
	return NewSCIClientHealthReportService(ss.apiClient)
}

// ReportClientHealthReport144ClientHealthSummary
//
// Operation ID: report_ClientHealthReport_144_clientHealthSummary
//
// Client Health Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport144ClientHealthSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport144clientHealthSummary200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport144clientHealthSummary200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport144ClientHealthSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport144clientHealthSummary200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientHealthReport148ClientConnectionHealth
//
// Operation ID: report_ClientHealthReport_148_clientConnectionHealth
//
// Client Health Report - Client Connection Health
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport148ClientConnectionHealth(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport148clientConnectionHealth200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport148clientConnectionHealth200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport148ClientConnectionHealth, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport148clientConnectionHealth200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientHealthReport149ClientHealthMetricTrends
//
// Operation ID: report_ClientHealthReport_149_clientHealthMetricTrends
//
// Client Health Report - Health Metric Trends
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport149ClientHealthMetricTrends(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport149ClientHealthMetricTrends, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport149clientHealthMetricTrends200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientHealthReport150TopClientHealthScoreByGroup
//
// Operation ID: report_ClientHealthReport_150_topClientHealthScoreByGroup
//
// Client Health Report - Health By Group
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport150TopClientHealthScoreByGroup(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport150TopClientHealthScoreByGroup, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
