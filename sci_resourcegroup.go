package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
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

type SCIResourceGroupBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIResourceGroupBatchDelete200ResponseType() *SCIResourceGroupBatchDelete200ResponseType {
	m := new(SCIResourceGroupBatchDelete200ResponseType)
	return m
}

type SCIResourceGroupFind200ResponseType []*SCIModelsResourceGroup

func MakeSCIResourceGroupFind200ResponseType() SCIResourceGroupFind200ResponseType {
	m := make(SCIResourceGroupFind200ResponseType, 0)
	return m
}

// ResourceGroupBatchDelete
//
// Delete multiple Resource Groups
//
// Request Body:
//	 - body string
func (s *SCIResourceGroupService) ResourceGroupBatchDelete(ctx context.Context, ids string, mutators ...RequestMutator) (*SCIResourceGroupBatchDelete200ResponseType, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIResourceGroupBatchDelete, false)
	if err = req.SetBody(ids); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIResourceGroupBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ResourceGroupCreate
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
	req = NewAPIRequest(http.MethodPost, RouteSCIResourceGroupCreate, false)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ResourceGroupFind
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
	req = NewAPIRequest(http.MethodGet, RouteSCIResourceGroupFind, false)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIResourceGroupFind200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// ResourceGroupFindById
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
	req = NewAPIRequest(http.MethodGet, RouteSCIResourceGroupFindById, false)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ResourceGroupPrototypeUpdateAttributes
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
	req = NewAPIRequest(http.MethodPut, RouteSCIResourceGroupPrototypeUpdateAttributes, false)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
