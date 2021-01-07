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
// Operation ID: addRkszonesSplitTunnelProfilesByZoneId
//
// Create a split tunnel profile.
//
// Request Body:
//	 - body *WSGSplitTunnelCreateSplitTunnelProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) AddRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesSplitTunnelProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesSplitTunnelProfiles
//
// Operation ID: deleteRkszonesSplitTunnelProfiles
//
// Use this API command to delete bulk split tunnel profiles.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesSplitTunnelProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesSplitTunnelProfilesById
//
// Operation ID: deleteRkszonesSplitTunnelProfilesById
//
// Use this API command to delete a split tunnel profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesSplitTunnelProfilesById
//
// Operation ID: findRkszonesSplitTunnelProfilesById
//
// Get a split tunnel profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGSplitTunnelProfileAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSplitTunnelProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSplitTunnelProfile()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesSplitTunnelProfilesByQueryCriteria
//
// Operation ID: findRkszonesSplitTunnelProfilesByQueryCriteria
//
// Use this API command to retrieve a list of split tunnel profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSplitTunnelProfileQueryAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSplitTunnelProfileQuery
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRkszonesSplitTunnelProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSplitTunnelProfileQuery()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesSplitTunnelProfilesByZoneId
//
// Operation ID: findRkszonesSplitTunnelProfilesByZoneId
//
// Get a ID list of split tunnel profile in this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGSplitTunnelProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSplitTunnelProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesSplitTunnelProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSplitTunnelProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesSplitTunnelProfilesById
//
// Operation ID: partialUpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify a split tunnel profile.
//
// Request Body:
//	 - body *WSGSplitTunnelModifySplitTunnelProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) PartialUpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelModifySplitTunnelProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesSplitTunnelProfilesById
//
// Operation ID: updateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify entire information of a split tunnel profile.
//
// Request Body:
//	 - body *WSGSplitTunnelCreateSplitTunnelProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) UpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesSplitTunnelProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
