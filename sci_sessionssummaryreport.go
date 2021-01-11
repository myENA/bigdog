package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
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
// Definition: SessionsSummaryReport_SessionsSummaryReport_33_topTable_Data
type SCISessionsSummaryReport33topTableData []*SCISessionsSummaryReport33topTableDataType

func MakeSCISessionsSummaryReport33topTableData() SCISessionsSummaryReport33topTableData {
	m := make(SCISessionsSummaryReport33topTableData, 0)
	return m
}

// SCISessionsSummaryReport33topTableDataType
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_33_topTable_DataType
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

	RxBytes *float64 `json:"rxBytes,omitempty"`

	SessionDuration *float64 `json:"sessionDuration,omitempty"`

	SessionId *string `json:"sessionId,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	Username *string `json:"username,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISessionsSummaryReport33topTableDataType) UnmarshalJSON(b []byte) error {
	type _SCISessionsSummaryReport33topTableDataType SCISessionsSummaryReport33topTableDataType
	tmpType := new(_SCISessionsSummaryReport33topTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "apSerial")
	delete(tmpType.XAdditionalProperties, "clientMac")
	delete(tmpType.XAdditionalProperties, "disconnectTime")
	delete(tmpType.XAdditionalProperties, "firstConnection")
	delete(tmpType.XAdditionalProperties, "hostname")
	delete(tmpType.XAdditionalProperties, "manufacturer")
	delete(tmpType.XAdditionalProperties, "osType")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "sessionDuration")
	delete(tmpType.XAdditionalProperties, "sessionId")
	delete(tmpType.XAdditionalProperties, "timestamp")
	delete(tmpType.XAdditionalProperties, "txBytes")
	delete(tmpType.XAdditionalProperties, "username")
	*t = SCISessionsSummaryReport33topTableDataType(*tmpType)
	return nil
}

func (t *SCISessionsSummaryReport33topTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.ApSerial != nil {
		tmp["apSerial"] = t.ApSerial
	}
	if t.ClientMac != nil {
		tmp["clientMac"] = t.ClientMac
	}
	if t.DisconnectTime != nil {
		tmp["disconnectTime"] = t.DisconnectTime
	}
	if t.FirstConnection != nil {
		tmp["firstConnection"] = t.FirstConnection
	}
	if t.Hostname != nil {
		tmp["hostname"] = t.Hostname
	}
	if t.Manufacturer != nil {
		tmp["manufacturer"] = t.Manufacturer
	}
	if t.OsType != nil {
		tmp["osType"] = t.OsType
	}
	if t.RxBytes != nil {
		tmp["rxBytes"] = t.RxBytes
	}
	if t.SessionDuration != nil {
		tmp["sessionDuration"] = t.SessionDuration
	}
	if t.SessionId != nil {
		tmp["sessionId"] = t.SessionId
	}
	if t.Timestamp != nil {
		tmp["timestamp"] = t.Timestamp
	}
	if t.TxBytes != nil {
		tmp["txBytes"] = t.TxBytes
	}
	if t.Username != nil {
		tmp["username"] = t.Username
	}
	return json.Marshal(tmp)
}

func NewSCISessionsSummaryReport33topTableDataType() *SCISessionsSummaryReport33topTableDataType {
	m := new(SCISessionsSummaryReport33topTableDataType)
	return m
}

