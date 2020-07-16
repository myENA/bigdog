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

// SCINetworkWirelessReport20overviewData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_20_overview_Data
type SCINetworkWirelessReport20overviewData []*SCINetworkWirelessReport20overviewDataType

func MakeSCINetworkWirelessReport20overviewData() SCINetworkWirelessReport20overviewData {
	m := make(SCINetworkWirelessReport20overviewData, 0)
	return m
}

// SCINetworkWirelessReport20overviewDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_20_overview_DataType
type SCINetworkWirelessReport20overviewDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	TotalAps *int `json:"totalAps,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`
}

func NewSCINetworkWirelessReport20overviewDataType() *SCINetworkWirelessReport20overviewDataType {
	m := new(SCINetworkWirelessReport20overviewDataType)
	return m
}

// SCINetworkWirelessReport21trafficDistributionData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_Data
type SCINetworkWirelessReport21trafficDistributionData []*SCINetworkWirelessReport21trafficDistributionDataType

func MakeSCINetworkWirelessReport21trafficDistributionData() SCINetworkWirelessReport21trafficDistributionData {
	m := make(SCINetworkWirelessReport21trafficDistributionData, 0)
	return m
}

// SCINetworkWirelessReport21trafficDistributionDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_DataType
type SCINetworkWirelessReport21trafficDistributionDataType struct {
	Children []*SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType `json:"children,omitempty"`

	Key *string `json:"key,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSCINetworkWirelessReport21trafficDistributionDataType() *SCINetworkWirelessReport21trafficDistributionDataType {
	m := new(SCINetworkWirelessReport21trafficDistributionDataType)
	return m
}

// SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_DataTypeChildrenType
type SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType struct {
	Children []*SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType `json:"children,omitempty"`

	Display *string `json:"display,omitempty"`

	Key *string `json:"key,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSCINetworkWirelessReport21trafficDistributionDataTypeChildrenType() *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType {
	m := new(SCINetworkWirelessReport21trafficDistributionDataTypeChildrenType)
	return m
}

// SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_21_trafficDistribution_DataTypeChildrenTypeChildrenType
type SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType struct {
	Display *string `json:"display,omitempty"`

	Key *string `json:"key,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int `json:"size,omitempty"`
}

func NewSCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType() *SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType {
	m := new(SCINetworkWirelessReport21trafficDistributionDataTypeChildrenTypeChildrenType)
	return m
}

// SCINetworkWirelessReport22trafficTrendData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_22_trafficTrend_Data
type SCINetworkWirelessReport22trafficTrendData []*SCINetworkWirelessReport22trafficTrendDataType

func MakeSCINetworkWirelessReport22trafficTrendData() SCINetworkWirelessReport22trafficTrendData {
	m := make(SCINetworkWirelessReport22trafficTrendData, 0)
	return m
}

// SCINetworkWirelessReport22trafficTrendDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_22_trafficTrend_DataType
type SCINetworkWirelessReport22trafficTrendDataType struct {
	End *string `json:"end,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`
}

