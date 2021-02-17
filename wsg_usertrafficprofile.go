package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGUserTrafficProfileService struct {
	apiClient *VSZClient
}

func NewWSGUserTrafficProfileService(c *VSZClient) *WSGUserTrafficProfileService {
	s := new(WSGUserTrafficProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGUserTrafficProfileService() *WSGUserTrafficProfileService {
	return NewWSGUserTrafficProfileService(ss.apiClient)
}

// AddProfilesUtp
//
// Use this API command to create a new user traffic profile.
//
// Operation ID: addProfilesUtp
// Operation path: /profiles/utp
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateUserTrafficProfile
func (s *WSGUserTrafficProfileService) AddProfilesUtp(ctx context.Context, body *WSGProfileCreateUserTrafficProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesUtp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddProfilesUtpCloneById
//
// Use this API command to copy a traffic profile.
//
// Operation ID: addProfilesUtpCloneById
// Operation path: /profiles/utp/clone/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGProfileClone
//
// Required parameters:
// - id string
//		- required
func (s *WSGUserTrafficProfileService) AddProfilesUtpCloneById(ctx context.Context, body *WSGProfileClone, id string, mutators ...RequestMutator) (*WSGProfileCloneAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileCloneAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesUtpCloneById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileCloneAPIResponse), err
}

// DeleteProfilesUtp
//
// Use this API command to delete a list of traffic profile.
//
// Operation ID: deleteProfilesUtp
// Operation path: /profiles/utp
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileDeleteBulkUserTrafficProfile
func (s *WSGUserTrafficProfileService) DeleteProfilesUtp(ctx context.Context, body *WSGProfileDeleteBulkUserTrafficProfile, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesUtp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesUtpById
//
// Use this API command to delete an user traffic profile.
//
// Operation ID: deleteProfilesUtpById
// Operation path: /profiles/utp/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesUtpById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesUtpDownlinkRateLimitingById
//
// Use this API command to disable downlink rate limiting of user traffic profile.
//
// Operation ID: deleteProfilesUtpDownlinkRateLimitingById
// Operation path: /profiles/utp/{id}/downlinkRateLimiting
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpDownlinkRateLimitingById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesUtpDownlinkRateLimitingById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesUtpUplinkRateLimitingById
//
// Use this API command to disable uplink rateLimiting of user traffic profile.
//
// Operation ID: deleteProfilesUtpUplinkRateLimitingById
// Operation path: /profiles/utp/{id}/uplinkRateLimiting
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpUplinkRateLimitingById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesUtpUplinkRateLimitingById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesUtp
//
// Use this API command to retrieve a list of user traffic profile.
//
// Operation ID: findProfilesUtp
// Operation path: /profiles/utp
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGUserTrafficProfileService) FindProfilesUtp(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesUtp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesUtpById
//
// Use this API command to retrieve an user traffic profile.
//
// Operation ID: findProfilesUtpById
// Operation path: /profiles/utp/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGUserTrafficProfileService) FindProfilesUtpById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileUserTrafficProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileUserTrafficProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesUtpById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileUserTrafficProfileAPIResponse), err
}

// FindProfilesUtpByQueryCriteria
//
// Use this API command to retrieve a list of User Traffic Profile by query criteria.
//
// Operation ID: findProfilesUtpByQueryCriteria
// Operation path: /profiles/utp/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGUserTrafficProfileService) FindProfilesUtpByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileUserTrafficProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileUserTrafficProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesUtpByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileUserTrafficProfileListAPIResponse), err
}

// PartialUpdateProfilesUtpById
//
// Use this API command to modify the configuration of user traffic profile.
//
// Operation ID: partialUpdateProfilesUtpById
// Operation path: /profiles/utp/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyUserTrafficProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGUserTrafficProfileService) PartialUpdateProfilesUtpById(ctx context.Context, body *WSGProfileModifyUserTrafficProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesUtpById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
