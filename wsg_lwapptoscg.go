package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGLWAPPTOSCGService struct {
	apiClient *VSZClient
}

func NewWSGLWAPPTOSCGService(c *VSZClient) *WSGLWAPPTOSCGService {
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
//
// Operation ID: findLwapp2scg
// Operation path: /lwapp2scg
// Success code: 200 (OK)
func (s *WSGLWAPPTOSCGService) FindLwapp2scg(ctx context.Context, mutators ...RequestMutator) (*WSGSystemLwapp2scgConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSystemLwapp2scgConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindLwapp2scg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemLwapp2scgConfigurationAPIResponse), err
}

// PartialUpdateLwapp2scg
//
// Use this API command to modify the basic information of the Lwapp Config.
//
// Operation ID: partialUpdateLwapp2scg
// Operation path: /lwapp2scg
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scg(ctx context.Context, body *WSGSystemModifyLwapp2scg, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateLwapp2scg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Operation ID: partialUpdateLwapp2scgApList
// Operation path: /lwapp2scg/apList
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList(ctx context.Context, body *WSGSystemModifyLwapp2scg, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateLwapp2scgApList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
