package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addProfilesZoneAffinity
//
// Use this API command to create zone affinity profile.
//
// Request Body:
//	 - body *WSGProfileCreateZoneAffinityProfile
func (s *WSGZoneAffinityProfileService) AddProfilesZoneAffinity(ctx context.Context, body *WSGProfileCreateZoneAffinityProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesZoneAffinity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesZoneAffinityById
//
// Operation ID: deleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) DeleteProfilesZoneAffinityById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesZoneAffinityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindProfilesZoneAffinity
//
// Operation ID: findProfilesZoneAffinity
//
// Use this API command to get all zone affinity profiles.
//
// Optional Parameters:
// - vdpId string
//		- nullable
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileZoneAffinityProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileZoneAffinityProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileZoneAffinityProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesZoneAffinity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["vdpId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("vdpId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileZoneAffinityProfileListAPIResponse), err
}

// FindProfilesZoneAffinityById
//
// Operation ID: findProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileReturnZoneAffinityProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileReturnZoneAffinityProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileReturnZoneAffinityProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesZoneAffinityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileReturnZoneAffinityProfileAPIResponse), err
}

// PartialUpdateProfilesZoneAffinityById
//
// Operation ID: partialUpdateProfilesZoneAffinityById
//
// Use this API command to modify zone affinity profile.
//
// Request Body:
//	 - body *WSGProfileModifyZoneAffinityProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *WSGProfileModifyZoneAffinityProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesZoneAffinityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
