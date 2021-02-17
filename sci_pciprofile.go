package bigdog

// API Version: 1.0.0

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
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

func newSCIPCIProfileBatchDelete200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfileBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfileBatchDelete200ResponseTypeAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SCIPCIProfileBatchDelete200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
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

func newSCIPCIProfileFind200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfileFind200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfileFind200ResponseTypeAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := make(SCIPCIProfileFind200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return err
	}
	r.Data = data
	return nil
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

func newSCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SCIPCIProfilePrototypecountreports200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
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

func newSCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := make(SCIPCIProfilePrototypegetreports200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return err
	}
	r.Data = data
	return nil
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfileBatchDelete200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIPciProfileBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIPciProfileCreateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfileFind200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIPciProfileFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfilePrototypecountreports200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIPciProfilePrototypeCountReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("where", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsPciReportAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIPciProfilePrototypeCreateReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodDelete, RouteSCIPciProfilePrototypeDeleteReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodDelete, RouteSCIPciProfilePrototypeDestroyByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsPciReportAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIPciProfilePrototypeFindByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIPCIProfilePrototypegetreports200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIPciProfilePrototypeGetReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsPciReportAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIPciProfilePrototypeUpdateByIdReports, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("fk", fk)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIPciProfileUpdateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
