package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIWirelessClientsReportService struct {
	apiClient *SCIClient
}

func NewSCIWirelessClientsReportService(c *SCIClient) *SCIWirelessClientsReportService {
	s := new(SCIWirelessClientsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWirelessClientsReportService() *SCIWirelessClientsReportService {
	return NewSCIWirelessClientsReportService(ss.apiClient)
}

// ReportWirelessClientsReport12Overview
//
// Operation ID: report_WirelessClientsReport_12_overview
//
// Wireless - Clients Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport12Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport12overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport12overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport12Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport12overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport13TopChart
//
// Operation ID: report_WirelessClientsReport_13_topChart
//
// Wireless - Clients Report - Top 10 Unique Clients by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport13TopChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport13topChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport13topChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport13TopChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport13topChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport14TopTable
//
// Operation ID: report_WirelessClientsReport_14_topTable
//
// Wireless - Clients Report - Clients Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport14TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport14topTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport14topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport14TopTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport14topTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport15TrendChart
//
// Operation ID: report_WirelessClientsReport_15_trendChart
//
// Wireless - Clients Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport15TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport15trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport15trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport15TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport15trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport16TrendTable
//
// Operation ID: report_WirelessClientsReport_16_trendTable
//
// Wireless - Clients Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport16TrendTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport16trendTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport16trendTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport16TrendTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport16trendTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport17TopPercentile
//
// Operation ID: report_WirelessClientsReport_17_topPercentile
//
// Wireless - Clients Report - Top Clients by Traffic Percentile
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport17TopPercentile(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport17topPercentile200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport17topPercentile200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport17TopPercentile, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport17topPercentile200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport18TopNOSByClientCount
//
// Operation ID: report_WirelessClientsReport_18_topNOSByClientCount
//
// Wireless - Clients Report - Top 10 OS by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport18TopNOSByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport18TopNOSByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport18topNOSByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport19Top10ManufacturersByClientCount
//
// Operation ID: report_WirelessClientsReport_19_top10ManufacturersByClientCount
//
// Wireless - Clients Report - Top 10 Manufacturers by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport19Top10ManufacturersByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport19Top10ManufacturersByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount
//
// Operation ID: report_WirelessClientsReport_112_top10AuthenticationMechanismByClientCount
//
// Wireless - Clients Report - Top 10 Authentication Methods by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
