package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
		resp     APIResponse
		err      error

		respFn = newWSGAAATestAAAServerResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAAATestAAAServerResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemAaaTest, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAAATestAAAServerResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAATestAAAServerResultAPIResponse), err
}
