package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGTestAAAServerService struct {
	apiClient *VSZClient
}

func NewWSGTestAAAServerService(c *VSZClient) *WSGTestAAAServerService {
	s := new(WSGTestAAAServerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTestAAAServerService() *WSGTestAAAServerService {
	return NewWSGTestAAAServerService(ss.apiClient)
}

// AddSystemAaaTest
//
// Use this API command to test AAA server.
//
// Operation ID: addSystemAaaTest
// Operation path: /system/aaa/test
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAAATestAuthenticationServer
func (s *WSGTestAAAServerService) AddSystemAaaTest(ctx context.Context, body *WSGAAATestAuthenticationServer, mutators ...RequestMutator) (*WSGAAATestAAAServerResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAATestAAAServerResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSystemAaaTest, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAAATestAAAServerResultAPIResponse), err
}
