package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL2AccessControlinDomainLevelService struct {
	apiClient *APIClient
}

func NewWSGL2AccessControlinDomainLevelService(c *APIClient) *WSGL2AccessControlinDomainLevelService {
	s := new(WSGL2AccessControlinDomainLevelService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlinDomainLevelService() *WSGL2AccessControlinDomainLevelService {
	return NewWSGL2AccessControlinDomainLevelService(ss.apiClient)
}

// AddL2AccessControls
//
// Use this API command to create a new L2 Access Control.
//
// Request Body:
//	 - body *WSGL2AccessControlCreateL2AccessControl
func (s *WSGL2AccessControlinDomainLevelService) AddL2AccessControls(ctx context.Context, body *WSGL2AccessControlCreateL2AccessControl) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddL2AccessControls, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteL2AccessControls
//
// Use this API command to delete a list of L2 Access Control.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL2AccessControlinDomainLevelService) DeleteL2AccessControls(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteL2AccessControls, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteL2AccessControlsById
//
// Use this API command to delete an L2 Access Control.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) DeleteL2AccessControlsById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteL2AccessControlsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindL2AccessControls
//
// Use this API command to retrieve a list of L2 Access Control.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControls(ctx context.Context, optionalParams map[string][]string) (*WSGL2AccessControlList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGL2AccessControlList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindL2AccessControls, true)
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
	resp = NewWSGL2AccessControlList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindL2AccessControlsById
//
// Use this API command to retrieve an L2 Access Control.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControlsById(ctx context.Context, id string) (*WSGL2AccessControl, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGL2AccessControl
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindL2AccessControlsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGL2AccessControl()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindL2AccessControlsByQueryCriteria
//
// Query L2 Access Control with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControlsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGL2AccessControlList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGL2AccessControlList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindL2AccessControlsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGL2AccessControlList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateL2AccessControlsById
//
// Use this API command to modify a specific L2 Access Control.
//
// Request Body:
//	 - body *WSGL2AccessControlModifyL2AccessControl
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) UpdateL2AccessControlsById(ctx context.Context, body *WSGL2AccessControlModifyL2AccessControl, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateL2AccessControlsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
