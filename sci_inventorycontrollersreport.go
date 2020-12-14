package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
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
// Definition: InventoryControllersReport_InventoryControllersReport_96_krack_Data
type SCIInventoryControllersReport96krackData []*SCIInventoryControllersReport96krackDataType

func MakeSCIInventoryControllersReport96krackData() SCIInventoryControllersReport96krackData {
	m := make(SCIInventoryControllersReport96krackData, 0)
	return m
}

// SCIInventoryControllersReport96krackDataType
//
// Definition: InventoryControllersReport_InventoryControllersReport_96_krack_DataType
type SCIInventoryControllersReport96krackDataType struct {
	Number *string `json:"number,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	Recommendation *string `json:"recommendation,omitempty"`

	System *string `json:"system,omitempty"`

	SystemUrl *string `json:"systemUrl,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryControllersReport96krackDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryControllersReport96krackDataType SCIInventoryControllersReport96krackDataType
	tmpType := new(_SCIInventoryControllersReport96krackDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "number")
	delete(tmpType.XAdditionalProperties, "percentage")
	delete(tmpType.XAdditionalProperties, "recommendation")
	delete(tmpType.XAdditionalProperties, "system")
	delete(tmpType.XAdditionalProperties, "systemUrl")
	delete(tmpType.XAdditionalProperties, "zoneName")
	*t = SCIInventoryControllersReport96krackDataType(*tmpType)
	return nil
}

