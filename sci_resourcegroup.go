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

func newSCIResourceGroupBatchDelete200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIResourceGroupBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIResourceGroupBatchDelete200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIResourceGroupBatchDelete200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIResourceGroupBatchDelete200ResponseType() *SCIResourceGroupBatchDelete200ResponseType {
	m := new(SCIResourceGroupBatchDelete200ResponseType)
	return m
}

// SCIResourceGroupFind200ResponseType
//
// Definition: resourceGroup_find200ResponseType
type SCIResourceGroupFind200ResponseType []*SCIModelsResourceGroup

func MakeSCIResourceGroupFind200ResponseType() SCIResourceGroupFind200ResponseType {
	m := make(SCIResourceGroupFind200ResponseType, 0)
	return m
}

// ResourceGroupBatchDelete
//
// Operation ID: resourceGroup_batchDelete
//
// Delete multiple Resource Groups
//
// Form Data Parameters:
// - ids string
//		- required
func (s *SCIResourceGroupService) ResourceGroupBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIResourceGroupBatchDelete200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIResourceGroupBatchDelete200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIResourceGroupBatchDelete200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIResourceGroupBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIResourceGroupBatchDelete200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIResourceGroupBatchDelete200ResponseTypeAPIResponse), err
}

// ResourceGroupCreate
//
// Operation ID: resourceGroup_create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsResourceGroup
func (s *SCIResourceGroupService) ResourceGroupCreate(ctx context.Context, data *SCIModelsResourceGroup, mutators ...RequestMutator) (*SCIModelsResourceGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsResourceGroupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsResourceGroupAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIResourceGroupCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsResourceGroupAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsResourceGroupAPIResponse), err
}

// ResourceGroupFind
//
// Operation ID: resourceGroup_find
//
// Find all instances of the model matched by filter from the data source.
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCIResourceGroupService) ResourceGroupFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIResourceGroupFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// ResourceGroupFindById
//
// Operation ID: resourceGroup_findById
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
func (s *SCIResourceGroupService) ResourceGroupFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsResourceGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsResourceGroupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsResourceGroupAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIResourceGroupFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsResourceGroupAPIResponse), err
}

// ResourceGroupPrototypeUpdateAttributes
//
// Operation ID: resourceGroup_prototype_updateAttributes
//
// Update attributes for a model instance and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsResourceGroup
//
// Required Parameters:
// - id string
//		- required
func (s *SCIResourceGroupService) ResourceGroupPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsResourceGroup, id string, mutators ...RequestMutator) (*SCIModelsResourceGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsResourceGroupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsResourceGroupAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIResourceGroupPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsResourceGroupAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsResourceGroupAPIResponse), err
}
