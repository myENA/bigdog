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

// SCIInventorySwitchesReport113overviewData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.113.overview.Data
type SCIInventorySwitchesReport113overviewData []*SCIInventorySwitchesReport113overviewDataType

func MakeSCIInventorySwitchesReport113overviewData() SCIInventorySwitchesReport113overviewData {
	m := make(SCIInventorySwitchesReport113overviewData, 0)
	return m
}

// SCIInventorySwitchesReport113overviewDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.113.overview.DataType
type SCIInventorySwitchesReport113overviewDataType struct {
	DownPortCount *int `json:"downPortCount,omitempty"`

	OfflineSwitchCount *int `json:"offlineSwitchCount,omitempty"`

	OnlineSwitchCount *int `json:"onlineSwitchCount,omitempty"`

	OtherSwitchCount *int `json:"otherSwitchCount,omitempty"`

	PortCount *int `json:"portCount,omitempty"`

	SwitchCount *int `json:"switchCount,omitempty"`

	SwitchUnitCount *int `json:"switchUnitCount,omitempty"`

	UpPortCount *int `json:"upPortCount,omitempty"`
}

func NewSCIInventorySwitchesReport113overviewDataType() *SCIInventorySwitchesReport113overviewDataType {
	m := new(SCIInventorySwitchesReport113overviewDataType)
	return m
}

// SCIInventorySwitchesReport116switchCountTrendData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.116.switchCountTrend.Data
type SCIInventorySwitchesReport116switchCountTrendData []*SCIInventorySwitchesReport116switchCountTrendDataType

func MakeSCIInventorySwitchesReport116switchCountTrendData() SCIInventorySwitchesReport116switchCountTrendData {
	m := make(SCIInventorySwitchesReport116switchCountTrendData, 0)
	return m
}

// SCIInventorySwitchesReport116switchCountTrendDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.116.switchCountTrend.DataType
type SCIInventorySwitchesReport116switchCountTrendDataType struct {
	Count *int `json:"count,omitempty"`

	End *string `json:"end,omitempty"`

	OfflineCount *int `json:"offlineCount,omitempty"`

	OnlineCount *int `json:"onlineCount,omitempty"`

	Start *string `json:"start,omitempty"`

	SwitchUnitCount *int `json:"switchUnitCount,omitempty"`
}

