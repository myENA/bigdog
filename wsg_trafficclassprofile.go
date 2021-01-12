package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGTrafficClassProfileService struct {
	apiClient *VSZClient
}

func NewWSGTrafficClassProfileService(c *VSZClient) *WSGTrafficClassProfileService {
	s := new(WSGTrafficClassProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTrafficClassProfileService() *WSGTrafficClassProfileService {
	return NewWSGTrafficClassProfileService(ss.apiClient)
}

// AddRkszonesTrafficClassProfileByZoneId
//
// Use this API command to create a new Traffic Class Profile of a zone.
//
// Operation ID: addRkszonesTrafficClassProfileByZoneId
// Operation path: /rkszones/{zoneId}/trafficClassProfile
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateTrafficClassProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) AddRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *WSGProfileCreateTrafficClassProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesTrafficClassProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesTrafficClassProfileById
//
// Use this API command to delete a Traffic Class Profile of a zone.
//
// Operation ID: deleteRkszonesTrafficClassProfileById
// Operation path: /rkszones/{zoneId}/trafficClassProfile/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesTrafficClassProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesTrafficClassProfileByZoneId
//
// Use this API command to bulk delete Traffic Class Profiles of a zone.
//
// Operation ID: deleteRkszonesTrafficClassProfileByZoneId
// Operation path: /rkszones/{zoneId}/trafficClassProfile
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesTrafficClassProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesTrafficClassProfileById
//
// Use this API command to retrieve a Traffic Class Profile of zone.
//
// Operation ID: findRkszonesTrafficClassProfileById
// Operation path: /rkszones/{zoneId}/trafficClassProfile/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGCommonTrafficClassProfileRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonTrafficClassProfileRefAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesTrafficClassProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonTrafficClassProfileRefAPIResponse), err
}

// FindRkszonesTrafficClassProfileByZoneId
//
// Use this API command to retrieve a list of Traffic Class Profile of a zone.
//
// Operation ID: findRkszonesTrafficClassProfileByZoneId
// Operation path: /rkszones/{zoneId}/trafficClassProfile
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileTrafficClassProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileTrafficClassProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesTrafficClassProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
}

// FindServicesTrafficClassProfileByQueryCriteria
//
// Retrieve a list of Traffic Class Profile.
//
// Operation ID: findServicesTrafficClassProfileByQueryCriteria
// Operation path: /query/services/trafficClassProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGTrafficClassProfileService) FindServicesTrafficClassProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileTrafficClassProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileTrafficClassProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesTrafficClassProfileByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
}

// PartialUpdateRkszonesTrafficClassProfileById
//
// Use this API command to modify Traffic Class Profile of a zone.
//
// Operation ID: partialUpdateRkszonesTrafficClassProfileById
// Operation path: /rkszones/{zoneId}/trafficClassProfile/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileCreateTrafficClassProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) PartialUpdateRkszonesTrafficClassProfileById(ctx context.Context, body *WSGProfileCreateTrafficClassProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesTrafficClassProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
