package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type SCIPCIProfileService struct {
	apiClient *SCIClient
}

func NewSCIPCIProfileService(c *SCIClient) *SCIPCIProfileService {
	s := new(SCIPCIProfileService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIPCIProfileService() *SCIPCIProfileService {
	return NewSCIPCIProfileService(ss.apiClient)
}

// SCIPCIProfileBatchDelete200ResponseType
//
// Definition: pciProfile_batchDelete200ResponseType
type SCIPCIProfileBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIPCIProfileBatchDelete200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIPCIProfileBatchDelete200ResponseType
}

func newSCIPCIProfileBatchDelete200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfileBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfileBatchDelete200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIPCIProfileBatchDelete200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIPCIProfileBatchDelete200ResponseType() *SCIPCIProfileBatchDelete200ResponseType {
	m := new(SCIPCIProfileBatchDelete200ResponseType)
	return m
}

// SCIPCIProfileFind200ResponseType
//
// Definition: pciProfile_find200ResponseType
type SCIPCIProfileFind200ResponseType []*SCIModelsPciProfile

type SCIPCIProfileFind200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIPCIProfileFind200ResponseType
}

func newSCIPCIProfileFind200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfileFind200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfileFind200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = make(SCIPCIProfileFind200ResponseType, 0)
	return json.NewDecoder(r).Decode(&r.Data)
}
func MakeSCIPCIProfileFind200ResponseType() SCIPCIProfileFind200ResponseType {
	m := make(SCIPCIProfileFind200ResponseType, 0)
	return m
}

// SCIPCIProfilePrototypecountreports200ResponseType
//
// Definition: pciProfile_prototype_count_reports200ResponseType
type SCIPCIProfilePrototypecountreports200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIPCIProfilePrototypecountreports200ResponseType
}

func newSCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIPCIProfilePrototypecountreports200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIPCIProfilePrototypecountreports200ResponseType() *SCIPCIProfilePrototypecountreports200ResponseType {
	m := new(SCIPCIProfilePrototypecountreports200ResponseType)
	return m
}

// SCIPCIProfilePrototypegetreports200ResponseType
//
// Definition: pciProfile_prototype_get_reports200ResponseType
type SCIPCIProfilePrototypegetreports200ResponseType []*SCIModelsPciReport

type SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIPCIProfilePrototypegetreports200ResponseType
}

func newSCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = make(SCIPCIProfilePrototypegetreports200ResponseType, 0)
	return json.NewDecoder(r).Decode(&r.Data)
}
func MakeSCIPCIProfilePrototypegetreports200ResponseType() SCIPCIProfilePrototypegetreports200ResponseType {
	m := make(SCIPCIProfilePrototypegetreports200ResponseType, 0)
	return m
}

// PciProfileBatchDelete
//
// Delete multiple PCI Profiles
//
// Operation ID: pciProfile_batchDelete
// Operation path: /pciProfiles/batchDelete
// Success code: 200 (OK)
//
// Form data parameters:
// - ids string
//		- required
func (s *SCIPCIProfileService) PciProfileBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIPCIProfileBatchDelete200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfileBatchDelete200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIPCIProfileBatchDelete200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciProfileBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIPCIProfileBatchDelete200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIPCIProfileBatchDelete200ResponseTypeAPIResponse), err
}

// PciProfileCreateWithRelations
//
// Create a PCI profile and generate a report.
//
// Operation ID: pciProfile_createWithRelations
// Operation path: /pciProfiles/createWithRelations
// Success code: 200 (OK)
//
// Form data parameters:
// - answers string
//		- required
// - name string
//		- required
// - ssids string
//		- required
func (s *SCIPCIProfileService) PciProfileCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciProfileCreateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PciProfileFind
//
// Find all instances of the model matched by filter from the data source.
//
// Operation ID: pciProfile_find
// Operation path: /pciProfiles
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIPCIProfileService) PciProfileFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIPCIProfileFind200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfileFind200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIPCIProfileFind200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfileFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIPCIProfileFind200ResponseTypeAPIResponse), err
}

// PciProfilePrototypeCountReports
//
// Counts reports of pciProfile.
//
// Operation ID: pciProfile_prototype_count_reports
// Operation path: /pciProfiles/{id}/reports/count
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - where string
//		- nullable
func (s *SCIPCIProfileService) PciProfilePrototypeCountReports(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfilePrototypeCountReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse), err
}

// PciProfilePrototypeCreateReports
//
// Creates a new instance in reports of this model.
//
// Operation ID: pciProfile_prototype_create_reports
// Operation path: /pciProfiles/{id}/reports
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsPciReport
//
// Required parameters:
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeCreateReports(ctx context.Context, data *SCIModelsPciReport, id string, mutators ...RequestMutator) (*SCIModelsPciReportAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsPciReportAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsPciReportAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciProfilePrototypeCreateReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsPciReportAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsPciReportAPIResponse), err
}

// PciProfilePrototypeDeleteReports
//
// Deletes all reports of this model.
//
// Operation ID: pciProfile_prototype_delete_reports
// Operation path: /pciProfiles/{id}/reports
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeDeleteReports(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSCIPciProfilePrototypeDeleteReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PciProfilePrototypeDestroyByIdReports
//
// Delete a related item by id for reports.
//
// Operation ID: pciProfile_prototype_destroyById_reports
// Operation path: /pciProfiles/{id}/reports/{fk}
// Success code: 204 (No Content)
//
// Required parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeDestroyByIdReports(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSCIPciProfilePrototypeDestroyByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PciProfilePrototypeFindByIdReports
//
// Find a related item by id for reports.
//
// Operation ID: pciProfile_prototype_findById_reports
// Operation path: /pciProfiles/{id}/reports/{fk}
// Success code: 200 (OK)
//
// Required parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeFindByIdReports(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*SCIModelsPciReportAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsPciReportAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsPciReportAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfilePrototypeFindByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsPciReportAPIResponse), err
}

// PciProfilePrototypeGetReports
//
// Queries reports of pciProfile.
//
// Operation ID: pciProfile_prototype_get_reports
// Operation path: /pciProfiles/{id}/reports
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIPCIProfileService) PciProfilePrototypeGetReports(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfilePrototypeGetReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse), err
}

// PciProfilePrototypeUpdateByIdReports
//
// Update a related item by id for reports.
//
// Operation ID: pciProfile_prototype_updateById_reports
// Operation path: /pciProfiles/{id}/reports/{fk}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsPciReport
//
// Required parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeUpdateByIdReports(ctx context.Context, data *SCIModelsPciReport, fk string, id string, mutators ...RequestMutator) (*SCIModelsPciReportAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsPciReportAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsPciReportAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIPciProfilePrototypeUpdateByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsPciReportAPIResponse), err
	}
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsPciReportAPIResponse), err
}

// PciProfileUpdateWithRelations
//
// Update a PCI profile and generate a report.
//
// Operation ID: pciProfile_updateWithRelations
// Operation path: /pciProfiles/{id}/updateWithRelations
// Success code: 200 (OK)
//
// Form data parameters:
// - answers string
//		- required
// - name string
//		- required
// - ssids string
//		- required
//
// Required parameters:
// - id float64
//		- required
func (s *SCIPCIProfileService) PciProfileUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIPciProfileUpdateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