func NewSCIInventorySwitchesReport116switchCountTrendDataType() *SCIInventorySwitchesReport116switchCountTrendDataType {
	m := new(SCIInventorySwitchesReport116switchCountTrendDataType)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.117.top10SwitchVersionChart.Data
type SCIInventorySwitchesReport117top10SwitchVersionChartData []*SCIInventorySwitchesReport117top10SwitchVersionChartDataType

func MakeSCIInventorySwitchesReport117top10SwitchVersionChartData() SCIInventorySwitchesReport117top10SwitchVersionChartData {
	m := make(SCIInventorySwitchesReport117top10SwitchVersionChartData, 0)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.117.top10SwitchVersionChart.DataType
type SCIInventorySwitchesReport117top10SwitchVersionChartDataType struct {
	SwitchCount *int `json:"switchCount,omitempty"`

	SwitchFirmware *string `json:"switchFirmware,omitempty"`
}

func NewSCIInventorySwitchesReport117top10SwitchVersionChartDataType() *SCIInventorySwitchesReport117top10SwitchVersionChartDataType {
	m := new(SCIInventorySwitchesReport117top10SwitchVersionChartDataType)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartMetaData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.117.top10SwitchVersionChart.MetaData
type SCIInventorySwitchesReport117top10SwitchVersionChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventorySwitchesReport117top10SwitchVersionChartMetaData() *SCIInventorySwitchesReport117top10SwitchVersionChartMetaData {
	m := new(SCIInventorySwitchesReport117top10SwitchVersionChartMetaData)
	return m
}

// SCIInventorySwitchesReport118topSwitchVersionsData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.118.topSwitchVersions.Data
type SCIInventorySwitchesReport118topSwitchVersionsData []*SCIInventorySwitchesReport118topSwitchVersionsDataType

func MakeSCIInventorySwitchesReport118topSwitchVersionsData() SCIInventorySwitchesReport118topSwitchVersionsData {
	m := make(SCIInventorySwitchesReport118topSwitchVersionsData, 0)
	return m
}

// SCIInventorySwitchesReport118topSwitchVersionsDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.118.topSwitchVersions.DataType
type SCIInventorySwitchesReport118topSwitchVersionsDataType struct {
	Index *int `json:"index,omitempty"`

	SwitchCount *int `json:"switchCount,omitempty"`

	SwitchFirmware *string `json:"switchFirmware,omitempty"`

	SwitchPercent *float64 `json:"switchPercent,omitempty"`
}

func NewSCIInventorySwitchesReport118topSwitchVersionsDataType() *SCIInventorySwitchesReport118topSwitchVersionsDataType {
	m := new(SCIInventorySwitchesReport118topSwitchVersionsDataType)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.121.topSwitchModelsChart.Data
type SCIInventorySwitchesReport121topSwitchModelsChartData []*SCIInventorySwitchesReport121topSwitchModelsChartDataType

func MakeSCIInventorySwitchesReport121topSwitchModelsChartData() SCIInventorySwitchesReport121topSwitchModelsChartData {
	m := make(SCIInventorySwitchesReport121topSwitchModelsChartData, 0)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.121.topSwitchModelsChart.DataType
type SCIInventorySwitchesReport121topSwitchModelsChartDataType struct {
	SwitchCount *int `json:"switchCount,omitempty"`

	SwitchModel *string `json:"switchModel,omitempty"`
}

func NewSCIInventorySwitchesReport121topSwitchModelsChartDataType() *SCIInventorySwitchesReport121topSwitchModelsChartDataType {
	m := new(SCIInventorySwitchesReport121topSwitchModelsChartDataType)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartMetaData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.121.topSwitchModelsChart.MetaData
type SCIInventorySwitchesReport121topSwitchModelsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventorySwitchesReport121topSwitchModelsChartMetaData() *SCIInventorySwitchesReport121topSwitchModelsChartMetaData {
	m := new(SCIInventorySwitchesReport121topSwitchModelsChartMetaData)
	return m
}

// SCIInventorySwitchesReport122topSwitchModelsData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.122.topSwitchModels.Data
type SCIInventorySwitchesReport122topSwitchModelsData []*SCIInventorySwitchesReport122topSwitchModelsDataType

func MakeSCIInventorySwitchesReport122topSwitchModelsData() SCIInventorySwitchesReport122topSwitchModelsData {
	m := make(SCIInventorySwitchesReport122topSwitchModelsData, 0)
	return m
}

// SCIInventorySwitchesReport122topSwitchModelsDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.122.topSwitchModels.DataType
type SCIInventorySwitchesReport122topSwitchModelsDataType struct {
	Index *int `json:"index,omitempty"`

	SwitchCount *int `json:"switchCount,omitempty"`

	SwitchModel *string `json:"switchModel,omitempty"`

	SwitchPercent *float64 `json:"switchPercent,omitempty"`
}

func NewSCIInventorySwitchesReport122topSwitchModelsDataType() *SCIInventorySwitchesReport122topSwitchModelsDataType {
	m := new(SCIInventorySwitchesReport122topSwitchModelsDataType)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.132.portStatusTrend.Data
type SCIInventorySwitchesReport132portStatusTrendData []*SCIInventorySwitchesReport132portStatusTrendDataType

func MakeSCIInventorySwitchesReport132portStatusTrendData() SCIInventorySwitchesReport132portStatusTrendData {
	m := make(SCIInventorySwitchesReport132portStatusTrendData, 0)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendDataType
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.132.portStatusTrend.DataType
type SCIInventorySwitchesReport132portStatusTrendDataType struct {
	PortCount *int `json:"portCount,omitempty"`

	PortStatus *string `json:"portStatus,omitempty"`
}

func NewSCIInventorySwitchesReport132portStatusTrendDataType() *SCIInventorySwitchesReport132portStatusTrendDataType {
	m := new(SCIInventorySwitchesReport132portStatusTrendDataType)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendMetaData
//
// Definition: InventorySwitchesReport.InventorySwitchesReport.132.portStatusTrend.MetaData
type SCIInventorySwitchesReport132portStatusTrendMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventorySwitchesReport132portStatusTrendMetaData() *SCIInventorySwitchesReport132portStatusTrendMetaData {
	m := new(SCIInventorySwitchesReport132portStatusTrendMetaData)
	return m
}

// ReportInventorySwitchesReport113Overview
//
// Operation ID: report.InventorySwitchesReport.113.overview
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
// Operation ID: report.InventorySwitchesReport.116.switchCountTrend
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
// Operation ID: report.InventorySwitchesReport.117.top10SwitchVersionChart
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
// Operation ID: report.InventorySwitchesReport.118.topSwitchVersions
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
// Operation ID: report.InventorySwitchesReport.121.topSwitchModelsChart
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
// Operation ID: report.InventorySwitchesReport.122.topSwitchModels
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
// Operation ID: report.InventorySwitchesReport.132.portStatusTrend
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
