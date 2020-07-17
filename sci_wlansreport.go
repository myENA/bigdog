package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIWLANsReportService struct {
	apiClient *SCIClient
}

func NewSCIWLANsReportService(c *SCIClient) *SCIWLANsReportService {
	s := new(SCIWLANsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWLANsReportService() *SCIWLANsReportService {
	return NewSCIWLANsReportService(ss.apiClient)
}

// ReportWLANsReport35Overview
//
// Operation ID: report_WLANsReport_35_overview
//
// WLANs Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport35Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport35overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport35overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport35Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport35overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport36Top10SsidsByTraffic
//
// Operation ID: report_WLANsReport_36_top10SsidsByTraffic
//
// WLANs Report - Top SSIDs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport36Top10SsidsByTraffic(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport36top10SsidsByTraffic200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport36top10SsidsByTraffic200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport36Top10SsidsByTraffic, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport36top10SsidsByTraffic200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport37ActiveSsidsTrend
//
// Operation ID: report_WLANsReport_37_activeSsidsTrend
//
// WLANs Report - Active SSIDs Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport37ActiveSsidsTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport37activeSsidsTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport37activeSsidsTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport37ActiveSsidsTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport37activeSsidsTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport38Top10SsidsByClientCount
//
// Operation ID: report_WLANsReport_38_top10SsidsByClientCount
//
// WLANs Report - Top SSIDs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport38Top10SsidsByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport38top10SsidsByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport38top10SsidsByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport38Top10SsidsByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport38top10SsidsByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport39SsidChangesOverTime
//
// Operation ID: report_WLANsReport_39_ssidChangesOverTime
//
// WLANs Report - SSID Changes Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport39SsidChangesOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport39ssidChangesOverTime200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport39ssidChangesOverTime200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport39SsidChangesOverTime, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport39ssidChangesOverTime200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport40TopSsidsByTrafficTable
//
// Operation ID: report_WLANsReport_40_topSsidsByTrafficTable
//
// WLANs Report - Top SSIDs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport40TopSsidsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport40TopSsidsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport40topSsidsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport41TopSsidsByClientsTable
//
// Operation ID: report_WLANsReport_41_topSsidsByClientsTable
//
// WLANs Report - Top SSIDs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport41TopSsidsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport41topSsidsByClientsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport41topSsidsByClientsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport41TopSsidsByClientsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport41topSsidsByClientsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
