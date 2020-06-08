package ruckus

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIPciprofileService struct {
	apiClient *SCIClient
}

func NewSCIPciprofileService(c *SCIClient) *SCIPciprofileService {
	s := new(SCIPciprofileService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIPciprofileService() *SCIPciprofileService {
	return NewSCIPciprofileService(ss.apiClient)
}

type SCIPciprofileBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIPciprofileBatchDelete200ResponseType() *SCIPciprofileBatchDelete200ResponseType {
	m := new(SCIPciprofileBatchDelete200ResponseType)
	return m
}

type SCIPciprofileCreateWithRelations200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIPciprofileCreateWithRelations200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIPciprofileCreateWithRelations200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIPciprofileCreateWithRelations200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIPciprofileCreateWithRelations200ResponseType() *SCIPciprofileCreateWithRelations200ResponseType {
	m := new(SCIPciprofileCreateWithRelations200ResponseType)
	return m
}

type SCIPciprofileFind200ResponseType []*SCIModelsPciProfile

func MakeSCIPciprofileFind200ResponseType() SCIPciprofileFind200ResponseType {
	m := make(SCIPciprofileFind200ResponseType, 0)
	return m
}

type SCIPciprofileUpdateWithRelations200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIPciprofileUpdateWithRelations200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIPciprofileUpdateWithRelations200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIPciprofileUpdateWithRelations200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIPciprofileUpdateWithRelations200ResponseType() *SCIPciprofileUpdateWithRelations200ResponseType {
	m := new(SCIPciprofileUpdateWithRelations200ResponseType)
	return m
}

// PciProfileBatchDelete
//
// Delete multiple PCI Profiles
//
// Request Body:
//	 - body string
func (s *SCIPciprofileService) PciProfileBatchDelete(ctx context.Context, body string) (*SCIPciprofileBatchDelete200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPciprofileBatchDelete200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIPciProfileBatchDelete, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIPciprofileBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileCreateWithRelations
//
// Create a PCI profile and generate a report.
//
// Request Body:
//	 - body string
func (s *SCIPciprofileService) PciProfileCreateWithRelations(ctx context.Context, body string) (*SCIPciprofileCreateWithRelations200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPciprofileCreateWithRelations200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIPciProfileCreateWithRelations, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIPciprofileCreateWithRelations200ResponseType()
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
func (s *SCIPciprofileService) PciProfileFind(ctx context.Context, optionalParams map[string][]string) (SCIPciprofileFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIPciprofileFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfileFind, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIPciprofileFind200ResponseType()
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
func (s *SCIPciprofileService) PciProfilePrototypeCountReports(ctx context.Context, id string, optionalParams map[string][]string) (*SCIPciprofileprototypecountreports200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPciprofileprototypecountreports200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfilePrototypeCountReports, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["where"]; ok && len(v) > 0 {
		req.SetQueryParameter("where", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIPciprofileprototypecountreports200ResponseType()
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
func (s *SCIPciprofileService) PciProfilePrototypeCreateReports(ctx context.Context, body *SCIModelsPciReport, id string) (*SCIModelsPciReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIPciProfilePrototypeCreateReports, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SCIPciprofileService) PciProfilePrototypeDeleteReports(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIPciProfilePrototypeDeleteReports, false)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SCIPciprofileService) PciProfilePrototypeDestroyByIdReports(ctx context.Context, fk string, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCIPciProfilePrototypeDestroyByIdReports, false)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SCIPciprofileService) PciProfilePrototypeFindByIdReports(ctx context.Context, fk string, id string) (*SCIModelsPciReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfilePrototypeFindByIdReports, false)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SCIPciprofileService) PciProfilePrototypeGetReports(ctx context.Context, id string, optionalParams map[string][]string) (SCIPciprofileprototypegetreports200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIPciprofileprototypegetreports200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIPciProfilePrototypeGetReports, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIPciprofileprototypegetreports200ResponseType()
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
func (s *SCIPciprofileService) PciProfilePrototypeUpdateByIdReports(ctx context.Context, body *SCIModelsPciReport, fk string, id string) (*SCIModelsPciReport, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCIPciProfilePrototypeUpdateByIdReports, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("fk", fk)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsPciReport()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PciProfileUpdateWithRelations
//
// Update a PCI profile and generate a report.
//
// Request Body:
//	 - body string
//
// Required Parameters:
// - id float64
//		- required
func (s *SCIPciprofileService) PciProfileUpdateWithRelations(ctx context.Context, body string, id string) (*SCIPciprofileUpdateWithRelations200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIPciprofileUpdateWithRelations200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIPciProfileUpdateWithRelations, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIPciprofileUpdateWithRelations200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
