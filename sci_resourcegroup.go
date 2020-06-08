package ruckus

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIResourcegroupService struct {
	apiClient *SCIClient
}

func NewSCIResourcegroupService(c *SCIClient) *SCIResourcegroupService {
	s := new(SCIResourcegroupService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIResourcegroupService() *SCIResourcegroupService {
	return NewSCIResourcegroupService(ss.apiClient)
}

type SCIResourcegroupBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIResourcegroupBatchDelete200ResponseType() *SCIResourcegroupBatchDelete200ResponseType {
	m := new(SCIResourcegroupBatchDelete200ResponseType)
	return m
}

type SCIResourcegroupFind200ResponseType []*SCIModelsResourceGroup

func MakeSCIResourcegroupFind200ResponseType() SCIResourcegroupFind200ResponseType {
	m := make(SCIResourcegroupFind200ResponseType, 0)
	return m
}

// ResourceGroupBatchDelete
//
// Delete multiple Resource Groups
//
// Request Body:
//	 - body string
func (s *SCIResourcegroupService) ResourceGroupBatchDelete(ctx context.Context, body string) (*SCIResourcegroupBatchDelete200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIResourcegroupBatchDelete200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIResourceGroupBatchDelete, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIResourcegroupBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ResourceGroupCreate
//
// Create a new instance of the model and persist it into the data source.
//
// Request Body:
//	 - body *SCIModelsResourceGroup
func (s *SCIResourcegroupService) ResourceGroupCreate(ctx context.Context, body *SCIModelsResourceGroup) (*SCIModelsResourceGroup, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SCIResourcegroupService) ResourceGroupFind(ctx context.Context, optionalParams map[string][]string) (SCIResourcegroupFind200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIResourcegroupFind200ResponseType
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSCIResourcegroupFind200ResponseType()
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
func (s *SCIResourcegroupService) ResourceGroupFindById(ctx context.Context, id string, optionalParams map[string][]string) (*SCIModelsResourceGroup, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *SCIResourcegroupService) ResourceGroupPrototypeUpdateAttributes(ctx context.Context, body *SCIModelsResourceGroup, id string) (*SCIModelsResourceGroup, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIModelsResourceGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
