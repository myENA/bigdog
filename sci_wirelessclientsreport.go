package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIWirelessClientsReportService struct {
	apiClient *SCIClient
}

func NewSCIWirelessClientsReportService(c *SCIClient) *SCIWirelessClientsReportService {
	s := new(SCIWirelessClientsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWirelessClientsReportService() *SCIWirelessClientsReportService {
	return NewSCIWirelessClientsReportService(ss.apiClient)
}

// SCIWirelessClientsReport12overviewData
//
// Definition: WirelessClientsReport.WirelessClientsReport.12.overview.Data
type SCIWirelessClientsReport12overviewData []*SCIWirelessClientsReport12overviewDataType

func MakeSCIWirelessClientsReport12overviewData() SCIWirelessClientsReport12overviewData {
	m := make(SCIWirelessClientsReport12overviewData, 0)
	return m
}

// SCIWirelessClientsReport12overviewDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.12.overview.DataType
type SCIWirelessClientsReport12overviewDataType struct {
	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`
}

func NewSCIWirelessClientsReport12overviewDataType() *SCIWirelessClientsReport12overviewDataType {
	m := new(SCIWirelessClientsReport12overviewDataType)
	return m
}

// SCIWirelessClientsReport14topTableData
//
// Definition: WirelessClientsReport.WirelessClientsReport.14.topTable.Data
type SCIWirelessClientsReport14topTableData []*SCIWirelessClientsReport14topTableDataType

func MakeSCIWirelessClientsReport14topTableData() SCIWirelessClientsReport14topTableData {
	m := make(SCIWirelessClientsReport14topTableData, 0)
	return m
}

// SCIWirelessClientsReport14topTableDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.14.topTable.DataType
type SCIWirelessClientsReport14topTableDataType struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Index *int `json:"index,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	OsType *string `json:"osType,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewSCIWirelessClientsReport14topTableDataType() *SCIWirelessClientsReport14topTableDataType {
	m := new(SCIWirelessClientsReport14topTableDataType)
	return m
}

// SCIWirelessClientsReport14topTableMetaData
//
// Definition: WirelessClientsReport.WirelessClientsReport.14.topTable.MetaData
type SCIWirelessClientsReport14topTableMetaData struct {
	MaxValues *SCIWirelessClientsReport14topTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIWirelessClientsReport14topTableMetaData() *SCIWirelessClientsReport14topTableMetaData {
	m := new(SCIWirelessClientsReport14topTableMetaData)
	return m
}

// SCIWirelessClientsReport14topTableMetaDataMaxValuesType
//
// Definition: WirelessClientsReport.WirelessClientsReport.14.topTable.MetaDataMaxValuesType
type SCIWirelessClientsReport14topTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIWirelessClientsReport14topTableMetaDataMaxValuesType() *SCIWirelessClientsReport14topTableMetaDataMaxValuesType {
	m := new(SCIWirelessClientsReport14topTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessClientsReport15trendChartData
//
// Definition: WirelessClientsReport.WirelessClientsReport.15.trendChart.Data
type SCIWirelessClientsReport15trendChartData []*SCIWirelessClientsReport15trendChartDataType

func MakeSCIWirelessClientsReport15trendChartData() SCIWirelessClientsReport15trendChartData {
	m := make(SCIWirelessClientsReport15trendChartData, 0)
	return m
}

// SCIWirelessClientsReport15trendChartDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.15.trendChart.DataType
type SCIWirelessClientsReport15trendChartDataType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	Uniqueusers *float64 `json:"unique_users,omitempty"`

	Uniqueusers24 *float64 `json:"unique_users_2-4,omitempty"`

	Uniqueusers5 *float64 `json:"unique_users_5,omitempty"`
}

func NewSCIWirelessClientsReport15trendChartDataType() *SCIWirelessClientsReport15trendChartDataType {
	m := new(SCIWirelessClientsReport15trendChartDataType)
	return m
}

// SCIWirelessClientsReport16trendTableData
//
// Definition: WirelessClientsReport.WirelessClientsReport.16.trendTable.Data
type SCIWirelessClientsReport16trendTableData []*SCIWirelessClientsReport16trendTableDataType

func MakeSCIWirelessClientsReport16trendTableData() SCIWirelessClientsReport16trendTableData {
	m := make(SCIWirelessClientsReport16trendTableData, 0)
	return m
}

// SCIWirelessClientsReport16trendTableDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.16.trendTable.DataType
type SCIWirelessClientsReport16trendTableDataType struct {
	End *string `json:"end,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	Start *string `json:"start,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	Uniqueusers *float64 `json:"unique_users,omitempty"`

	Uniqueusers24 *float64 `json:"unique_users_2-4,omitempty"`

	Uniqueusers5 *float64 `json:"unique_users_5,omitempty"`
}

func NewSCIWirelessClientsReport16trendTableDataType() *SCIWirelessClientsReport16trendTableDataType {
	m := new(SCIWirelessClientsReport16trendTableDataType)
	return m
}

// SCIWirelessClientsReport16trendTableMetaData
//
// Definition: WirelessClientsReport.WirelessClientsReport.16.trendTable.MetaData
type SCIWirelessClientsReport16trendTableMetaData struct {
	MaxValues *SCIWirelessClientsReport16trendTableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIWirelessClientsReport16trendTableMetaData() *SCIWirelessClientsReport16trendTableMetaData {
	m := new(SCIWirelessClientsReport16trendTableMetaData)
	return m
}

// SCIWirelessClientsReport16trendTableMetaDataMaxValuesType
//
// Definition: WirelessClientsReport.WirelessClientsReport.16.trendTable.MetaDataMaxValuesType
type SCIWirelessClientsReport16trendTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	Uniqueusers *float64 `json:"unique_users,omitempty"`

	Uniqueusers24 *float64 `json:"unique_users_2-4,omitempty"`

	Uniqueusers5 *float64 `json:"unique_users_5,omitempty"`
}

func NewSCIWirelessClientsReport16trendTableMetaDataMaxValuesType() *SCIWirelessClientsReport16trendTableMetaDataMaxValuesType {
	m := new(SCIWirelessClientsReport16trendTableMetaDataMaxValuesType)
	return m
}

// SCIWirelessClientsReport17topPercentileData
//
// Definition: WirelessClientsReport.WirelessClientsReport.17.topPercentile.Data
type SCIWirelessClientsReport17topPercentileData []*SCIWirelessClientsReport17topPercentileDataType

func MakeSCIWirelessClientsReport17topPercentileData() SCIWirelessClientsReport17topPercentileData {
	m := make(SCIWirelessClientsReport17topPercentileData, 0)
	return m
}

// SCIWirelessClientsReport17topPercentileDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.17.topPercentile.DataType
type SCIWirelessClientsReport17topPercentileDataType struct {
	ClientCount *int `json:"clientCount,omitempty"`

	Percent *string `json:"percent,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`
}

func NewSCIWirelessClientsReport17topPercentileDataType() *SCIWirelessClientsReport17topPercentileDataType {
	m := new(SCIWirelessClientsReport17topPercentileDataType)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountData
//
// Definition: WirelessClientsReport.WirelessClientsReport.18.topNOSByClientCount.Data
type SCIWirelessClientsReport18topNOSByClientCountData []*SCIWirelessClientsReport18topNOSByClientCountDataType

func MakeSCIWirelessClientsReport18topNOSByClientCountData() SCIWirelessClientsReport18topNOSByClientCountData {
	m := make(SCIWirelessClientsReport18topNOSByClientCountData, 0)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.18.topNOSByClientCount.DataType
type SCIWirelessClientsReport18topNOSByClientCountDataType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	OsType *string `json:"osType,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIWirelessClientsReport18topNOSByClientCountDataType() *SCIWirelessClientsReport18topNOSByClientCountDataType {
	m := new(SCIWirelessClientsReport18topNOSByClientCountDataType)
	return m
}

// SCIWirelessClientsReport18topNOSByClientCountMetaData
//
// Definition: WirelessClientsReport.WirelessClientsReport.18.topNOSByClientCount.MetaData
type SCIWirelessClientsReport18topNOSByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`
}

func NewSCIWirelessClientsReport18topNOSByClientCountMetaData() *SCIWirelessClientsReport18topNOSByClientCountMetaData {
	m := new(SCIWirelessClientsReport18topNOSByClientCountMetaData)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountData
//
// Definition: WirelessClientsReport.WirelessClientsReport.19.top10ManufacturersByClientCount.Data
type SCIWirelessClientsReport19top10ManufacturersByClientCountData []*SCIWirelessClientsReport19top10ManufacturersByClientCountDataType

func MakeSCIWirelessClientsReport19top10ManufacturersByClientCountData() SCIWirelessClientsReport19top10ManufacturersByClientCountData {
	m := make(SCIWirelessClientsReport19top10ManufacturersByClientCountData, 0)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.19.top10ManufacturersByClientCount.DataType
type SCIWirelessClientsReport19top10ManufacturersByClientCountDataType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIWirelessClientsReport19top10ManufacturersByClientCountDataType() *SCIWirelessClientsReport19top10ManufacturersByClientCountDataType {
	m := new(SCIWirelessClientsReport19top10ManufacturersByClientCountDataType)
	return m
}

// SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData
//
// Definition: WirelessClientsReport.WirelessClientsReport.19.top10ManufacturersByClientCount.MetaData
type SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`
}

func NewSCIWirelessClientsReport19top10ManufacturersByClientCountMetaData() *SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData {
	m := new(SCIWirelessClientsReport19top10ManufacturersByClientCountMetaData)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData
//
// Definition: WirelessClientsReport.WirelessClientsReport.112.top10AuthenticationMechanismByClientCount.Data
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData []*SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType

func MakeSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData() SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData {
	m := make(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountData, 0)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType
//
// Definition: WirelessClientsReport.WirelessClientsReport.112.top10AuthenticationMechanismByClientCount.DataType
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType struct {
	AuthMethod *string `json:"authMethod,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType() *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType {
	m := new(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountDataType)
	return m
}

// SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData
//
// Definition: WirelessClientsReport.WirelessClientsReport.112.top10AuthenticationMechanismByClientCount.MetaData
type SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`
}

func NewSCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData() *SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData {
	m := new(SCIWirelessClientsReport112top10AuthenticationMechanismByClientCountMetaData)
	return m
}

// ReportWirelessClientsReport12Overview
//
// Operation ID: report.WirelessClientsReport.12.overview
//
// Wireless - Clients Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport12Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport12overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport12overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport12Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport12overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport13TopChart
//
// Operation ID: report.WirelessClientsReport.13.topChart
//
// Wireless - Clients Report - Top 10 Unique Clients by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport13TopChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport13topChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport13topChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport13TopChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport13topChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport14TopTable
//
// Operation ID: report.WirelessClientsReport.14.topTable
//
// Wireless - Clients Report - Clients Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport14TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport14topTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport14topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport14TopTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport14topTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport15TrendChart
//
// Operation ID: report.WirelessClientsReport.15.trendChart
//
// Wireless - Clients Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport15TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport15trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport15trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport15TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport15trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport16TrendTable
//
// Operation ID: report.WirelessClientsReport.16.trendTable
//
// Wireless - Clients Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport16TrendTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport16trendTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport16trendTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport16TrendTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport16trendTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport17TopPercentile
//
// Operation ID: report.WirelessClientsReport.17.topPercentile
//
// Wireless - Clients Report - Top Clients by Traffic Percentile
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport17TopPercentile(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport17topPercentile200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport17topPercentile200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport17TopPercentile, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport17topPercentile200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport18TopNOSByClientCount
//
// Operation ID: report.WirelessClientsReport.18.topNOSByClientCount
//
// Wireless - Clients Report - Top 10 OS by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport18TopNOSByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport18topNOSByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport18TopNOSByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport18topNOSByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport19Top10ManufacturersByClientCount
//
// Operation ID: report.WirelessClientsReport.19.top10ManufacturersByClientCount
//
// Wireless - Clients Report - Top 10 Manufacturers by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport19Top10ManufacturersByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport19Top10ManufacturersByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport19top10ManufacturersByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount
//
// Operation ID: report.WirelessClientsReport.112.top10AuthenticationMechanismByClientCount
//
// Wireless - Clients Report - Top 10 Authentication Methods by Client Count
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWirelessClientsReportService) ReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWirelessClientsReport112Top10AuthenticationMechanismByClientCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportWirelessClientsReport112top10AuthenticationMechanismByClientCount200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
