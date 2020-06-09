package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGConnectivityToolsService struct {
	apiClient *VSZClient
}

func NewWSGConnectivityToolsService(c *VSZClient) *WSGConnectivityToolsService {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddToolSpeedflex, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGToolTestResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, apMac string, targetIP string) (*WSGModelsFindToolPing200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGModelsFindToolPing200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolPing, true)
	req.SetQueryParameter("apMac", []string{apMac})
	req.SetQueryParameter("targetIP", []string{targetIP})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGModelsFindToolPing200ResponseType()
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolSpeedflexByWcid, true)
	req.SetPathParameter("wcid", wcid)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGToolTestResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, apMac string, targetIP string, optionalParams map[string][]string) (*WSGModelsFindToolTraceRoute200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGModelsFindToolTraceRoute200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolTraceRoute, true)
	req.SetQueryParameter("apMac", []string{apMac})
	req.SetQueryParameter("targetIP", []string{targetIP})
	if v, ok := optionalParams["timeoutInSec"]; ok && len(v) > 0 {
		req.SetQueryParameter("timeoutInSec", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGModelsFindToolTraceRoute200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
