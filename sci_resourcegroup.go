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

type SCIResourceGroupService struct {
	apiClient *SCIClient
}

func NewSCIResourceGroupService(c *SCIClient) *SCIResourceGroupService {
	s := new(SCIResourceGroupService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIResourceGroupService() *SCIResourceGroupService {
	return NewSCIResourceGroupService(ss.apiClient)
}

// SCIResourceGroupBatchDelete200ResponseType
//
// Definition: resourceGroup_batchDelete200ResponseType
type SCIResourceGroupBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIResourceGroupBatchDelete200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIResourceGroupBatchDelete200ResponseType
}

func newSCIResourceGroupBatchDelete200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIResourceGroupBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIResourceGroupBatchDelete200ResponseTypeAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SCIResourceGroupBatchDelete200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSCIResourceGroupBatchDelete200ResponseType() *SCIResourceGroupBatchDelete200ResponseType {
	m := new(SCIResourceGroupBatchDelete200ResponseType)
	return m
}

// SCIResourceGroupFind200ResponseType
//
// Definition: resourceGroup_find200ResponseType
type SCIResourceGroupFind200ResponseType []*SCIModelsResourceGroup

type SCIResourceGroupFind200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data SCIResourceGroupFind200ResponseType
}

func newSCIResourceGroupFind200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIResourceGroupFind200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIResourceGroupFind200ResponseTypeAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := make(SCIResourceGroupFind200ResponseType, 0)
	if err := r.doHydrate(&data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func MakeSCIResourceGroupFind200ResponseType() SCIResourceGroupFind200ResponseType {
	m := make(SCIResourceGroupFind200ResponseType, 0)
	return m
}

// ResourceGroupBatchDelete
//
// Delete multiple Resource Groups
//
// Operation ID: resourceGroup_batchDelete
// Operation path: /resourceGroups/batchDelete
// Success code: 200 (OK)
//
// Form data parameters:
// - ids string
//		- required
func (s *SCIResourceGroupService) ResourceGroupBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIResourceGroupBatchDelete200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIResourceGroupBatchDelete200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIResourceGroupBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIResourceGroupBatchDelete200ResponseTypeAPIResponse), err
}

// ResourceGroupCreate
//
// Create a new instance of the model and persist it into the data source.
//
// Operation ID: resourceGroup_create
// Operation path: /resourceGroups
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsResourceGroup
func (s *SCIResourceGroupService) ResourceGroupCreate(ctx context.Context, data *SCIModelsResourceGroup, mutators ...RequestMutator) (*SCIModelsResourceGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsResourceGroupAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIResourceGroupCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsResourceGroupAPIResponse), err
}

// ResourceGroupFind
//
// Find all instances of the model matched by filter from the data source.
//
// Operation ID: resourceGroup_find
// Operation path: /resourceGroups
// Success code: 200 (OK)
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIResourceGroupService) ResourceGroupFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIResourceGroupFind200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIResourceGroupFind200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIResourceGroupFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIResourceGroupFind200ResponseTypeAPIResponse), err
}

// ResourceGroupFindById
//
// Find a model instance by id from the data source.
//
// Operation ID: resourceGroup_findById
// Operation path: /resourceGroups/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCIResourceGroupService) ResourceGroupFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsResourceGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsResourceGroupAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIResourceGroupFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsResourceGroupAPIResponse), err
}

// ResourceGroupPrototypeUpdateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Operation ID: resourceGroup_prototype_updateAttributes
// Operation path: /resourceGroups/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsResourceGroup
//
// Required parameters:
// - id string
//		- required
func (s *SCIResourceGroupService) ResourceGroupPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsResourceGroup, id string, mutators ...RequestMutator) (*SCIModelsResourceGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIModelsResourceGroupAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIResourceGroupPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(data)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsResourceGroupAPIResponse), err
}
