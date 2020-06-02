package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGConnectivitytoolsService struct {
	apiClient *APIClient
}

func NewWSGConnectivitytoolsService(c *APIClient) *WSGConnectivitytoolsService {
	s := new(WSGConnectivitytoolsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGConnectivitytoolsService() *WSGConnectivitytoolsService {
	return NewWSGConnectivitytoolsService(ss.apiClient)
}

// AddToolSpeedflex
//
// Use this API command to start the SpeedFlex test.
//
// Request Body:
//	 - body *WSGToolSpeedFlex
func (s *WSGConnectivitytoolsService) AddToolSpeedflex(ctx context.Context, body *WSGToolSpeedFlex) (*WSGToolTestResult, *APIResponseMeta, error) {
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
func (s *WSGConnectivitytoolsService) FindToolPing(ctx context.Context, apMac string, targetIP string) (*string, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolPing, true)
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
func (s *WSGConnectivitytoolsService) FindToolSpeedflexByWcid(ctx context.Context, wcid string) (*WSGToolTestResult, *APIResponseMeta, error) {
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
func (s *WSGConnectivitytoolsService) FindToolTraceRoute(ctx context.Context, apMac string, targetIP string, optionalParams map[string][]string) (*string, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindToolTraceRoute, true)
	req.SetQueryParameter("apMac", []string{apMac})
	req.SetQueryParameter("targetIP", []string{targetIP})
	if v, ok := optionalParams["timeoutInSec"]; ok && len(v) > 0 {
		req.SetQueryParameter("timeoutInSec", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(string)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
