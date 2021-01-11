package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGSplitTunnelProfileService struct {
	apiClient *VSZClient
}

func NewWSGSplitTunnelProfileService(c *VSZClient) *WSGSplitTunnelProfileService {
	s := new(WSGSplitTunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSplitTunnelProfileService() *WSGSplitTunnelProfileService {
	return NewWSGSplitTunnelProfileService(ss.apiClient)
}

// AddRkszonesSplitTunnelProfilesByZoneId
//
// Create a split tunnel profile.
//
// Operation ID: addRkszonesSplitTunnelProfilesByZoneId
// Operation path: /rkszones/{zoneId}/splitTunnelProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSplitTunnelCreateSplitTunnelProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) AddRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesSplitTunnelProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesSplitTunnelProfiles
//
// Use this API command to delete bulk split tunnel profiles.
//
// Operation ID: deleteRkszonesSplitTunnelProfiles
// Operation path: /rkszones/splitTunnelProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesSplitTunnelProfiles, true)
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

// DeleteRkszonesSplitTunnelProfilesById
//
// Use this API command to delete a split tunnel profile by ID.
//
// Operation ID: deleteRkszonesSplitTunnelProfilesById
// Operation path: /rkszones/{zoneId}/splitTunnelProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesSplitTunnelProfilesById
//
// Get a split tunnel profile by ID.
//
// Operation ID: findRkszonesSplitTunnelProfilesById
// Operation path: /rkszones/{zoneId}/splitTunnelProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGSplitTunnelProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSplitTunnelProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSplitTunnelProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSplitTunnelProfileAPIResponse), err
}

// FindRkszonesSplitTunnelProfilesByQueryCriteria
//
// Use this API command to retrieve a list of split tunnel profile by query criteria.
//
// Operation ID: findRkszonesSplitTunnelProfilesByQueryCriteria
// Operation path: /rkszones/splitTunnelProfiles/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSplitTunnelProfileQueryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSplitTunnelProfileQueryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSplitTunnelProfileQueryAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRkszonesSplitTunnelProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSplitTunnelProfileQueryAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSplitTunnelProfileQueryAPIResponse), err
}

// FindRkszonesSplitTunnelProfilesByZoneId
//
// Get a ID list of split tunnel profile in this Zone.
//
// Operation ID: findRkszonesSplitTunnelProfilesByZoneId
// Operation path: /rkszones/{zoneId}/splitTunnelProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGSplitTunnelProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSplitTunnelProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSplitTunnelProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesSplitTunnelProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSplitTunnelProfileListAPIResponse), err
}

// PartialUpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify a split tunnel profile.
//
// Operation ID: partialUpdateRkszonesSplitTunnelProfilesById
// Operation path: /rkszones/{zoneId}/splitTunnelProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSplitTunnelModifySplitTunnelProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) PartialUpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelModifySplitTunnelProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify entire information of a split tunnel profile.
//
// Operation ID: updateRkszonesSplitTunnelProfilesById
// Operation path: /rkszones/{zoneId}/splitTunnelProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSplitTunnelCreateSplitTunnelProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) UpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
