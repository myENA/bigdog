package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDiffServService struct {
	apiClient *VSZClient
}

func NewWSGDiffServService(c *VSZClient) *WSGDiffServService {
	s := new(WSGDiffServService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDiffServService() *WSGDiffServService {
	return NewWSGDiffServService(ss.apiClient)
}

// AddRkszonesDiffservByZoneId
//
// Operation ID: addRkszonesDiffservByZoneId
//
// Use this API command to create DiffServ profile.
//
// Request Body:
//	 - body *WSGZoneCreateDiffServProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDiffServService) AddRkszonesDiffservByZoneId(ctx context.Context, body *WSGZoneCreateDiffServProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDiffservByZoneId, true)
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

// DeleteRkszonesDiffservById
//
// Operation ID: deleteRkszonesDiffservById
//
// Use this API command to delete DiffServ profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDiffServService) DeleteRkszonesDiffservById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDiffservById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindRkszonesDiffservById
//
// Operation ID: findRkszonesDiffservById
//
// Use this API command to retrieve DiffServ profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGZoneDiffServConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneDiffServConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGZoneDiffServConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDiffservById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneDiffServConfigurationAPIResponse), err
}

// FindRkszonesDiffservByZoneId
//
// Operation ID: findRkszonesDiffservByZoneId
//
// Use this API command to retrieve a list of DiffServ profile.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGZoneDiffServListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneDiffServListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGZoneDiffServListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDiffservByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneDiffServListAPIResponse), err
}

// FindServicesDscpProfileByQueryCriteria
//
// Operation ID: findServicesDscpProfileByQueryCriteria
//
// Query DSCP Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDiffServService) FindServicesDscpProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesDscpProfileByQueryCriteria, true)
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

// PartialUpdateRkszonesDiffservById
//
// Operation ID: partialUpdateRkszonesDiffservById
//
// Use this API command to modify the configuration of DiffServ profile.
//
// Request Body:
//	 - body *WSGZoneModifyDiffServProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDiffServService) PartialUpdateRkszonesDiffservById(ctx context.Context, body *WSGZoneModifyDiffServProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesDiffservById, true)
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
