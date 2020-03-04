package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
func (s *WSGLWAPPTOSCGService) FindLwapp2scg(ctx context.Context) (*WSGSystemLwapp2scgConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemLwapp2scgConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindLwapp2scg), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemLwapp2scgConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateLwapp2scg
//
// Use this API command to modify the basic information of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scg(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateLwapp2scg), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateLwapp2scgApList), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
