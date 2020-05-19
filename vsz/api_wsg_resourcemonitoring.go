package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGResourceMonitoringService struct {
	apiClient *APIClient
}

func NewWSGResourceMonitoringService(c *APIClient) *WSGResourceMonitoringService {
	s := new(WSGResourceMonitoringService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGResourceMonitoringService() *WSGResourceMonitoringService {
	return NewWSGResourceMonitoringService(ss.apiClient)
}

// FindResourceMonitoringSummaryByQueryCriteria
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
func (s *WSGResourceMonitoringService) FindResourceMonitoringSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string) (*WSGCommonMonitoringSummary, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, resource, "required,oneof"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindResourceMonitoringSummaryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonMonitoringSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
