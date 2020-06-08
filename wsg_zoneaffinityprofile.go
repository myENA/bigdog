package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGZoneaffinityprofileService struct {
	apiClient *VSZClient
}

func NewWSGZoneaffinityprofileService(c *VSZClient) *WSGZoneaffinityprofileService {
	s := new(WSGZoneaffinityprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZoneaffinityprofileService() *WSGZoneaffinityprofileService {
	return NewWSGZoneaffinityprofileService(ss.apiClient)
}

// AddProfilesZoneAffinity
//
// Use this API command to create zone affinity profile.
//
// Request Body:
//	 - body *WSGProfileCreateZoneAffinityProfile
func (s *WSGZoneaffinityprofileService) AddProfilesZoneAffinity(ctx context.Context, body *WSGProfileCreateZoneAffinityProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesZoneAffinity, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneaffinityprofileService) DeleteProfilesZoneAffinityById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesZoneAffinityById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesZoneAffinity
//
// Use this API command to get all zone affinity profiles.
//
// Optional Parameters:
// - vdpId string
//		- nullable
func (s *WSGZoneaffinityprofileService) FindProfilesZoneAffinity(ctx context.Context, optionalParams map[string][]string) (*WSGProfileZoneAffinityProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileZoneAffinityProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesZoneAffinity, true)
	if v, ok := optionalParams["vdpId"]; ok && len(v) > 0 {
		req.SetQueryParameter("vdpId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileZoneAffinityProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneaffinityprofileService) FindProfilesZoneAffinityById(ctx context.Context, id string) (*WSGProfileReturnZoneAffinityProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileReturnZoneAffinityProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesZoneAffinityById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileReturnZoneAffinityProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesZoneAffinityById
//
// Use this API command to modify zone affinity profile.
//
// Request Body:
//	 - body *WSGProfileModifyZoneAffinityProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneaffinityprofileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *WSGProfileModifyZoneAffinityProfile, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesZoneAffinityById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
