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
// Operation ID: report_SwitchDetailsReport_125_switchSummary
//
// Switch Details Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport125SwitchSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport125SwitchSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport125switchSummary200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport126SwitchResourceUtilization
//
// Operation ID: report_SwitchDetailsReport_126_switchResourceUtilization
//
// Switch Details Report - Resource Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport126SwitchResourceUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport126SwitchResourceUtilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport126switchResourceUtilization200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport129TopSwitchPortsByTrafficChart
//
// Operation ID: report_SwitchDetailsReport_129_topSwitchPortsByTrafficChart
//
// Switch Details Report - Top Ports by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport129TopSwitchPortsByTrafficChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport129TopSwitchPortsByTrafficChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport129topSwitchPortsByTrafficChart200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport130TopSwitchPortsByTrafficTable
//
// Operation ID: report_SwitchDetailsReport_130_topSwitchPortsByTrafficTable
//
// Switch Details Report - Top Ports by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport130TopSwitchPortsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport130TopSwitchPortsByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport130topSwitchPortsByTrafficTable200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport131SwitchTrafficTrend
//
// Operation ID: report_SwitchDetailsReport_131_switchTrafficTrend
//
// Switch Details Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport131SwitchTrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport131SwitchTrafficTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport131switchTrafficTrend200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport137LldpNeighborTable
//
// Operation ID: report_SwitchDetailsReport_137_lldpNeighborTable
//
// Switch Details Report - LLDP Neighbor List
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport137LldpNeighborTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport137LldpNeighborTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport137lldpNeighborTable200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport138SwitchUptimeHistory
//
// Operation ID: report_SwitchDetailsReport_138_switchUptimeHistory
//
// Switch Details Report - Uptime History
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport138SwitchUptimeHistory(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport138SwitchUptimeHistory, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport138switchUptimeHistory200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport139SwitchDetails
//
// Operation ID: report_SwitchDetailsReport_139_switchDetails
//
// Switch Details Report - Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport139SwitchDetails(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport139SwitchDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport139switchDetails200ResponseTypeAPIResponse), err
}

// ReportSwitchDetailsReport152PerSwitchErrorTrend
//
// Operation ID: report_SwitchDetailsReport_152_perSwitchErrorTrend
//
// Switch Details Report - Error Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIDetailsReportsService) ReportSwitchDetailsReport152PerSwitchErrorTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSwitchDetailsReport152PerSwitchErrorTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSwitchDetailsReport152perSwitchErrorTrend200ResponseTypeAPIResponse), err
}
