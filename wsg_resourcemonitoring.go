package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGResourceMonitoringService struct {
	apiClient *VSZClient
}

func NewWSGResourceMonitoringService(c *VSZClient) *WSGResourceMonitoringService {
	s := new(WSGResourceMonitoringService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGResourceMonitoringService() *WSGResourceMonitoringService {
	return NewWSGResourceMonitoringService(ss.apiClient)
}

// FindResourceMonitoringSummaryByQueryCriteria
//
// Operation ID: findResourceMonitoringSummaryByQueryCriteria
//
// Retrieve monitoring status of resource
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[client,ap]
func (s *WSGResourceMonitoringService) FindResourceMonitoringSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, mutators ...RequestMutator) (*WSGCommonMonitoringSummaryAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindResourceMonitoringSummaryByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonMonitoringSummary()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
