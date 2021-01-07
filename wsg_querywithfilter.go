package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: findGgsnGtpcConStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN Connection.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindGgsnGtpcConStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsGgsnGtpcConListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRACStatsGgsnGtpcConList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindGgsnGtpcConStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGRACStatsGgsnGtpcConList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGgsnGtpStatsByQueryCriteria
//
// Operation ID: findGgsnGtpStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN/PGW GTP-C Sessions.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindGgsnGtpStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsGgsnGtpListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRACStatsGgsnGtpList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindGgsnGtpStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGRACStatsGgsnGtpList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRadiusProxyStatsByQueryCriteria
//
// Operation ID: findRadiusProxyStatsByQueryCriteria
//
// Use this API command to retrieve a list of Radius Proxy.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindRadiusProxyStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGRACStatsRadiusProxyListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRACStatsRadiusProxyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRadiusProxyStatsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGRACStatsRadiusProxyList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
