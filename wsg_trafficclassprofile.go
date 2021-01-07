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
// Operation ID: addRkszonesTrafficClassProfileByZoneId
//
// Use this API command to create a new Traffic Class Profile of a zone.
//
// Request Body:
//	 - body *WSGProfileCreateTrafficClassProfile
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesTrafficClassProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesTrafficClassProfileById
//
// Operation ID: deleteRkszonesTrafficClassProfileById
//
// Use this API command to delete a Traffic Class Profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesTrafficClassProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesTrafficClassProfileByZoneId
//
// Operation ID: deleteRkszonesTrafficClassProfileByZoneId
//
// Use this API command to bulk delete Traffic Class Profiles of a zone.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesTrafficClassProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindRkszonesTrafficClassProfileById
//
// Operation ID: findRkszonesTrafficClassProfileById
//
// Use this API command to retrieve a Traffic Class Profile of zone.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonTrafficClassProfileRefAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesTrafficClassProfileById, true)
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
// Operation ID: findRkszonesTrafficClassProfileByZoneId
//
// Use this API command to retrieve a list of Traffic Class Profile of a zone.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesTrafficClassProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
}

// FindServicesTrafficClassProfileByQueryCriteria
//
// Operation ID: findServicesTrafficClassProfileByQueryCriteria
//
// Retrieve a list of Traffic Class Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGTrafficClassProfileService) FindServicesTrafficClassProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileTrafficClassProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileTrafficClassProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesTrafficClassProfileByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileTrafficClassProfileListAPIResponse), err
}

// PartialUpdateRkszonesTrafficClassProfileById
//
// Operation ID: partialUpdateRkszonesTrafficClassProfileById
//
// Use this API command to modify Traffic Class Profile of a zone.
//
// Request Body:
//	 - body *WSGProfileCreateTrafficClassProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGTrafficClassProfileService) PartialUpdateRkszonesTrafficClassProfileById(ctx context.Context, body *WSGProfileCreateTrafficClassProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesTrafficClassProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
