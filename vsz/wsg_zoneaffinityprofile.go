package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGZoneAffinityProfileService struct {
	apiClient *APIClient
}

func NewWSGZoneAffinityProfileService(c *APIClient) *WSGZoneAffinityProfileService {
	s := new(WSGZoneAffinityProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZoneAffinityProfileService() *WSGZoneAffinityProfileService {
	return NewWSGZoneAffinityProfileService(ss.apiClient)
}

// AddProfilesZoneAffinity
//
// Use this API command to create zone affinity profile.
//
// Request Body:
//	 - body *WSGProfileCreateZoneAffinityProfile
func (s *WSGZoneAffinityProfileService) AddProfilesZoneAffinity(ctx context.Context, body *WSGProfileCreateZoneAffinityProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddProfilesZoneAffinity), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) DeleteProfilesZoneAffinityById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteProfilesZoneAffinityById), true)
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
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, optionalParams map[string][]string) (*WSGProfileZoneAffinityProfileList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesZoneAffinity), true)
	if v, ok := optionalParams["vdpId"]; ok {
		req.AddQueryParameter("vdpId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileZoneAffinityProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, id string) (*WSGProfileReturnZoneAffinityProfile, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesZoneAffinityById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileReturnZoneAffinityProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGZoneAffinityProfileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *WSGProfileModifyZoneAffinityProfile, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateProfilesZoneAffinityById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
