package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGClientIsolationWhitelistService struct {
	apiClient *APIClient
}

func NewWSGClientIsolationWhitelistService(c *APIClient) *WSGClientIsolationWhitelistService {
	s := new(WSGClientIsolationWhitelistService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService() *WSGClientIsolationWhitelistService {
	return NewWSGClientIsolationWhitelistService(ss.apiClient)
}

// AddRkszonesClientIsolationWhitelistByZoneId
//
// Create a new ClientIsolationWhitelist.
//
// Request Body:
//	 - body *WSGProfileCreateClientIsolationWhitelist
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) AddRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, body *WSGProfileCreateClientIsolationWhitelist, zoneId string) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesClientIsolationWhitelistByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesClientIsolationWhitelist
//
// Use this API command to delete Bulk Client Isolation Whitelist.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelist, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteRkszonesClientIsolationWhitelistById
//
// Delete a Client Isolation Whitelist.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelistById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindRkszonesClientIsolationWhitelistById
//
// Retrieve an Client Isolation Whitelist.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById(ctx context.Context, id string, zoneId string) (*WSGProfileClientIsolationWhitelist, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileClientIsolationWhitelist
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClientIsolationWhitelist()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesClientIsolationWhitelistByZoneId
//
// Retrieve a list of Client Isolation Whitelist.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, zoneId string) (*WSGProfileClientIsolationWhitelistArray, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileClientIsolationWhitelistArray
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClientIsolationWhitelistArray()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesClientIsolationWhitelistByQueryCriteria
//
// Retrieve a list of Client Isolation Whitelist.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileClientIsolationWhitelistArray, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileClientIsolationWhitelistArray
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesClientIsolationWhitelistByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClientIsolationWhitelistArray()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateRkszonesClientIsolationWhitelistById
//
// Modify a specific Client Isolation Whitelist basic.
//
// Request Body:
//	 - body *WSGProfileModifyClientIsolationWhitelist
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) PartialUpdateRkszonesClientIsolationWhitelistById(ctx context.Context, body *WSGProfileModifyClientIsolationWhitelist, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesClientIsolationWhitelistById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
