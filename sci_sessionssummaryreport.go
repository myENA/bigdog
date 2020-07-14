package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCISessionsSummaryReportService struct {
	apiClient *SCIClient
}

func NewSCISessionsSummaryReportService(c *SCIClient) *SCISessionsSummaryReportService {
	s := new(SCISessionsSummaryReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCISessionsSummaryReportService() *SCISessionsSummaryReportService {
	return NewSCISessionsSummaryReportService(ss.apiClient)
}

// SCISessionsSummaryReport33topTableData
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.33.topTable.Data
type SCISessionsSummaryReport33topTableData []*SCISessionsSummaryReport33topTableDataType

func MakeSCISessionsSummaryReport33topTableData() SCISessionsSummaryReport33topTableData {
	m := make(SCISessionsSummaryReport33topTableData, 0)
	return m
}

// SCISessionsSummaryReport33topTableDataType
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.33.topTable.DataType
type SCISessionsSummaryReport33topTableDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ApSerial *string `json:"apSerial,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	DisconnectTime *string `json:"disconnectTime,omitempty"`

	FirstConnection *string `json:"firstConnection,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	OsType *string `json:"osType,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	SessionDuration *int `json:"sessionDuration,omitempty"`

	SessionId *string `json:"sessionId,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewSCISessionsSummaryReport33topTableDataType() *SCISessionsSummaryReport33topTableDataType {
	m := new(SCISessionsSummaryReport33topTableDataType)
	return m
}

// SCISessionsSummaryReport33topTableMetaData
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.33.topTable.MetaData
type SCISessionsSummaryReport33topTableMetaData struct {
	PagingIdentifiers *SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType `json:"pagingIdentifiers,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSCISessionsSummaryReport33topTableMetaData() *SCISessionsSummaryReport33topTableMetaData {
	m := new(SCISessionsSummaryReport33topTableMetaData)
	return m
}

// SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.33.topTable.MetaDataPagingIdentifiersType
type SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType struct {
	Type *string `json:"type,omitempty"`
}

func NewSCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType() *SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType {
	m := new(SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType)
	return m
}

// SCISessionsSummaryReport34overviewData
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.34.overview.Data
type SCISessionsSummaryReport34overviewData []*SCISessionsSummaryReport34overviewDataType

func MakeSCISessionsSummaryReport34overviewData() SCISessionsSummaryReport34overviewData {
	m := make(SCISessionsSummaryReport34overviewData, 0)
	return m
}

// SCISessionsSummaryReport34overviewDataType
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.34.overview.DataType
type SCISessionsSummaryReport34overviewDataType struct {
	SessionDurationAvg *float64 `json:"sessionDurationAvg,omitempty"`
}

func NewSCISessionsSummaryReport34overviewDataType() *SCISessionsSummaryReport34overviewDataType {
	m := new(SCISessionsSummaryReport34overviewDataType)
	return m
}

// SCISessionsSummaryReport42durationPercentileData
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.42.durationPercentile.Data
type SCISessionsSummaryReport42durationPercentileData []*SCISessionsSummaryReport42durationPercentileDataType

func MakeSCISessionsSummaryReport42durationPercentileData() SCISessionsSummaryReport42durationPercentileData {
	m := make(SCISessionsSummaryReport42durationPercentileData, 0)
	return m
}

// SCISessionsSummaryReport42durationPercentileDataType
//
// Definition: SessionsSummaryReport.SessionsSummaryReport.42.durationPercentile.DataType
type SCISessionsSummaryReport42durationPercentileDataType struct {
	Duration *int `json:"duration,omitempty"`

	PercentRank *int `json:"percentRank,omitempty"`
}

func NewSCISessionsSummaryReport42durationPercentileDataType() *SCISessionsSummaryReport42durationPercentileDataType {
	m := new(SCISessionsSummaryReport42durationPercentileDataType)
	return m
}

// ReportSessionsSummaryReport33TopTable
//
// Operation ID: report.SessionsSummaryReport.33.topTable
//
// Sessions Summary Report - Top Sessions Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport33TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport33topTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSessionsSummaryReport33topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSessionsSummaryReport33TopTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSessionsSummaryReport33topTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSessionsSummaryReport34Overview
//
// Operation ID: report.SessionsSummaryReport.34.overview
//
// Sessions Summary Report - Average Durations
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport34Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport34overview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSessionsSummaryReport34overview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSessionsSummaryReport34Overview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSessionsSummaryReport34overview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportSessionsSummaryReport42DurationPercentile
//
// Operation ID: report.SessionsSummaryReport.42.durationPercentile
//
// Sessions Summary Report - Session Duration Percentiles
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport42DurationPercentile(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport42durationPercentile200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportSessionsSummaryReport42durationPercentile200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportSessionsSummaryReport42DurationPercentile, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportSessionsSummaryReport42durationPercentile200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
