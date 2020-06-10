package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGResourceHealthService struct {
	apiClient *VSZClient
}

func NewWSGResourceHealthService(c *VSZClient) *WSGResourceHealthService {
	s := new(WSGResourceHealthService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGResourceHealthService() *WSGResourceHealthService {
	return NewWSGResourceHealthService(ss.apiClient)
}

// FindResourceHealthSummaryByQueryCriteria
//
// Get current connected count of requested resource
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[client,ap]
func (s *WSGResourceHealthService) FindResourceHealthSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, mutators ...RequestMutator) (*WSGCommonMonitoringSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonMonitoringSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindResourceHealthSummaryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonMonitoringSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
