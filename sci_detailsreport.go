package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
	"time"
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
// Switch Details Report - Summary
//
// Operation ID: report_SwitchDetailsReport_125_switchSummary
// Operation path: /reports/18/sections/125/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport125SwitchSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport125SwitchSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport126SwitchResourceUtilization
//
// Switch Details Report - Resource Utilization
//
// Operation ID: report_SwitchDetailsReport_126_switchResourceUtilization
// Operation path: /reports/18/sections/126/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport126SwitchResourceUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport126SwitchResourceUtilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport129TopSwitchPortsByTrafficChart
//
// Switch Details Report - Top Ports by Traffic
//
// Operation ID: report_SwitchDetailsReport_129_topSwitchPortsByTrafficChart
// Operation path: /reports/18/sections/129/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport129TopSwitchPortsByTrafficChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport129TopSwitchPortsByTrafficChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport130TopSwitchPortsByTrafficTable
//
// Switch Details Report - Top Ports by Traffic
//
// Operation ID: report_SwitchDetailsReport_130_topSwitchPortsByTrafficTable
// Operation path: /reports/18/sections/130/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport130TopSwitchPortsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport130TopSwitchPortsByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport131SwitchTrafficTrend
//
// Switch Details Report - Traffic Trend
//
// Operation ID: report_SwitchDetailsReport_131_switchTrafficTrend
// Operation path: /reports/18/sections/131/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport131SwitchTrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport131SwitchTrafficTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport137LldpNeighborTable
//
// Switch Details Report - LLDP Neighbor List
//
// Operation ID: report_SwitchDetailsReport_137_lldpNeighborTable
// Operation path: /reports/18/sections/137/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport137LldpNeighborTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport137LldpNeighborTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport138SwitchUptimeHistory
//
// Switch Details Report - Uptime History
//
// Operation ID: report_SwitchDetailsReport_138_switchUptimeHistory
// Operation path: /reports/18/sections/138/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport138SwitchUptimeHistory(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport138SwitchUptimeHistory, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport139SwitchDetails
//
// Switch Details Report - Details
//
// Operation ID: report_SwitchDetailsReport_139_switchDetails
// Operation path: /reports/18/sections/139/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport139SwitchDetails(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport139SwitchDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport152PerSwitchErrorTrend
//
// Switch Details Report - Error Trend
//
// Operation ID: report_SwitchDetailsReport_152_perSwitchErrorTrend
// Operation path: /reports/18/sections/152/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport152PerSwitchErrorTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportSwitchDetailsReport152PerSwitchErrorTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse), err
}
