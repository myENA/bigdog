package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGHotspotServiceService struct {
	apiClient *VSZClient
}

func NewWSGHotspotServiceService(c *VSZClient) *WSGHotspotServiceService {
	s := new(WSGHotspotServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspotServiceService() *WSGHotspotServiceService {
	return NewWSGHotspotServiceService(ss.apiClient)
}

// AddRkszonesPortalsHotspotExternalByZoneId
//
// Operation ID: addRkszonesPortalsHotspotExternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with external logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotExternal
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotExternal, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsHotspotExternalByZoneId, true)
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

// AddRkszonesPortalsHotspotInternalByZoneId
//
// Operation ID: addRkszonesPortalsHotspotInternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with internal logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotInternal
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotInternal, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsHotspotInternalByZoneId, true)
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

// AddRkszonesPortalsHotspotSmartClientOnlyByZoneId
//
// Operation ID: addRkszonesPortalsHotspotSmartClientOnlyByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with smart client only of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotSmartClientOnly
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotSmartClientOnly, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsHotspotSmartClientOnlyByZoneId, true)
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

// DeleteRkszonesPortalsHotspotById
//
// Operation ID: deleteRkszonesPortalsHotspotById
//
// Use this API command to delete a Hotspot (WISPr) of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) DeleteRkszonesPortalsHotspotById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalsHotspotById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesPortalsHotspotById
//
// Operation ID: findRkszonesPortalsHotspotById
//
// Use this API command to retrieve a Hotspot (WISPr) of zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceHotspot, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceHotspot
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsHotspotById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPortalServiceHotspot()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesPortalsHotspotByZoneId
//
// Operation ID: findRkszonesPortalsHotspotByZoneId
//
// Use this API command to retrieve a list of Hotspot (WISPr) of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsHotspotByZoneId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPortalServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesHotspotByQueryCriteria
//
// Operation ID: findServicesHotspotByQueryCriteria
//
// Query Hotspot Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspotServiceService) FindServicesHotspotByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesHotspotByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesPortalsHotspotById
//
// Operation ID: partialUpdateRkszonesPortalsHotspotById
//
// Use this API command to modify the configuration on Hotspot (WISPr) of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceModifyHotspot
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) PartialUpdateRkszonesPortalsHotspotById(ctx context.Context, body *WSGPortalServiceModifyHotspot, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsHotspotById, true)
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
