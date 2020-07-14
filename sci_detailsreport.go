package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIDetailsReportsService struct {
	apiClient *SCIClient
}

func NewSCIDetailsReportsService(c *SCIClient) *SCIDetailsReportsService {
	s := new(SCIDetailsReportsService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIDetailsReportsService() *SCIDetailsReportsService {
	return NewSCIDetailsReportsService(ss.apiClient)
}

// ReportSwitchDetailsReport125SwitchSummary
//
// Operation ID: report.SwitchDetailsReport.125.switchSummary
//
// Switch Details Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport125SwitchSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport125switchSummary200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport125switchSummary200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport125SwitchSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport125switchSummary200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport126SwitchResourceUtilization
//
// Operation ID: report.SwitchDetailsReport.126.switchResourceUtilization
//
// Switch Details Report - Resource Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport126SwitchResourceUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport126SwitchResourceUtilization, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport126switchResourceUtilization200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport129TopSwitchPortsByTrafficChart
//
// Operation ID: report.SwitchDetailsReport.129.topSwitchPortsByTrafficChart
//
// Switch Details Report - Top Ports by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport129TopSwitchPortsByTrafficChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport129TopSwitchPortsByTrafficChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport130TopSwitchPortsByTrafficTable
//
// Operation ID: report.SwitchDetailsReport.130.topSwitchPortsByTrafficTable
//
// Switch Details Report - Top Ports by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport130TopSwitchPortsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport130TopSwitchPortsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport131SwitchTrafficTrend
//
// Operation ID: report.SwitchDetailsReport.131.switchTrafficTrend
//
// Switch Details Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport131SwitchTrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport131SwitchTrafficTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport131switchTrafficTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport137LldpNeighborTable
//
// Operation ID: report.SwitchDetailsReport.137.lldpNeighborTable
//
// Switch Details Report - LLDP Neighbor List
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport137LldpNeighborTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport137LldpNeighborTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport137lldpNeighborTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport138SwitchUptimeHistory
//
// Operation ID: report.SwitchDetailsReport.138.switchUptimeHistory
//
// Switch Details Report - Uptime History
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport138SwitchUptimeHistory(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport138SwitchUptimeHistory, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport138switchUptimeHistory200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSwitchDetailsReport139SwitchDetails
//
// Operation ID: report.SwitchDetailsReport.139.switchDetails
//
// Switch Details Report - Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport139SwitchDetails(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport139switchDetails200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSwitchDetailsReport139switchDetails200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSwitchDetailsReport139SwitchDetails, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSwitchDetailsReport139switchDetails200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
