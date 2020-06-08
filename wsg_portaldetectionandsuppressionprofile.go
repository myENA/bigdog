package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGPortaldetectionandsuppressionprofileService struct {
	apiClient *VSZClient
}

func NewWSGPortaldetectionandsuppressionprofileService(c *VSZClient) *WSGPortaldetectionandsuppressionprofileService {
	s := new(WSGPortaldetectionandsuppressionprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGPortaldetectionandsuppressionprofileService() *WSGPortaldetectionandsuppressionprofileService {
	return NewWSGPortaldetectionandsuppressionprofileService(ss.apiClient)
}

// AddRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to create portal detection and suppression profile.
//
// Request Body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) AddRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalDetectionProfilesByZoneId, true)
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

// DeleteRkszonesPortalDetectionProfilesById
//
// Use this API command to delete portal detection and suppression profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) DeleteRkszonesPortalDetectionProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalDetectionProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to delete multiple portal detection and suppression profiles.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) DeleteRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalDetectionProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesPortalDetectionProfilesById
//
// Use this API command to get portal detection and suppression profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) FindRkszonesPortalDetectionProfilesById(ctx context.Context, id string, zoneId string) (*WSGPortalDetectionProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalDetectionProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalDetectionProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalDetectionProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesPortalDetectionProfilesByQueryCriteria
//
// Query portal detection and suppression profile with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPortaldetectionandsuppressionprofileService) FindRkszonesPortalDetectionProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGPortalDetectionProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalDetectionProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesPortalDetectionProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalDetectionProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to get portal detection and suppression profile list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) FindRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, zoneId string) (*WSGPortalDetectionProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalDetectionProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalDetectionProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalDetectionProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Request Body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) PartialUpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalDetectionProfilesById, true)
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

// UpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Request Body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGPortaldetectionandsuppressionprofileService) UpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesPortalDetectionProfilesById, true)
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
