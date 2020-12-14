package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
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
func (s *SCIResourceGroupService) ResourceGroupBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIResourceGroupBatchDelete200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIResourceGroupBatchDelete200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIResourceGroupBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIResourceGroupBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ResourceGroupCreate
//
// Operation ID: resourceGroup_create
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsResourceGroup
func (s *SCIResourceGroupService) ResourceGroupCreate(ctx context.Context, data *SCIModelsResourceGroup, mutators ...RequestMutator) (*SCIModelsResourceGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsResourceGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIResourceGroupCreate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIResourceGroupService) ResourceGroupFind(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (SCIResourceGroupFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIResourceGroupFind200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIResourceGroupFind, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("filter", vv)
		}
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIResourceGroupFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SCIResourceGroupService) ResourceGroupFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsResourceGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsResourceGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCIResourceGroupFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("filter", vv)
		}
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SCIResourceGroupService) ResourceGroupPrototypeUpdateAttributes(ctx context.Context, data *SCIModelsResourceGroup, id string, mutators ...RequestMutator) (*SCIModelsResourceGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsResourceGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIResourceGroupPrototypeUpdateAttributes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
