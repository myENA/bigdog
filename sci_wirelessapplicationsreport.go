package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIWirelessApplicationsReportService struct {
	apiClient *SCIClient
}

func NewSCIWirelessApplicationsReportService(c *SCIClient) *SCIWirelessApplicationsReportService {
	s := new(SCIWirelessApplicationsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWirelessApplicationsReportService() *SCIWirelessApplicationsReportService {
	return NewSCIWirelessApplicationsReportService(ss.apiClient)
}

// ReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume
//
// Operation ID: report_WirelessApplicationsReport_7_top10ApplicationsByTrafficVolume
//
// Wireless - Applications Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport8TopAppsByTrafficTable
//
// Operation ID: report_WirelessApplicationsReport_8_topAppsByTrafficTable
//
// Wireless - Applications Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport8TopAppsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport9TopAppsByClientsTable
//
// Operation ID: report_WirelessApplicationsReport_9_topAppsByClientsTable
//
// Wireless - Applications Report - Top Applications by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport9TopAppsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport9TopAppsByClientsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport10Overview
//
// Operation ID: report_WirelessApplicationsReport_10_overview
//
// Wireless - Applications Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport10Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport10overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport10overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport10Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport10overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport11Top10ApplicationsByClientCount
//
// Operation ID: report_WirelessApplicationsReport_11_top10ApplicationsByClientCount
//
// Wireless - Applications Report - Top Applications by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport11Top10ApplicationsByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport11Top10ApplicationsByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
