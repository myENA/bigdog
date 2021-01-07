package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
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

func newSCIPCIProfileBatchDelete200ResponseTypeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SCIPCIProfileBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
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

func newSCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
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

func MakeSCIPCIProfilePrototypegetreports200ResponseType() SCIPCIProfilePrototypegetreports200ResponseType {
	m := make(SCIPCIProfilePrototypegetreports200ResponseType, 0)
	return m
}

// PciProfileBatchDelete
//
// Operation ID: pciProfile_batchDelete
//
// Delete multiple PCI Profiles
//
// Form Data Parameters:
// - ids string
//		- required
func (s *SCIPCIProfileService) PciProfileBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIPCIProfileBatchDelete200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPCIProfileBatchDelete200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciProfileBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIPCIProfileBatchDelete200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileCreateWithRelations
//
// Operation ID: pciProfile_createWithRelations
//
// Create a PCI profile and generate a report.
//
// Form Data Parameters:
// - answers string
//		- required
// - name string
//		- required
// - ssids string
//		- required
func (s *SCIPCIProfileService) PciProfileCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciProfileCreateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileFind
//
// Operation ID: pciProfile_find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIPCIProfileService) PciProfileFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIPCIProfileFind200ResponseType, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIPCIProfileFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfileFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIPCIProfileFind200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PciProfilePrototypeCountReports
//
// Operation ID: pciProfile_prototype_count_reports
//
// Counts reports of pciProfile.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - where string
//		- nullable
func (s *SCIPCIProfileService) PciProfilePrototypeCountReports(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPCIProfilePrototypecountreports200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfilePrototypeCountReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIPCIProfilePrototypecountreports200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfilePrototypeCreateReports
//
// Operation ID: pciProfile_prototype_create_reports
//
// Creates a new instance in reports of this model.
//
// Request Body:
//	 - body *SCIModelsPciReport
//
// Required Parameters:
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeCreateReports(ctx context.Context, data *SCIModelsPciReport, id string, mutators ...RequestMutator) (*SCIModelsPciReportAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsPciReport
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciProfilePrototypeCreateReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsPciReport()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfilePrototypeDeleteReports
//
// Operation ID: pciProfile_prototype_delete_reports
//
// Deletes all reports of this model.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeDeleteReports(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSCIPciProfilePrototypeDeleteReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PciProfilePrototypeDestroyByIdReports
//
// Operation ID: pciProfile_prototype_destroyById_reports
//
// Delete a related item by id for reports.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeDestroyByIdReports(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSCIPciProfilePrototypeDestroyByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PciProfilePrototypeFindByIdReports
//
// Operation ID: pciProfile_prototype_findById_reports
//
// Find a related item by id for reports.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeFindByIdReports(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*SCIModelsPciReportAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsPciReport
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfilePrototypeFindByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsPciReport()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfilePrototypeGetReports
//
// Operation ID: pciProfile_prototype_get_reports
//
// Queries reports of pciProfile.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIPCIProfileService) PciProfilePrototypeGetReports(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIPCIProfilePrototypegetreports200ResponseType, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIPCIProfilePrototypegetreports200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIPciProfilePrototypeGetReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIPCIProfilePrototypegetreports200ResponseType()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PciProfilePrototypeUpdateByIdReports
//
// Operation ID: pciProfile_prototype_updateById_reports
//
// Update a related item by id for reports.
//
// Request Body:
//	 - body *SCIModelsPciReport
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeUpdateByIdReports(ctx context.Context, data *SCIModelsPciReport, fk string, id string, mutators ...RequestMutator) (*SCIModelsPciReportAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsPciReport
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIPciProfilePrototypeUpdateByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsPciReport()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileUpdateWithRelations
//
// Operation ID: pciProfile_updateWithRelations
//
// Update a PCI profile and generate a report.
//
// Form Data Parameters:
// - answers string
//		- required
// - name string
//		- required
// - ssids string
//		- required
//
// Required Parameters:
// - id float64
//		- required
func (s *SCIPCIProfileService) PciProfileUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIPciProfileUpdateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
