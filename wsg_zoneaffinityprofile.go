package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGZoneAffinityProfileService struct {
	apiClient *VSZClient
}

func NewWSGZoneAffinityProfileService(c *VSZClient) *WSGZoneAffinityProfileService {
	s := new(WSGZoneAffinityProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZoneAffinityProfileService() *WSGZoneAffinityProfileService {
	return NewWSGZoneAffinityProfileService(ss.apiClient)
}

// AddProfilesZoneAffinity
//
// Use this API command to create zone affinity profile.
//
// Operation ID: addProfilesZoneAffinity
// Operation path: /profiles/zoneAffinity
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateZoneAffinityProfile
func (s *WSGZoneAffinityProfileService) AddProfilesZoneAffinity(ctx context.Context, body *WSGProfileCreateZoneAffinityProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesZoneAffinity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Operation ID: deleteProfilesZoneAffinityById
// Operation path: /profiles/zoneAffinity/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) DeleteProfilesZoneAffinityById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesZoneAffinityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesZoneAffinity
//
// Use this API command to get all zone affinity profiles.
//
// Operation ID: findProfilesZoneAffinity
// Operation path: /profiles/zoneAffinity
// Success code: 200 (OK)
//
// Optional parameters:
// - vdpId string
//		- nullable
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileZoneAffinityProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileZoneAffinityProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesZoneAffinity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["vdpId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("vdpId", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileZoneAffinityProfileListAPIResponse), err
}

// FindProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Operation ID: findProfilesZoneAffinityById
// Operation path: /profiles/zoneAffinity/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileReturnZoneAffinityProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileReturnZoneAffinityProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesZoneAffinityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileReturnZoneAffinityProfileAPIResponse), err
}

// PartialUpdateProfilesZoneAffinityById
//
// Use this API command to modify zone affinity profile.
//
// Operation ID: partialUpdateProfilesZoneAffinityById
// Operation path: /profiles/zoneAffinity/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyZoneAffinityProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *WSGProfileModifyZoneAffinityProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesZoneAffinityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
