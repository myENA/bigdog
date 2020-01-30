package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWechatService struct {
	apiClient *APIClient
}

func NewWSGWechatService(c *APIClient) *WSGWechatService {
	s := new(WSGWechatService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWechatService() *WSGWechatService {
	return NewWSGWechatService(ss.apiClient)
}

// AddRkszonesPortalsWechatByZoneId
//
// Use this API command to create wechat portal.
//
// Request Body:
//	 - body *WSGPortalServiceCreateWechat
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWechatService) AddRkszonesPortalsWechatByZoneId(ctx context.Context, body *WSGPortalServiceCreateWechat, zoneId string) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsWechatByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesPortalsWechatById
//
// Use this API command to delete wechat portal.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWechatService) DeleteRkszonesPortalsWechatById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalsWechatById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindRkszonesPortalsWechatById
//
// Use this API command to retrieve wechat portal by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWechatService) FindRkszonesPortalsWechatById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceWechatConfiguration, error) {
	var (
		req      *APIRequest
		resp     *WSGPortalServiceWechatConfiguration
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsWechatById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceWechatConfiguration()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesPortalsWechatByZoneId
//
// Use this API command to retrieve a list of wechat portal.
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
func (s *WSGWechatService) FindRkszonesPortalsWechatByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGPortalServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGPortalServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsWechatByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesWechatProfileByQueryCriteria
//
// Query Wechat Portals with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWechatService) FindServicesWechatProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesWechatProfileByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// PartialUpdateRkszonesPortalsWechatById
//
// Use this API command to modify the configuration of wechat portal.
//
// Request Body:
//	 - body *WSGPortalServiceModifyWechat
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWechatService) PartialUpdateRkszonesPortalsWechatById(ctx context.Context, body *WSGPortalServiceModifyWechat, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsWechatById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
