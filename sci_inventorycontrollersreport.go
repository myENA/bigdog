package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIInventoryControllersReportService struct {
	apiClient *SCIClient
}

func NewSCIInventoryControllersReportService(c *SCIClient) *SCIInventoryControllersReportService {
	s := new(SCIInventoryControllersReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIInventoryControllersReportService() *SCIInventoryControllersReportService {
	return NewSCIInventoryControllersReportService(ss.apiClient)
}

// SCIInventoryControllersReport96krackData
//
// Definition: InventoryControllersReport.InventoryControllersReport.96.krack.Data
type SCIInventoryControllersReport96krackData []*SCIInventoryControllersReport96krackDataType

func MakeSCIInventoryControllersReport96krackData() SCIInventoryControllersReport96krackData {
	m := make(SCIInventoryControllersReport96krackData, 0)
	return m
}

// SCIInventoryControllersReport96krackDataType
//
// Definition: InventoryControllersReport.InventoryControllersReport.96.krack.DataType
type SCIInventoryControllersReport96krackDataType struct {
	Number *string `json:"number,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	Recommendation *string `json:"recommendation,omitempty"`

	System *string `json:"system,omitempty"`

	SystemUrl *string `json:"systemUrl,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`
}

func NewSCIInventoryControllersReport96krackDataType() *SCIInventoryControllersReport96krackDataType {
	m := new(SCIInventoryControllersReport96krackDataType)
	return m
}

// SCIInventoryControllersReport96krackMetaData
//
// Definition: InventoryControllersReport.InventoryControllersReport.96.krack.MetaData
type SCIInventoryControllersReport96krackMetaData struct {
	Number *string `json:"number,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`
}

func NewSCIInventoryControllersReport96krackMetaData() *SCIInventoryControllersReport96krackMetaData {
	m := new(SCIInventoryControllersReport96krackMetaData)
	return m
}

// SCIInventoryControllersReport98resourceUtilizationData
//
// Definition: InventoryControllersReport.InventoryControllersReport.98.resourceUtilization.Data
type SCIInventoryControllersReport98resourceUtilizationData []*SCIInventoryControllersReport98resourceUtilizationDataType

func MakeSCIInventoryControllersReport98resourceUtilizationData() SCIInventoryControllersReport98resourceUtilizationData {
	m := make(SCIInventoryControllersReport98resourceUtilizationData, 0)
	return m
}

// SCIInventoryControllersReport98resourceUtilizationDataType
//
// Definition: InventoryControllersReport.InventoryControllersReport.98.resourceUtilization.DataType
type SCIInventoryControllersReport98resourceUtilizationDataType struct {
	CpuUtilization *float64 `json:"cpuUtilization,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	DiskUtilization *float64 `json:"diskUtilization,omitempty"`

	MemoryUtilization *float64 `json:"memoryUtilization,omitempty"`
}

func NewSCIInventoryControllersReport98resourceUtilizationDataType() *SCIInventoryControllersReport98resourceUtilizationDataType {
	m := new(SCIInventoryControllersReport98resourceUtilizationDataType)
	return m
}

// SCIInventoryControllersReport99licenseUtilizationData
//
// Definition: InventoryControllersReport.InventoryControllersReport.99.licenseUtilization.Data
type SCIInventoryControllersReport99licenseUtilizationData []*SCIInventoryControllersReport99licenseUtilizationDataType

func MakeSCIInventoryControllersReport99licenseUtilizationData() SCIInventoryControllersReport99licenseUtilizationData {
	m := make(SCIInventoryControllersReport99licenseUtilizationData, 0)
	return m
}

// SCIInventoryControllersReport99licenseUtilizationDataType
//
// Definition: InventoryControllersReport.InventoryControllersReport.99.licenseUtilization.DataType
type SCIInventoryControllersReport99licenseUtilizationDataType struct {
	ApsDown *int `json:"apsDown,omitempty"`

	ApsManaged *int `json:"apsManaged,omitempty"`

	ApsUp *int `json:"apsUp,omitempty"`

	LicenseAvailable *int `json:"licenseAvailable,omitempty"`

	LicenseConsumed *int `json:"licenseConsumed,omitempty"`

	LicenseCount *int `json:"licenseCount,omitempty"`

	LicenseUtilization *float64 `json:"licenseUtilization,omitempty"`

	System *string `json:"system,omitempty"`

	SystemUrl *string `json:"systemUrl,omitempty"`
}

func NewSCIInventoryControllersReport99licenseUtilizationDataType() *SCIInventoryControllersReport99licenseUtilizationDataType {
	m := new(SCIInventoryControllersReport99licenseUtilizationDataType)
	return m
}

// SCIInventoryControllersReport114controllerInventoryOverviewData
//
// Definition: InventoryControllersReport.InventoryControllersReport.114.controllerInventoryOverview.Data
type SCIInventoryControllersReport114controllerInventoryOverviewData []*SCIInventoryControllersReport114controllerInventoryOverviewDataType

func MakeSCIInventoryControllersReport114controllerInventoryOverviewData() SCIInventoryControllersReport114controllerInventoryOverviewData {
	m := make(SCIInventoryControllersReport114controllerInventoryOverviewData, 0)
	return m
}

// SCIInventoryControllersReport114controllerInventoryOverviewDataType
//
// Definition: InventoryControllersReport.InventoryControllersReport.114.controllerInventoryOverview.DataType
type SCIInventoryControllersReport114controllerInventoryOverviewDataType struct {
	Offline *int `json:"offline,omitempty"`

	Online *int `json:"online,omitempty"`

	SzCount *int `json:"szCount,omitempty"`

	Total *int `json:"total,omitempty"`

	ZdCount *int `json:"zdCount,omitempty"`
}

func NewSCIInventoryControllersReport114controllerInventoryOverviewDataType() *SCIInventoryControllersReport114controllerInventoryOverviewDataType {
	m := new(SCIInventoryControllersReport114controllerInventoryOverviewDataType)
	return m
}

// ReportInventoryControllersReport96Krack
//
// Operation ID: report.InventoryControllersReport.96.krack
//
// Inventory - Controllers Report - KRACK Assessment
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport96Krack(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport96krack200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport96krack200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport96Krack, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport96krack200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryControllersReport98ResourceUtilization
//
// Operation ID: report.InventoryControllersReport.98.resourceUtilization
//
// Inventory - Controllers Report - Resource Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport98ResourceUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport98resourceUtilization200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport98resourceUtilization200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport98ResourceUtilization, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport98resourceUtilization200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryControllersReport99LicenseUtilization
//
// Operation ID: report.InventoryControllersReport.99.licenseUtilization
//
// Inventory - Controllers Report - License Utilization
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport99LicenseUtilization(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport99licenseUtilization200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport99licenseUtilization200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport99LicenseUtilization, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport99licenseUtilization200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportInventoryControllersReport114ControllerInventoryOverview
//
// Operation ID: report.InventoryControllersReport.114.controllerInventoryOverview
//
// Inventory - Controllers Report - Overview
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventoryControllersReportService) ReportInventoryControllersReport114ControllerInventoryOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportInventoryControllersReport114ControllerInventoryOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
