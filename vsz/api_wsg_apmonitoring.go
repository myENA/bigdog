package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAPMonitoringService struct {
	apiClient *APIClient
}

func NewWSGAPMonitoringService(c *APIClient) *WSGAPMonitoringService {
	s := new(WSGAPMonitoringService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAPMonitoringService() *WSGAPMonitoringService {
	return NewWSGAPMonitoringService(ss.apiClient)
}

type WSGAPMonitoringSummary struct {
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	OfflineCount *int `json:"offlineCount,omitempty"`

	OnlineCount *int `json:"onlineCount,omitempty"`
}

func NewWSGAPMonitoringSummary() *WSGAPMonitoringSummary {
	m := new(WSGAPMonitoringSummary)
	return m
}

// FindApMonitoringSummaryByQueryCriteria
//
// Get up / down status of AP's
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAPMonitoringService) FindApMonitoringSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAPMonitoringSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPMonitoringSummary
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApMonitoringSummaryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPMonitoringSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
