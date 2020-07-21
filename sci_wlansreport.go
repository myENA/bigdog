package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIWLANsReportService struct {
	apiClient *SCIClient
}

func NewSCIWLANsReportService(c *SCIClient) *SCIWLANsReportService {
	s := new(SCIWLANsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWLANsReportService() *SCIWLANsReportService {
	return NewSCIWLANsReportService(ss.apiClient)
}

// SCIWLANsReport35overviewData
//
// Definition: WLANsReport_WLANsReport_35_overview_Data
type SCIWLANsReport35overviewData []*SCIWLANsReport35overviewDataType

func MakeSCIWLANsReport35overviewData() SCIWLANsReport35overviewData {
	m := make(SCIWLANsReport35overviewData, 0)
	return m
}

// SCIWLANsReport35overviewDataType
//
// Definition: WLANsReport_WLANsReport_35_overview_DataType
type SCIWLANsReport35overviewDataType struct {
	TotalAddedClientCount *float64 `json:"totalAddedClientCount,omitempty"`

	TotalAddedSsids *int `json:"totalAddedSsids,omitempty"`

	TotalClientCount *float64 `json:"totalClientCount,omitempty"`

	TotalDeletedClientCount *float64 `json:"totalDeletedClientCount,omitempty"`

	TotalDeletedSsids *float64 `json:"totalDeletedSsids,omitempty"`

	TotalSsids *int `json:"totalSsids,omitempty"`
}

func NewSCIWLANsReport35overviewDataType() *SCIWLANsReport35overviewDataType {
	m := new(SCIWLANsReport35overviewDataType)
	return m
}

// SCIWLANsReport37activeSsidsTrendData
//
// Definition: WLANsReport_WLANsReport_37_activeSsidsTrend_Data
type SCIWLANsReport37activeSsidsTrendData [][]*SCIWLANsReport37activeSsidsTrendDataTypeType

func MakeSCIWLANsReport37activeSsidsTrendData() SCIWLANsReport37activeSsidsTrendData {
	m := make(SCIWLANsReport37activeSsidsTrendData, 0)
	return m
}

// SCIWLANsReport37activeSsidsTrendDataType
//
// Definition: WLANsReport_WLANsReport_37_activeSsidsTrend_DataType
type SCIWLANsReport37activeSsidsTrendDataType []*SCIWLANsReport37activeSsidsTrendDataTypeType

func MakeSCIWLANsReport37activeSsidsTrendDataType() SCIWLANsReport37activeSsidsTrendDataType {
	m := make(SCIWLANsReport37activeSsidsTrendDataType, 0)
	return m
}

// SCIWLANsReport37activeSsidsTrendDataTypeType
//
// Definition: WLANsReport_WLANsReport_37_activeSsidsTrend_DataTypeType
type SCIWLANsReport37activeSsidsTrendDataTypeType struct {
	End *string `json:"end,omitempty"`

	Ssid *float64 `json:"ssid,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIWLANsReport37activeSsidsTrendDataTypeType() *SCIWLANsReport37activeSsidsTrendDataTypeType {
	m := new(SCIWLANsReport37activeSsidsTrendDataTypeType)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountData
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_Data
type SCIWLANsReport38top10SsidsByClientCountData [][]*SCIWLANsReport38top10SsidsByClientCountDataTypeType

func MakeSCIWLANsReport38top10SsidsByClientCountData() SCIWLANsReport38top10SsidsByClientCountData {
	m := make(SCIWLANsReport38top10SsidsByClientCountData, 0)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountDataType
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_DataType
type SCIWLANsReport38top10SsidsByClientCountDataType []*SCIWLANsReport38top10SsidsByClientCountDataTypeType

func MakeSCIWLANsReport38top10SsidsByClientCountDataType() SCIWLANsReport38top10SsidsByClientCountDataType {
	m := make(SCIWLANsReport38top10SsidsByClientCountDataType, 0)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountDataTypeType
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_DataTypeType
type SCIWLANsReport38top10SsidsByClientCountDataTypeType struct {
	End *string `json:"end,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Start *string `json:"start,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`
}

func NewSCIWLANsReport38top10SsidsByClientCountDataTypeType() *SCIWLANsReport38top10SsidsByClientCountDataTypeType {
	m := new(SCIWLANsReport38top10SsidsByClientCountDataTypeType)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountMetaData
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_MetaData
type SCIWLANsReport38top10SsidsByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`
}

func NewSCIWLANsReport38top10SsidsByClientCountMetaData() *SCIWLANsReport38top10SsidsByClientCountMetaData {
	m := new(SCIWLANsReport38top10SsidsByClientCountMetaData)
	return m
}

// SCIWLANsReport39ssidChangesOverTimeData
//
// Definition: WLANsReport_WLANsReport_39_ssidChangesOverTime_Data
type SCIWLANsReport39ssidChangesOverTimeData []*SCIWLANsReport39ssidChangesOverTimeDataType

func MakeSCIWLANsReport39ssidChangesOverTimeData() SCIWLANsReport39ssidChangesOverTimeData {
	m := make(SCIWLANsReport39ssidChangesOverTimeData, 0)
	return m
}

// SCIWLANsReport39ssidChangesOverTimeDataType
//
// Definition: WLANsReport_WLANsReport_39_ssidChangesOverTime_DataType
type SCIWLANsReport39ssidChangesOverTimeDataType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Status *string `json:"status,omitempty"`

	Time *int `json:"time,omitempty"`
}

func NewSCIWLANsReport39ssidChangesOverTimeDataType() *SCIWLANsReport39ssidChangesOverTimeDataType {
	m := new(SCIWLANsReport39ssidChangesOverTimeDataType)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableData
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_Data
type SCIWLANsReport40topSsidsByTrafficTableData []*SCIWLANsReport40topSsidsByTrafficTableDataType

func MakeSCIWLANsReport40topSsidsByTrafficTableData() SCIWLANsReport40topSsidsByTrafficTableData {
	m := make(SCIWLANsReport40topSsidsByTrafficTableData, 0)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableDataType
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_DataType
type SCIWLANsReport40topSsidsByTrafficTableDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *int `json:"index,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIWLANsReport40topSsidsByTrafficTableDataType() *SCIWLANsReport40topSsidsByTrafficTableDataType {
	m := new(SCIWLANsReport40topSsidsByTrafficTableDataType)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableMetaData
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_MetaData
type SCIWLANsReport40topSsidsByTrafficTableMetaData struct {
	MaxValues *SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIWLANsReport40topSsidsByTrafficTableMetaData() *SCIWLANsReport40topSsidsByTrafficTableMetaData {
	m := new(SCIWLANsReport40topSsidsByTrafficTableMetaData)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_MetaDataMaxValuesType
type SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType struct {
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

func NewSCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType() *SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableData
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_Data
type SCIWLANsReport41topSsidsByClientsTableData []*SCIWLANsReport41topSsidsByClientsTableDataType

func MakeSCIWLANsReport41topSsidsByClientsTableData() SCIWLANsReport41topSsidsByClientsTableData {
	m := make(SCIWLANsReport41topSsidsByClientsTableData, 0)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableDataType
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_DataType
type SCIWLANsReport41topSsidsByClientsTableDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *int `json:"index,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIWLANsReport41topSsidsByClientsTableDataType() *SCIWLANsReport41topSsidsByClientsTableDataType {
	m := new(SCIWLANsReport41topSsidsByClientsTableDataType)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableMetaData
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_MetaData
type SCIWLANsReport41topSsidsByClientsTableMetaData struct {
	MaxValues *SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIWLANsReport41topSsidsByClientsTableMetaData() *SCIWLANsReport41topSsidsByClientsTableMetaData {
	m := new(SCIWLANsReport41topSsidsByClientsTableMetaData)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_MetaDataMaxValuesType
type SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType struct {
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

func NewSCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType() *SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType {
	m := new(SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType)
	return m
}

// ReportWLANsReport35Overview
//
// Operation ID: report_WLANsReport_35_overview
//
// WLANs Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport35Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport35overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport35overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport35Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport35overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport36Top10SsidsByTraffic
//
// Operation ID: report_WLANsReport_36_top10SsidsByTraffic
//
// WLANs Report - Top SSIDs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport36Top10SsidsByTraffic(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport36top10SsidsByTraffic200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport36top10SsidsByTraffic200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport36Top10SsidsByTraffic, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport36top10SsidsByTraffic200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport37ActiveSsidsTrend
//
// Operation ID: report_WLANsReport_37_activeSsidsTrend
//
// WLANs Report - Active SSIDs Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport37ActiveSsidsTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport37activeSsidsTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport37activeSsidsTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport37ActiveSsidsTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport37activeSsidsTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport38Top10SsidsByClientCount
//
// Operation ID: report_WLANsReport_38_top10SsidsByClientCount
//
// WLANs Report - Top SSIDs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport38Top10SsidsByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport38top10SsidsByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport38top10SsidsByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport38Top10SsidsByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport38top10SsidsByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport39SsidChangesOverTime
//
// Operation ID: report_WLANsReport_39_ssidChangesOverTime
//
// WLANs Report - SSID Changes Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport39SsidChangesOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport39ssidChangesOverTime200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport39ssidChangesOverTime200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport39SsidChangesOverTime, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport39ssidChangesOverTime200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport40TopSsidsByTrafficTable
//
// Operation ID: report_WLANsReport_40_topSsidsByTrafficTable
//
// WLANs Report - Top SSIDs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport40TopSsidsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport40topSsidsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport40TopSsidsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport40topSsidsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWLANsReport41TopSsidsByClientsTable
//
// Operation ID: report_WLANsReport_41_topSsidsByClientsTable
//
// WLANs Report - Top SSIDs by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport41TopSsidsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport41topSsidsByClientsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWLANsReport41topSsidsByClientsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWLANsReport41TopSsidsByClientsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWLANsReport41topSsidsByClientsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
