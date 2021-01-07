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
// Operation ID: addPrecedence
//
// Use this API command to create Precedence Profile.
//
// Request Body:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddPrecedence, true)
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
// Operation ID: deletePrecedence
//
// Use this API command to Bulk Delete Precedence Profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkPrecedenceProfile
func (s *WSGPrecedenceProfileService) DeletePrecedence(ctx context.Context, body *WSGProfileDeleteBulkPrecedenceProfile, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeletePrecedence, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeletePrecedenceById
//
// Operation ID: deletePrecedenceById
//
// Use this API command to Delete Precedence Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) DeletePrecedenceById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeletePrecedenceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindPrecedence
//
// Operation ID: findPrecedence
//
// Use this API command to Get Precedence Profile list.
//
// Optional Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindPrecedence, true)
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
// Operation ID: findPrecedenceById
//
// Use this API command to Get Precedence Profile by profile's ID.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindPrecedenceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileCreatePrecedenceProfileAPIResponse), err
}

// FindPrecedenceByQueryCriteria
//
// Operation ID: findPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Request Body:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindPrecedenceByQueryCriteria, true)
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
// Operation ID: partialUpdatePrecedenceById
//
// Use this API command to Modify Precedence Profile by profile's ID.
//
// Request Body:
//	 - body *WSGProfileUpdatePrecedenceProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) PartialUpdatePrecedenceById(ctx context.Context, body *WSGProfileUpdatePrecedenceProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdatePrecedenceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
