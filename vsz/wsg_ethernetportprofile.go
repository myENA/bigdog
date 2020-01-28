package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGEthernetPortProfileService struct {
	apiClient *APIClient
}

func NewWSGEthernetPortProfileService(c *APIClient) *WSGEthernetPortProfileService {
	s := new(WSGEthernetPortProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEthernetPortProfileService() *WSGEthernetPortProfileService {
	return NewWSGEthernetPortProfileService(ss.apiClient)
}

// AddRkszonesProfileEthernetPortByZoneId
//
// Create a new Ethernet Port Porfile.
//
// Request Body:
//	 - body *WSGEthernetPortCreateEthernetPortProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) AddRkszonesProfileEthernetPortByZoneId(ctx context.Context, body *WSGEthernetPortCreateEthernetPortProfile, zoneId string) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesProfileEthernetPortByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesProfileEthernetPortById
//
// Delete Ethernet Port Porfile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) DeleteRkszonesProfileEthernetPortById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesProfileEthernetPortById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindRkszonesProfileEthernetPortById
//
// Retrieve a Ethernet Port Porfile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById(ctx context.Context, id string, zoneId string) (*WSGEthernetPortProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGEthernetPortProfile
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesProfileEthernetPortById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEthernetPortProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesProfileEthernetPortByZoneId
//
// Retrieve a list of Ethernet Port Porfiles within a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGEthernetPortProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGEthernetPortProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesProfileEthernetPortByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEthernetPortProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateRkszonesProfileEthernetPortById
//
// Modify a specific Ethernet Port Porfile.
//
// Request Body:
//	 - body *WSGEthernetPortModifyEthernetPortProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) PartialUpdateRkszonesProfileEthernetPortById(ctx context.Context, body *WSGEthernetPortModifyEthernetPortProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesProfileEthernetPortById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
