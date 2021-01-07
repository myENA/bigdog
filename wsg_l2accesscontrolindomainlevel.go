package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGL2AccessControlinDomainLevelService struct {
	apiClient *VSZClient
}

func NewWSGL2AccessControlinDomainLevelService(c *VSZClient) *WSGL2AccessControlinDomainLevelService {
	s := new(WSGL2AccessControlinDomainLevelService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlinDomainLevelService() *WSGL2AccessControlinDomainLevelService {
	return NewWSGL2AccessControlinDomainLevelService(ss.apiClient)
}

// AddL2AccessControls
//
// Operation ID: addL2AccessControls
//
// Use this API command to create a new L2 Access Control.
//
// Request Body:
//	 - body *WSGL2AccessControlCreateL2AccessControl
func (s *WSGL2AccessControlinDomainLevelService) AddL2AccessControls(ctx context.Context, body *WSGL2AccessControlCreateL2AccessControl, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddL2AccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteL2AccessControls
//
// Operation ID: deleteL2AccessControls
//
// Use this API command to delete a list of L2 Access Control.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL2AccessControlinDomainLevelService) DeleteL2AccessControls(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteL2AccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteL2AccessControlsById
//
// Operation ID: deleteL2AccessControlsById
//
// Use this API command to delete an L2 Access Control.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) DeleteL2AccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteL2AccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindL2AccessControls
//
// Operation ID: findL2AccessControls
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
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControls(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGL2AccessControlListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindL2AccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
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
	resp = NewWSGL2AccessControlList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindL2AccessControlsById
//
// Operation ID: findL2AccessControlsById
//
// Use this API command to retrieve an L2 Access Control.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGL2AccessControlAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindL2AccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGL2AccessControl()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindL2AccessControlsByQueryCriteria
//
// Operation ID: findL2AccessControlsByQueryCriteria
//
// Query L2 Access Control with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControlsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGL2AccessControlListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindL2AccessControlsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGL2AccessControlList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateL2AccessControlsById
//
// Operation ID: updateL2AccessControlsById
//
// Use this API command to modify a specific L2 Access Control.
//
// Request Body:
//	 - body *WSGL2AccessControlModifyL2AccessControl
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) UpdateL2AccessControlsById(ctx context.Context, body *WSGL2AccessControlModifyL2AccessControl, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateL2AccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
