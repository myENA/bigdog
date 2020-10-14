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
func (s *WSGPrecedenceProfileService) AddPrecedence(ctx context.Context, body *WSGProfileCreatePrecedenceProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddPrecedence, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeletePrecedence
//
// Operation ID: deletePrecedence
//
// Use this API command to Bulk Delete Precedence Profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkPrecedenceProfile
func (s *WSGPrecedenceProfileService) DeletePrecedence(ctx context.Context, body *WSGProfileDeleteBulkPrecedenceProfile, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeletePrecedence, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGPrecedenceProfileService) DeletePrecedenceById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeletePrecedenceById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGPrecedenceProfileService) FindPrecedence(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfilePrecedenceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfilePrecedenceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPrecedence, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfilePrecedenceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGPrecedenceProfileService) FindPrecedenceById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileCreatePrecedenceProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileCreatePrecedenceProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPrecedenceById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileCreatePrecedenceProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPrecedenceByQueryCriteria
//
// Operation ID: findPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPrecedenceProfileService) FindPrecedenceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfilePrecedenceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfilePrecedenceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindPrecedenceByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfilePrecedenceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGPrecedenceProfileService) PartialUpdatePrecedenceById(ctx context.Context, body *WSGProfileUpdatePrecedenceProfile, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePrecedenceById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
