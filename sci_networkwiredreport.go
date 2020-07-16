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

// SCINetworkWiredReport123topSwitchPOEUtilChartData
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_Data
type SCINetworkWiredReport123topSwitchPOEUtilChartData []*SCINetworkWiredReport123topSwitchPOEUtilChartDataType

func MakeSCINetworkWiredReport123topSwitchPOEUtilChartData() SCINetworkWiredReport123topSwitchPOEUtilChartData {
	m := make(SCINetworkWiredReport123topSwitchPOEUtilChartData, 0)
	return m
}

// SCINetworkWiredReport123topSwitchPOEUtilChartDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_DataType
type SCINetworkWiredReport123topSwitchPOEUtilChartDataType struct {
	PoeUtilization *int `json:"poeUtilization,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`
}

func NewSCINetworkWiredReport123topSwitchPOEUtilChartDataType() *SCINetworkWiredReport123topSwitchPOEUtilChartDataType {
	m := new(SCINetworkWiredReport123topSwitchPOEUtilChartDataType)
	return m
}

// SCINetworkWiredReport123topSwitchPOEUtilChartMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_123_topSwitchPOEUtilChart_MetaData
type SCINetworkWiredReport123topSwitchPOEUtilChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCINetworkWiredReport123topSwitchPOEUtilChartMetaData() *SCINetworkWiredReport123topSwitchPOEUtilChartMetaData {
	m := new(SCINetworkWiredReport123topSwitchPOEUtilChartMetaData)
	return m
}

// SCINetworkWiredReport124topSwitchPOEUtilsData
//
// Definition: NetworkWiredReport_NetworkWiredReport_124_topSwitchPOEUtils_Data
type SCINetworkWiredReport124topSwitchPOEUtilsData []*SCINetworkWiredReport124topSwitchPOEUtilsDataType

func MakeSCINetworkWiredReport124topSwitchPOEUtilsData() SCINetworkWiredReport124topSwitchPOEUtilsData {
	m := make(SCINetworkWiredReport124topSwitchPOEUtilsData, 0)
	return m
}

// SCINetworkWiredReport124topSwitchPOEUtilsDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_124_topSwitchPOEUtils_DataType
type SCINetworkWiredReport124topSwitchPOEUtilsDataType struct {
	Index *int `json:"index,omitempty"`

	PoeTotal *int `json:"poeTotal,omitempty"`

	PoeUtilization *int `json:"poeUtilization,omitempty"`

	PoeUtilPercent *float64 `json:"poeUtilPercent,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`
}