func NewSCINetworkWirelessReport22trafficTrendDataType() *SCINetworkWirelessReport22trafficTrendDataType {
	m := new(SCINetworkWirelessReport22trafficTrendDataType)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_Data
type SCINetworkWirelessReport23trafficOverTimeTableData []*SCINetworkWirelessReport23trafficOverTimeTableDataType

func MakeSCINetworkWirelessReport23trafficOverTimeTableData() SCINetworkWirelessReport23trafficOverTimeTableData {
	m := make(SCINetworkWirelessReport23trafficOverTimeTableData, 0)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_DataType
type SCINetworkWirelessReport23trafficOverTimeTableDataType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	End *string `json:"end,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount interface{} `json:"sessionCount,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTraffic24 *int `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *int `json:"totalTraffic_5,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UniqueUsers interface{} `json:"uniqueUsers,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCINetworkWirelessReport23trafficOverTimeTableDataType() *SCINetworkWirelessReport23trafficOverTimeTableDataType {
	m := new(SCINetworkWirelessReport23trafficOverTimeTableDataType)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_MetaData
type SCINetworkWirelessReport23trafficOverTimeTableMetaData struct {
	MaxValues *SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCINetworkWirelessReport23trafficOverTimeTableMetaData() *SCINetworkWirelessReport23trafficOverTimeTableMetaData {
	m := new(SCINetworkWirelessReport23trafficOverTimeTableMetaData)
	return m
}

// SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_23_trafficOverTimeTable_MetaDataMaxValuesType
type SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTraffic24 *int `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *int `json:"totalTraffic_5,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UniqueUsers interface{} `json:"uniqueUsers,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType() *SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType {
	m := new(SCINetworkWirelessReport23trafficOverTimeTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_Data
type SCINetworkWirelessReport24topAPsByTrafficTableData []*SCINetworkWirelessReport24topAPsByTrafficTableDataType

func MakeSCINetworkWirelessReport24topAPsByTrafficTableData() SCINetworkWirelessReport24topAPsByTrafficTableData {
	m := make(SCINetworkWirelessReport24topAPsByTrafficTableData, 0)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_DataType
type SCINetworkWirelessReport24topAPsByTrafficTableDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

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

func NewSCINetworkWirelessReport24topAPsByTrafficTableDataType() *SCINetworkWirelessReport24topAPsByTrafficTableDataType {
	m := new(SCINetworkWirelessReport24topAPsByTrafficTableDataType)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_MetaData
type SCINetworkWirelessReport24topAPsByTrafficTableMetaData struct {
	MaxValues *SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCINetworkWirelessReport24topAPsByTrafficTableMetaData() *SCINetworkWirelessReport24topAPsByTrafficTableMetaData {
	m := new(SCINetworkWirelessReport24topAPsByTrafficTableMetaData)
	return m
}

// SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_24_topAPsByTrafficTable_MetaDataMaxValuesType
type SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

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

func NewSCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType() *SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType {
	m := new(SCINetworkWirelessReport24topAPsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_Data
type SCINetworkWirelessReport25topAPsByClientsTableData []*SCINetworkWirelessReport25topAPsByClientsTableDataType

func MakeSCINetworkWirelessReport25topAPsByClientsTableData() SCINetworkWirelessReport25topAPsByClientsTableData {
	m := make(SCINetworkWirelessReport25topAPsByClientsTableData, 0)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_DataType
type SCINetworkWirelessReport25topAPsByClientsTableDataType struct {
	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

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

func NewSCINetworkWirelessReport25topAPsByClientsTableDataType() *SCINetworkWirelessReport25topAPsByClientsTableDataType {
	m := new(SCINetworkWirelessReport25topAPsByClientsTableDataType)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_MetaData
type SCINetworkWirelessReport25topAPsByClientsTableMetaData struct {
	MaxValues *SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCINetworkWirelessReport25topAPsByClientsTableMetaData() *SCINetworkWirelessReport25topAPsByClientsTableMetaData {
	m := new(SCINetworkWirelessReport25topAPsByClientsTableMetaData)
	return m
}

// SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_25_topAPsByClientsTable_MetaDataMaxValuesType
type SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType struct {
	AvgRateTotalRxTraffic *float64 `json:"avgRateTotalRxTraffic,omitempty"`

	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgRateTotalTxTraffic *float64 `json:"avgRateTotalTxTraffic,omitempty"`

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

func NewSCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType() *SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType {
	m := new(SCINetworkWirelessReport25topAPsByClientsTableMetaDataMaxValuesType)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_Data
type SCINetworkWirelessReport27top10ApByClientCountData []*SCINetworkWirelessReport27top10ApByClientCountDataType

func MakeSCINetworkWirelessReport27top10ApByClientCountData() SCINetworkWirelessReport27top10ApByClientCountData {
	m := make(SCINetworkWirelessReport27top10ApByClientCountData, 0)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountDataType
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_DataType
type SCINetworkWirelessReport27top10ApByClientCountDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`
}

func NewSCINetworkWirelessReport27top10ApByClientCountDataType() *SCINetworkWirelessReport27top10ApByClientCountDataType {
	m := new(SCINetworkWirelessReport27top10ApByClientCountDataType)
	return m
}

// SCINetworkWirelessReport27top10ApByClientCountMetaData
//
// Definition: NetworkWirelessReport_NetworkWirelessReport_27_top10ApByClientCount_MetaData
type SCINetworkWirelessReport27top10ApByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`
}

func NewSCINetworkWirelessReport27top10ApByClientCountMetaData() *SCINetworkWirelessReport27top10ApByClientCountMetaData {
	m := new(SCINetworkWirelessReport27top10ApByClientCountMetaData)
	return m
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
