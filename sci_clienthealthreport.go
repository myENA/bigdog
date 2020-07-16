package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIClientHealthReportService struct {
	apiClient *SCIClient
}

func NewSCIClientHealthReportService(c *SCIClient) *SCIClientHealthReportService {
	s := new(SCIClientHealthReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIClientHealthReportService() *SCIClientHealthReportService {
	return NewSCIClientHealthReportService(ss.apiClient)
}

// SCIClientHealthReport144clientHealthSummaryData
//
// Definition: ClientHealthReport_ClientHealthReport_144_clientHealthSummary_Data
type SCIClientHealthReport144clientHealthSummaryData []*SCIClientHealthReport144clientHealthSummaryDataType

func MakeSCIClientHealthReport144clientHealthSummaryData() SCIClientHealthReport144clientHealthSummaryData {
	m := make(SCIClientHealthReport144clientHealthSummaryData, 0)
	return m
}

// SCIClientHealthReport144clientHealthSummaryDataType
//
// Definition: ClientHealthReport_ClientHealthReport_144_clientHealthSummary_DataType
type SCIClientHealthReport144clientHealthSummaryDataType struct {
	AvgRFHealthClientCount *int `json:"avgRFHealthClientCount,omitempty"`

	AvgRFHealthRatio *int `json:"avgRFHealthRatio,omitempty"`

	GoodRFHealthClientCount *int `json:"goodRFHealthClientCount,omitempty"`

	GoodRFHealthRatio *int `json:"goodRFHealthRatio,omitempty"`

	PoorRFHealthClientCount *int `json:"poorRFHealthClientCount,omitempty"`

	PoorRFHealthRatio *int `json:"poorRFHealthRatio,omitempty"`

	TotalClients *int `json:"totalClients,omitempty"`
}

func NewSCIClientHealthReport144clientHealthSummaryDataType() *SCIClientHealthReport144clientHealthSummaryDataType {
	m := new(SCIClientHealthReport144clientHealthSummaryDataType)
	return m
}

// SCIClientHealthReport148clientConnectionHealthData
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_Data
type SCIClientHealthReport148clientConnectionHealthData []*SCIClientHealthReport148clientConnectionHealthDataType

func MakeSCIClientHealthReport148clientConnectionHealthData() SCIClientHealthReport148clientConnectionHealthData {
	m := make(SCIClientHealthReport148clientConnectionHealthData, 0)
	return m
}

// SCIClientHealthReport148clientConnectionHealthDataType
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_DataType
type SCIClientHealthReport148clientConnectionHealthDataType struct {
	Avg *int `json:"avg,omitempty"`

	Good *int `json:"good,omitempty"`

	Poor *int `json:"poor,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`
}

func NewSCIClientHealthReport148clientConnectionHealthDataType() *SCIClientHealthReport148clientConnectionHealthDataType {
	m := new(SCIClientHealthReport148clientConnectionHealthDataType)
	return m
}

// SCIClientHealthReport148clientConnectionHealthMetaData
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_MetaData
type SCIClientHealthReport148clientConnectionHealthMetaData struct {
	CategoryKey *string `json:"categoryKey,omitempty"`

	Granularity *string `json:"granularity,omitempty"`
}

func NewSCIClientHealthReport148clientConnectionHealthMetaData() *SCIClientHealthReport148clientConnectionHealthMetaData {
	m := new(SCIClientHealthReport148clientConnectionHealthMetaData)
	return m
}

// SCIClientHealthReport149clientHealthMetricTrendsData
//
// Definition: ClientHealthReport_ClientHealthReport_149_clientHealthMetricTrends_Data
type SCIClientHealthReport149clientHealthMetricTrendsData []*SCIClientHealthReport149clientHealthMetricTrendsDataType

func MakeSCIClientHealthReport149clientHealthMetricTrendsData() SCIClientHealthReport149clientHealthMetricTrendsData {
	m := make(SCIClientHealthReport149clientHealthMetricTrendsData, 0)
	return m
}

// SCIClientHealthReport149clientHealthMetricTrendsDataType
//
// Definition: ClientHealthReport_ClientHealthReport_149_clientHealthMetricTrends_DataType
type SCIClientHealthReport149clientHealthMetricTrendsDataType struct {
	AvgSnr *float64 `json:"avgSnr,omitempty"`

	End *string `json:"end,omitempty"`

	MaxSnr *int `json:"maxSnr,omitempty"`

	MinSnr *int `json:"minSnr,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIClientHealthReport149clientHealthMetricTrendsDataType() *SCIClientHealthReport149clientHealthMetricTrendsDataType {
	m := new(SCIClientHealthReport149clientHealthMetricTrendsDataType)
	return m
}

// ReportClientHealthReport144ClientHealthSummary
//
// Operation ID: report_ClientHealthReport_144_clientHealthSummary
//
// Client Health Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport144ClientHealthSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport144clientHealthSummary200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport144clientHealthSummary200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport144ClientHealthSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport144clientHealthSummary200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientHealthReport148ClientConnectionHealth
//
// Operation ID: report_ClientHealthReport_148_clientConnectionHealth
//
// Client Health Report - Client Connection Health
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport148ClientConnectionHealth(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport148clientConnectionHealth200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport148clientConnectionHealth200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport148ClientConnectionHealth, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport148clientConnectionHealth200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientHealthReport149ClientHealthMetricTrends
//
// Operation ID: report_ClientHealthReport_149_clientHealthMetricTrends
//
// Client Health Report - Health Metric Trends
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport149ClientHealthMetricTrends(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport149clientHealthMetricTrends200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport149ClientHealthMetricTrends, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport149clientHealthMetricTrends200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportClientHealthReport150TopClientHealthScoreByGroup
//
// Operation ID: report_ClientHealthReport_150_topClientHealthScoreByGroup
//
// Client Health Report - Health By Group
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIClientHealthReportService) ReportClientHealthReport150TopClientHealthScoreByGroup(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportClientHealthReport150TopClientHealthScoreByGroup, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportClientHealthReport150topClientHealthScoreByGroup200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
