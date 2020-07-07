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

// SCISCNetworkTrafficReport93scNetworkTrafficData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.93.scNetworkTraffic.Data
type SCISCNetworkTrafficReport93scNetworkTrafficData []interface{}

func MakeSCISCNetworkTrafficReport93scNetworkTrafficData() SCISCNetworkTrafficReport93scNetworkTrafficData {
	m := make(SCISCNetworkTrafficReport93scNetworkTrafficData, 0)
	return m
}

// SCISCNetworkTrafficReport94scNetworkTrendData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.94.scNetworkTrend.Data
type SCISCNetworkTrafficReport94scNetworkTrendData []interface{}

func MakeSCISCNetworkTrafficReport94scNetworkTrendData() SCISCNetworkTrafficReport94scNetworkTrendData {
	m := make(SCISCNetworkTrafficReport94scNetworkTrendData, 0)
	return m
}

// SCISCNetworkTrafficReport100droppedCallRateData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.100.droppedCallRate.Data
type SCISCNetworkTrafficReport100droppedCallRateData []interface{}

func MakeSCISCNetworkTrafficReport100droppedCallRateData() SCISCNetworkTrafficReport100droppedCallRateData {
	m := make(SCISCNetworkTrafficReport100droppedCallRateData, 0)
	return m
}

// SCISCNetworkTrafficReport101connectionSetupSuccessRateData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.101.connectionSetupSuccessRate.Data
type SCISCNetworkTrafficReport101connectionSetupSuccessRateData []interface{}

func MakeSCISCNetworkTrafficReport101connectionSetupSuccessRateData() SCISCNetworkTrafficReport101connectionSetupSuccessRateData {
	m := make(SCISCNetworkTrafficReport101connectionSetupSuccessRateData, 0)
	return m
}

// SCISCNetworkTrafficReport102handoverSuccessRateData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.102.handoverSuccessRate.Data
type SCISCNetworkTrafficReport102handoverSuccessRateData []interface{}

func MakeSCISCNetworkTrafficReport102handoverSuccessRateData() SCISCNetworkTrafficReport102handoverSuccessRateData {
	m := make(SCISCNetworkTrafficReport102handoverSuccessRateData, 0)
	return m
}

// SCISCNetworkTrafficReport103avgThroughputData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.103.avgThroughput.Data
type SCISCNetworkTrafficReport103avgThroughputData []interface{}

func MakeSCISCNetworkTrafficReport103avgThroughputData() SCISCNetworkTrafficReport103avgThroughputData {
	m := make(SCISCNetworkTrafficReport103avgThroughputData, 0)
	return m
}

// SCISCNetworkTrafficReport104scAvailabilityData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.104.scAvailability.Data
type SCISCNetworkTrafficReport104scAvailabilityData []interface{}

func MakeSCISCNetworkTrafficReport104scAvailabilityData() SCISCNetworkTrafficReport104scAvailabilityData {
	m := make(SCISCNetworkTrafficReport104scAvailabilityData, 0)
	return m
}

// SCISCNetworkTrafficReport105rscConnectionStatsData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.105.rscConnectionStats.Data
type SCISCNetworkTrafficReport105rscConnectionStatsData []interface{}

func MakeSCISCNetworkTrafficReport105rscConnectionStatsData() SCISCNetworkTrafficReport105rscConnectionStatsData {
	m := make(SCISCNetworkTrafficReport105rscConnectionStatsData, 0)
	return m
}

// SCISCNetworkTrafficReport106rscGpsStatsData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.106.rscGpsStats.Data
type SCISCNetworkTrafficReport106rscGpsStatsData []interface{}

func MakeSCISCNetworkTrafficReport106rscGpsStatsData() SCISCNetworkTrafficReport106rscGpsStatsData {
	m := make(SCISCNetworkTrafficReport106rscGpsStatsData, 0)
	return m
}

// SCISCNetworkTrafficReport107trafficVolumeData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.107.trafficVolume.Data
type SCISCNetworkTrafficReport107trafficVolumeData []interface{}

func MakeSCISCNetworkTrafficReport107trafficVolumeData() SCISCNetworkTrafficReport107trafficVolumeData {
	m := make(SCISCNetworkTrafficReport107trafficVolumeData, 0)
	return m
}

// SCISCNetworkTrafficReport108phaseSyncLossData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.108.phaseSyncLoss.Data
type SCISCNetworkTrafficReport108phaseSyncLossData []interface{}

func MakeSCISCNetworkTrafficReport108phaseSyncLossData() SCISCNetworkTrafficReport108phaseSyncLossData {
	m := make(SCISCNetworkTrafficReport108phaseSyncLossData, 0)
	return m
}

// SCISCNetworkTrafficReport109frequencySyncLossData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.109.frequencySyncLoss.Data
type SCISCNetworkTrafficReport109frequencySyncLossData []interface{}

