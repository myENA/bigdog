package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIInventorySwitchesReportService struct {
	apiClient *SCIClient
}

func NewSCIInventorySwitchesReportService(c *SCIClient) *SCIInventorySwitchesReportService {
	s := new(SCIInventorySwitchesReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIInventorySwitchesReportService() *SCIInventorySwitchesReportService {
	return NewSCIInventorySwitchesReportService(ss.apiClient)
}

// ReportInventorySwitchesReport113Overview
//
// Operation ID: report_InventorySwitchesReport_113_overview
//
// Inventory - Switches Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport113Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport113overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport113overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport113Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport113overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventorySwitchesReport116SwitchCountTrend
//
// Operation ID: report_InventorySwitchesReport_116_switchCountTrend
//
// Inventory - Switches Report - Switch Count Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport116SwitchCountTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport116switchCountTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport116switchCountTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport116SwitchCountTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport116switchCountTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventorySwitchesReport117Top10SwitchVersionChart
//
// Operation ID: report_InventorySwitchesReport_117_top10SwitchVersionChart
//
// Inventory - Switches Report - Top Switch Software Versions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport117Top10SwitchVersionChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport117Top10SwitchVersionChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventorySwitchesReport118TopSwitchVersions
//
// Operation ID: report_InventorySwitchesReport_118_topSwitchVersions
//
// Inventory - Switches Report - Top Switch Software Versions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport118TopSwitchVersions(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport118topSwitchVersions200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport118TopSwitchVersions, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport118topSwitchVersions200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventorySwitchesReport121TopSwitchModelsChart
//
// Operation ID: report_InventorySwitchesReport_121_topSwitchModelsChart
//
// Inventory - Switches Report - Top Switch Models
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport121TopSwitchModelsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport121TopSwitchModelsChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventorySwitchesReport122TopSwitchModels
//
// Operation ID: report_InventorySwitchesReport_122_topSwitchModels
//
// Inventory - Switches Report - Top Switch Models
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport122TopSwitchModels(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport122topSwitchModels200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport122topSwitchModels200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport122TopSwitchModels, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport122topSwitchModels200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventorySwitchesReport132PortStatusTrend
//
// Operation ID: report_InventorySwitchesReport_132_portStatusTrend
//
// Inventory - Switches Report - Port Status Trends
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport132PortStatusTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport132portStatusTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventorySwitchesReport132portStatusTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventorySwitchesReport132PortStatusTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventorySwitchesReport132portStatusTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
