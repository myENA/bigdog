package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL3RoamingService struct {
	apiClient *APIClient
}

func NewWSGL3RoamingService(c *APIClient) *WSGL3RoamingService {
	s := new(WSGL3RoamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL3RoamingService() *WSGL3RoamingService {
	return NewWSGL3RoamingService(ss.apiClient)
}

// FindProfilesTunnelL3Roaming
//
// Use this API command to retrieve L3 Roaming basic configuration.
func (s *WSGL3RoamingService) FindProfilesTunnelL3Roaming(ctx context.Context) (*WSGProfileGetL3RoamingConfig, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileGetL3RoamingConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelL3Roaming, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileGetL3RoamingConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateProfilesTunnelL3Roaming
//
// Use this API command to modify L3 Roaming basic configuration.
//
// Request Body:
//	 - body *WSGProfileUpdateL3RoamingConfig
func (s *WSGL3RoamingService) PartialUpdateProfilesTunnelL3Roaming(ctx context.Context, body *WSGProfileUpdateL3RoamingConfig) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelL3Roaming, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