func MakeSCISCNetworkTrafficReport109frequencySyncLossData() SCISCNetworkTrafficReport109frequencySyncLossData {
	m := make(SCISCNetworkTrafficReport109frequencySyncLossData, 0)
	return m
}

// SCISCNetworkTrafficReport111rscTrafficStatsData
//
// Definition: SCNetworkTrafficReport.SCNetworkTrafficReport.111.rscTrafficStats.Data
type SCISCNetworkTrafficReport111rscTrafficStatsData []interface{}

func MakeSCISCNetworkTrafficReport111rscTrafficStatsData() SCISCNetworkTrafficReport111rscTrafficStatsData {
	m := make(SCISCNetworkTrafficReport111rscTrafficStatsData, 0)
	return m
}

// ReportSCNetworkTrafficReport93ScNetworkTraffic
//
// Operation ID: report.SCNetworkTrafficReport.93.scNetworkTraffic
//
// SC Network Traffic Report - Summary
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport93ScNetworkTraffic(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport93ScNetworkTraffic, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport93scNetworkTraffic200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport94ScNetworkTrend
//
// Operation ID: report.SCNetworkTrafficReport.94.scNetworkTrend
//
// SC Network Traffic Report - SmartCell Trend Over Time
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport94ScNetworkTrend(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport94ScNetworkTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport94scNetworkTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport100DroppedCallRate
//
// Operation ID: report.SCNetworkTrafficReport.100.droppedCallRate
//
// SC Network Traffic Report - SmartCell Dropped Call Rate
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport100DroppedCallRate(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport100DroppedCallRate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport100droppedCallRate200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport101ConnectionSetupSuccessRate
//
// Operation ID: report.SCNetworkTrafficReport.101.connectionSetupSuccessRate
//
// SC Network Traffic Report - SmartCell Connection Setup Success Rate
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport101ConnectionSetupSuccessRate(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport101ConnectionSetupSuccessRate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport101connectionSetupSuccessRate200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport102HandoverSuccessRate
//
// Operation ID: report.SCNetworkTrafficReport.102.handoverSuccessRate
//
// SC Network Traffic Report - SmartCell Handover Success Rate
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport102HandoverSuccessRate(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport102HandoverSuccessRate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport102handoverSuccessRate200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport103AvgThroughput
//
// Operation ID: report.SCNetworkTrafficReport.103.avgThroughput
//
// SC Network Traffic Report - SmartCell Average Throughput
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport103AvgThroughput(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport103avgThroughput200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport103AvgThroughput, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport103avgThroughput200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport104ScAvailability
//
// Operation ID: report.SCNetworkTrafficReport.104.scAvailability
//
// SC Network Traffic Report - SmartCell Availability
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport104ScAvailability(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport104scAvailability200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport104ScAvailability, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport104scAvailability200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport105RscConnectionStats
//
// Operation ID: report.SCNetworkTrafficReport.105.rscConnectionStats
//
// SC Network Traffic Report - SmartCell Connection Statistics
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport105RscConnectionStats(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport105RscConnectionStats, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport105rscConnectionStats200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport106RscGpsStats
//
// Operation ID: report.SCNetworkTrafficReport.106.rscGpsStats
//
// SC Network Traffic Report - RSC GPS Statistics
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport106RscGpsStats(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport106RscGpsStats, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport106rscGpsStats200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport107TrafficVolume
//
// Operation ID: report.SCNetworkTrafficReport.107.trafficVolume
//
// SC Network Traffic Report - SmartCell Traffic Volume
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport107TrafficVolume(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport107trafficVolume200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport107TrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport107trafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport108PhaseSyncLoss
//
// Operation ID: report.SCNetworkTrafficReport.108.phaseSyncLoss
//
// SC Network Traffic Report - SmartCell Phase Sync Loss
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport108PhaseSyncLoss(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport108PhaseSyncLoss, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport108phaseSyncLoss200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport109FrequencySyncLoss
//
// Operation ID: report.SCNetworkTrafficReport.109.frequencySyncLoss
//
// SC Network Traffic Report - SmartCell Frequency Sync Loss
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport109FrequencySyncLoss(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport109FrequencySyncLoss, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport109frequencySyncLoss200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSCNetworkTrafficReport111RscTrafficStats
//
// Operation ID: report.SCNetworkTrafficReport.111.rscTrafficStats
//
// SC Network Traffic Report - RSC Traffic Statistics
//
// Request Body:
//	 - body *SCIReportQueryBody
func (s *SCISCNetworkTrafficReportService) ReportSCNetworkTrafficReport111RscTrafficStats(ctx context.Context, body *SCIReportQueryBody, mutators ...RequestMutator) (*SCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSCNetworkTrafficReport111RscTrafficStats, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSCNetworkTrafficReport111rscTrafficStats200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
