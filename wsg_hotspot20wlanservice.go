package bigdog

// API Version: v9_0

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
func (s *WSGHotspot20WLANServiceService) AddRkszonesHs20sByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20WlanProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesHs20sByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGHotspot20WLANServiceService) DeleteRkszonesHs20sById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesHs20sById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceHotspot20WlanProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceHotspot20WlanProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesHs20sById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPortalServiceHotspot20WlanProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesHs20sByZoneId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPortalServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesHotspot20ProfileByQueryCriteria
//
// Operation ID: findServicesHotspot20ProfileByQueryCriteria
//
// Query Hotspot20 Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20WLANServiceService) FindServicesHotspot20ProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesHotspot20ProfileByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGHotspot20WLANServiceService) PartialUpdateRkszonesHs20sById(ctx context.Context, body *WSGPortalServiceModifyHotspot20WlanProfile, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesHs20sById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
