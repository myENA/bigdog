package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIAirtimeUtilizationReportService struct {
	apiClient *SCIClient
}

func NewSCIAirtimeUtilizationReportService(c *SCIClient) *SCIAirtimeUtilizationReportService {
	s := new(SCIAirtimeUtilizationReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAirtimeUtilizationReportService() *SCIAirtimeUtilizationReportService {
	return NewSCIAirtimeUtilizationReportService(ss.apiClient)
}

// SCIAirtimeUtilizationReport1overviewData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_1_overview_Data
type SCIAirtimeUtilizationReport1overviewData []*SCIAirtimeUtilizationReport1overviewDataType

func MakeSCIAirtimeUtilizationReport1overviewData() SCIAirtimeUtilizationReport1overviewData {
	m := make(SCIAirtimeUtilizationReport1overviewData, 0)
	return m
}

// SCIAirtimeUtilizationReport1overviewDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_1_overview_DataType
type SCIAirtimeUtilizationReport1overviewDataType struct {
	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	Radio *string `json:"radio,omitempty"`
}

func NewSCIAirtimeUtilizationReport1overviewDataType() *SCIAirtimeUtilizationReport1overviewDataType {
	m := new(SCIAirtimeUtilizationReport1overviewDataType)
	return m
}

// SCIAirtimeUtilizationReport2topChartData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_2_topChart_Data
type SCIAirtimeUtilizationReport2topChartData []*SCIAirtimeUtilizationReport2topChartDataType

func MakeSCIAirtimeUtilizationReport2topChartData() SCIAirtimeUtilizationReport2topChartData {
	m := make(SCIAirtimeUtilizationReport2topChartData, 0)
	return m
}

// SCIAirtimeUtilizationReport2topChartDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_2_topChart_DataType
type SCIAirtimeUtilizationReport2topChartDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewSCIAirtimeUtilizationReport2topChartDataType() *SCIAirtimeUtilizationReport2topChartDataType {
	m := new(SCIAirtimeUtilizationReport2topChartDataType)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_Data
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableData []*SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType

func MakeSCIAirtimeUtilizationReport3topAPsByAirtime24TableData() SCIAirtimeUtilizationReport3topAPsByAirtime24TableData {
	m := make(SCIAirtimeUtilizationReport3topAPsByAirtime24TableData, 0)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_DataType
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	Index *int `json:"index,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType() *SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType {
	m := new(SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_MetaData
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData struct {
	MaxValues *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData() *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData {
	m := new(SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_MetaDataMaxValuesType
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType struct {
	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType() *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType {
	m := new(SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_Data
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableData []*SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType

func MakeSCIAirtimeUtilizationReport4topAPsByAirtime5TableData() SCIAirtimeUtilizationReport4topAPsByAirtime5TableData {
	m := make(SCIAirtimeUtilizationReport4topAPsByAirtime5TableData, 0)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_DataType
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	Index *int `json:"index,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType() *SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType {
	m := new(SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_MetaData
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData struct {
	MaxValues *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData() *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData {
	m := new(SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_MetaDataMaxValuesType
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType struct {
	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType() *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType {
	m := new(SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType)
	return m
}

// SCIAirtimeUtilizationReport5trendChartData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_5_trendChart_Data
type SCIAirtimeUtilizationReport5trendChartData [][]*SCIAirtimeUtilizationReport5trendChartDataTypeType

func MakeSCIAirtimeUtilizationReport5trendChartData() SCIAirtimeUtilizationReport5trendChartData {
	m := make(SCIAirtimeUtilizationReport5trendChartData, 0)
	return m
}

// SCIAirtimeUtilizationReport5trendChartDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_5_trendChart_DataType
type SCIAirtimeUtilizationReport5trendChartDataType []*SCIAirtimeUtilizationReport5trendChartDataTypeType

func MakeSCIAirtimeUtilizationReport5trendChartDataType() SCIAirtimeUtilizationReport5trendChartDataType {
	m := make(SCIAirtimeUtilizationReport5trendChartDataType, 0)
	return m
}

// SCIAirtimeUtilizationReport5trendChartDataTypeType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_5_trendChart_DataTypeType
type SCIAirtimeUtilizationReport5trendChartDataTypeType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIAirtimeUtilizationReport5trendChartDataTypeType() *SCIAirtimeUtilizationReport5trendChartDataTypeType {
	m := new(SCIAirtimeUtilizationReport5trendChartDataTypeType)
	return m
}

// SCIAirtimeUtilizationReport6trendTableData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_Data
type SCIAirtimeUtilizationReport6trendTableData []*SCIAirtimeUtilizationReport6trendTableDataType

func MakeSCIAirtimeUtilizationReport6trendTableData() SCIAirtimeUtilizationReport6trendTableData {
	m := make(SCIAirtimeUtilizationReport6trendTableData, 0)
	return m
}

// SCIAirtimeUtilizationReport6trendTableDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_DataType
type SCIAirtimeUtilizationReport6trendTableDataType struct {
	AirtimeBusyAvg24 *float64 `json:"airtimeBusyAvg_2-4,omitempty"`

	AirtimeBusyAvg5 *float64 `json:"airtimeBusyAvg_5,omitempty"`

	AirtimeIdleAvg24 *float64 `json:"airtimeIdleAvg_2-4,omitempty"`

	AirtimeIdleAvg5 *float64 `json:"airtimeIdleAvg_5,omitempty"`

	AirtimeRxAvg24 *float64 `json:"airtimeRxAvg_2-4,omitempty"`

	AirtimeRxAvg5 *float64 `json:"airtimeRxAvg_5,omitempty"`

	AirtimeTxAvg24 *float64 `json:"airtimeTxAvg_2-4,omitempty"`

	AirtimeTxAvg5 *float64 `json:"airtimeTxAvg_5,omitempty"`

	AirtimeUtilizationAvg24 *float64 `json:"airtimeUtilizationAvg_2-4,omitempty"`

	AirtimeUtilizationAvg5 *float64 `json:"airtimeUtilizationAvg_5,omitempty"`

	ClientCount interface{} `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount interface{} `json:"sessionCount,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport6trendTableDataType() *SCIAirtimeUtilizationReport6trendTableDataType {
	m := new(SCIAirtimeUtilizationReport6trendTableDataType)
	return m
}

// SCIAirtimeUtilizationReport6trendTableMetaData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_MetaData
type SCIAirtimeUtilizationReport6trendTableMetaData struct {
	MaxValues *SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIAirtimeUtilizationReport6trendTableMetaData() *SCIAirtimeUtilizationReport6trendTableMetaData {
	m := new(SCIAirtimeUtilizationReport6trendTableMetaData)
	return m
}

// SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_MetaDataMaxValuesType
type SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType struct {
	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType() *SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType {
	m := new(SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType)
	return m
}

// ReportAirtimeUtilizationReport1Overview
//
// Operation ID: report_AirtimeUtilizationReport_1_overview
//
// Airtime Utilization Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport1Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport1overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport1overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport1Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport1overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport2TopChart
//
// Operation ID: report_AirtimeUtilizationReport_2_topChart
//
// Airtime Utilization Report - Top 10 APs by Airtime Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport2TopChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport2topChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport2topChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport2TopChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport2topChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport3TopAPsByAirtime24Table
//
// Operation ID: report_AirtimeUtilizationReport_3_topAPsByAirtime24Table
//
// Airtime Utilization Report - Top APs by Airtime Utilization for 2.4 GHz
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport3TopAPsByAirtime24Table(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport3TopAPsByAirtime24Table, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport4TopAPsByAirtime5Table
//
// Operation ID: report_AirtimeUtilizationReport_4_topAPsByAirtime5Table
//
// Airtime Utilization Report - Top APs by Airtime Utilization for 5 GHz
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport4TopAPsByAirtime5Table(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport4TopAPsByAirtime5Table, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport5TrendChart
//
// Operation ID: report_AirtimeUtilizationReport_5_trendChart
//
// Airtime Utilization Report - Airtime Utilization Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport5TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport5trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport5trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport5TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport5trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAirtimeUtilizationReport6TrendTable
//
// Operation ID: report_AirtimeUtilizationReport_6_trendTable
//
// Airtime Utilization Report - Airtime Utilization Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport6TrendTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport6trendTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAirtimeUtilizationReport6trendTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAirtimeUtilizationReport6TrendTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAirtimeUtilizationReport6trendTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
