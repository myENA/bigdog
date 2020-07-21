package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIClientDetailsReportService struct {
	apiClient *SCIClient
}

func NewSCIClientDetailsReportService(c *SCIClient) *SCIClientDetailsReportService {
	s := new(SCIClientDetailsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIClientDetailsReportService() *SCIClientDetailsReportService {
	return NewSCIClientDetailsReportService(ss.apiClient)
}

// SCIClientDetailsReport8topAppsByTrafficTableData
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_Data
type SCIClientDetailsReport8topAppsByTrafficTableData []*SCIClientDetailsReport8topAppsByTrafficTableDataType

func MakeSCIClientDetailsReport8topAppsByTrafficTableData() SCIClientDetailsReport8topAppsByTrafficTableData {
	m := make(SCIClientDetailsReport8topAppsByTrafficTableData, 0)
	return m
}

// SCIClientDetailsReport8topAppsByTrafficTableDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_DataType
type SCIClientDetailsReport8topAppsByTrafficTableDataType struct {
	App *string `json:"app,omitempty"`

	Index *int `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIClientDetailsReport8topAppsByTrafficTableDataType() *SCIClientDetailsReport8topAppsByTrafficTableDataType {
	m := new(SCIClientDetailsReport8topAppsByTrafficTableDataType)
	return m
}

// SCIClientDetailsReport8topAppsByTrafficTableMetaData
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_MetaData
type SCIClientDetailsReport8topAppsByTrafficTableMetaData struct {
	MaxValues *SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIClientDetailsReport8topAppsByTrafficTableMetaData() *SCIClientDetailsReport8topAppsByTrafficTableMetaData {
	m := new(SCIClientDetailsReport8topAppsByTrafficTableMetaData)
	return m
}

// SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType
//
// Definition: ClientDetailsReport_ClientDetailsReport_8_topAppsByTrafficTable_MetaDataMaxValuesType
type SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType() *SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIClientDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIClientDetailsReport82rssTrendData
//
// Definition: ClientDetailsReport_ClientDetailsReport_82_rssTrend_Data
type SCIClientDetailsReport82rssTrendData []*SCIClientDetailsReport82rssTrendDataType

func MakeSCIClientDetailsReport82rssTrendData() SCIClientDetailsReport82rssTrendData {
	m := make(SCIClientDetailsReport82rssTrendData, 0)
	return m
}

// SCIClientDetailsReport82rssTrendDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_82_rssTrend_DataType
type SCIClientDetailsReport82rssTrendDataType struct {
	AvgRss *float64 `json:"avgRss,omitempty"`

	End *string `json:"end,omitempty"`

	MaxRss *float64 `json:"maxRss,omitempty"`

	MinRss *float64 `json:"minRss,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIClientDetailsReport82rssTrendDataType() *SCIClientDetailsReport82rssTrendDataType {
	m := new(SCIClientDetailsReport82rssTrendDataType)
	return m
}

// SCIClientDetailsReport83snrTrendData
//
// Definition: ClientDetailsReport_ClientDetailsReport_83_snrTrend_Data
type SCIClientDetailsReport83snrTrendData []*SCIClientDetailsReport83snrTrendDataType

func MakeSCIClientDetailsReport83snrTrendData() SCIClientDetailsReport83snrTrendData {
	m := make(SCIClientDetailsReport83snrTrendData, 0)
	return m
}

// SCIClientDetailsReport83snrTrendDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_83_snrTrend_DataType
type SCIClientDetailsReport83snrTrendDataType struct {
	AvgSnr *float64 `json:"avgSnr,omitempty"`

	End *string `json:"end,omitempty"`

	MaxSnr *float64 `json:"maxSnr,omitempty"`

	MinSnr *float64 `json:"minSnr,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIClientDetailsReport83snrTrendDataType() *SCIClientDetailsReport83snrTrendDataType {
	m := new(SCIClientDetailsReport83snrTrendDataType)
	return m
}

// SCIClientDetailsReport86summaryData
//
// Definition: ClientDetailsReport_ClientDetailsReport_86_summary_Data
type SCIClientDetailsReport86summaryData []*SCIClientDetailsReport86summaryDataType

func MakeSCIClientDetailsReport86summaryData() SCIClientDetailsReport86summaryData {
	m := make(SCIClientDetailsReport86summaryData, 0)
	return m
}

// SCIClientDetailsReport86summaryDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_86_summary_DataType
type SCIClientDetailsReport86summaryDataType struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	LastApMac *string `json:"lastApMac,omitempty"`

	LastApName *string `json:"lastApName,omitempty"`

	LastStatus *string `json:"lastStatus,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	OsType *string `json:"osType,omitempty"`
}

func NewSCIClientDetailsReport86summaryDataType() *SCIClientDetailsReport86summaryDataType {
	m := new(SCIClientDetailsReport86summaryDataType)
	return m
}

// SCIClientDetailsReport87clientStatsData
//
// Definition: ClientDetailsReport_ClientDetailsReport_87_clientStats_Data
type SCIClientDetailsReport87clientStatsData []*SCIClientDetailsReport87clientStatsDataType

func MakeSCIClientDetailsReport87clientStatsData() SCIClientDetailsReport87clientStatsData {
	m := make(SCIClientDetailsReport87clientStatsData, 0)
	return m
}

// SCIClientDetailsReport87clientStatsDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_87_clientStats_DataType
type SCIClientDetailsReport87clientStatsDataType struct {
	AvgRateTotalTraffic *float64 `json:"avgRateTotalTraffic,omitempty"`

	AvgSessionDuration *float64 `json:"avgSessionDuration,omitempty"`

	NumApps *float64 `json:"numApps,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalAps *int `json:"totalAps,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTraffic24 *int `json:"totalTraffic_2-4,omitempty"`

	TotalTraffic5 *int `json:"totalTraffic_5,omitempty"`
}

func NewSCIClientDetailsReport87clientStatsDataType() *SCIClientDetailsReport87clientStatsDataType {
	m := new(SCIClientDetailsReport87clientStatsDataType)
	return m
}

// SCIClientDetailsReport89trafficTrendData
//
// Definition: ClientDetailsReport_ClientDetailsReport_89_trafficTrend_Data
type SCIClientDetailsReport89trafficTrendData [][]*SCIClientDetailsReport89trafficTrendDataTypeType

func MakeSCIClientDetailsReport89trafficTrendData() SCIClientDetailsReport89trafficTrendData {
	m := make(SCIClientDetailsReport89trafficTrendData, 0)
	return m
}

// SCIClientDetailsReport89trafficTrendDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_89_trafficTrend_DataType
type SCIClientDetailsReport89trafficTrendDataType []*SCIClientDetailsReport89trafficTrendDataTypeType

func MakeSCIClientDetailsReport89trafficTrendDataType() SCIClientDetailsReport89trafficTrendDataType {
	m := make(SCIClientDetailsReport89trafficTrendDataType, 0)
	return m
}

// SCIClientDetailsReport89trafficTrendDataTypeType
//
// Definition: ClientDetailsReport_ClientDetailsReport_89_trafficTrend_DataTypeType
type SCIClientDetailsReport89trafficTrendDataTypeType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	UserRxTraffic *int `json:"userRxTraffic,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxTraffic *int `json:"userTxTraffic,omitempty"`
}

func NewSCIClientDetailsReport89trafficTrendDataTypeType() *SCIClientDetailsReport89trafficTrendDataTypeType {
	m := new(SCIClientDetailsReport89trafficTrendDataTypeType)
	return m
}

// SCIClientDetailsReport92sessionsTableData
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_Data
type SCIClientDetailsReport92sessionsTableData []*SCIClientDetailsReport92sessionsTableDataType

func MakeSCIClientDetailsReport92sessionsTableData() SCIClientDetailsReport92sessionsTableData {
	m := make(SCIClientDetailsReport92sessionsTableData, 0)
	return m
}

// SCIClientDetailsReport92sessionsTableDataType
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_DataType
type SCIClientDetailsReport92sessionsTableDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AuthMethod *string `json:"authMethod,omitempty"`

	DisconnectTime *string `json:"disconnectTime,omitempty"`

	FirstConnection *string `json:"firstConnection,omitempty"`

	Radio *string `json:"radio,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	SessionDuration *int `json:"sessionDuration,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIClientDetailsReport92sessionsTableDataType() *SCIClientDetailsReport92sessionsTableDataType {
	m := new(SCIClientDetailsReport92sessionsTableDataType)
	return m
}

// SCIClientDetailsReport92sessionsTableMetaData
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_MetaData
type SCIClientDetailsReport92sessionsTableMetaData struct {
	MaxValues *SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIClientDetailsReport92sessionsTableMetaData() *SCIClientDetailsReport92sessionsTableMetaData {
	m := new(SCIClientDetailsReport92sessionsTableMetaData)
	return m
}

// SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType
//
// Definition: ClientDetailsReport_ClientDetailsReport_92_sessionsTable_MetaDataMaxValuesType
type SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIClientDetailsReport92sessionsTableMetaDataMaxValuesType() *SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType {
	m := new(SCIClientDetailsReport92sessionsTableMetaDataMaxValuesType)
	return m
}

// ReportClientDetailsReport7Top10ApplicationsByTrafficVolume
//
// Operation ID: report_ClientDetailsReport_7_top10ApplicationsByTrafficVolume
//
// Client Details Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport7Top10ApplicationsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport7top10ApplicationsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport8TopAppsByTrafficTable
//
// Operation ID: report_ClientDetailsReport_8_topAppsByTrafficTable
//
// Client Details Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport8TopAppsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport8topAppsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport82RssTrend
//
// Operation ID: report_ClientDetailsReport_82_rssTrend
//
// Client Details Report - RSS Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport82RssTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport82rssTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport82rssTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport82RssTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport82rssTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport83SnrTrend
//
// Operation ID: report_ClientDetailsReport_83_snrTrend
//
// Client Details Report - SNR Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport83SnrTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport83snrTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport83snrTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport83SnrTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport83snrTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport86Summary
//
// Operation ID: report_ClientDetailsReport_86_summary
//
// Client Details Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport86Summary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport86summary200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport86summary200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport86Summary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport86summary200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport87ClientStats
//
// Operation ID: report_ClientDetailsReport_87_clientStats
//
// Client Details Report - Stats
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport87ClientStats(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport87clientStats200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport87clientStats200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport87ClientStats, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport87clientStats200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport89TrafficTrend
//
// Operation ID: report_ClientDetailsReport_89_trafficTrend
//
// Client Details Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport89TrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport89trafficTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport89trafficTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport89TrafficTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport89trafficTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientDetailsReport92SessionsTable
//
// Operation ID: report_ClientDetailsReport_92_sessionsTable
//
// Client Details Report - Sessions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientDetailsReportService) ReportClientDetailsReport92SessionsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientDetailsReport92sessionsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientDetailsReport92sessionsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientDetailsReport92SessionsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientDetailsReport92sessionsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
