package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGTrafficAnalysisService struct {
	apiClient *VSZClient
}

func NewWSGTrafficAnalysisService(c *VSZClient) *WSGTrafficAnalysisService {
	s := new(WSGTrafficAnalysisService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTrafficAnalysisService() *WSGTrafficAnalysisService {
	return NewWSGTrafficAnalysisService(ss.apiClient)
}

// WSGTrafficAnalysisResults
//
// Definition: trafficanalysis_trafficAnalysisResults
//
// Results from a Traffic Analysis query.  Values may be float, integer, or string in type
type WSGTrafficAnalysisResults struct {
	// FirstIndex
	// Can probably be ignored
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Can probably be ignored
	HasMore *bool `json:"hasMore,omitempty"`

	List []interface{} `json:"list,omitempty"`

	// RawDataTotalCount
	// Always seems to be zero, not sure.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Number of items in "list"
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGTrafficAnalysisResults() *WSGTrafficAnalysisResults {
	m := new(WSGTrafficAnalysisResults)
	return m
}

// FindTrafficAnalysisAggregatesByQueryCriteria
//
// Operation ID: findTrafficAnalysisAggregatesByQueryCriteria
//
// View traffic analysis aggregates
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[clients,usage]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisAggregatesByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisAggregatesGroupedByQueryCriteria
//
// Operation ID: findTrafficAnalysisAggregatesGroupedByQueryCriteria
//
// View grouped traffic analysis aggregates
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[clients,usage]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisAggregatesGroupedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisAggregatesGroupedByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisClientResourceByQueryCriteria
//
// Operation ID: findTrafficAnalysisClientResourceByQueryCriteria
//
// View client resource analytics
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[usage,os,app]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisClientResourceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisClientResourceByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindTrafficAnalysisLineRatesByQueryCriteria
//
// Operation ID: findTrafficAnalysisLineRatesByQueryCriteria
//
// View line rate aggregation data
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[usage,clients,aggclients]
// - source string
//		- required
//		- oneof:[ap,wlan]
func (s *WSGTrafficAnalysisService) FindTrafficAnalysisLineRatesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, source string, mutators ...RequestMutator) (*WSGTrafficAnalysisResults, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGTrafficAnalysisResults
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindTrafficAnalysisLineRatesByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	req.SetPathParameter("source", source)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGTrafficAnalysisResults()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
