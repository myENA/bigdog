package vsz

// API Version: v8_1

import (
	"context"
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
func (s *WSGZoneAffinityProfileService) AddProfilesZoneAffinity(ctx context.Context, body *WSGProfileCreateZoneAffinityProfile) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesZoneAffinity, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) DeleteProfilesZoneAffinityById(ctx context.Context, id string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesZoneAffinityById, true)
}

// FindProfilesZoneAffinity
//
// Use this API command to get all zone affinity profiles.
//
// Optional Parameters:
// - vdpId string
//		- nullable
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, optionalParams map[string]interface{}) (*WSGProfileZoneAffinityProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileZoneAffinityProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesZoneAffinity, true)
}

// FindProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, id string) (*WSGProfileReturnZoneAffinityProfile, error) {
	var (
		req  *APIRequest
		resp *WSGProfileReturnZoneAffinityProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesZoneAffinityById, true)
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
func (s *WSGZoneAffinityProfileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *WSGProfileModifyZoneAffinityProfile, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesZoneAffinityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
