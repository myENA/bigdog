package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addRkszonesRestrictedApAccessProfilesByZoneId
//
// Create a Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGProfileCreateRestrictedApAccessProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) AddRkszonesRestrictedApAccessProfilesByZoneId(ctx context.Context, body *WSGProfileCreateRestrictedApAccessProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesRestrictedApAccessProfilesByZoneId, true)
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

// DeleteRkszonesRestrictedApAccessProfiles
//
// Operation ID: deleteRkszonesRestrictedApAccessProfiles
//
// Use this API command to delete Bulk Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRestrictedAPAccessProfileService) DeleteRkszonesRestrictedApAccessProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesRestrictedApAccessProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesRestrictedApAccessProfilesById
//
// Operation ID: deleteRkszonesRestrictedApAccessProfilesById
//
// Delete a Restricted AP Access Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) DeleteRkszonesRestrictedApAccessProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesRestrictedApAccessProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindRkszonesRestrictedApAccessProfilesById
//
// Operation ID: findRkszonesRestrictedApAccessProfilesById
//
// Retrieve a Restricted AP Access Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileRestrictedApAccessProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileRestrictedApAccessProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileRestrictedApAccessProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesRestrictedApAccessProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileRestrictedApAccessProfileAPIResponse), err
}

// FindRkszonesRestrictedApAccessProfilesByQueryCriteria
//
// Operation ID: findRkszonesRestrictedApAccessProfilesByQueryCriteria
//
// Retrieve a list of Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileRestrictedApAccessProfileArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileRestrictedApAccessProfileArrayAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileRestrictedApAccessProfileArrayAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRkszonesRestrictedApAccessProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileRestrictedApAccessProfileArrayAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileRestrictedApAccessProfileArrayAPIResponse), err
}

// FindRkszonesRestrictedApAccessProfilesByZoneId
//
// Operation ID: findRkszonesRestrictedApAccessProfilesByZoneId
//
// Retrieve Restricted AP Access Profile list.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileIdListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileIdListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileIdListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesRestrictedApAccessProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileIdListAPIResponse), err
}

// UpdateRkszonesRestrictedApAccessProfilesById
//
// Operation ID: updateRkszonesRestrictedApAccessProfilesById
//
// Modify a Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGProfileModifyRestrictedApAccessProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) UpdateRkszonesRestrictedApAccessProfilesById(ctx context.Context, body *WSGProfileModifyRestrictedApAccessProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesRestrictedApAccessProfilesById, true)
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
