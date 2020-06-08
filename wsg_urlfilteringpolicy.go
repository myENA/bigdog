package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGUrlfilteringpolicyService struct {
	apiClient *VSZClient
}

func NewWSGUrlfilteringpolicyService(c *VSZClient) *WSGUrlfilteringpolicyService {
	s := new(WSGUrlfilteringpolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGUrlfilteringpolicyService() *WSGUrlfilteringpolicyService {
	return NewWSGUrlfilteringpolicyService(ss.apiClient)
}

// AddUrlFilteringUrlFilteringPolicy
//
// Use this API command to create a URL Filtering policy.
//
// Request Body:
//	 - body *WSGURLFilteringCreateUrlFilteringPolicy
func (s *WSGUrlfilteringpolicyService) AddUrlFilteringUrlFilteringPolicy(ctx context.Context, body *WSGURLFilteringCreateUrlFilteringPolicy) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUrlFilteringUrlFilteringPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteUrlFilteringUrlFilteringPolicy
//
// Use this API command to delete bulk URL Filtering policies.
//
// Request Body:
//	 - body *WSGURLFilteringDeleteBulk
func (s *WSGUrlfilteringpolicyService) DeleteUrlFilteringUrlFilteringPolicy(ctx context.Context, body *WSGURLFilteringDeleteBulk) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUrlFilteringUrlFilteringPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteUrlFilteringUrlFilteringPolicyById
//
// Use this API command to delete a URL Filtering policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGUrlfilteringpolicyService) DeleteUrlFilteringUrlFilteringPolicyById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUrlFilteringUrlFilteringPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindUrlFilteringBlockCategories
//
// Use this API command to retrieve the block categories of URL Filtering.
func (s *WSGUrlfilteringpolicyService) FindUrlFilteringBlockCategories(ctx context.Context) (*WSGURLFilteringBlockCategoriesList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGURLFilteringBlockCategoriesList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUrlFilteringBlockCategories, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGURLFilteringBlockCategoriesList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUrlFilteringByQueryCriteria
//
// Use this API command to retrieve a list of URL Filtering policies by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGUrlfilteringpolicyService) FindUrlFilteringByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGURLFilteringPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGURLFilteringPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUrlFilteringByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGURLFilteringPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUrlFilteringUrlFilteringPolicy
//
// Use this API command to retrieve list of URL Filtering policies.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGUrlfilteringpolicyService) FindUrlFilteringUrlFilteringPolicy(ctx context.Context, optionalParams map[string][]string) (*WSGURLFilteringPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGURLFilteringPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUrlFilteringUrlFilteringPolicy, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGURLFilteringPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUrlFilteringUrlFilteringPolicyById
//
// Use this API command to retrieve an URL Filtering policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGUrlfilteringpolicyService) FindUrlFilteringUrlFilteringPolicyById(ctx context.Context, id string) (*WSGURLFilteringPolicy, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGURLFilteringPolicy
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUrlFilteringUrlFilteringPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGURLFilteringPolicy()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateUrlFilteringUrlFilteringPolicyById
//
// Use this API command to patch a URL Filtering policy.
//
// Request Body:
//	 - body *WSGURLFilteringModifyUrlFilteringPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGUrlfilteringpolicyService) PartialUpdateUrlFilteringUrlFilteringPolicyById(ctx context.Context, body *WSGURLFilteringModifyUrlFilteringPolicy, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUrlFilteringUrlFilteringPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
