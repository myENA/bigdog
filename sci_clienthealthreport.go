package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
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

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientHealthReport144clientHealthSummaryDataType) UnmarshalJSON(b []byte) error {
	type _SCIClientHealthReport144clientHealthSummaryDataType SCIClientHealthReport144clientHealthSummaryDataType
	tmpType := new(_SCIClientHealthReport144clientHealthSummaryDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgRFHealthClientCount")
	delete(tmpType.XAdditionalProperties, "avgRFHealthRatio")
	delete(tmpType.XAdditionalProperties, "goodRFHealthClientCount")
	delete(tmpType.XAdditionalProperties, "goodRFHealthRatio")
	delete(tmpType.XAdditionalProperties, "poorRFHealthClientCount")
	delete(tmpType.XAdditionalProperties, "poorRFHealthRatio")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCIClientHealthReport144clientHealthSummaryDataType(*tmpType)
	return nil
}

func (t *SCIClientHealthReport144clientHealthSummaryDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AvgRFHealthClientCount != nil {
		tmp["avgRFHealthClientCount"] = t.AvgRFHealthClientCount
	}
	if t.AvgRFHealthRatio != nil {
		tmp["avgRFHealthRatio"] = t.AvgRFHealthRatio
	}
	if t.GoodRFHealthClientCount != nil {
		tmp["goodRFHealthClientCount"] = t.GoodRFHealthClientCount
	}
	if t.GoodRFHealthRatio != nil {
		tmp["goodRFHealthRatio"] = t.GoodRFHealthRatio
	}
	if t.PoorRFHealthClientCount != nil {
		tmp["poorRFHealthClientCount"] = t.PoorRFHealthClientCount
	}
	if t.PoorRFHealthRatio != nil {
		tmp["poorRFHealthRatio"] = t.PoorRFHealthRatio
	}
	if t.TotalClients != nil {
		tmp["totalClients"] = t.TotalClients
	}
	return json.Marshal(tmp)
}

func NewSCIClientHealthReport144clientHealthSummaryDataType() *SCIClientHealthReport144clientHealthSummaryDataType {
	m := new(SCIClientHealthReport144clientHealthSummaryDataType)
	return m
}

// SCIClientHealthReport148clientConnectionHealthData
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_Data
type SCIClientHealthReport148clientConnectionHealthData [][]*SCIClientHealthReport148clientConnectionHealthDataTypeType

func MakeSCIClientHealthReport148clientConnectionHealthData() SCIClientHealthReport148clientConnectionHealthData {
	m := make(SCIClientHealthReport148clientConnectionHealthData, 0)
	return m
}

// SCIClientHealthReport148clientConnectionHealthDataType
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_DataType
type SCIClientHealthReport148clientConnectionHealthDataType []*SCIClientHealthReport148clientConnectionHealthDataTypeType

func MakeSCIClientHealthReport148clientConnectionHealthDataType() SCIClientHealthReport148clientConnectionHealthDataType {
	m := make(SCIClientHealthReport148clientConnectionHealthDataType, 0)
	return m
}

// SCIClientHealthReport148clientConnectionHealthDataTypeType
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_DataTypeType
type SCIClientHealthReport148clientConnectionHealthDataTypeType struct {
	Avg *int `json:"avg,omitempty"`

	Good *int `json:"good,omitempty"`

	Poor *int `json:"poor,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientHealthReport148clientConnectionHealthDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIClientHealthReport148clientConnectionHealthDataTypeType SCIClientHealthReport148clientConnectionHealthDataTypeType
	tmpType := new(_SCIClientHealthReport148clientConnectionHealthDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avg")
	delete(tmpType.XAdditionalProperties, "good")
	delete(tmpType.XAdditionalProperties, "poor")
	delete(tmpType.XAdditionalProperties, "timestamp")
	*t = SCIClientHealthReport148clientConnectionHealthDataTypeType(*tmpType)
	return nil
}

func (t *SCIClientHealthReport148clientConnectionHealthDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Avg != nil {
		tmp["avg"] = t.Avg
	}
	if t.Good != nil {
		tmp["good"] = t.Good
	}
	if t.Poor != nil {
		tmp["poor"] = t.Poor
	}
	if t.Timestamp != nil {
		tmp["timestamp"] = t.Timestamp
	}
	return json.Marshal(tmp)
}

func NewSCIClientHealthReport148clientConnectionHealthDataTypeType() *SCIClientHealthReport148clientConnectionHealthDataTypeType {
	m := new(SCIClientHealthReport148clientConnectionHealthDataTypeType)
	return m
}

// SCIClientHealthReport148clientConnectionHealthMetaData
//
// Definition: ClientHealthReport_ClientHealthReport_148_clientConnectionHealth_MetaData
type SCIClientHealthReport148clientConnectionHealthMetaData struct {
	CategoryKey *string `json:"categoryKey,omitempty"`

	Granularity *string `json:"granularity,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientHealthReport148clientConnectionHealthMetaData) UnmarshalJSON(b []byte) error {
	type _SCIClientHealthReport148clientConnectionHealthMetaData SCIClientHealthReport148clientConnectionHealthMetaData
	tmpType := new(_SCIClientHealthReport148clientConnectionHealthMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "categoryKey")
	delete(tmpType.XAdditionalProperties, "granularity")
	*t = SCIClientHealthReport148clientConnectionHealthMetaData(*tmpType)
	return nil
}

func (t *SCIClientHealthReport148clientConnectionHealthMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.CategoryKey != nil {
		tmp["categoryKey"] = t.CategoryKey
	}
	if t.Granularity != nil {
		tmp["granularity"] = t.Granularity
	}
	return json.Marshal(tmp)
}

func NewSCIClientHealthReport148clientConnectionHealthMetaData() *SCIClientHealthReport148clientConnectionHealthMetaData {
	m := new(SCIClientHealthReport148clientConnectionHealthMetaData)
	return m
}

// SCIClientHealthReport149clientHealthMetricTrendsData
//
// Definition: ClientHealthReport_ClientHealthReport_149_clientHealthMetricTrends_Data
type SCIClientHealthReport149clientHealthMetricTrendsData [][]*SCIClientHealthReport149clientHealthMetricTrendsDataTypeType

func MakeSCIClientHealthReport149clientHealthMetricTrendsData() SCIClientHealthReport149clientHealthMetricTrendsData {
	m := make(SCIClientHealthReport149clientHealthMetricTrendsData, 0)
	return m
}

// SCIClientHealthReport149clientHealthMetricTrendsDataType
//
// Definition: ClientHealthReport_ClientHealthReport_149_clientHealthMetricTrends_DataType
type SCIClientHealthReport149clientHealthMetricTrendsDataType []*SCIClientHealthReport149clientHealthMetricTrendsDataTypeType

func MakeSCIClientHealthReport149clientHealthMetricTrendsDataType() SCIClientHealthReport149clientHealthMetricTrendsDataType {
	m := make(SCIClientHealthReport149clientHealthMetricTrendsDataType, 0)
	return m
}

// SCIClientHealthReport149clientHealthMetricTrendsDataTypeType
//
// Definition: ClientHealthReport_ClientHealthReport_149_clientHealthMetricTrends_DataTypeType
type SCIClientHealthReport149clientHealthMetricTrendsDataTypeType struct {
	AvgMedianTxMCSRate *float64 `json:"avgMedianTxMCSRate,omitempty"`

	AvgRss *float64 `json:"avgRss,omitempty"`

	AvgSnr *float64 `json:"avgSnr,omitempty"`

	AvgThroughputEstimate *float64 `json:"avgThroughputEstimate,omitempty"`

	End *string `json:"end,omitempty"`

	MaxMedianTxMCSRate *int `json:"maxMedianTxMCSRate,omitempty"`

	MaxRss *int `json:"maxRss,omitempty"`

	MaxSnr *int `json:"maxSnr,omitempty"`

	MaxThroughputEstimate *int `json:"maxThroughputEstimate,omitempty"`

	MinMedianTxMCSRate *int `json:"minMedianTxMCSRate,omitempty"`

	MinRss *int `json:"minRss,omitempty"`

	MinSnr *int `json:"minSnr,omitempty"`

	MinThroughputEstimate *int `json:"minThroughputEstimate,omitempty"`

	Name *string `json:"name,omitempty"`

	Start *string `json:"start,omitempty"`

	Value *string `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIClientHealthReport149clientHealthMetricTrendsDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIClientHealthReport149clientHealthMetricTrendsDataTypeType SCIClientHealthReport149clientHealthMetricTrendsDataTypeType
	tmpType := new(_SCIClientHealthReport149clientHealthMetricTrendsDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "avgMedianTxMCSRate")
	delete(tmpType.XAdditionalProperties, "avgRss")
	delete(tmpType.XAdditionalProperties, "avgSnr")
	delete(tmpType.XAdditionalProperties, "avgThroughputEstimate")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "maxMedianTxMCSRate")
	delete(tmpType.XAdditionalProperties, "maxRss")
	delete(tmpType.XAdditionalProperties, "maxSnr")
	delete(tmpType.XAdditionalProperties, "maxThroughputEstimate")
	delete(tmpType.XAdditionalProperties, "minMedianTxMCSRate")
	delete(tmpType.XAdditionalProperties, "minRss")
	delete(tmpType.XAdditionalProperties, "minSnr")
	delete(tmpType.XAdditionalProperties, "minThroughputEstimate")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "value")
	*t = SCIClientHealthReport149clientHealthMetricTrendsDataTypeType(*tmpType)
	return nil
}

func (t *SCIClientHealthReport149clientHealthMetricTrendsDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AvgMedianTxMCSRate != nil {
		tmp["avgMedianTxMCSRate"] = t.AvgMedianTxMCSRate
	}
	if t.AvgRss != nil {
		tmp["avgRss"] = t.AvgRss
	}
	if t.AvgSnr != nil {
		tmp["avgSnr"] = t.AvgSnr
	}
	if t.AvgThroughputEstimate != nil {
		tmp["avgThroughputEstimate"] = t.AvgThroughputEstimate
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.MaxMedianTxMCSRate != nil {
		tmp["maxMedianTxMCSRate"] = t.MaxMedianTxMCSRate
	}
	if t.MaxRss != nil {
		tmp["maxRss"] = t.MaxRss
	}
	if t.MaxSnr != nil {
		tmp["maxSnr"] = t.MaxSnr
	}
	if t.MaxThroughputEstimate != nil {
		tmp["maxThroughputEstimate"] = t.MaxThroughputEstimate
	}
	if t.MinMedianTxMCSRate != nil {
		tmp["minMedianTxMCSRate"] = t.MinMedianTxMCSRate
	}
	if t.MinRss != nil {
		tmp["minRss"] = t.MinRss
	}
	if t.MinSnr != nil {
		tmp["minSnr"] = t.MinSnr
	}
	if t.MinThroughputEstimate != nil {
		tmp["minThroughputEstimate"] = t.MinThroughputEstimate
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.Value != nil {
		tmp["value"] = t.Value
	}
	return json.Marshal(tmp)
}

func NewSCIClientHealthReport149clientHealthMetricTrendsDataTypeType() *SCIClientHealthReport149clientHealthMetricTrendsDataTypeType {
	m := new(SCIClientHealthReport149clientHealthMetricTrendsDataTypeType)
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
