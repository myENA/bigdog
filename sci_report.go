package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type SCIreportService struct {
	apiClient *SCIClient
}

func NewSCIreportService(c *SCIClient) *SCIreportService {
	s := new(SCIreportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIreportService() *SCIreportService {
	return NewSCIreportService(ss.apiClient)
}

type SCIreportFind200ResponseType []*SCIModelsReport

func MakeSCIreportFind200ResponseType() SCIreportFind200ResponseType {
	m := make(SCIreportFind200ResponseType, 0)
	return m
}

type SCIreportGetData200ResponseType struct {
	Data []*SCIreportGetData200ResponseTypeDataType `json:"data,omitempty"`

	Metadata *SCIreportGetData200ResponseTypeMetadataType `json:"metadata,omitempty"`
}

func NewSCIreportGetData200ResponseType() *SCIreportGetData200ResponseType {
	m := new(SCIreportGetData200ResponseType)
	return m
}

type SCIreportGetData200ResponseTypeDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIreportGetData200ResponseTypeDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIreportGetData200ResponseTypeDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIreportGetData200ResponseTypeDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIreportGetData200ResponseTypeDataType() *SCIreportGetData200ResponseTypeDataType {
	m := new(SCIreportGetData200ResponseTypeDataType)
	return m
}

type SCIreportGetData200ResponseTypeMetadataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIreportGetData200ResponseTypeMetadataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIreportGetData200ResponseTypeMetadataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIreportGetData200ResponseTypeMetadataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIreportGetData200ResponseTypeMetadataType() *SCIreportGetData200ResponseTypeMetadataType {
	m := new(SCIreportGetData200ResponseTypeMetadataType)
	return m
}

type SCIreportWithRelations200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIreportWithRelations200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIreportWithRelations200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIreportWithRelations200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIreportWithRelations200ResponseType() *SCIreportWithRelations200ResponseType {
	m := new(SCIreportWithRelations200ResponseType)
	return m
}

// ReportDownloadReport
//
// Form Data Parameters:
// - state string
//		- required
// - timezone string
//		- required
//
// Required Parameters:
// - format string
//		- required
// - id string
//		- required
func (s *SCIreportService) ReportDownloadReport(ctx context.Context, formValues url.Values, format string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportDownloadReport, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return rm, err
	}
	req.SetPathParameter("format", format)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *SCIreportService) ReportFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIreportFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIreportFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIreportFind200ResponseType()
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
func (s *SCIreportService) ReportFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFindById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportGetData
//
// For the <b><code>filter</code></b> field below, an example would be <pre><code class="json">{ "type": "or", "fields": [{ "type": "selector", "dimension": "apMac", "value": "000000000000" }]}</code></pre>
//
// Form Data Parameters:
// - end string
//		- required
// - filter string
//		- nullable
// - granularity string
//		- nullable
// - limit float64
//		- nullable
// - metric string
//		- nullable
// - pagingIdentifiers string
//		- nullable
// - start string
//		- required
// - switchFilter string
//		- nullable
//
// Required Parameters:
// - id string
//		- required
// - sectionId string
//		- required
func (s *SCIreportService) ReportGetData(ctx context.Context, formValues url.Values, id string, sectionId string, mutators ...RequestMutator) (*SCIreportGetData200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIreportGetData200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportGetData, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("sectionId", sectionId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIreportGetData200ResponseType()
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
func (s *SCIreportService) ReportPrototypeGetSections(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIReportgetsections200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIReportgetsections200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIReportPrototypeGetSections, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIReportgetsections200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ReportWithRelations
//
// For the <b><code>urlSegmentName</code></b> field below, examples could be <code>overview</code>, <code>network</code>, <code>ap</code>, <code>clients</code>
//
// Form Data Parameters:
// - urlSegmentName string
//		- required
func (s *SCIreportService) ReportWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIreportWithRelations200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIreportWithRelations200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWithRelations, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIreportWithRelations200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
