package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIInventoryControllersReportService struct {
	apiClient *SCIClient
}

func NewSCIInventoryControllersReportService(c *SCIClient) *SCIInventoryControllersReportService {
	s := new(SCIInventoryControllersReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIInventoryControllersReportService() *SCIInventoryControllersReportService {
	return NewSCIInventoryControllersReportService(ss.apiClient)
}

// ReportInventoryControllersReport96Krack
//
// Operation ID: report_InventoryControllersReport_96_krack
//
// Inventory - Controllers Report - KRACK Assessment
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport96Krack(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport96krack200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport96krack200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport96Krack, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport96krack200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryControllersReport98ResourceUtilization
//
// Operation ID: report_InventoryControllersReport_98_resourceUtilization
//
// Inventory - Controllers Report - Resource Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport98ResourceUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport98resourceUtilization200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport98resourceUtilization200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport98ResourceUtilization, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport98resourceUtilization200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryControllersReport99LicenseUtilization
//
// Operation ID: report_InventoryControllersReport_99_licenseUtilization
//
// Inventory - Controllers Report - License Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport99LicenseUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport99licenseUtilization200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport99licenseUtilization200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport99LicenseUtilization, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport99licenseUtilization200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryControllersReport114ControllerInventoryOverview
//
// Operation ID: report_InventoryControllersReport_114_controllerInventoryOverview
//
// Inventory - Controllers Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport114ControllerInventoryOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport114ControllerInventoryOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
