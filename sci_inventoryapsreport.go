package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIInventoryAPsReportService struct {
	apiClient *SCIClient
}

func NewSCIInventoryAPsReportService(c *SCIClient) *SCIInventoryAPsReportService {
	s := new(SCIInventoryAPsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIInventoryAPsReportService() *SCIInventoryAPsReportService {
	return NewSCIInventoryAPsReportService(ss.apiClient)
}

// ReportInventoryAPsReport46ApInventoryOverview
//
// Operation ID: report_InventoryAPsReport_46_apInventoryOverview
//
// Inventory - APs Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport46ApInventoryOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport46apInventoryOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport46apInventoryOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport46ApInventoryOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport46apInventoryOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport47TopApsDisconnection
//
// Operation ID: report_InventoryAPsReport_47_topApsDisconnection
//
// Inventory - APs Report - Top APs by Offline Duration
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport47TopApsDisconnection(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport47topApsDisconnection200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport47topApsDisconnection200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport47TopApsDisconnection, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport47topApsDisconnection200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport48ApCountTrend
//
// Operation ID: report_InventoryAPsReport_48_apCountTrend
//
// Inventory - APs Report - AP Count Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport48ApCountTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport48apCountTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport48apCountTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport48ApCountTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport48apCountTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport49ApStatusTrend
//
// Operation ID: report_InventoryAPsReport_49_apStatusTrend
//
// Inventory - APs Report - AP Status Trends
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport49ApStatusTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport49apStatusTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport49apStatusTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport49ApStatusTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport49apStatusTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport50TopApsModelsChart
//
// Operation ID: report_InventoryAPsReport_50_topApsModelsChart
//
// Inventory - APs Report - Top AP Models
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport50TopApsModelsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport50topApsModelsChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport50topApsModelsChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport50TopApsModelsChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport50topApsModelsChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport51Top10ApVersionsChart
//
// Operation ID: report_InventoryAPsReport_51_top10ApVersionsChart
//
// Inventory - APs Report - Top AP Software Versions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport51Top10ApVersionsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport51Top10ApVersionsChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport51top10ApVersionsChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport52TopApsRebootReasons
//
// Operation ID: report_InventoryAPsReport_52_topApsRebootReasons
//
// Inventory - APs Report - Top 10 AP Reboot Reasons
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport52TopApsRebootReasons(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport52topApsRebootReasons200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport52TopApsRebootReasons, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport52topApsRebootReasons200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport53Top10ApsRebootCounts
//
// Operation ID: report_InventoryAPsReport_53_top10ApsRebootCounts
//
// Inventory - APs Report - Top APs by Reboot Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport53Top10ApsRebootCounts(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport53Top10ApsRebootCounts, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport53top10ApsRebootCounts200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport54TopApAlarmTypes
//
// Operation ID: report_InventoryAPsReport_54_topApAlarmTypes
//
// Inventory - APs Report - Top 10 AP Alarm Types
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport54TopApAlarmTypes(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport54topApAlarmTypes200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport54TopApAlarmTypes, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport54topApAlarmTypes200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport55TopAPModels
//
// Operation ID: report_InventoryAPsReport_55_topAPModels
//
// Inventory - APs Report - Top AP Models
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport55TopAPModels(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport55topAPModels200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport55topAPModels200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport55TopAPModels, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport55topAPModels200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport56TopAPVersions
//
// Operation ID: report_InventoryAPsReport_56_topAPVersions
//
// Inventory - APs Report - Top AP Software Versions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport56TopAPVersions(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport56topAPVersions200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport56topAPVersions200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport56TopAPVersions, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport56topAPVersions200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport57TopAPsOffline
//
// Operation ID: report_InventoryAPsReport_57_topAPsOffline
//
// Inventory - APs Report - Top APs by Offline Duration
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport57TopAPsOffline(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport57topAPsOffline200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport57topAPsOffline200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport57TopAPsOffline, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport57topAPsOffline200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport58TopAPsByReboots
//
// Operation ID: report_InventoryAPsReport_58_topAPsByReboots
//
// Inventory - APs Report - Top APs by Reboot Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport58TopAPsByReboots(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport58topAPsByReboots200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport58topAPsByReboots200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport58TopAPsByReboots, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport58topAPsByReboots200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport59ApsConfiguredInMultiCtrl
//
// Operation ID: report_InventoryAPsReport_59_apsConfiguredInMultiCtrl
//
// Inventory - APs Report - APs Configured in Multiple Systems
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport59ApsConfiguredInMultiCtrl(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport59ApsConfiguredInMultiCtrl, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport59apsConfiguredInMultiCtrl200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport60ApDetailsOnOfflineStatus
//
// Operation ID: report_InventoryAPsReport_60_apDetailsOnOfflineStatus
//
// Inventory - APs Report - AP Details for Online/Offline Status
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport60ApDetailsOnOfflineStatus(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport60ApDetailsOnOfflineStatus, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport60apDetailsOnOfflineStatus200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryAPsReport61ApDetailsOtherStatus
//
// Operation ID: report_InventoryAPsReport_61_apDetailsOtherStatus
//
// Inventory - APs Report - AP Details for Other Statuses
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryAPsReportService) ReportInventoryAPsReport61ApDetailsOtherStatus(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryAPsReport61ApDetailsOtherStatus, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryAPsReport61apDetailsOtherStatus200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
