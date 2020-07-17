package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCINetworkWiredReportService struct {
	apiClient *SCIClient
}

func NewSCINetworkWiredReportService(c *SCIClient) *SCINetworkWiredReportService {
	s := new(SCINetworkWiredReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCINetworkWiredReportService() *SCINetworkWiredReportService {
	return NewSCINetworkWiredReportService(ss.apiClient)
}

// ReportNetworkWiredReport123TopSwitchPOEUtilChart
//
// Operation ID: report_NetworkWiredReport_123_topSwitchPOEUtilChart
//
// Network - Wired Report - Top Switches by PoE Usage
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport123TopSwitchPOEUtilChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport123TopSwitchPOEUtilChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport123topSwitchPOEUtilChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport124TopSwitchPOEUtils
//
// Operation ID: report_NetworkWiredReport_124_topSwitchPOEUtils
//
// Network - Wired Report - Top Switches by PoE Usage
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport124TopSwitchPOEUtils(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport124TopSwitchPOEUtils, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport124topSwitchPOEUtils200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport127Top10SwitchesByTrafficVolume
//
// Operation ID: report_NetworkWiredReport_127_top10SwitchesByTrafficVolume
//
// Network - Wired Report - Top Switches by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport127Top10SwitchesByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport127Top10SwitchesByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport127top10SwitchesByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport128TopSwitchesByTrafficTable
//
// Operation ID: report_NetworkWiredReport_128_topSwitchesByTrafficTable
//
// Network - Wired Report - Top Switches by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport128TopSwitchesByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport128TopSwitchesByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport128topSwitchesByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport134WiredOverview
//
// Operation ID: report_NetworkWiredReport_134_wiredOverview
//
// Network - Wired Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport134WiredOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport134wiredOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport134wiredOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport134WiredOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport134wiredOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport135WiredTrafficDistribution
//
// Operation ID: report_NetworkWiredReport_135_wiredTrafficDistribution
//
// Network - Wired Report - Traffic Distribution by Switch Model and Port Speed
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport135WiredTrafficDistribution(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport135WiredTrafficDistribution, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport135wiredTrafficDistribution200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport136SwitchTrafficTrend
//
// Operation ID: report_NetworkWiredReport_136_switchTrafficTrend
//
// Network - Wired Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport136SwitchTrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport136switchTrafficTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport136SwitchTrafficTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport136switchTrafficTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport141SwitchErrorTrend
//
// Operation ID: report_NetworkWiredReport_141_switchErrorTrend
//
// Network - Wired Report - Error Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport141SwitchErrorTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport141switchErrorTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport141switchErrorTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport141SwitchErrorTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport141switchErrorTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport142TopSwitchesByErrorsChart
//
// Operation ID: report_NetworkWiredReport_142_topSwitchesByErrorsChart
//
// Network - Wired Report - Top Switches by Errors
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport142TopSwitchesByErrorsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport142TopSwitchesByErrorsChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport142topSwitchesByErrorsChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWiredReport143TopSwitchesByErrorsTable
//
// Operation ID: report_NetworkWiredReport_143_topSwitchesByErrorsTable
//
// Network - Wired Report - Top Switches by Errors
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWiredReportService) ReportNetworkWiredReport143TopSwitchesByErrorsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWiredReport143TopSwitchesByErrorsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWiredReport143topSwitchesByErrorsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
