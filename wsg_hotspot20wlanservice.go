package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGHotspot20WLANServiceService struct {
	apiClient *VSZClient
}

func NewWSGHotspot20WLANServiceService(c *VSZClient) *WSGHotspot20WLANServiceService {
	s := new(WSGHotspot20WLANServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20WLANServiceService() *WSGHotspot20WLANServiceService {
	return NewWSGHotspot20WLANServiceService(ss.apiClient)
}

// AddRkszonesHs20sByZoneId
//
// Operation ID: addRkszonesHs20sByZoneId
//
// Use this API command to create a new Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspot20WlanProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) AddRkszonesHs20sByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20WlanProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesHs20sByZoneId, true)
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

// DeleteRkszonesHs20sById
//
// Operation ID: deleteRkszonesHs20sById
//
// Use this API command to delete a Hotspot 2.0 WLAN Profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) DeleteRkszonesHs20sById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesHs20sById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindRkszonesHs20sById
//
// Operation ID: findRkszonesHs20sById
//
// Use this API command to retrieve a Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceHotspot20WlanProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceHotspot20WlanProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceHotspot20WlanProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesHs20sById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceHotspot20WlanProfileAPIResponse), err
}

// FindRkszonesHs20sByZoneId
//
// Operation ID: findRkszonesHs20sByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 WLAN profiles of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesHs20sByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceListAPIResponse), err
}

// FindServicesHotspot20ProfileByQueryCriteria
//
// Operation ID: findServicesHotspot20ProfileByQueryCriteria
//
// Query Hotspot20 Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20WLANServiceService) FindServicesHotspot20ProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesHotspot20ProfileByQueryCriteria, true)
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

// PartialUpdateRkszonesHs20sById
//
// Operation ID: partialUpdateRkszonesHs20sById
//
// Use this API command to modify the configuration on Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceModifyHotspot20WlanProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) PartialUpdateRkszonesHs20sById(ctx context.Context, body *WSGPortalServiceModifyHotspot20WlanProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesHs20sById, true)
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
