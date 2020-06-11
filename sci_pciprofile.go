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

type SCIPCIProfileBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIPCIProfileBatchDelete200ResponseType() *SCIPCIProfileBatchDelete200ResponseType {
	m := new(SCIPCIProfileBatchDelete200ResponseType)
	return m
}

type SCIPCIProfileCreateWithRelations200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIPCIProfileCreateWithRelations200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIPCIProfileCreateWithRelations200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIPCIProfileCreateWithRelations200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIPCIProfileCreateWithRelations200ResponseType() *SCIPCIProfileCreateWithRelations200ResponseType {
	m := new(SCIPCIProfileCreateWithRelations200ResponseType)
	return m
}

type SCIPCIProfileFind200ResponseType []*SCIModelsPciProfile

func MakeSCIPCIProfileFind200ResponseType() SCIPCIProfileFind200ResponseType {
	m := make(SCIPCIProfileFind200ResponseType, 0)
	return m
}

type SCIPCIProfilecountreports200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIPCIProfilecountreports200ResponseType() *SCIPCIProfilecountreports200ResponseType {
	m := new(SCIPCIProfilecountreports200ResponseType)
	return m
}

type SCIPCIProfilegetreports200ResponseType []*SCIModelsPciReport

func MakeSCIPCIProfilegetreports200ResponseType() SCIPCIProfilegetreports200ResponseType {
	m := make(SCIPCIProfilegetreports200ResponseType, 0)
	return m
}

type SCIPCIProfileUpdateWithRelations200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIPCIProfileUpdateWithRelations200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIPCIProfileUpdateWithRelations200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIPCIProfileUpdateWithRelations200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIPCIProfileUpdateWithRelations200ResponseType() *SCIPCIProfileUpdateWithRelations200ResponseType {
	m := new(SCIPCIProfileUpdateWithRelations200ResponseType)
	return m
}

// PciProfileBatchDelete
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
// Create a PCI profile and generate a report.
//
// Form Data Parameters:
// - answers string
//		- required
// - name string
//		- required
// - ssids string
//		- required
func (s *SCIPCIProfileService) PciProfileCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIPCIProfileCreateWithRelations200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPCIProfileCreateWithRelations200ResponseType
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
	resp = NewSCIPCIProfileCreateWithRelations200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileFind
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PciProfilePrototypeDestroyByIdReports
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PciProfilePrototypeFindByIdReports
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
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
func (s *SCIPCIProfileService) PciProfileUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*SCIPCIProfileUpdateWithRelations200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPCIProfileUpdateWithRelations200ResponseType
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
	resp = NewSCIPCIProfileUpdateWithRelations200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
