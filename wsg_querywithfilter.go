package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGQueryWithFilterService struct {
	apiClient *VSZClient
}

func NewWSGQueryWithFilterService(c *VSZClient) *WSGQueryWithFilterService {
	s := new(WSGQueryWithFilterService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGQueryWithFilterService() *WSGQueryWithFilterService {
	return NewWSGQueryWithFilterService(ss.apiClient)
}

// FindGgsnGtpcConStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN Connection.
//
// Operation ID: findGgsnGtpcConStatsByQueryCriteria
// Operation path: /query/ggsnGtpcCon/stats
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindGgsnGtpcConStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsGgsnGtpcConListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGRACStatsGgsnGtpcConListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindGgsnGtpcConStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRACStatsGgsnGtpcConListAPIResponse), err
}

// FindGgsnGtpStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN/PGW GTP-C Sessions.
//
// Operation ID: findGgsnGtpStatsByQueryCriteria
// Operation path: /query/ggsnGtp/stats
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindGgsnGtpStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsGgsnGtpListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGRACStatsGgsnGtpListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindGgsnGtpStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRACStatsGgsnGtpListAPIResponse), err
}

// FindRadiusProxyStatsByQueryCriteria
//
// Use this API command to retrieve a list of Radius Proxy.
//
// Operation ID: findRadiusProxyStatsByQueryCriteria
// Operation path: /query/radiusProxy/stats
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindRadiusProxyStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsRadiusProxyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGRACStatsRadiusProxyListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindRadiusProxyStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGRACStatsRadiusProxyListAPIResponse), err
}
