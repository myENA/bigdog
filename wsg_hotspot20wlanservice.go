package ruckus

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
// Use this API command to create a new Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspot20WlanProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) AddRkszonesHs20sByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20WlanProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesHs20sById
//
// Use this API command to delete a Hotspot 2.0 WLAN Profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) DeleteRkszonesHs20sById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesHs20sById
//
// Use this API command to retrieve a Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceHotspot20WlanProfile, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceHotspot20WlanProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesHs20sByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 WLAN profiles of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, zoneId string) (*WSGPortalServiceList, *APIResponseMeta, error) {
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
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesHotspot20ProfileByQueryCriteria
//
// Query Hotspot20 Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20WLANServiceService) FindServicesHotspot20ProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesHotspot20ProfileByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesHs20sById
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
func (s *WSGHotspot20WLANServiceService) PartialUpdateRkszonesHs20sById(ctx context.Context, body *WSGPortalServiceModifyHotspot20WlanProfile, id string, zoneId string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
