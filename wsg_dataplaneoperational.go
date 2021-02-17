package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Use this API command to switchover DP to another cluster
//
// Operation ID: addDpsSwitchoverCluster
// Operation path: /dps/switchoverCluster
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDPSwitchoverDp
func (s *WSGDataPlaneOperationalService) AddDpsSwitchoverCluster(ctx context.Context, body *WSGDPSwitchoverDp, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpsSwitchoverCluster, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
