package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIAPsRebootReportService struct {
	apiClient *SCIClient
}

func NewSCIAPsRebootReportService(c *SCIClient) *SCIAPsRebootReportService {
	s := new(SCIAPsRebootReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAPsRebootReportService() *SCIAPsRebootReportService {
	return NewSCIAPsRebootReportService(ss.apiClient)
}

// SCIAPsRebootReport43totalRebootsData
//
// Definition: APsRebootReport_APsRebootReport_43_totalReboots_Data
type SCIAPsRebootReport43totalRebootsData []*SCIAPsRebootReport43totalRebootsDataType

func MakeSCIAPsRebootReport43totalRebootsData() SCIAPsRebootReport43totalRebootsData {
	m := make(SCIAPsRebootReport43totalRebootsData, 0)
	return m
}

// SCIAPsRebootReport43totalRebootsDataType
//
// Definition: APsRebootReport_APsRebootReport_43_totalReboots_DataType
type SCIAPsRebootReport43totalRebootsDataType struct {
	TotalRebootedAps *float64 `json:"totalRebootedAps,omitempty"`

	TotalReboots *float64 `json:"totalReboots,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPsRebootReport43totalRebootsDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPsRebootReport43totalRebootsDataType SCIAPsRebootReport43totalRebootsDataType
	tmpType := new(_SCIAPsRebootReport43totalRebootsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalRebootedAps")
	delete(tmpType.XAdditionalProperties, "totalReboots")
	*t = SCIAPsRebootReport43totalRebootsDataType(*tmpType)
	return nil
}

func (t *SCIAPsRebootReport43totalRebootsDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.TotalRebootedAps != nil {
		tmp["totalRebootedAps"] = t.TotalRebootedAps
	}
	if t.TotalReboots != nil {
		tmp["totalReboots"] = t.TotalReboots
	}
	return json.Marshal(tmp)
}

func NewSCIAPsRebootReport43totalRebootsDataType() *SCIAPsRebootReport43totalRebootsDataType {
	m := new(SCIAPsRebootReport43totalRebootsDataType)
	return m
}

// SCIAPsRebootReport44topApRebootsTableData
//
// Definition: APsRebootReport_APsRebootReport_44_topApRebootsTable_Data
type SCIAPsRebootReport44topApRebootsTableData []*SCIAPsRebootReport44topApRebootsTableDataType

func MakeSCIAPsRebootReport44topApRebootsTableData() SCIAPsRebootReport44topApRebootsTableData {
	m := make(SCIAPsRebootReport44topApRebootsTableData, 0)
	return m
}

// SCIAPsRebootReport44topApRebootsTableDataType
//
// Definition: APsRebootReport_APsRebootReport_44_topApRebootsTable_DataType
type SCIAPsRebootReport44topApRebootsTableDataType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApModel *string `json:"apModel,omitempty"`

	ApName *string `json:"apName,omitempty"`

	Count *float64 `json:"count,omitempty"`

	Index *float64 `json:"index,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPsRebootReport44topApRebootsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPsRebootReport44topApRebootsTableDataType SCIAPsRebootReport44topApRebootsTableDataType
	tmpType := new(_SCIAPsRebootReport44topApRebootsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apModel")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "count")
	delete(tmpType.XAdditionalProperties, "index")
	*t = SCIAPsRebootReport44topApRebootsTableDataType(*tmpType)
	return nil
}

func (t *SCIAPsRebootReport44topApRebootsTableDataType) MarshalJSON() ([]byte, error) {
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
	if t.ApModel != nil {
		tmp["apModel"] = t.ApModel
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.Count != nil {
		tmp["count"] = t.Count
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	return json.Marshal(tmp)
}

func NewSCIAPsRebootReport44topApRebootsTableDataType() *SCIAPsRebootReport44topApRebootsTableDataType {
	m := new(SCIAPsRebootReport44topApRebootsTableDataType)
	return m
}

// SCIAPsRebootReport44topApRebootsTableMetaData
//
// Definition: APsRebootReport_APsRebootReport_44_topApRebootsTable_MetaData
type SCIAPsRebootReport44topApRebootsTableMetaData struct {
	TotalCount *float64 `json:"totalCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPsRebootReport44topApRebootsTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAPsRebootReport44topApRebootsTableMetaData SCIAPsRebootReport44topApRebootsTableMetaData
	tmpType := new(_SCIAPsRebootReport44topApRebootsTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalCount")
	*t = SCIAPsRebootReport44topApRebootsTableMetaData(*tmpType)
	return nil
}

func (t *SCIAPsRebootReport44topApRebootsTableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.TotalCount != nil {
		tmp["totalCount"] = t.TotalCount
	}
	return json.Marshal(tmp)
}

func NewSCIAPsRebootReport44topApRebootsTableMetaData() *SCIAPsRebootReport44topApRebootsTableMetaData {
	m := new(SCIAPsRebootReport44topApRebootsTableMetaData)
	return m
}

// SCIAPsRebootReport45topApRebootsOverTimeData
//
// Definition: APsRebootReport_APsRebootReport_45_topApRebootsOverTime_Data
type SCIAPsRebootReport45topApRebootsOverTimeData []*SCIAPsRebootReport45topApRebootsOverTimeDataType

func MakeSCIAPsRebootReport45topApRebootsOverTimeData() SCIAPsRebootReport45topApRebootsOverTimeData {
	m := make(SCIAPsRebootReport45topApRebootsOverTimeData, 0)
	return m
}

// SCIAPsRebootReport45topApRebootsOverTimeDataType
//
// Definition: APsRebootReport_APsRebootReport_45_topApRebootsOverTime_DataType
type SCIAPsRebootReport45topApRebootsOverTimeDataType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAPsRebootReport45topApRebootsOverTimeDataType) UnmarshalJSON(b []byte) error {
	type _SCIAPsRebootReport45topApRebootsOverTimeDataType SCIAPsRebootReport45topApRebootsOverTimeDataType
	tmpType := new(_SCIAPsRebootReport45topApRebootsOverTimeDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIAPsRebootReport45topApRebootsOverTimeDataType(*tmpType)
	return nil
}

func (t *SCIAPsRebootReport45topApRebootsOverTimeDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIAPsRebootReport45topApRebootsOverTimeDataType() *SCIAPsRebootReport45topApRebootsOverTimeDataType {
	m := new(SCIAPsRebootReport45topApRebootsOverTimeDataType)
	return m
}

// ReportAPsRebootReport43TotalReboots
//
// Operation ID: report_APsRebootReport_43_totalReboots
//
// APs Reboot Report - Total Reboots
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport43TotalReboots(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport43totalReboots200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPsRebootReport43totalReboots200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportAPsRebootReport43TotalReboots, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport43totalReboots200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// ReportAPsRebootReport44TopApRebootsTable
//
// Operation ID: report_APsRebootReport_44_topApRebootsTable
//
// APs Reboot Report - Top AP Reboots
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport44TopApRebootsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport44topApRebootsTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPsRebootReport44topApRebootsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportAPsRebootReport44TopApRebootsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport44topApRebootsTable200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// ReportAPsRebootReport45TopApRebootsOverTime
//
// Operation ID: report_APsRebootReport_45_topApRebootsOverTime
//
// APs Reboot Report - AP Reboots
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPsRebootReportService) ReportAPsRebootReport45TopApRebootsOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport45topApRebootsOverTime200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportAPsRebootReport45TopApRebootsOverTime, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport45topApRebootsOverTime200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
