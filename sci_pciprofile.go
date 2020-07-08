package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
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
// Definition: pciProfile.batchDelete200ResponseType
type SCIPCIProfileBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIPCIProfileBatchDelete200ResponseType() *SCIPCIProfileBatchDelete200ResponseType {
	m := new(SCIPCIProfileBatchDelete200ResponseType)
	return m
}

// SCIPCIProfileFind200ResponseType
//
// Definition: pciProfile.find200ResponseType
type SCIPCIProfileFind200ResponseType []*SCIModelsPciProfile

func MakeSCIPCIProfileFind200ResponseType() SCIPCIProfileFind200ResponseType {
	m := make(SCIPCIProfileFind200ResponseType, 0)
	return m
}

// SCIPCIProfilecountreports200ResponseType
//
// Definition: pciProfile.prototype.__count__reports200ResponseType
type SCIPCIProfilecountreports200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIPCIProfilecountreports200ResponseType() *SCIPCIProfilecountreports200ResponseType {
	m := new(SCIPCIProfilecountreports200ResponseType)
	return m
}

// SCIPCIProfilegetreports200ResponseType
//
// Definition: pciProfile.prototype.__get__reports200ResponseType
type SCIPCIProfilegetreports200ResponseType []*SCIModelsPciReport

func MakeSCIPCIProfilegetreports200ResponseType() SCIPCIProfilegetreports200ResponseType {
	m := make(SCIPCIProfilegetreports200ResponseType, 0)
	return m
}

// PciProfileBatchDelete
//
// Operation ID: pciProfile.batchDelete
//
// Delete multiple PCI Profiles
//
// Form Data Parameters:
// - ids string
//		- required
func (s *SCIPCIProfileService) PciProfileBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIPCIProfileBatchDelete200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIPciProfileBatchDelete, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIPCIProfileBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileCreateWithRelations
//
// Operation ID: pciProfile.createWithRelations
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
func (s *SCIPCIProfileService) PciProfileCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIPciProfileCreateWithRelations, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PciProfileFind
//
// Operation ID: pciProfile.find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIPCIProfileService) PciProfileFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIPCIProfileFind200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfileFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIPCIProfileFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PciProfilePrototypeCountReports
//
// Operation ID: pciProfile.prototype.__count__reports
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
func (s *SCIPCIProfileService) PciProfilePrototypeCountReports(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIPCIProfilecountreports200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPCIProfilecountreports200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfilePrototypeCountReports, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIPCIProfilecountreports200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfilePrototypeCreateReports
//
// Operation ID: pciProfile.prototype.__create__reports
//
// Creates a new instance in reports of this model.
//
// Request Body:
//	 - body *SCIModelsPciReport
//
// Required Parameters:
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeCreateReports(ctx context.Context, data *SCIModelsPciReport, id string, mutators ...RequestMutator) (*SCIModelsPciReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIPciProfilePrototypeCreateReports, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsPciReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfilePrototypeDeleteReports
//
// Operation ID: pciProfile.prototype.__delete__reports
//
// Deletes all reports of this model.
//
// Required Parameters:
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeDeleteReports(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIPciProfilePrototypeDeleteReports, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PciProfilePrototypeDestroyByIdReports
//
// Operation ID: pciProfile.prototype.__destroyById__reports
//
// Delete a related item by id for reports.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeDestroyByIdReports(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIPciProfilePrototypeDestroyByIdReports, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PciProfilePrototypeFindByIdReports
//
// Operation ID: pciProfile.prototype.__findById__reports
//
// Find a related item by id for reports.
//
// Required Parameters:
// - fk string
//		- required
// - id string
//		- required
func (s *SCIPCIProfileService) PciProfilePrototypeFindByIdReports(ctx context.Context, fk string, id string, mutators ...RequestMutator) (*SCIModelsPciReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfilePrototypeFindByIdReports, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsPciReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfilePrototypeGetReports
//
// Operation ID: pciProfile.prototype.__get__reports
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
func (s *SCIPCIProfileService) PciProfilePrototypeGetReports(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (SCIPCIProfilegetreports200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIPCIProfilegetreports200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfilePrototypeGetReports, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIPCIProfilegetreports200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PciProfilePrototypeUpdateByIdReports
//
// Operation ID: pciProfile.prototype.__updateById__reports
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
func (s *SCIPCIProfileService) PciProfilePrototypeUpdateByIdReports(ctx context.Context, data *SCIModelsPciReport, fk string, id string, mutators ...RequestMutator) (*SCIModelsPciReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIPciProfilePrototypeUpdateByIdReports, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsPciReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileUpdateWithRelations
//
// Operation ID: pciProfile.updateWithRelations
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
func (s *SCIPCIProfileService) PciProfileUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIPciProfileUpdateWithRelations, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
