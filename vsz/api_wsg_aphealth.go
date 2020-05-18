package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAPHealthService struct {
	apiClient *APIClient
}

func NewWSGAPHealthService(c *APIClient) *WSGAPHealthService {
	s := new(WSGAPHealthService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAPHealthService() *WSGAPHealthService {
	return NewWSGAPHealthService(ss.apiClient)
}

type WSGAPHealthSummary struct {
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	OfflineCount *int `json:"offlineCount,omitempty"`

	OnlineCount *int `json:"onlineCount,omitempty"`
}

func NewWSGAPHealthSummary() *WSGAPHealthSummary {
	m := new(WSGAPHealthSummary)
	return m
}

// FindApHealthSummaryByQueryCriteria
//
// Get up / down status of AP's
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAPHealthService) FindApHealthSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAPHealthSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPHealthSummary
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApHealthSummaryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPHealthSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
