package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL3roamingService struct {
	apiClient *APIClient
}

func NewWSGL3roamingService(c *APIClient) *WSGL3roamingService {
	s := new(WSGL3roamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL3roamingService() *WSGL3roamingService {
	return NewWSGL3roamingService(ss.apiClient)
}

// FindProfilesTunnelL3Roaming
//
// Use this API command to retrieve L3 Roaming basic configuration.
func (s *WSGL3roamingService) FindProfilesTunnelL3Roaming(ctx context.Context) (*WSGProfileGetL3RoamingConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileGetL3RoamingConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelL3Roaming, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileGetL3RoamingConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesTunnelL3Roaming
//
// Use this API command to modify L3 Roaming basic configuration.
//
// Request Body:
//	 - body *WSGProfileUpdateL3RoamingConfig
func (s *WSGL3roamingService) PartialUpdateProfilesTunnelL3Roaming(ctx context.Context, body *WSGProfileUpdateL3RoamingConfig) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelL3Roaming, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
