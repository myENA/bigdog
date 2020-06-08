package ruckus

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIReportService struct {
	apiClient *SCIClient
}

func NewSCIReportService(c *SCIClient) *SCIReportService {
	s := new(SCIReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIReportService() *SCIReportService {
	return NewSCIReportService(ss.apiClient)
}

type SCIReportFind200ResponseType []*SCIModelsReport

func MakeSCIReportFind200ResponseType() SCIReportFind200ResponseType {
	m := make(SCIReportFind200ResponseType, 0)
	return m
}

type SCIReportGetData200ResponseType struct {
	Data []**SCIReportGetData200ResponseType `json:"data,omitempty"`

	Metadata **SCIReportGetData200ResponseType `json:"metadata,omitempty"`
}

func NewSCIReportGetData200ResponseType() *SCIReportGetData200ResponseType {
	m := new(SCIReportGetData200ResponseType)
	return m
}

type SCIReportGetData200ResponseTypeDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportGetData200ResponseTypeDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIReportGetData200ResponseTypeDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIReportGetData200ResponseTypeDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIReportGetData200ResponseTypeDataType() *SCIReportGetData200ResponseTypeDataType {
	m := new(SCIReportGetData200ResponseTypeDataType)
	return m
}

type SCIReportGetData200ResponseTypeMetadataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportGetData200ResponseTypeMetadataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIReportGetData200ResponseTypeMetadataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIReportGetData200ResponseTypeMetadataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIReportGetData200ResponseTypeMetadataType() *SCIReportGetData200ResponseTypeMetadataType {
	m := new(SCIReportGetData200ResponseTypeMetadataType)
	return m
}

type SCIReportWithRelations200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIReportWithRelations200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIReportWithRelations200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIReportWithRelations200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIReportWithRelations200ResponseType() *SCIReportWithRelations200ResponseType {
	m := new(SCIReportWithRelations200ResponseType)
	return m
}

// ReportDownloadReport
//
// Request Body:
//	 - body string
//
// Required Parameters:
// - format string
//		- required
// - id string
//		- required
func (s *SCIReportService) ReportDownloadReport(ctx context.Context, body string, format string, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportDownloadReport, false)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("format", format)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// ReportFind
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportFind(ctx context.Context, optionalParams map[string][]string) (SCIReportFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIReportFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFind, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIReportFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ReportFindById
//
// Find a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportFindById(ctx context.Context, id string, optionalParams map[string][]string) (*SCIModelsReport, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsReport
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFindById, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportGetData
//
// For the <b><code>filter</code></b> field below, an example would be <pre><code class="json">{ "type": "or", "fields": [{ "type": "selector", "dimension": "apMac", "value": "000000000000" }]}</code></pre>
//
// Request Body:
//	 - body string
//
// Required Parameters:
// - id string
//		- required
// - sectionId string
//		- required
func (s *SCIReportService) ReportGetData(ctx context.Context, body string, id string, sectionId string) (*SCIReportGetData200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportGetData200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportGetData, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("sectionId", sectionId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIReportGetData200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportPrototypeGetSections
//
// Queries sections of report.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportPrototypeGetSections(ctx context.Context, id string, optionalParams map[string][]string) (SCIReportprototypegetsections200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIReportprototypegetsections200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportPrototypeGetSections, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIReportprototypegetsections200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ReportWithRelations
//
// For the <b><code>urlSegmentName</code></b> field below, examples could be <code>overview</code>, <code>network</code>, <code>ap</code>, <code>clients</code>
//
// Request Body:
//	 - body string
func (s *SCIReportService) ReportWithRelations(ctx context.Context, body string) (*SCIReportWithRelations200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportWithRelations200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWithRelations, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIReportWithRelations200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
