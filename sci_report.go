package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
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

// SCIReportFind200ResponseType
//
// Definition: report.find200ResponseType
type SCIReportFind200ResponseType []*SCIModelsReport

func MakeSCIReportFind200ResponseType() SCIReportFind200ResponseType {
	m := make(SCIReportFind200ResponseType, 0)
	return m
}

// SCIReportGetData200ResponseType
//
// Definition: report.getData200ResponseType
type SCIReportGetData200ResponseType struct {
	Data []interface{} `json:"data,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

func NewSCIReportGetData200ResponseType() *SCIReportGetData200ResponseType {
	m := new(SCIReportGetData200ResponseType)
	return m
}

// SCIReportQueryFilter
//
// Definition: report.queryFilter
//
// Data query filter body
type SCIReportQueryFilter struct {
	AlphaNumeric *bool `json:"alphaNumeric,omitempty"`

	Dimension *string `json:"dimension,omitempty"`

	Filter []*SCIReportQueryFilter `json:"filter,omitempty"`

	Function *string `json:"function,omitempty"`

	Lower *string `json:"lower,omitempty"`

	LowerStrict *bool `json:"lowerStrict,omitempty"`

	Pattern *string `json:"pattern,omitempty"`

	// Type
	// Type of filter
	// Constraints:
	//    - required
	Type *string `json:"type"`

	Upper *string `json:"upper,omitempty"`

	UpperStrict *bool `json:"upperStrict,omitempty"`

	Value *string `json:"value,omitempty"`

	Values []string `json:"values,omitempty"`
}

func NewSCIReportQueryFilter() *SCIReportQueryFilter {
	m := new(SCIReportQueryFilter)
	return m
}

// SCIReportSearchQuery
//
// Definition: report.searchQuery
//
// Search query
type SCIReportSearchQuery struct {
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// Type
	// Constraints:
	//    - required
	Type *string `json:"type"`

	Value *string `json:"value,omitempty"`

	Values []string `json:"values,omitempty"`
}

func NewSCIReportSearchQuery() *SCIReportSearchQuery {
	m := new(SCIReportSearchQuery)
	return m
}

// ReportDownloadReport
//
// Operation ID: report.downloadReport
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
func (s *SCIReportService) ReportDownloadReport(ctx context.Context, formValues url.Values, format string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
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
// Operation ID: report.find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIReportService) ReportFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIReportFind200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIReportFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIReportFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ReportFindById
//
// Operation ID: report.findById
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
func (s *SCIReportService) ReportFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsReport, *APIResponseMeta, error) {
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
// Operation ID: report.getData
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
func (s *SCIReportService) ReportGetData(ctx context.Context, formValues url.Values, id string, sectionId string, mutators ...RequestMutator) (*SCIReportGetData200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIReportGetData, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("sectionId", sectionId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportGetData200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportPrototypeGetSections
//
// Operation ID: report.prototype.__get__sections
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
func (s *SCIReportService) ReportPrototypeGetSections(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIReportgetsections200ResponseType, *APIResponseMeta, error) {
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
// Operation ID: report.withRelations
//
// For the <b><code>urlSegmentName</code></b> field below, examples could be <code>overview</code>, <code>network</code>, <code>ap</code>, <code>clients</code>
//
// Required Parameters:
// - urlSegmentName string
//		- required
func (s *SCIReportService) ReportWithRelations(ctx context.Context, urlSegmentName string, mutators ...RequestMutator) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportWithRelations, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetQueryParameter("urlSegmentName", []string{urlSegmentName})
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