func NewSCINetworkWiredReport124topSwitchPOEUtilsDataType() *SCINetworkWiredReport124topSwitchPOEUtilsDataType {
	m := new(SCINetworkWiredReport124topSwitchPOEUtilsDataType)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableData
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_Data
type SCINetworkWiredReport128topSwitchesByTrafficTableData []*SCINetworkWiredReport128topSwitchesByTrafficTableDataType

func MakeSCINetworkWiredReport128topSwitchesByTrafficTableData() SCINetworkWiredReport128topSwitchesByTrafficTableData {
	m := make(SCINetworkWiredReport128topSwitchesByTrafficTableData, 0)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_DataType
type SCINetworkWiredReport128topSwitchesByTrafficTableDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	Index *int `json:"index,omitempty"`

	PortCount *int `json:"portCount,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`

	SwitchUnitCount *int `json:"switchUnitCount,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCINetworkWiredReport128topSwitchesByTrafficTableDataType() *SCINetworkWiredReport128topSwitchesByTrafficTableDataType {
	m := new(SCINetworkWiredReport128topSwitchesByTrafficTableDataType)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_MetaData
type SCINetworkWiredReport128topSwitchesByTrafficTableMetaData struct {
	MaxValues *SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCINetworkWiredReport128topSwitchesByTrafficTableMetaData() *SCINetworkWiredReport128topSwitchesByTrafficTableMetaData {
	m := new(SCINetworkWiredReport128topSwitchesByTrafficTableMetaData)
	return m
}

// SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType
//
// Definition: NetworkWiredReport_NetworkWiredReport_128_topSwitchesByTrafficTable_MetaDataMaxValuesType
type SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType() *SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType {
	m := new(SCINetworkWiredReport128topSwitchesByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWiredReport134wiredOverviewData
//
// Definition: NetworkWiredReport_NetworkWiredReport_134_wiredOverview_Data
type SCINetworkWiredReport134wiredOverviewData []*SCINetworkWiredReport134wiredOverviewDataType

func MakeSCINetworkWiredReport134wiredOverviewData() SCINetworkWiredReport134wiredOverviewData {
	m := make(SCINetworkWiredReport134wiredOverviewData, 0)
	return m
}

// SCINetworkWiredReport134wiredOverviewDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_134_wiredOverview_DataType
type SCINetworkWiredReport134wiredOverviewDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalSwitches *int `json:"totalSwitches,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCINetworkWiredReport134wiredOverviewDataType() *SCINetworkWiredReport134wiredOverviewDataType {
	m := new(SCINetworkWiredReport134wiredOverviewDataType)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionData
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_Data
type SCINetworkWiredReport135wiredTrafficDistributionData []*SCINetworkWiredReport135wiredTrafficDistributionDataType

func MakeSCINetworkWiredReport135wiredTrafficDistributionData() SCINetworkWiredReport135wiredTrafficDistributionData {
	m := make(SCINetworkWiredReport135wiredTrafficDistributionData, 0)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_DataType
type SCINetworkWiredReport135wiredTrafficDistributionDataType struct {
	Categories []string `json:"categories,omitempty"`

	Series []*SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType `json:"series,omitempty"`
}

func NewSCINetworkWiredReport135wiredTrafficDistributionDataType() *SCINetworkWiredReport135wiredTrafficDistributionDataType {
	m := new(SCINetworkWiredReport135wiredTrafficDistributionDataType)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_DataTypeSeriesType
type SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType struct {
	Data []float64 `json:"data,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType() *SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType {
	m := new(SCINetworkWiredReport135wiredTrafficDistributionDataTypeSeriesType)
	return m
}

// SCINetworkWiredReport135wiredTrafficDistributionMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_135_wiredTrafficDistribution_MetaData
type SCINetworkWiredReport135wiredTrafficDistributionMetaData struct {
	ChartType *string `json:"chartType,omitempty"`
}

func NewSCINetworkWiredReport135wiredTrafficDistributionMetaData() *SCINetworkWiredReport135wiredTrafficDistributionMetaData {
	m := new(SCINetworkWiredReport135wiredTrafficDistributionMetaData)
	return m
}

// SCINetworkWiredReport136switchTrafficTrendData
//
// Definition: NetworkWiredReport_NetworkWiredReport_136_switchTrafficTrend_Data
type SCINetworkWiredReport136switchTrafficTrendData []*SCINetworkWiredReport136switchTrafficTrendDataType

func MakeSCINetworkWiredReport136switchTrafficTrendData() SCINetworkWiredReport136switchTrafficTrendData {
	m := make(SCINetworkWiredReport136switchTrafficTrendData, 0)
	return m
}

// SCINetworkWiredReport136switchTrafficTrendDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_136_switchTrafficTrend_DataType
type SCINetworkWiredReport136switchTrafficTrendDataType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCINetworkWiredReport136switchTrafficTrendDataType() *SCINetworkWiredReport136switchTrafficTrendDataType {
	m := new(SCINetworkWiredReport136switchTrafficTrendDataType)
	return m
}

// SCINetworkWiredReport141switchErrorTrendData
//
// Definition: NetworkWiredReport_NetworkWiredReport_141_switchErrorTrend_Data
type SCINetworkWiredReport141switchErrorTrendData []*SCINetworkWiredReport141switchErrorTrendDataType

func MakeSCINetworkWiredReport141switchErrorTrendData() SCINetworkWiredReport141switchErrorTrendData {
	m := make(SCINetworkWiredReport141switchErrorTrendData, 0)
	return m
}

// SCINetworkWiredReport141switchErrorTrendDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_141_switchErrorTrend_DataType
type SCINetworkWiredReport141switchErrorTrendDataType struct {
	CrcErrors *int `json:"crcErrors,omitempty"`

	End *string `json:"end,omitempty"`

	InDiscards *int `json:"inDiscards,omitempty"`

	InErrors *int `json:"inErrors,omitempty"`

	OutErrors *int `json:"outErrors,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCINetworkWiredReport141switchErrorTrendDataType() *SCINetworkWiredReport141switchErrorTrendDataType {
	m := new(SCINetworkWiredReport141switchErrorTrendDataType)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartData
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_Data
type SCINetworkWiredReport142topSwitchesByErrorsChartData []*SCINetworkWiredReport142topSwitchesByErrorsChartDataType

func MakeSCINetworkWiredReport142topSwitchesByErrorsChartData() SCINetworkWiredReport142topSwitchesByErrorsChartData {
	m := make(SCINetworkWiredReport142topSwitchesByErrorsChartData, 0)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_DataType
type SCINetworkWiredReport142topSwitchesByErrorsChartDataType struct {
	Error *int `json:"error,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`
}

func NewSCINetworkWiredReport142topSwitchesByErrorsChartDataType() *SCINetworkWiredReport142topSwitchesByErrorsChartDataType {
	m := new(SCINetworkWiredReport142topSwitchesByErrorsChartDataType)
	return m
}

// SCINetworkWiredReport142topSwitchesByErrorsChartMetaData
//
// Definition: NetworkWiredReport_NetworkWiredReport_142_topSwitchesByErrorsChart_MetaData
type SCINetworkWiredReport142topSwitchesByErrorsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCINetworkWiredReport142topSwitchesByErrorsChartMetaData() *SCINetworkWiredReport142topSwitchesByErrorsChartMetaData {
	m := new(SCINetworkWiredReport142topSwitchesByErrorsChartMetaData)
	return m
}

// SCINetworkWiredReport143topSwitchesByErrorsTableData
//
// Definition: NetworkWiredReport_NetworkWiredReport_143_topSwitchesByErrorsTable_Data
type SCINetworkWiredReport143topSwitchesByErrorsTableData []*SCINetworkWiredReport143topSwitchesByErrorsTableDataType

func MakeSCINetworkWiredReport143topSwitchesByErrorsTableData() SCINetworkWiredReport143topSwitchesByErrorsTableData {
	m := make(SCINetworkWiredReport143topSwitchesByErrorsTableData, 0)
	return m
}

// SCINetworkWiredReport143topSwitchesByErrorsTableDataType
//
// Definition: NetworkWiredReport_NetworkWiredReport_143_topSwitchesByErrorsTable_DataType
type SCINetworkWiredReport143topSwitchesByErrorsTableDataType struct {
	Error *int `json:"error,omitempty"`

	Index *int `json:"index,omitempty"`

	InErr *int `json:"inErr,omitempty"`

	OutErr *int `json:"outErr,omitempty"`

	SwitchId *string `json:"switchId,omitempty"`

	SwitchName *string `json:"switchName,omitempty"`
}

func NewSCINetworkWiredReport143topSwitchesByErrorsTableDataType() *SCINetworkWiredReport143topSwitchesByErrorsTableDataType {
	m := new(SCINetworkWiredReport143topSwitchesByErrorsTableDataType)
	return m
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
