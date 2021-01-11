package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGL3RoamingService struct {
	apiClient *VSZClient
}

func NewWSGL3RoamingService(c *VSZClient) *WSGL3RoamingService {
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
//
// Operation ID: findProfilesTunnelL3Roaming
// Operation path: /profiles/tunnel/l3Roaming
// Success code: 200 (OK)
func (s *WSGL3RoamingService) FindProfilesTunnelL3Roaming(ctx context.Context, mutators ...RequestMutator) (*WSGProfileGetL3RoamingConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileGetL3RoamingConfigAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileGetL3RoamingConfigAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesTunnelL3Roaming, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileGetL3RoamingConfigAPIResponse), err
}

// PartialUpdateProfilesTunnelL3Roaming
//
// Use this API command to modify L3 Roaming basic configuration.
//
// Operation ID: partialUpdateProfilesTunnelL3Roaming
// Operation path: /profiles/tunnel/l3Roaming
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileUpdateL3RoamingConfig
func (s *WSGL3RoamingService) PartialUpdateProfilesTunnelL3Roaming(ctx context.Context, body *WSGProfileUpdateL3RoamingConfig, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelL3Roaming, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
