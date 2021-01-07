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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport93ScNetworkTraffic, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport94ScNetworkTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport100DroppedCallRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport101ConnectionSetupSuccessRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport102HandoverSuccessRate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport103AvgThroughput, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport103avgThroughput200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport104scAvailability200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport104ScAvailability, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport104scAvailability200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport105RscConnectionStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport106RscGpsStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport107TrafficVolume, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport107trafficVolume200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport108PhaseSyncLoss, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport109FrequencySyncLoss, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
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
		rm       *APIResponseMeta
		resp     *SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSCNetworkTrafficReport111RscTrafficStats, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
