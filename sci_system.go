package bigdog

// API Version: 1.0.0

import (
	"context"
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

// SCISystemFind200ResponseType
//
// Definition: system_find200ResponseType
type SCISystemFind200ResponseType []*SCIModelsSystem

func MakeSCISystemFind200ResponseType() SCISystemFind200ResponseType {
	m := make(SCISystemFind200ResponseType, 0)
	return m
}

// SCISystemGetSsids200ResponseType
//
// Definition: system_getSsids200ResponseType
type SCISystemGetSsids200ResponseType []interface{}

func MakeSCISystemGetSsids200ResponseType() SCISystemGetSsids200ResponseType {
	m := make(SCISystemGetSsids200ResponseType, 0)
	return m
}

// SystemCreate
//
// Operation ID: system_create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsSystem
func (s *SCISystemService) SystemCreate(ctx context.Context, data *SCIModelsSystem, mutators ...RequestMutator) (*SCIModelsSystemAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsSystemAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsSystemAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCISystemCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsSystemAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSystemAPIResponse), err
}

// SystemDeleteById
//
// Operation ID: system_deleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCISystemService) SystemDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSCISystemDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// SystemFind
//
// Operation ID: system_find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCISystemService) SystemFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCISystemFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// SystemFindById
//
// Operation ID: system_findById
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
func (s *SCISystemService) SystemFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsSystemAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsSystemAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsSystemAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCISystemFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSystemAPIResponse), err
}

// SystemGetSsids
//
// Operation ID: system_getSsids
//
// authenticate user with specific role
func (s *SCISystemService) SystemGetSsids(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCISystemGetSsids, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// SystemPrototypeUpdateAttributes
//
// Operation ID: system_prototype_updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsSystem
//
// Required Parameters:
// - id string
//		- required
func (s *SCISystemService) SystemPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsSystem, id string, mutators ...RequestMutator) (*SCIModelsSystemAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsSystemAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsSystemAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCISystemPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsSystemAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSystemAPIResponse), err
}
