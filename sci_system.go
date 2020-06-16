package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCISystemService struct {
	apiClient *SCIClient
}

func NewSCISystemService(c *SCIClient) *SCISystemService {
	s := new(SCISystemService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCISystemService() *SCISystemService {
	return NewSCISystemService(ss.apiClient)
}

// SCISystemDeleteById200ResponseType
//
// Definition: system.deleteById200ResponseType
type SCISystemDeleteById200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISystemDeleteById200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCISystemDeleteById200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCISystemDeleteById200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCISystemDeleteById200ResponseType() *SCISystemDeleteById200ResponseType {
	m := new(SCISystemDeleteById200ResponseType)
	return m
}

// SCISystemFind200ResponseType
//
// Definition: system.find200ResponseType
type SCISystemFind200ResponseType []*SCIModelsSystem

func MakeSCISystemFind200ResponseType() SCISystemFind200ResponseType {
	m := make(SCISystemFind200ResponseType, 0)
	return m
}

// SCISystemGetSsids200ResponseType
//
// Definition: system.getSsids200ResponseType
type SCISystemGetSsids200ResponseType []interface{}

func MakeSCISystemGetSsids200ResponseType() SCISystemGetSsids200ResponseType {
	m := make(SCISystemGetSsids200ResponseType, 0)
	return m
}

// SystemCreate
//
// Operation ID: system.create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsSystem
func (s *SCISystemService) SystemCreate(ctx context.Context, data *SCIModelsSystem, mutators ...RequestMutator) (*SCIModelsSystem, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSystem
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCISystemCreate, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSystem()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SystemDeleteById
//
// Operation ID: system.deleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCISystemService) SystemDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*SCISystemDeleteById200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCISystemDeleteById200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCISystemDeleteById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCISystemDeleteById200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SystemFind
//
// Operation ID: system.find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCISystemService) SystemFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCISystemFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCISystemFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCISystemFind, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCISystemFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// SystemFindById
//
// Operation ID: system.findById
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
func (s *SCISystemService) SystemFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsSystem, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSystem
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCISystemFindById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSystem()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SystemGetSsids
//
// Operation ID: system.getSsids
//
// authenticate user with specific role
func (s *SCISystemService) SystemGetSsids(ctx context.Context, mutators ...RequestMutator) (SCISystemGetSsids200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCISystemGetSsids200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCISystemGetSsids, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCISystemGetSsids200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// SystemPrototypeUpdateAttributes
//
// Operation ID: system.prototype.updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsSystem
//
// Required Parameters:
// - id string
//		- required
func (s *SCISystemService) SystemPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsSystem, id string, mutators ...RequestMutator) (*SCIModelsSystem, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSystem
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCISystemPrototypeUpdateAttributes, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSystem()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
