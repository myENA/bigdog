package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGLWAPPTOSCGService struct {
	apiClient *APIClient
}

func NewWSGLWAPPTOSCGService(c *APIClient) *WSGLWAPPTOSCGService {
	s := new(WSGLWAPPTOSCGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLWAPPTOSCGService() *WSGLWAPPTOSCGService {
	return NewWSGLWAPPTOSCGService(ss.apiClient)
}

// FindLwapp2scg
//
// Use this API command to retrieve Lwapp Config.
func (s *WSGLWAPPTOSCGService) FindLwapp2scg(ctx context.Context) (*WSGSystemLwapp2scgConfiguration, error) {
	var (
		req      *APIRequest
		resp     *WSGSystemLwapp2scgConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLwapp2scg, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemLwapp2scgConfiguration()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateLwapp2scg
//
// Use this API command to modify the basic information of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scg(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateLwapp2scg, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateLwapp2scgApList, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
