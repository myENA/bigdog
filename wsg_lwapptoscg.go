package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGLwapptoscgService struct {
	apiClient *VSZClient
}

func NewWSGLwapptoscgService(c *VSZClient) *WSGLwapptoscgService {
	s := new(WSGLwapptoscgService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLwapptoscgService() *WSGLwapptoscgService {
	return NewWSGLwapptoscgService(ss.apiClient)
}

// FindLwapp2scg
//
// Use this API command to retrieve Lwapp Config.
func (s *WSGLwapptoscgService) FindLwapp2scg(ctx context.Context) (*WSGSystemLwapp2scgConfiguration, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLwapp2scg, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemLwapp2scgConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateLwapp2scg
//
// Use this API command to modify the basic information of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLwapptoscgService) PartialUpdateLwapp2scg(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateLwapp2scg, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLwapptoscgService) PartialUpdateLwapp2scgApList(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateLwapp2scgApList, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