func (t *SCIInventoryControllersReport96krackDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Number != nil {
		tmp["number"] = t.Number
	}
	if t.Percentage != nil {
		tmp["percentage"] = t.Percentage
	}
	if t.Recommendation != nil {
		tmp["recommendation"] = t.Recommendation
	}
	if t.System != nil {
		tmp["system"] = t.System
	}
	if t.SystemUrl != nil {
		tmp["systemUrl"] = t.SystemUrl
	}
	if t.ZoneName != nil {
		tmp["zoneName"] = t.ZoneName
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryControllersReport96krackDataType() *SCIInventoryControllersReport96krackDataType {
	m := new(SCIInventoryControllersReport96krackDataType)
	return m
}

// SCIInventoryControllersReport96krackMetaData
//
// Definition: InventoryControllersReport_InventoryControllersReport_96_krack_MetaData
type SCIInventoryControllersReport96krackMetaData struct {
	Number *string `json:"number,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryControllersReport96krackMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventoryControllersReport96krackMetaData SCIInventoryControllersReport96krackMetaData
	tmpType := new(_SCIInventoryControllersReport96krackMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "number")
	delete(tmpType.XAdditionalProperties, "percentage")
	*t = SCIInventoryControllersReport96krackMetaData(*tmpType)
	return nil
}

func (t *SCIInventoryControllersReport96krackMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Number != nil {
		tmp["number"] = t.Number
	}
	if t.Percentage != nil {
		tmp["percentage"] = t.Percentage
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryControllersReport96krackMetaData() *SCIInventoryControllersReport96krackMetaData {
	m := new(SCIInventoryControllersReport96krackMetaData)
	return m
}

// SCIInventoryControllersReport98resourceUtilizationData
//
// Definition: InventoryControllersReport_InventoryControllersReport_98_resourceUtilization_Data
type SCIInventoryControllersReport98resourceUtilizationData []*SCIInventoryControllersReport98resourceUtilizationDataType

func MakeSCIInventoryControllersReport98resourceUtilizationData() SCIInventoryControllersReport98resourceUtilizationData {
	m := make(SCIInventoryControllersReport98resourceUtilizationData, 0)
	return m
}

// SCIInventoryControllersReport98resourceUtilizationDataType
//
// Definition: InventoryControllersReport_InventoryControllersReport_98_resourceUtilization_DataType
type SCIInventoryControllersReport98resourceUtilizationDataType struct {
	CpuUtilization *float64 `json:"cpuUtilization,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	DiskUtilization *float64 `json:"diskUtilization,omitempty"`

	MemoryUtilization *float64 `json:"memoryUtilization,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryControllersReport98resourceUtilizationDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryControllersReport98resourceUtilizationDataType SCIInventoryControllersReport98resourceUtilizationDataType
	tmpType := new(_SCIInventoryControllersReport98resourceUtilizationDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "cpuUtilization")
	delete(tmpType.XAdditionalProperties, "ctrlMac")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "ctrlSerial")
	delete(tmpType.XAdditionalProperties, "diskUtilization")
	delete(tmpType.XAdditionalProperties, "memoryUtilization")
	*t = SCIInventoryControllersReport98resourceUtilizationDataType(*tmpType)
	return nil
}

func (t *SCIInventoryControllersReport98resourceUtilizationDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.CpuUtilization != nil {
		tmp["cpuUtilization"] = t.CpuUtilization
	}
	if t.CtrlMac != nil {
		tmp["ctrlMac"] = t.CtrlMac
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.CtrlSerial != nil {
		tmp["ctrlSerial"] = t.CtrlSerial
	}
	if t.DiskUtilization != nil {
		tmp["diskUtilization"] = t.DiskUtilization
	}
	if t.MemoryUtilization != nil {
		tmp["memoryUtilization"] = t.MemoryUtilization
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryControllersReport98resourceUtilizationDataType() *SCIInventoryControllersReport98resourceUtilizationDataType {
	m := new(SCIInventoryControllersReport98resourceUtilizationDataType)
	return m
}

// SCIInventoryControllersReport99licenseUtilizationData
//
// Definition: InventoryControllersReport_InventoryControllersReport_99_licenseUtilization_Data
type SCIInventoryControllersReport99licenseUtilizationData []*SCIInventoryControllersReport99licenseUtilizationDataType

func MakeSCIInventoryControllersReport99licenseUtilizationData() SCIInventoryControllersReport99licenseUtilizationData {
	m := make(SCIInventoryControllersReport99licenseUtilizationData, 0)
	return m
}

// SCIInventoryControllersReport99licenseUtilizationDataType
//
// Definition: InventoryControllersReport_InventoryControllersReport_99_licenseUtilization_DataType
type SCIInventoryControllersReport99licenseUtilizationDataType struct {
	ApsDown *float64 `json:"apsDown,omitempty"`

	ApsManaged *float64 `json:"apsManaged,omitempty"`

	ApsUp *float64 `json:"apsUp,omitempty"`

	LicenseAvailable *float64 `json:"licenseAvailable,omitempty"`

	LicenseConsumed *float64 `json:"licenseConsumed,omitempty"`

	LicenseCount *float64 `json:"licenseCount,omitempty"`

	LicenseUtilization *float64 `json:"licenseUtilization,omitempty"`

	System *string `json:"system,omitempty"`

	SystemUrl *string `json:"systemUrl,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryControllersReport99licenseUtilizationDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryControllersReport99licenseUtilizationDataType SCIInventoryControllersReport99licenseUtilizationDataType
	tmpType := new(_SCIInventoryControllersReport99licenseUtilizationDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apsDown")
	delete(tmpType.XAdditionalProperties, "apsManaged")
	delete(tmpType.XAdditionalProperties, "apsUp")
	delete(tmpType.XAdditionalProperties, "licenseAvailable")
	delete(tmpType.XAdditionalProperties, "licenseConsumed")
	delete(tmpType.XAdditionalProperties, "licenseCount")
	delete(tmpType.XAdditionalProperties, "licenseUtilization")
	delete(tmpType.XAdditionalProperties, "system")
	delete(tmpType.XAdditionalProperties, "systemUrl")
	*t = SCIInventoryControllersReport99licenseUtilizationDataType(*tmpType)
	return nil
}

func (t *SCIInventoryControllersReport99licenseUtilizationDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApsDown != nil {
		tmp["apsDown"] = t.ApsDown
	}
	if t.ApsManaged != nil {
		tmp["apsManaged"] = t.ApsManaged
	}
	if t.ApsUp != nil {
		tmp["apsUp"] = t.ApsUp
	}
	if t.LicenseAvailable != nil {
		tmp["licenseAvailable"] = t.LicenseAvailable
	}
	if t.LicenseConsumed != nil {
		tmp["licenseConsumed"] = t.LicenseConsumed
	}
	if t.LicenseCount != nil {
		tmp["licenseCount"] = t.LicenseCount
	}
	if t.LicenseUtilization != nil {
		tmp["licenseUtilization"] = t.LicenseUtilization
	}
	if t.System != nil {
		tmp["system"] = t.System
	}
	if t.SystemUrl != nil {
		tmp["systemUrl"] = t.SystemUrl
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryControllersReport99licenseUtilizationDataType() *SCIInventoryControllersReport99licenseUtilizationDataType {
	m := new(SCIInventoryControllersReport99licenseUtilizationDataType)
	return m
}

// SCIInventoryControllersReport114controllerInventoryOverviewData
//
// Definition: InventoryControllersReport_InventoryControllersReport_114_controllerInventoryOverview_Data
type SCIInventoryControllersReport114controllerInventoryOverviewData []*SCIInventoryControllersReport114controllerInventoryOverviewDataType

func MakeSCIInventoryControllersReport114controllerInventoryOverviewData() SCIInventoryControllersReport114controllerInventoryOverviewData {
	m := make(SCIInventoryControllersReport114controllerInventoryOverviewData, 0)
	return m
}

// SCIInventoryControllersReport114controllerInventoryOverviewDataType
//
// Definition: InventoryControllersReport_InventoryControllersReport_114_controllerInventoryOverview_DataType
type SCIInventoryControllersReport114controllerInventoryOverviewDataType struct {
	Offline *float64 `json:"offline,omitempty"`

	Online *float64 `json:"online,omitempty"`

	SzCount *float64 `json:"szCount,omitempty"`

	Total *float64 `json:"total,omitempty"`

	ZdCount *float64 `json:"zdCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventoryControllersReport114controllerInventoryOverviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventoryControllersReport114controllerInventoryOverviewDataType SCIInventoryControllersReport114controllerInventoryOverviewDataType
	tmpType := new(_SCIInventoryControllersReport114controllerInventoryOverviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "offline")
	delete(tmpType.XAdditionalProperties, "online")
	delete(tmpType.XAdditionalProperties, "szCount")
	delete(tmpType.XAdditionalProperties, "total")
	delete(tmpType.XAdditionalProperties, "zdCount")
	*t = SCIInventoryControllersReport114controllerInventoryOverviewDataType(*tmpType)
	return nil
}

func (t *SCIInventoryControllersReport114controllerInventoryOverviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Offline != nil {
		tmp["offline"] = t.Offline
	}
	if t.Online != nil {
		tmp["online"] = t.Online
	}
	if t.SzCount != nil {
		tmp["szCount"] = t.SzCount
	}
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	if t.ZdCount != nil {
		tmp["zdCount"] = t.ZdCount
	}
	return json.Marshal(tmp)
}

func NewSCIInventoryControllersReport114controllerInventoryOverviewDataType() *SCIInventoryControllersReport114controllerInventoryOverviewDataType {
	m := new(SCIInventoryControllersReport114controllerInventoryOverviewDataType)
	return m
}

// ReportInventoryControllersReport96Krack
//
// Operation ID: report_InventoryControllersReport_96_krack
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportInventoryControllersReport96Krack, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_InventoryControllersReport_98_resourceUtilization
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportInventoryControllersReport98ResourceUtilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_InventoryControllersReport_99_licenseUtilization
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportInventoryControllersReport99LicenseUtilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
// Operation ID: report_InventoryControllersReport_114_controllerInventoryOverview
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportInventoryControllersReport114ControllerInventoryOverview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportInventoryControllersReport114controllerInventoryOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
