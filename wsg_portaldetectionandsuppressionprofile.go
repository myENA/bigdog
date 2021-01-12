package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGPortalDetectionandSuppressionProfileService struct {
	apiClient *VSZClient
}

func NewWSGPortalDetectionandSuppressionProfileService(c *VSZClient) *WSGPortalDetectionandSuppressionProfileService {
	s := new(WSGPortalDetectionandSuppressionProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGPortalDetectionandSuppressionProfileService() *WSGPortalDetectionandSuppressionProfileService {
	return NewWSGPortalDetectionandSuppressionProfileService(ss.apiClient)
}

// AddRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to create portal detection and suppression profile.
//
// Operation ID: addRkszonesPortalDetectionProfilesByZoneId
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) AddRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesPortalDetectionProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesPortalDetectionProfilesById
//
// Use this API command to delete portal detection and suppression profile by profile's ID.
//
// Operation ID: deleteRkszonesPortalDetectionProfilesById
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesPortalDetectionProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to delete multiple portal detection and suppression profiles.
//
// Operation ID: deleteRkszonesPortalDetectionProfilesByZoneId
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesPortalDetectionProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesPortalDetectionProfilesById
//
// Use this API command to get portal detection and suppression profile by profile's ID.
//
// Operation ID: findRkszonesPortalDetectionProfilesById
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalDetectionProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalDetectionProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesPortalDetectionProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalDetectionProfileAPIResponse), err
}

// FindRkszonesPortalDetectionProfilesByQueryCriteria
//
// Query portal detection and suppression profile with specified filters.
//
// Operation ID: findRkszonesPortalDetectionProfilesByQueryCriteria
// Operation path: /rkszones/portalDetectionProfiles/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGPortalDetectionProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalDetectionProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindRkszonesPortalDetectionProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalDetectionProfileListAPIResponse), err
}

// FindRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to get portal detection and suppression profile list.
//
// Operation ID: findRkszonesPortalDetectionProfilesByZoneId
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalDetectionProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalDetectionProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesPortalDetectionProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalDetectionProfileListAPIResponse), err
}

// PartialUpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Operation ID: partialUpdateRkszonesPortalDetectionProfilesById
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) PartialUpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalDetectionProfilesById, true)
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

// UpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Operation ID: updateRkszonesPortalDetectionProfilesById
// Operation path: /rkszones/{zoneId}/portalDetectionProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) UpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesPortalDetectionProfilesById, true)
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
