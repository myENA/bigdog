package bigdog

// API Version: 1.0.0

import (
	"context"
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

	TotalReboots *int `json:"totalReboots,omitempty"`
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

	Count *int `json:"count,omitempty"`

	Index *int `json:"index,omitempty"`
}

func NewSCIAPsRebootReport44topApRebootsTableDataType() *SCIAPsRebootReport44topApRebootsTableDataType {
	m := new(SCIAPsRebootReport44topApRebootsTableDataType)
	return m
}

// SCIAPsRebootReport44topApRebootsTableMetaData
//
// Definition: APsRebootReport_APsRebootReport_44_topApRebootsTable_MetaData
type SCIAPsRebootReport44topApRebootsTableMetaData struct {
	TotalCount *int `json:"totalCount,omitempty"`
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

	R320297594CB1CD0615F0 *int `json:"r320.29759 (4C:B1:CD:06:15:F0),omitempty"`

	R500E0107F25FC10 *int `json:"R500 (E0:10:7F:25:FC:10),omitempty"`

	R610297591C3A600152C0 *int `json:"r610.29759 (1C:3A:60:01:52:C0),omitempty"`

	R710297592CC5D3098780 *int `json:"r710.29759 (2C:C5:D3:09:87:80),omitempty"`

	R75029759C803F5046120 *int `json:"r750.29759 (C8:03:F5:04:61:20),omitempty"`

	RuckusAP34FA9F1B7F80 *int `json:"RuckusAP (34:FA:9F:1B:7F:80),omitempty"`

	RuckusAP34FA9F1B81C0 *int `json:"RuckusAP (34:FA:9F:1B:81:C0),omitempty"`

	RuckusAPC803F50468F0 *int `json:"RuckusAP (C8:03:F5:04:68:F0),omitempty"`

	RuckusAPE0107F288420 *int `json:"RuckusAP (E0:10:7F:28:84:20),omitempty"`

	Start *string `json:"start,omitempty"`

	T61029759187C0B0C9D60 *int `json:"t610.29759 (18:7C:0B:0C:9D:60),omitempty"`
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
func (s *SCIAPsRebootReportService) ReportAPsRebootReport43TotalReboots(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport43totalReboots200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPsRebootReport43TotalReboots, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport43totalReboots200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *SCIAPsRebootReportService) ReportAPsRebootReport44TopApRebootsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport44topApRebootsTable200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPsRebootReport44TopApRebootsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport44topApRebootsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *SCIAPsRebootReportService) ReportAPsRebootReport45TopApRebootsOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPsRebootReport45topApRebootsOverTime200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPsRebootReport45TopApRebootsOverTime, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPsRebootReport45topApRebootsOverTime200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
