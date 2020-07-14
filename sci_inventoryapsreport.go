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

// SCIInventoryAPsReport46apInventoryOverviewData
//
// Definition: InventoryAPsReport.InventoryAPsReport.46.apInventoryOverview.Data
type SCIInventoryAPsReport46apInventoryOverviewData []*SCIInventoryAPsReport46apInventoryOverviewDataType

func MakeSCIInventoryAPsReport46apInventoryOverviewData() SCIInventoryAPsReport46apInventoryOverviewData {
	m := make(SCIInventoryAPsReport46apInventoryOverviewData, 0)
	return m
}

// SCIInventoryAPsReport46apInventoryOverviewDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.46.apInventoryOverview.DataType
type SCIInventoryAPsReport46apInventoryOverviewDataType struct {
	Offline *int `json:"offline,omitempty"`

	Online *int `json:"online,omitempty"`

	Others *int `json:"others,omitempty"`

	Reboots *int `json:"reboots,omitempty"`

	Total *int `json:"total,omitempty"`

	TotalApsWAlarm *int `json:"totalApsWAlarm,omitempty"`

	TotalApsWReboot *int `json:"totalApsWReboot,omitempty"`
}

func NewSCIInventoryAPsReport46apInventoryOverviewDataType() *SCIInventoryAPsReport46apInventoryOverviewDataType {
	m := new(SCIInventoryAPsReport46apInventoryOverviewDataType)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionData
//
// Definition: InventoryAPsReport.InventoryAPsReport.47.topApsDisconnection.Data
type SCIInventoryAPsReport47topApsDisconnectionData []*SCIInventoryAPsReport47topApsDisconnectionDataType

func MakeSCIInventoryAPsReport47topApsDisconnectionData() SCIInventoryAPsReport47topApsDisconnectionData {
	m := make(SCIInventoryAPsReport47topApsDisconnectionData, 0)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.47.topApsDisconnection.DataType
type SCIInventoryAPsReport47topApsDisconnectionDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *int `json:"value,omitempty"`
}

func NewSCIInventoryAPsReport47topApsDisconnectionDataType() *SCIInventoryAPsReport47topApsDisconnectionDataType {
	m := new(SCIInventoryAPsReport47topApsDisconnectionDataType)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.47.topApsDisconnection.MetaData
type SCIInventoryAPsReport47topApsDisconnectionMetaData struct {
	Legend *SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType `json:"legend,omitempty"`
}

func NewSCIInventoryAPsReport47topApsDisconnectionMetaData() *SCIInventoryAPsReport47topApsDisconnectionMetaData {
	m := new(SCIInventoryAPsReport47topApsDisconnectionMetaData)
	return m
}

// SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType
//
// Definition: InventoryAPsReport.InventoryAPsReport.47.topApsDisconnection.MetaDataLegendType
type SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType struct {
	On *string `json:"on,omitempty"`
}

func NewSCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType() *SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType {
	m := new(SCIInventoryAPsReport47topApsDisconnectionMetaDataLegendType)
	return m
}

// SCIInventoryAPsReport48apCountTrendData
//
// Definition: InventoryAPsReport.InventoryAPsReport.48.apCountTrend.Data
type SCIInventoryAPsReport48apCountTrendData []*SCIInventoryAPsReport48apCountTrendDataType

func MakeSCIInventoryAPsReport48apCountTrendData() SCIInventoryAPsReport48apCountTrendData {
	m := make(SCIInventoryAPsReport48apCountTrendData, 0)
	return m
}

// SCIInventoryAPsReport48apCountTrendDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.48.apCountTrend.DataType
type SCIInventoryAPsReport48apCountTrendDataType struct {
	ConnectedAPs *float64 `json:"connectedAPs,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalAPs *float64 `json:"totalAPs,omitempty"`
}

func NewSCIInventoryAPsReport48apCountTrendDataType() *SCIInventoryAPsReport48apCountTrendDataType {
	m := new(SCIInventoryAPsReport48apCountTrendDataType)
	return m
}

// SCIInventoryAPsReport49apStatusTrendData
//
// Definition: InventoryAPsReport.InventoryAPsReport.49.apStatusTrend.Data
type SCIInventoryAPsReport49apStatusTrendData []*SCIInventoryAPsReport49apStatusTrendDataType

func MakeSCIInventoryAPsReport49apStatusTrendData() SCIInventoryAPsReport49apStatusTrendData {
	m := make(SCIInventoryAPsReport49apStatusTrendData, 0)
	return m
}

// SCIInventoryAPsReport49apStatusTrendDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.49.apStatusTrend.DataType
type SCIInventoryAPsReport49apStatusTrendDataType struct {
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	Count *int `json:"count,omitempty"`
}

func NewSCIInventoryAPsReport49apStatusTrendDataType() *SCIInventoryAPsReport49apStatusTrendDataType {
	m := new(SCIInventoryAPsReport49apStatusTrendDataType)
	return m
}

// SCIInventoryAPsReport49apStatusTrendMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.49.apStatusTrend.MetaData
type SCIInventoryAPsReport49apStatusTrendMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventoryAPsReport49apStatusTrendMetaData() *SCIInventoryAPsReport49apStatusTrendMetaData {
	m := new(SCIInventoryAPsReport49apStatusTrendMetaData)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartData
//
// Definition: InventoryAPsReport.InventoryAPsReport.50.topApsModelsChart.Data
type SCIInventoryAPsReport50topApsModelsChartData []*SCIInventoryAPsReport50topApsModelsChartDataType

func MakeSCIInventoryAPsReport50topApsModelsChartData() SCIInventoryAPsReport50topApsModelsChartData {
	m := make(SCIInventoryAPsReport50topApsModelsChartData, 0)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.50.topApsModelsChart.DataType
type SCIInventoryAPsReport50topApsModelsChartDataType struct {
	ApCount *int `json:"apCount,omitempty"`

	ApModel *string `json:"apModel,omitempty"`
}

func NewSCIInventoryAPsReport50topApsModelsChartDataType() *SCIInventoryAPsReport50topApsModelsChartDataType {
	m := new(SCIInventoryAPsReport50topApsModelsChartDataType)
	return m
}

// SCIInventoryAPsReport50topApsModelsChartMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.50.topApsModelsChart.MetaData
type SCIInventoryAPsReport50topApsModelsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventoryAPsReport50topApsModelsChartMetaData() *SCIInventoryAPsReport50topApsModelsChartMetaData {
	m := new(SCIInventoryAPsReport50topApsModelsChartMetaData)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartData
//
// Definition: InventoryAPsReport.InventoryAPsReport.51.top10ApVersionsChart.Data
type SCIInventoryAPsReport51top10ApVersionsChartData []*SCIInventoryAPsReport51top10ApVersionsChartDataType

func MakeSCIInventoryAPsReport51top10ApVersionsChartData() SCIInventoryAPsReport51top10ApVersionsChartData {
	m := make(SCIInventoryAPsReport51top10ApVersionsChartData, 0)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.51.top10ApVersionsChart.DataType
type SCIInventoryAPsReport51top10ApVersionsChartDataType struct {
	ApCount *int `json:"apCount,omitempty"`

	ApFwVersion *string `json:"apFwVersion,omitempty"`
}

func NewSCIInventoryAPsReport51top10ApVersionsChartDataType() *SCIInventoryAPsReport51top10ApVersionsChartDataType {
	m := new(SCIInventoryAPsReport51top10ApVersionsChartDataType)
	return m
}

// SCIInventoryAPsReport51top10ApVersionsChartMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.51.top10ApVersionsChart.MetaData
type SCIInventoryAPsReport51top10ApVersionsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventoryAPsReport51top10ApVersionsChartMetaData() *SCIInventoryAPsReport51top10ApVersionsChartMetaData {
	m := new(SCIInventoryAPsReport51top10ApVersionsChartMetaData)
	return m
}

// SCIInventoryAPsReport52topApsRebootReasonsData
//
// Definition: InventoryAPsReport.InventoryAPsReport.52.topApsRebootReasons.Data
type SCIInventoryAPsReport52topApsRebootReasonsData []*SCIInventoryAPsReport52topApsRebootReasonsDataType

func MakeSCIInventoryAPsReport52topApsRebootReasonsData() SCIInventoryAPsReport52topApsRebootReasonsData {
	m := make(SCIInventoryAPsReport52topApsRebootReasonsData, 0)
	return m
}

// SCIInventoryAPsReport52topApsRebootReasonsDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.52.topApsRebootReasons.DataType
type SCIInventoryAPsReport52topApsRebootReasonsDataType struct {
	Count *int `json:"count,omitempty"`

	Reason *string `json:"reason,omitempty"`
}

func NewSCIInventoryAPsReport52topApsRebootReasonsDataType() *SCIInventoryAPsReport52topApsRebootReasonsDataType {
	m := new(SCIInventoryAPsReport52topApsRebootReasonsDataType)
	return m
}

// SCIInventoryAPsReport52topApsRebootReasonsMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.52.topApsRebootReasons.MetaData
type SCIInventoryAPsReport52topApsRebootReasonsMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventoryAPsReport52topApsRebootReasonsMetaData() *SCIInventoryAPsReport52topApsRebootReasonsMetaData {
	m := new(SCIInventoryAPsReport52topApsRebootReasonsMetaData)
	return m
}

// SCIInventoryAPsReport53top10ApsRebootCountsData
//
// Definition: InventoryAPsReport.InventoryAPsReport.53.top10ApsRebootCounts.Data
type SCIInventoryAPsReport53top10ApsRebootCountsData []*SCIInventoryAPsReport53top10ApsRebootCountsDataType

func MakeSCIInventoryAPsReport53top10ApsRebootCountsData() SCIInventoryAPsReport53top10ApsRebootCountsData {
	m := make(SCIInventoryAPsReport53top10ApsRebootCountsData, 0)
	return m
}

// SCIInventoryAPsReport53top10ApsRebootCountsDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.53.top10ApsRebootCounts.DataType
type SCIInventoryAPsReport53top10ApsRebootCountsDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	Count *int `json:"count,omitempty"`
}

func NewSCIInventoryAPsReport53top10ApsRebootCountsDataType() *SCIInventoryAPsReport53top10ApsRebootCountsDataType {
	m := new(SCIInventoryAPsReport53top10ApsRebootCountsDataType)
	return m
}

// SCIInventoryAPsReport53top10ApsRebootCountsMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.53.top10ApsRebootCounts.MetaData
type SCIInventoryAPsReport53top10ApsRebootCountsMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventoryAPsReport53top10ApsRebootCountsMetaData() *SCIInventoryAPsReport53top10ApsRebootCountsMetaData {
	m := new(SCIInventoryAPsReport53top10ApsRebootCountsMetaData)
	return m
}

// SCIInventoryAPsReport54topApAlarmTypesData
//
// Definition: InventoryAPsReport.InventoryAPsReport.54.topApAlarmTypes.Data
type SCIInventoryAPsReport54topApAlarmTypesData []*SCIInventoryAPsReport54topApAlarmTypesDataType

func MakeSCIInventoryAPsReport54topApAlarmTypesData() SCIInventoryAPsReport54topApAlarmTypesData {
	m := make(SCIInventoryAPsReport54topApAlarmTypesData, 0)
	return m
}

// SCIInventoryAPsReport54topApAlarmTypesDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.54.topApAlarmTypes.DataType
type SCIInventoryAPsReport54topApAlarmTypesDataType struct {
	AlarmType *string `json:"alarmType,omitempty"`

	Count *int `json:"count,omitempty"`
}

func NewSCIInventoryAPsReport54topApAlarmTypesDataType() *SCIInventoryAPsReport54topApAlarmTypesDataType {
	m := new(SCIInventoryAPsReport54topApAlarmTypesDataType)
	return m
}

// SCIInventoryAPsReport54topApAlarmTypesMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.54.topApAlarmTypes.MetaData
type SCIInventoryAPsReport54topApAlarmTypesMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCIInventoryAPsReport54topApAlarmTypesMetaData() *SCIInventoryAPsReport54topApAlarmTypesMetaData {
	m := new(SCIInventoryAPsReport54topApAlarmTypesMetaData)
	return m
}

// SCIInventoryAPsReport55topAPModelsData
//
// Definition: InventoryAPsReport.InventoryAPsReport.55.topAPModels.Data
type SCIInventoryAPsReport55topAPModelsData []*SCIInventoryAPsReport55topAPModelsDataType

func MakeSCIInventoryAPsReport55topAPModelsData() SCIInventoryAPsReport55topAPModelsData {
	m := make(SCIInventoryAPsReport55topAPModelsData, 0)
	return m
}

// SCIInventoryAPsReport55topAPModelsDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.55.topAPModels.DataType
type SCIInventoryAPsReport55topAPModelsDataType struct {
	ApCount *int `json:"apCount,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApPercent *float64 `json:"apPercent,omitempty"`

	Index *int `json:"index,omitempty"`
}

func NewSCIInventoryAPsReport55topAPModelsDataType() *SCIInventoryAPsReport55topAPModelsDataType {
	m := new(SCIInventoryAPsReport55topAPModelsDataType)
	return m
}

// SCIInventoryAPsReport56topAPVersionsData
//
// Definition: InventoryAPsReport.InventoryAPsReport.56.topAPVersions.Data
type SCIInventoryAPsReport56topAPVersionsData []*SCIInventoryAPsReport56topAPVersionsDataType

func MakeSCIInventoryAPsReport56topAPVersionsData() SCIInventoryAPsReport56topAPVersionsData {
	m := make(SCIInventoryAPsReport56topAPVersionsData, 0)
	return m
}

// SCIInventoryAPsReport56topAPVersionsDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.56.topAPVersions.DataType
type SCIInventoryAPsReport56topAPVersionsDataType struct {
	ApCount *int `json:"apCount,omitempty"`

	ApFwVersion *string `json:"apFwVersion,omitempty"`

	ApPercent *float64 `json:"apPercent,omitempty"`

	Index *int `json:"index,omitempty"`
}

func NewSCIInventoryAPsReport56topAPVersionsDataType() *SCIInventoryAPsReport56topAPVersionsDataType {
	m := new(SCIInventoryAPsReport56topAPVersionsDataType)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineData
//
// Definition: InventoryAPsReport.InventoryAPsReport.57.topAPsOffline.Data
type SCIInventoryAPsReport57topAPsOfflineData []*SCIInventoryAPsReport57topAPsOfflineDataType

func MakeSCIInventoryAPsReport57topAPsOfflineData() SCIInventoryAPsReport57topAPsOfflineData {
	m := make(SCIInventoryAPsReport57topAPsOfflineData, 0)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.57.topAPsOffline.DataType
type SCIInventoryAPsReport57topAPsOfflineDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApLocation *string `json:"apLocation,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApName *string `json:"apName,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	DisconnectDuration *int `json:"disconnectDuration,omitempty"`

	Index *int `json:"index,omitempty"`
}

func NewSCIInventoryAPsReport57topAPsOfflineDataType() *SCIInventoryAPsReport57topAPsOfflineDataType {
	m := new(SCIInventoryAPsReport57topAPsOfflineDataType)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineMetaData
//
// Definition: InventoryAPsReport.InventoryAPsReport.57.topAPsOffline.MetaData
type SCIInventoryAPsReport57topAPsOfflineMetaData struct {
	MaxValues *SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIInventoryAPsReport57topAPsOfflineMetaData() *SCIInventoryAPsReport57topAPsOfflineMetaData {
	m := new(SCIInventoryAPsReport57topAPsOfflineMetaData)
	return m
}

// SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType
//
// Definition: InventoryAPsReport.InventoryAPsReport.57.topAPsOffline.MetaDataMaxValuesType
type SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType struct {
	DisconnectDuration *int `json:"disconnectDuration,omitempty"`
}

func NewSCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType() *SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType {
	m := new(SCIInventoryAPsReport57topAPsOfflineMetaDataMaxValuesType)
	return m
}

// SCIInventoryAPsReport58topAPsByRebootsData
//
// Definition: InventoryAPsReport.InventoryAPsReport.58.topAPsByReboots.Data
type SCIInventoryAPsReport58topAPsByRebootsData []*SCIInventoryAPsReport58topAPsByRebootsDataType

func MakeSCIInventoryAPsReport58topAPsByRebootsData() SCIInventoryAPsReport58topAPsByRebootsData {
	m := make(SCIInventoryAPsReport58topAPsByRebootsData, 0)
	return m
}

// SCIInventoryAPsReport58topAPsByRebootsDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.58.topAPsByReboots.DataType
type SCIInventoryAPsReport58topAPsByRebootsDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	Count *int `json:"count,omitempty"`

	Index *int `json:"index,omitempty"`

	LastRebootTime *int `json:"lastRebootTime,omitempty"`
}

func NewSCIInventoryAPsReport58topAPsByRebootsDataType() *SCIInventoryAPsReport58topAPsByRebootsDataType {
	m := new(SCIInventoryAPsReport58topAPsByRebootsDataType)
	return m
}

// SCIInventoryAPsReport60apDetailsOnOfflineStatusData
//
// Definition: InventoryAPsReport.InventoryAPsReport.60.apDetailsOnOfflineStatus.Data
type SCIInventoryAPsReport60apDetailsOnOfflineStatusData []*SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType

func MakeSCIInventoryAPsReport60apDetailsOnOfflineStatusData() SCIInventoryAPsReport60apDetailsOnOfflineStatusData {
	m := make(SCIInventoryAPsReport60apDetailsOnOfflineStatusData, 0)
	return m
}

// SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType
//
// Definition: InventoryAPsReport.InventoryAPsReport.60.apDetailsOnOfflineStatus.DataType
type SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApLocation *string `json:"apLocation,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	Index *int `json:"index,omitempty"`

	LastStatusChangeTime *string `json:"lastStatusChangeTime,omitempty"`
}

func NewSCIInventoryAPsReport60apDetailsOnOfflineStatusDataType() *SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType {
	m := new(SCIInventoryAPsReport60apDetailsOnOfflineStatusDataType)
	return m
}

// ReportInventoryAPsReport46ApInventoryOverview
//
// Operation ID: report.InventoryAPsReport.46.apInventoryOverview
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
// Operation ID: report.InventoryAPsReport.47.topApsDisconnection
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
// Operation ID: report.InventoryAPsReport.48.apCountTrend
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
// Operation ID: report.InventoryAPsReport.49.apStatusTrend
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
// Operation ID: report.InventoryAPsReport.50.topApsModelsChart
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
// Operation ID: report.InventoryAPsReport.51.top10ApVersionsChart
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
// Operation ID: report.InventoryAPsReport.52.topApsRebootReasons
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
// Operation ID: report.InventoryAPsReport.53.top10ApsRebootCounts
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
// Operation ID: report.InventoryAPsReport.54.topApAlarmTypes
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
// Operation ID: report.InventoryAPsReport.55.topAPModels
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
// Operation ID: report.InventoryAPsReport.56.topAPVersions
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
// Operation ID: report.InventoryAPsReport.57.topAPsOffline
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
// Operation ID: report.InventoryAPsReport.58.topAPsByReboots
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
// Operation ID: report.InventoryAPsReport.59.apsConfiguredInMultiCtrl
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
// Operation ID: report.InventoryAPsReport.60.apDetailsOnOfflineStatus
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
// Operation ID: report.InventoryAPsReport.61.apDetailsOtherStatus
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
