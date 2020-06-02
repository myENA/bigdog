package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDataplaneoperationalService struct {
	apiClient *APIClient
}

func NewWSGDataplaneoperationalService(c *APIClient) *WSGDataplaneoperationalService {
	s := new(WSGDataplaneoperationalService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDataplaneoperationalService() *WSGDataplaneoperationalService {
	return NewWSGDataplaneoperationalService(ss.apiClient)
}

// AddDpsSwitchoverCluster
//
// Use this API command to switchover DP to another cluster
//
// Request Body:
//	 - body *WSGDPSwitchoverDp
func (s *WSGDataplaneoperationalService) AddDpsSwitchoverCluster(ctx context.Context, body *WSGDPSwitchoverDp) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpsSwitchoverCluster, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
