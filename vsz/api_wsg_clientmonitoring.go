package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGClientMonitoringService struct {
	apiClient *APIClient
}

func NewWSGClientMonitoringService(c *APIClient) *WSGClientMonitoringService {
	s := new(WSGClientMonitoringService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientMonitoringService() *WSGClientMonitoringService {
	return NewWSGClientMonitoringService(ss.apiClient)
}

type WSGClientMonitoringSummary struct {
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	OfflineCount *int `json:"offlineCount,omitempty"`

	OnlineCount *int `json:"onlineCount,omitempty"`
}

func NewWSGClientMonitoringSummary() *WSGClientMonitoringSummary {
	m := new(WSGClientMonitoringSummary)
	return m
}

// FindClientMonitoringSummaryByQueryCriteria
//
// Get current connected count of Clients
//
// Request Body:
//	 - body *WSGCommonQueryCriteria
func (s *WSGClientMonitoringService) FindClientMonitoringSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteria) (*WSGClientMonitoringSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientMonitoringSummary
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindClientMonitoringSummaryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientMonitoringSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