// SCISessionsSummaryReport33topTableMetaData
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_33_topTable_MetaData
type SCISessionsSummaryReport33topTableMetaData struct {
	PagingIdentifiers *SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType `json:"pagingIdentifiers,omitempty"`

	TotalCount *float64 `json:"totalCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISessionsSummaryReport33topTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCISessionsSummaryReport33topTableMetaData SCISessionsSummaryReport33topTableMetaData
	tmpType := new(_SCISessionsSummaryReport33topTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "pagingIdentifiers")
	delete(tmpType.XAdditionalProperties, "totalCount")
	*t = SCISessionsSummaryReport33topTableMetaData(*tmpType)
	return nil
}

func (t *SCISessionsSummaryReport33topTableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.PagingIdentifiers != nil {
		tmp["pagingIdentifiers"] = t.PagingIdentifiers
	}
	if t.TotalCount != nil {
		tmp["totalCount"] = t.TotalCount
	}
	return json.Marshal(tmp)
}

func NewSCISessionsSummaryReport33topTableMetaData() *SCISessionsSummaryReport33topTableMetaData {
	m := new(SCISessionsSummaryReport33topTableMetaData)
	return m
}

// SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_33_topTable_MetaDataPagingIdentifiersType
type SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType struct {
	Type *string `json:"type,omitempty"`
}

func NewSCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType() *SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType {
	m := new(SCISessionsSummaryReport33topTableMetaDataPagingIdentifiersType)
	return m
}

// SCISessionsSummaryReport34overviewData
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_34_overview_Data
type SCISessionsSummaryReport34overviewData []*SCISessionsSummaryReport34overviewDataType

func MakeSCISessionsSummaryReport34overviewData() SCISessionsSummaryReport34overviewData {
	m := make(SCISessionsSummaryReport34overviewData, 0)
	return m
}

// SCISessionsSummaryReport34overviewDataType
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_34_overview_DataType
type SCISessionsSummaryReport34overviewDataType struct {
	SessionDurationAvg *float64 `json:"sessionDurationAvg,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISessionsSummaryReport34overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCISessionsSummaryReport34overviewDataType SCISessionsSummaryReport34overviewDataType
	tmpType := new(_SCISessionsSummaryReport34overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "sessionDurationAvg")
	*t = SCISessionsSummaryReport34overviewDataType(*tmpType)
	return nil
}

func (t *SCISessionsSummaryReport34overviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.SessionDurationAvg != nil {
		tmp["sessionDurationAvg"] = t.SessionDurationAvg
	}
	return json.Marshal(tmp)
}

func NewSCISessionsSummaryReport34overviewDataType() *SCISessionsSummaryReport34overviewDataType {
	m := new(SCISessionsSummaryReport34overviewDataType)
	return m
}

// SCISessionsSummaryReport42durationPercentileData
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_42_durationPercentile_Data
type SCISessionsSummaryReport42durationPercentileData []*SCISessionsSummaryReport42durationPercentileDataType

func MakeSCISessionsSummaryReport42durationPercentileData() SCISessionsSummaryReport42durationPercentileData {
	m := make(SCISessionsSummaryReport42durationPercentileData, 0)
	return m
}

// SCISessionsSummaryReport42durationPercentileDataType
//
// Definition: SessionsSummaryReport_SessionsSummaryReport_42_durationPercentile_DataType
type SCISessionsSummaryReport42durationPercentileDataType struct {
	Duration *float64 `json:"duration,omitempty"`

	PercentRank *float64 `json:"percentRank,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISessionsSummaryReport42durationPercentileDataType) UnmarshalJSON(b []byte) error {
	type _SCISessionsSummaryReport42durationPercentileDataType SCISessionsSummaryReport42durationPercentileDataType
	tmpType := new(_SCISessionsSummaryReport42durationPercentileDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "duration")
	delete(tmpType.XAdditionalProperties, "percentRank")
	*t = SCISessionsSummaryReport42durationPercentileDataType(*tmpType)
	return nil
}

func (t *SCISessionsSummaryReport42durationPercentileDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Duration != nil {
		tmp["duration"] = t.Duration
	}
	if t.PercentRank != nil {
		tmp["percentRank"] = t.PercentRank
	}
	return json.Marshal(tmp)
}

func NewSCISessionsSummaryReport42durationPercentileDataType() *SCISessionsSummaryReport42durationPercentileDataType {
	m := new(SCISessionsSummaryReport42durationPercentileDataType)
	return m
}

// ReportSessionsSummaryReport33TopTable
//
// Sessions Summary Report - Top Sessions Summary
//
// Operation ID: report_SessionsSummaryReport_33_topTable
// Operation path: /reports/6/sections/33/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport33TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSessionsSummaryReport33TopTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSessionsSummaryReport33topTable200ResponseTypeAPIResponse), err
}

// ReportSessionsSummaryReport34Overview
//
// Sessions Summary Report - Average Durations
//
// Operation ID: report_SessionsSummaryReport_34_overview
// Operation path: /reports/6/sections/34/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport34Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSessionsSummaryReport34Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSessionsSummaryReport34overview200ResponseTypeAPIResponse), err
}

// ReportSessionsSummaryReport42DurationPercentile
//
// Sessions Summary Report - Session Duration Percentiles
//
// Operation ID: report_SessionsSummaryReport_42_durationPercentile
// Operation path: /reports/6/sections/42/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCISessionsSummaryReportService) ReportSessionsSummaryReport42DurationPercentile(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportSessionsSummaryReport42DurationPercentile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportSessionsSummaryReport42durationPercentile200ResponseTypeAPIResponse), err
}
