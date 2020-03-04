package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGConnectivityToolsService struct {
	apiClient *APIClient
}

func NewWSGConnectivityToolsService(c *APIClient) *WSGConnectivityToolsService {
	s := new(WSGConnectivityToolsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGConnectivityToolsService() *WSGConnectivityToolsService {
	return NewWSGConnectivityToolsService(ss.apiClient)
}

// AddToolSpeedflex
//
// Use this API command to start the SpeedFlex test.
//
// Request Body:
//	 - body *WSGToolSpeedFlex
func (s *WSGConnectivityToolsService) AddToolSpeedflex(ctx context.Context, body *WSGToolSpeedFlex) (*WSGToolTestResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGToolTestResult
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddToolSpeedflex), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGToolTestResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindToolPing
//
// Use this API command to run the PING test on an AP.
//
// Required Parameters:
// - apMac string
//		- required
// - targetIP string
//		- required
func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, apMac string, targetIP string) (*string, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *string
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, targetIP, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindToolPing), true)
	req.SetQueryParameter("apMac", []string{apMac})
	req.SetQueryParameter("targetIP", []string{targetIP})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(string)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindToolSpeedflexByWcid
//
// Use this API command to retrieve existing SpeedFlex test results.
//
// Required Parameters:
// - wcid string
//		- required
func (s *WSGConnectivityToolsService) FindToolSpeedflexByWcid(ctx context.Context, wcid string) (*WSGToolTestResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGToolTestResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, wcid, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindToolSpeedflexByWcid), true)
	req.SetPathParameter("wcid", wcid)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGToolTestResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindToolTraceRoute
//
// Use this API command to run the traceroute test on an AP.
//
// Required Parameters:
// - apMac string
//		- required
// - targetIP string
//		- required
//
// Optional Parameters:
// - timeoutInSec string
//		- nullable
func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, apMac string, targetIP string, optionalParams map[string][]string) (*string, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *string
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, targetIP, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindToolTraceRoute), true)
	req.SetQueryParameter("apMac", []string{apMac})
	req.SetQueryParameter("targetIP", []string{targetIP})
	if v, ok := optionalParams["timeoutInSec"]; ok {
		req.AddQueryParameter("timeoutInSec", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(string)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
