package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDataPlaneOperationalService struct {
	apiClient *VSZClient
}

func NewWSGDataPlaneOperationalService(c *VSZClient) *WSGDataPlaneOperationalService {
	s := new(WSGDataPlaneOperationalService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDataPlaneOperationalService() *WSGDataPlaneOperationalService {
	return NewWSGDataPlaneOperationalService(ss.apiClient)
}

// AddDpsSwitchoverCluster
//
// Operation ID: addDpsSwitchoverCluster
//
// Use this API command to switchover DP to another cluster
//
// Request Body:
//	 - body *WSGDPSwitchoverDp
func (s *WSGDataPlaneOperationalService) AddDpsSwitchoverCluster(ctx context.Context, body *WSGDPSwitchoverDp, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpsSwitchoverCluster, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
