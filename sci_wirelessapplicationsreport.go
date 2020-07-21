package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIWirelessApplicationsReportService struct {
	apiClient *SCIClient
}

func NewSCIWirelessApplicationsReportService(c *SCIClient) *SCIWirelessApplicationsReportService {
	s := new(SCIWirelessApplicationsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWirelessApplicationsReportService() *SCIWirelessApplicationsReportService {
	return NewSCIWirelessApplicationsReportService(ss.apiClient)
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_Data
type SCIWirelessApplicationsReport8topAppsByTrafficTableData []*SCIWirelessApplicationsReport8topAppsByTrafficTableDataType

func MakeSCIWirelessApplicationsReport8topAppsByTrafficTableData() SCIWirelessApplicationsReport8topAppsByTrafficTableData {
	m := make(SCIWirelessApplicationsReport8topAppsByTrafficTableData, 0)
	return m
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_DataType
type SCIWirelessApplicationsReport8topAppsByTrafficTableDataType struct {
	App *string `json:"app,omitempty"`

	Index *int `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIWirelessApplicationsReport8topAppsByTrafficTableDataType() *SCIWirelessApplicationsReport8topAppsByTrafficTableDataType {
	m := new(SCIWirelessApplicationsReport8topAppsByTrafficTableDataType)
	return m
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_MetaData
type SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData struct {
	MaxValues *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIWirelessApplicationsReport8topAppsByTrafficTableMetaData() *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData {
	m := new(SCIWirelessApplicationsReport8topAppsByTrafficTableMetaData)
	return m
}

// SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_8_topAppsByTrafficTable_MetaDataMaxValuesType
type SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType() *SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIWirelessApplicationsReport8topAppsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_Data
type SCIWirelessApplicationsReport9topAppsByClientsTableData []*SCIWirelessApplicationsReport9topAppsByClientsTableDataType

func MakeSCIWirelessApplicationsReport9topAppsByClientsTableData() SCIWirelessApplicationsReport9topAppsByClientsTableData {
	m := make(SCIWirelessApplicationsReport9topAppsByClientsTableData, 0)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_DataType
type SCIWirelessApplicationsReport9topAppsByClientsTableDataType struct {
	App *string `json:"app,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *int `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIWirelessApplicationsReport9topAppsByClientsTableDataType() *SCIWirelessApplicationsReport9topAppsByClientsTableDataType {
	m := new(SCIWirelessApplicationsReport9topAppsByClientsTableDataType)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_MetaData
type SCIWirelessApplicationsReport9topAppsByClientsTableMetaData struct {
	MaxValues *SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIWirelessApplicationsReport9topAppsByClientsTableMetaData() *SCIWirelessApplicationsReport9topAppsByClientsTableMetaData {
	m := new(SCIWirelessApplicationsReport9topAppsByClientsTableMetaData)
	return m
}

// SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_9_topAppsByClientsTable_MetaDataMaxValuesType
type SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType() *SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType {
	m := new(SCIWirelessApplicationsReport9topAppsByClientsTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessApplicationsReport10overviewData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_Data
type SCIWirelessApplicationsReport10overviewData [][]*SCIWirelessApplicationsReport10overviewDataTypeType

func MakeSCIWirelessApplicationsReport10overviewData() SCIWirelessApplicationsReport10overviewData {
	m := make(SCIWirelessApplicationsReport10overviewData, 0)
	return m
}

// SCIWirelessApplicationsReport10overviewDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_DataType
type SCIWirelessApplicationsReport10overviewDataType []*SCIWirelessApplicationsReport10overviewDataTypeType

func MakeSCIWirelessApplicationsReport10overviewDataType() SCIWirelessApplicationsReport10overviewDataType {
	m := make(SCIWirelessApplicationsReport10overviewDataType, 0)
	return m
}

// SCIWirelessApplicationsReport10overviewDataTypeType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_DataTypeType
type SCIWirelessApplicationsReport10overviewDataTypeType struct {
	Applications *float64 `json:"applications,omitempty"`
}

func NewSCIWirelessApplicationsReport10overviewDataTypeType() *SCIWirelessApplicationsReport10overviewDataTypeType {
	m := new(SCIWirelessApplicationsReport10overviewDataTypeType)
	return m
}

// SCIWirelessApplicationsReport10overviewMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_10_overview_MetaData
type SCIWirelessApplicationsReport10overviewMetaData struct {
	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIWirelessApplicationsReport10overviewMetaData() *SCIWirelessApplicationsReport10overviewMetaData {
	m := new(SCIWirelessApplicationsReport10overviewMetaData)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_Data
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountData [][]*SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType

func MakeSCIWirelessApplicationsReport11top10ApplicationsByClientCountData() SCIWirelessApplicationsReport11top10ApplicationsByClientCountData {
	m := make(SCIWirelessApplicationsReport11top10ApplicationsByClientCountData, 0)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_DataType
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType []*SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType

func MakeSCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType() SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType {
	m := make(SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataType, 0)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_DataTypeType
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType struct {
	App *string `json:"app,omitempty"`

	End *string `json:"end,omitempty"`

	Rx *int `json:"rx,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Tx *int `json:"tx,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`
}

func NewSCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType() *SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType {
	m := new(SCIWirelessApplicationsReport11top10ApplicationsByClientCountDataTypeType)
	return m
}

// SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData
//
// Definition: WirelessApplicationsReport_WirelessApplicationsReport_11_top10ApplicationsByClientCount_MetaData
type SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`
}

func NewSCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData() *SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData {
	m := new(SCIWirelessApplicationsReport11top10ApplicationsByClientCountMetaData)
	return m
}

// ReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume
//
// Operation ID: report_WirelessApplicationsReport_7_top10ApplicationsByTrafficVolume
//
// Wireless - Applications Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport7Top10ApplicationsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport7top10ApplicationsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport8TopAppsByTrafficTable
//
// Operation ID: report_WirelessApplicationsReport_8_topAppsByTrafficTable
//
// Wireless - Applications Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport8TopAppsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport8topAppsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport9TopAppsByClientsTable
//
// Operation ID: report_WirelessApplicationsReport_9_topAppsByClientsTable
//
// Wireless - Applications Report - Top Applications by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport9TopAppsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport9TopAppsByClientsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport9topAppsByClientsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport10Overview
//
// Operation ID: report_WirelessApplicationsReport_10_overview
//
// Wireless - Applications Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport10Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport10overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport10overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport10Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport10overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessApplicationsReport11Top10ApplicationsByClientCount
//
// Operation ID: report_WirelessApplicationsReport_11_top10ApplicationsByClientCount
//
// Wireless - Applications Report - Top Applications by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessApplicationsReportService) ReportWirelessApplicationsReport11Top10ApplicationsByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessApplicationsReport11Top10ApplicationsByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessApplicationsReport11top10ApplicationsByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
