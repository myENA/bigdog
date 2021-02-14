package bigdog

// API Version: 1.0.0

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

type SCISystemFind200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCISystemFind200ResponseType
}

func newSCISystemFind200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCISystemFind200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCISystemFind200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCISystemFind200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCISystemFind200ResponseType() SCISystemFind200ResponseType {
	m := make(SCISystemFind200ResponseType, 0)
	return m
}

// SCISystemGetSsids200ResponseType
//
// Definition: system_getSsids200ResponseType
type SCISystemGetSsids200ResponseType []interface{}

type SCISystemGetSsids200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCISystemGetSsids200ResponseType
}

func newSCISystemGetSsids200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCISystemGetSsids200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCISystemGetSsids200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(SCISystemGetSsids200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeSCISystemGetSsids200ResponseType() SCISystemGetSsids200ResponseType {
	m := make(SCISystemGetSsids200ResponseType, 0)
	return m
}

// SystemCreate
//
// Create a new instance of the model and persist it into the data source.
//
// Operation ID: system_create
// Operation path: /systems
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsSystem
func (s *SCISystemService) SystemCreate(ctx context.Context, data *SCIModelsSystem, mutators ...RequestMutator) (*SCIModelsSystemAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsSystemAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCISystemCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSystemAPIResponse), err
}

// SystemDeleteById
//
// Delete a model instance by id from the data source.
//
// Operation ID: system_deleteById
// Operation path: /systems/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SCISystemService) SystemDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodDelete, RouteSCISystemDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// SystemFind
//
// Find all instances of the model matched by filter from the data source.
//
// Operation ID: system_find
// Operation path: /systems
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCISystemService) SystemFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCISystemFind200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCISystemFind200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCISystemFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCISystemFind200ResponseTypeAPIResponse), err
}

// SystemFindById
//
// Find a model instance by id from the data source.
//
// Operation ID: system_findById
// Operation path: /systems/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCISystemService) SystemFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsSystemAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsSystemAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCISystemFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSystemAPIResponse), err
}

// SystemGetSsids
//
// authenticate user with specific role
//
// Operation ID: system_getSsids
// Operation path: /systems/getSsids
// Success code: 200 (OK)
func (s *SCISystemService) SystemGetSsids(ctx context.Context, mutators ...RequestMutator) (*SCISystemGetSsids200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCISystemGetSsids200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCISystemGetSsids, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCISystemGetSsids200ResponseTypeAPIResponse), err
}

// SystemPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Operation ID: system_prototype_updateAttributes
// Operation path: /systems/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsSystem
//
// Required parameters:
// - id string
//		- required
func (s *SCISystemService) SystemPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsSystem, id string, mutators ...RequestMutator) (*SCIModelsSystemAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsSystemAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCISystemPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSystemAPIResponse), err
}
