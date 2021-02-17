package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
	"time"
)

type SCIAPsRebootReportService struct {
	apiClient *SCIClient
}

func NewSCIAPsRebootReportService(c *SCIClient) *SCIAPsRebootReportService {
	s := new(SCIAPsRebootReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAPsRebootReportService() *SCIAPsRebootReportService {
	return NewSCIAPsRebootReportService(ss.apiClient)
}

// ReportAPsRebootReport43TotalReboots
//
// APs Reboot Report - Total Reboots
//
// Operation ID: report_APsRebootReport_43_totalReboots
// Operation path: /reports/8/sections/43/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport43TotalReboots(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAPsRebootReport43TotalReboots, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse), err
}

// ReportAPsRebootReport44TopApRebootsTable
//
// APs Reboot Report - Top AP Reboots
//
// Operation ID: report_APsRebootReport_44_topApRebootsTable
// Operation path: /reports/8/sections/44/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport44TopApRebootsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAPsRebootReport44TopApRebootsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse), err
}

// ReportAPsRebootReport45TopApRebootsOverTime
//
// APs Reboot Report - AP Reboots
//
// Operation ID: report_APsRebootReport_45_topApRebootsOverTime
// Operation path: /reports/8/sections/45/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport45TopApRebootsOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAPsRebootReport45TopApRebootsOverTime, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse), err
}
