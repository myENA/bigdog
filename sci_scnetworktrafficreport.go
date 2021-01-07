package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCISCNetworkTrafficReportService struct {
	apiClient *SCIClient
}

func NewSCISCNetworkTrafficReportService(c *SCIClient) *SCISCNetworkTrafficReportService {
	s := new(SCISCNetworkTrafficReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCISCNetworkTrafficReportService() *SCISCNetworkTrafficReportService {
	return NewSCISCNetworkTrafficReportService(ss.apiClient)
}

// ReportSCNetworkTrafficReport93ScNetworkTraffic
//
// Operation ID: report_SCNetworkTrafficReport_93_scNetworkTraffic
//
// SC Network Traffic Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport93ScNetworkTraffic(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport93ScNetworkTraffic, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport94ScNetworkTrend
//
// Operation ID: report_SCNetworkTrafficReport_94_scNetworkTrend
//
// SC Network Traffic Report - SmartCell Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport94ScNetworkTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport94ScNetworkTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport100DroppedCallRate
//
// Operation ID: report_SCNetworkTrafficReport_100_droppedCallRate
//
// SC Network Traffic Report - SmartCell Dropped Call Rate
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport100DroppedCallRate(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport100DroppedCallRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport101ConnectionSetupSuccessRate
//
// Operation ID: report_SCNetworkTrafficReport_101_connectionSetupSuccessRate
//
// SC Network Traffic Report - SmartCell Connection Setup Success Rate
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport101ConnectionSetupSuccessRate(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport101ConnectionSetupSuccessRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport102HandoverSuccessRate
//
// Operation ID: report_SCNetworkTrafficReport_102_handoverSuccessRate
//
// SC Network Traffic Report - SmartCell Handover Success Rate
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport102HandoverSuccessRate(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport102HandoverSuccessRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport103AvgThroughput
//
// Operation ID: report_SCNetworkTrafficReport_103_avgThroughput
//
// SC Network Traffic Report - SmartCell Average Throughput
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport103AvgThroughput(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport103AvgThroughput, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport103avgThroughput200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport104ScAvailability
//
// Operation ID: report_SCNetworkTrafficReport_104_scAvailability
//
// SC Network Traffic Report - SmartCell Availability
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport104ScAvailability(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport104ScAvailability, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport104scAvailability200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport105RscConnectionStats
//
// Operation ID: report_SCNetworkTrafficReport_105_rscConnectionStats
//
// SC Network Traffic Report - SmartCell Connection Statistics
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport105RscConnectionStats(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport105RscConnectionStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport106RscGpsStats
//
// Operation ID: report_SCNetworkTrafficReport_106_rscGpsStats
//
// SC Network Traffic Report - RSC GPS Statistics
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport106RscGpsStats(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport106RscGpsStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport107TrafficVolume
//
// Operation ID: report_SCNetworkTrafficReport_107_trafficVolume
//
// SC Network Traffic Report - SmartCell Traffic Volume
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport107TrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport107TrafficVolume, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport107trafficVolume200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport108PhaseSyncLoss
//
// Operation ID: report_SCNetworkTrafficReport_108_phaseSyncLoss
//
// SC Network Traffic Report - SmartCell Phase Sync Loss
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport108PhaseSyncLoss(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport108PhaseSyncLoss, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport109FrequencySyncLoss
//
// Operation ID: report_SCNetworkTrafficReport_109_frequencySyncLoss
//
// SC Network Traffic Report - SmartCell Frequency Sync Loss
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport109FrequencySyncLoss(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport109FrequencySyncLoss, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseTypeAPIResponse), err
}

// ReportSCNetworkTrafficReport111RscTrafficStats
//
// Operation ID: report_SCNetworkTrafficReport_111_rscTrafficStats
//
// SC Network Traffic Report - RSC Traffic Statistics
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport111RscTrafficStats(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport111RscTrafficStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseTypeAPIResponse), err
}
