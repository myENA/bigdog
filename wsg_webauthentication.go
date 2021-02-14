package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGWebAuthenticationService struct {
	apiClient *VSZClient
}

func NewWSGWebAuthenticationService(c *VSZClient) *WSGWebAuthenticationService {
	s := new(WSGWebAuthenticationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWebAuthenticationService() *WSGWebAuthenticationService {
	return NewWSGWebAuthenticationService(ss.apiClient)
}

// AddRkszonesPortalsWebauthByZoneId
//
// Use this API command to create a new web authentication of a zone.
//
// Operation ID: addRkszonesPortalsWebauthByZoneId
// Operation path: /rkszones/{zoneId}/portals/webauth
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGPortalServiceCreateWebAuthentication
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) AddRkszonesPortalsWebauthByZoneId(ctx context.Context, body *WSGPortalServiceCreateWebAuthentication, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesPortalsWebauthByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesPortalsWebauthById
//
// Use this API command to delete an web authentication of a zone.
//
// Operation ID: deleteRkszonesPortalsWebauthById
// Operation path: /rkszones/{zoneId}/portals/webauth/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesPortalsWebauthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesPortalsWebauthRedirectById
//
// Use this API command to set redirect to the URL that user intends to visit on web authentication of a zone.
//
// Operation ID: deleteRkszonesPortalsWebauthRedirectById
// Operation path: /rkszones/{zoneId}/portals/webauth/{id}/redirect
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthRedirectById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesPortalsWebauthRedirectById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesPortalsWebauthById
//
// Use this API command to retrieve a web authentication of a zone.
//
// Operation ID: findRkszonesPortalsWebauthById
// Operation path: /rkszones/{zoneId}/portals/webauth/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceWebAuthenticationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceWebAuthenticationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesPortalsWebauthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceWebAuthenticationAPIResponse), err
}

// FindRkszonesPortalsWebauthByZoneId
//
// Use this API command to retrieve a list of web authentication of a zone.
//
// Operation ID: findRkszonesPortalsWebauthByZoneId
// Operation path: /rkszones/{zoneId}/portals/webauth
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesPortalsWebauthByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceListAPIResponse), err
}

// FindServicesWebAuthenticationByQueryCriteria
//
// Query Web Authentications with specified filters.
//
// Operation ID: findServicesWebAuthenticationByQueryCriteria
// Operation path: /query/services/webAuthentication
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWebAuthenticationService) FindServicesWebAuthenticationByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesWebAuthenticationByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateRkszonesPortalsWebauthById
//
// Use this API command to modify the configuration on web authentication of a zone.
//
// Operation ID: partialUpdateRkszonesPortalsWebauthById
// Operation path: /rkszones/{zoneId}/portals/webauth/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGPortalServiceModifyWebAuthentication
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) PartialUpdateRkszonesPortalsWebauthById(ctx context.Context, body *WSGPortalServiceModifyWebAuthentication, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsWebauthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
