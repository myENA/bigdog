package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGPortalDetectionandSuppressionProfileService struct {
	apiClient *APIClient
}

func NewWSGPortalDetectionandSuppressionProfileService(c *APIClient) *WSGPortalDetectionandSuppressionProfileService {
	s := new(WSGPortalDetectionandSuppressionProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGPortalDetectionandSuppressionProfileService() *WSGPortalDetectionandSuppressionProfileService {
	return NewWSGPortalDetectionandSuppressionProfileService(ss.apiClient)
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
func (s *WSGPortalDetectionandSuppressionProfileService) AddRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalDetectionProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalDetectionProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalDetectionProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesById(ctx context.Context, id string, zoneId string) (*WSGPortalDetectionProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGPortalDetectionProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalDetectionProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalDetectionProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesPortalDetectionProfilesByQueryCriteria
//
// Query portal detection and suppression profile with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGPortalDetectionProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGPortalDetectionProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesPortalDetectionProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalDetectionProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to get portal detection and suppression profile list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, zoneId string) (*WSGPortalDetectionProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGPortalDetectionProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalDetectionProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalDetectionProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) PartialUpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalDetectionProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) UpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesPortalDetectionProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
