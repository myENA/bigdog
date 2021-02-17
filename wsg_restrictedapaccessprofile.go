package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGRestrictedAPAccessProfileService struct {
	apiClient *VSZClient
}

func NewWSGRestrictedAPAccessProfileService(c *VSZClient) *WSGRestrictedAPAccessProfileService {
	s := new(WSGRestrictedAPAccessProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRestrictedAPAccessProfileService() *WSGRestrictedAPAccessProfileService {
	return NewWSGRestrictedAPAccessProfileService(ss.apiClient)
}

// AddRkszonesRestrictedApAccessProfilesByZoneId
//
// Create a Restricted AP Access Profile.
//
// Operation ID: addRkszonesRestrictedApAccessProfilesByZoneId
// Operation path: /rkszones/{zoneId}/restrictedApAccessProfiles
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateRestrictedApAccessProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) AddRkszonesRestrictedApAccessProfilesByZoneId(ctx context.Context, body *WSGProfileCreateRestrictedApAccessProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesRestrictedApAccessProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesRestrictedApAccessProfiles
//
// Use this API command to delete Bulk Restricted AP Access Profile.
//
// Operation ID: deleteRkszonesRestrictedApAccessProfiles
// Operation path: /rkszones/restrictedApAccessProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRestrictedAPAccessProfileService) DeleteRkszonesRestrictedApAccessProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesRestrictedApAccessProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesRestrictedApAccessProfilesById
//
// Delete a Restricted AP Access Profile.
//
// Operation ID: deleteRkszonesRestrictedApAccessProfilesById
// Operation path: /rkszones/{zoneId}/restrictedApAccessProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) DeleteRkszonesRestrictedApAccessProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesRestrictedApAccessProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesRestrictedApAccessProfilesById
//
// Retrieve a Restricted AP Access Profile.
//
// Operation ID: findRkszonesRestrictedApAccessProfilesById
// Operation path: /rkszones/{zoneId}/restrictedApAccessProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileRestrictedApAccessProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileRestrictedApAccessProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesRestrictedApAccessProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileRestrictedApAccessProfileAPIResponse), err
}

// FindRkszonesRestrictedApAccessProfilesByQueryCriteria
//
// Retrieve a list of Restricted AP Access Profile.
//
// Operation ID: findRkszonesRestrictedApAccessProfilesByQueryCriteria
// Operation path: /rkszones/restrictedApAccessProfiles/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileRestrictedApAccessProfileArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileRestrictedApAccessProfileArrayAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindRkszonesRestrictedApAccessProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileRestrictedApAccessProfileArrayAPIResponse), err
}

// FindRkszonesRestrictedApAccessProfilesByZoneId
//
// Retrieve Restricted AP Access Profile list.
//
// Operation ID: findRkszonesRestrictedApAccessProfilesByZoneId
// Operation path: /rkszones/{zoneId}/restrictedApAccessProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileIdListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileIdListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesRestrictedApAccessProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileIdListAPIResponse), err
}

// UpdateRkszonesRestrictedApAccessProfilesById
//
// Modify a Restricted AP Access Profile.
//
// Operation ID: updateRkszonesRestrictedApAccessProfilesById
// Operation path: /rkszones/{zoneId}/restrictedApAccessProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyRestrictedApAccessProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) UpdateRkszonesRestrictedApAccessProfilesById(ctx context.Context, body *WSGProfileModifyRestrictedApAccessProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesRestrictedApAccessProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
