package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGPrecedenceProfileService struct {
	apiClient *VSZClient
}

func NewWSGPrecedenceProfileService(c *VSZClient) *WSGPrecedenceProfileService {
	s := new(WSGPrecedenceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGPrecedenceProfileService() *WSGPrecedenceProfileService {
	return NewWSGPrecedenceProfileService(ss.apiClient)
}

// AddPrecedence
//
// Use this API command to create Precedence Profile.
//
// Operation ID: addPrecedence
// Operation path: /precedence
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreatePrecedenceProfile
func (s *WSGPrecedenceProfileService) AddPrecedence(ctx context.Context, body *WSGProfileCreatePrecedenceProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddPrecedence, true)
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

// DeletePrecedence
//
// Use this API command to Bulk Delete Precedence Profile.
//
// Operation ID: deletePrecedence
// Operation path: /precedence
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileDeleteBulkPrecedenceProfile
func (s *WSGPrecedenceProfileService) DeletePrecedence(ctx context.Context, body *WSGProfileDeleteBulkPrecedenceProfile, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeletePrecedence, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeletePrecedenceById
//
// Use this API command to Delete Precedence Profile by profile's ID.
//
// Operation ID: deletePrecedenceById
// Operation path: /precedence/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) DeletePrecedenceById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeletePrecedenceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindPrecedence
//
// Use this API command to Get Precedence Profile list.
//
// Operation ID: findPrecedence
// Operation path: /precedence
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGPrecedenceProfileService) FindPrecedence(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfilePrecedenceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfilePrecedenceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfilePrecedenceListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindPrecedence, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfilePrecedenceListAPIResponse), err
}

// FindPrecedenceById
//
// Use this API command to Get Precedence Profile by profile's ID.
//
// Operation ID: findPrecedenceById
// Operation path: /precedence/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) FindPrecedenceById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileCreatePrecedenceProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileCreatePrecedenceProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileCreatePrecedenceProfileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindPrecedenceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileCreatePrecedenceProfileAPIResponse), err
}

// FindPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Operation ID: findPrecedenceByQueryCriteria
// Operation path: /precedence/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPrecedenceProfileService) FindPrecedenceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfilePrecedenceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfilePrecedenceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfilePrecedenceListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindPrecedenceByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfilePrecedenceListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfilePrecedenceListAPIResponse), err
}

// PartialUpdatePrecedenceById
//
// Use this API command to Modify Precedence Profile by profile's ID.
//
// Operation ID: partialUpdatePrecedenceById
// Operation path: /precedence/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGProfileUpdatePrecedenceProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) PartialUpdatePrecedenceById(ctx context.Context, body *WSGProfileUpdatePrecedenceProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdatePrecedenceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
