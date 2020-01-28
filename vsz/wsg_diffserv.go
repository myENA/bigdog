package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGDiffServService struct {
	apiClient *APIClient
}

func NewWSGDiffServService(c *APIClient) *WSGDiffServService {
	s := new(WSGDiffServService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDiffServService() *WSGDiffServService {
	return NewWSGDiffServService(ss.apiClient)
}

// AddRkszonesDiffservByZoneId
//
// Use this API command to create DiffServ profile.
//
// Request Body:
//	 - body *WSGZoneCreateDiffServProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDiffServService) AddRkszonesDiffservByZoneId(ctx context.Context, body *WSGZoneCreateDiffServProfile, zoneId string) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDiffservByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesDiffservById
//
// Use this API command to delete DiffServ profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDiffServService) DeleteRkszonesDiffservById(ctx context.Context, id string, zoneId string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDiffservById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindRkszonesDiffservById
//
// Use this API command to retrieve DiffServ profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservById(ctx context.Context, id string, zoneId string) (*WSGZoneDiffServConfiguration, error) {
	var (
		req      *APIRequest
		resp     *WSGZoneDiffServConfiguration
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDiffservById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGZoneDiffServConfiguration()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesDiffservByZoneId
//
// Use this API command to retrieve a list of DiffServ profile.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservByZoneId(ctx context.Context, zoneId string) (*WSGZoneDiffServList, error) {
	var (
		req      *APIRequest
		resp     *WSGZoneDiffServList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDiffservByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGZoneDiffServList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateRkszonesDiffservById
//
// Use this API command to modify the basic information of DiffServ profile.
//
// Request Body:
//	 - body *WSGZoneModifyDiffServProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDiffServService) PartialUpdateRkszonesDiffservById(ctx context.Context, body *WSGZoneModifyDiffServProfile, id string, zoneId string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesDiffservById, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}
