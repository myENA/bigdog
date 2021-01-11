package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGHotspot20VenueServiceService struct {
	apiClient *VSZClient
}

func NewWSGHotspot20VenueServiceService(c *VSZClient) *WSGHotspot20VenueServiceService {
	s := new(WSGHotspot20VenueServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20VenueServiceService() *WSGHotspot20VenueServiceService {
	return NewWSGHotspot20VenueServiceService(ss.apiClient)
}

// AddRkszonesHs20VenuesByZoneId
//
// Use this API command to create a new Hotspot 2.0 venue profile of a zone.
//
// Operation ID: addRkszonesHs20VenuesByZoneId
// Operation path: /rkszones/{zoneId}/hs20/venues
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGPortalServiceCreateHotspot20VenueProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) AddRkszonesHs20VenuesByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20VenueProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesHs20VenuesByZoneId, true)
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

// DeleteRkszonesHs20VenuesById
//
// Use this API command to delete Hotspot 2.0 venue profile of a zone.
//
// Operation ID: deleteRkszonesHs20VenuesById
// Operation path: /rkszones/{zoneId}/hs20/venues/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) DeleteRkszonesHs20VenuesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesHs20VenuesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesHs20VenuesById
//
// Use this API command to retrieve a Hotspot 2.0 venue profile of a zone.
//
// Operation ID: findRkszonesHs20VenuesById
// Operation path: /rkszones/{zoneId}/hs20/venues/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceHotspot20VeuneProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceHotspot20VeuneProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceHotspot20VeuneProfileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesHs20VenuesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceHotspot20VeuneProfileAPIResponse), err
}

// FindRkszonesHs20VenuesByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 venue profile of a zone.
//
// Operation ID: findRkszonesHs20VenuesByZoneId
// Operation path: /rkszones/{zoneId}/hs20/venues
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesHs20VenuesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceListAPIResponse), err
}

// FindServicesVenueProfileByQueryCriteria
//
// Query Venue Profiles with specified filters.
//
// Operation ID: findServicesVenueProfileByQueryCriteria
// Operation path: /query/services/venueProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20VenueServiceService) FindServicesVenueProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesVenueProfileByQueryCriteria, true)
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

// PartialUpdateRkszonesHs20VenuesById
//
// Use this API command to modify the configuration on Hotspot 2.0 venue profile of a zone.
//
// Operation ID: partialUpdateRkszonesHs20VenuesById
// Operation path: /rkszones/{zoneId}/hs20/venues/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGPortalServiceModifyHotspot20VenueProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) PartialUpdateRkszonesHs20VenuesById(ctx context.Context, body *WSGPortalServiceModifyHotspot20VenueProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesHs20VenuesById, true)
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
