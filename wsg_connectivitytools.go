package bigdog

// API Version: v9_1

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
// Operation ID: addToolSpeedflex
//
// Use this API command to start the SpeedFlex test.
//
// Request Body:
//	 - body *WSGToolSpeedFlex
func (s *WSGConnectivityToolsService) AddToolSpeedflex(ctx context.Context, body *WSGToolSpeedFlex, mutators ...RequestMutator) (*WSGToolTestResult, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGToolTestResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindToolPing
//
// Operation ID: findToolPing
//
// Use this API command to run the PING test on an AP.
//
// Required Parameters:
// - apMac string
//		- required
// - targetIP string
//		- required
func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, apMac string, targetIP string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolPing, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetQueryParameter("apMac", apMac)
	req.SetQueryParameter("targetIP", targetIP)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindToolSpeedflexByWcid
//
// Operation ID: findToolSpeedflexByWcid
//
// Use this API command to retrieve existing SpeedFlex test results.
//
// Required Parameters:
// - wcid string
//		- required
func (s *WSGConnectivityToolsService) FindToolSpeedflexByWcid(ctx context.Context, wcid string, mutators ...RequestMutator) (*WSGToolTestResult, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("wcid", wcid)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGToolTestResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindToolTraceRoute
//
// Operation ID: findToolTraceRoute
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
func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, apMac string, targetIP string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolTraceRoute, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetQueryParameter("apMac", apMac)
	req.SetQueryParameter("targetIP", targetIP)
	if v, ok := optionalParams["timeoutInSec"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("timeoutInSec", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
