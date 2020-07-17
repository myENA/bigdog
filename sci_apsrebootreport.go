package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
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
// Operation ID: report_APsRebootReport_43_totalReboots
//
// APs Reboot Report - Total Reboots
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport43TotalReboots(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport43totalReboots200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPsRebootReport43totalReboots200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPsRebootReport43TotalReboots, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport43totalReboots200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPsRebootReport44TopApRebootsTable
//
// Operation ID: report_APsRebootReport_44_topApRebootsTable
//
// APs Reboot Report - Top AP Reboots
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport44TopApRebootsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport44topApRebootsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPsRebootReport44topApRebootsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPsRebootReport44TopApRebootsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport44topApRebootsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPsRebootReport45TopApRebootsOverTime
//
// Operation ID: report_APsRebootReport_45_topApRebootsOverTime
//
// APs Reboot Report - AP Reboots
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport45TopApRebootsOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPsRebootReport45TopApRebootsOverTime, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport45topApRebootsOverTime200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
