package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCINetworkWirelessReportService struct {
	apiClient *SCIClient
}

func NewSCINetworkWirelessReportService(c *SCIClient) *SCINetworkWirelessReportService {
	s := new(SCINetworkWirelessReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCINetworkWirelessReportService() *SCINetworkWirelessReportService {
	return NewSCINetworkWirelessReportService(ss.apiClient)
}

// ReportNetworkWirelessReport20Overview
//
// Operation ID: report_NetworkWirelessReport_20_overview
//
// Network - Wireless Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport20Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport20overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport20overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport20Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport20overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport21TrafficDistribution
//
// Operation ID: report_NetworkWirelessReport_21_trafficDistribution
//
// Network - Wireless Report - Traffic Distribution
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport21TrafficDistribution(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport21trafficDistribution200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport21trafficDistribution200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport21TrafficDistribution, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport21trafficDistribution200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport22TrafficTrend
//
// Operation ID: report_NetworkWirelessReport_22_trafficTrend
//
// Network - Wireless Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport22TrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport22trafficTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport22trafficTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport22TrafficTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport22trafficTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport23TrafficOverTimeTable
//
// Operation ID: report_NetworkWirelessReport_23_trafficOverTimeTable
//
// Network - Wireless Report - Traffic Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport23TrafficOverTimeTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport23TrafficOverTimeTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport23trafficOverTimeTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport24TopAPsByTrafficTable
//
// Operation ID: report_NetworkWirelessReport_24_topAPsByTrafficTable
//
// Network - Wireless Report - Top APs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport24TopAPsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport24TopAPsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport24topAPsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport25TopAPsByClientsTable
//
// Operation ID: report_NetworkWirelessReport_25_topAPsByClientsTable
//
// Network - Wireless Report - Top APs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport25TopAPsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport25TopAPsByClientsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport25topAPsByClientsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport26Top10APsByTrafficVolume
//
// Operation ID: report_NetworkWirelessReport_26_top10APsByTrafficVolume
//
// Network - Wireless Report - Top APs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport26Top10APsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport26Top10APsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport26top10APsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportNetworkWirelessReport27Top10ApByClientCount
//
// Operation ID: report_NetworkWirelessReport_27_top10ApByClientCount
//
// Network - Wireless Report - Top APs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCINetworkWirelessReportService) ReportNetworkWirelessReport27Top10ApByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportNetworkWirelessReport27Top10ApByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportNetworkWirelessReport27top10ApByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
