package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGClientHealthService struct {
	apiClient *APIClient
}

func NewWSGClientHealthService(c *APIClient) *WSGClientHealthService {
	s := new(WSGClientHealthService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientHealthService() *WSGClientHealthService {
	return NewWSGClientHealthService(ss.apiClient)
}

type WSGClientHealthSummary struct {
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	OfflineCount *int `json:"offlineCount,omitempty"`

	OnlineCount *int `json:"onlineCount,omitempty"`
}

func NewWSGClientHealthSummary() *WSGClientHealthSummary {
	m := new(WSGClientHealthSummary)
	return m
}

// FindClientHealthSummaryByQueryCriteria
//
// Get current connected count of Clients
//
// Request Body:
//	 - body *WSGCommonQueryCriteria
func (s *WSGClientHealthService) FindClientHealthSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteria) (*WSGClientHealthSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientHealthSummary
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindClientHealthSummaryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientHealthSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
