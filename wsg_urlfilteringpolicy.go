package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGURLFilteringPolicyService struct {
	apiClient *VSZClient
}

func NewWSGURLFilteringPolicyService(c *VSZClient) *WSGURLFilteringPolicyService {
	s := new(WSGURLFilteringPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGURLFilteringPolicyService() *WSGURLFilteringPolicyService {
	return NewWSGURLFilteringPolicyService(ss.apiClient)
}

// AddUrlFilteringUrlFilteringPolicy
//
// Use this API command to create a URL Filtering policy.
//
// Operation ID: addUrlFilteringUrlFilteringPolicy
// Operation path: /urlFiltering/urlFilteringPolicy
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGURLFilteringCreateUrlFilteringPolicy
func (s *WSGURLFilteringPolicyService) AddUrlFilteringUrlFilteringPolicy(ctx context.Context, body *WSGURLFilteringCreateUrlFilteringPolicy, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddUrlFilteringUrlFilteringPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteUrlFilteringUrlFilteringPolicy
//
// Use this API command to delete bulk URL Filtering policies.
//
// Operation ID: deleteUrlFilteringUrlFilteringPolicy
// Operation path: /urlFiltering/urlFilteringPolicy
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGURLFilteringDeleteBulk
func (s *WSGURLFilteringPolicyService) DeleteUrlFilteringUrlFilteringPolicy(ctx context.Context, body *WSGURLFilteringDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUrlFilteringUrlFilteringPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteUrlFilteringUrlFilteringPolicyById
//
// Use this API command to delete a URL Filtering policy.
//
// Operation ID: deleteUrlFilteringUrlFilteringPolicyById
// Operation path: /urlFiltering/urlFilteringPolicy/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGURLFilteringPolicyService) DeleteUrlFilteringUrlFilteringPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUrlFilteringUrlFilteringPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindUrlFilteringBlockCategories
//
// Use this API command to retrieve the block categories of URL Filtering.
//
// Operation ID: findUrlFilteringBlockCategories
// Operation path: /urlFiltering/blockCategories
// Success code: 200 (OK)
func (s *WSGURLFilteringPolicyService) FindUrlFilteringBlockCategories(ctx context.Context, mutators ...RequestMutator) (*WSGURLFilteringBlockCategoriesListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGURLFilteringBlockCategoriesListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGURLFilteringBlockCategoriesListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUrlFilteringBlockCategories, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGURLFilteringBlockCategoriesListAPIResponse), err
}

// FindUrlFilteringByQueryCriteria
//
// Use this API command to retrieve a list of URL Filtering policies by query criteria.
//
// Operation ID: findUrlFilteringByQueryCriteria
// Operation path: /urlFiltering/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGURLFilteringPolicyService) FindUrlFilteringByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGURLFilteringPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGURLFilteringPolicyListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGURLFilteringPolicyListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindUrlFilteringByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGURLFilteringPolicyListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGURLFilteringPolicyListAPIResponse), err
}

// FindUrlFilteringUrlFilteringPolicy
//
// Use this API command to retrieve list of URL Filtering policies.
//
// Operation ID: findUrlFilteringUrlFilteringPolicy
// Operation path: /urlFiltering/urlFilteringPolicy
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicy(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGURLFilteringPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGURLFilteringPolicyListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGURLFilteringPolicyListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUrlFilteringUrlFilteringPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGURLFilteringPolicyListAPIResponse), err
}

// FindUrlFilteringUrlFilteringPolicyById
//
// Use this API command to retrieve an URL Filtering policy.
//
// Operation ID: findUrlFilteringUrlFilteringPolicyById
// Operation path: /urlFiltering/urlFilteringPolicy/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGURLFilteringPolicyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGURLFilteringPolicyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGURLFilteringPolicyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUrlFilteringUrlFilteringPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGURLFilteringPolicyAPIResponse), err
}

// PartialUpdateUrlFilteringUrlFilteringPolicyById
//
// Use this API command to patch a URL Filtering policy.
//
// Operation ID: partialUpdateUrlFilteringUrlFilteringPolicyById
// Operation path: /urlFiltering/urlFilteringPolicy/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGURLFilteringModifyUrlFilteringPolicy
//
// Required parameters:
// - id string
//		- required
func (s *WSGURLFilteringPolicyService) PartialUpdateUrlFilteringUrlFilteringPolicyById(ctx context.Context, body *WSGURLFilteringModifyUrlFilteringPolicy, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateUrlFilteringUrlFilteringPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
