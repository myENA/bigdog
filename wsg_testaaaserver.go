package bigdog

// API Version: v9_0

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
// Request Body:
//	 - body *WSGAAATestAuthenticationServer
func (s *WSGTestAAAServerService) AddSystemAaaTest(ctx context.Context, body *WSGAAATestAuthenticationServer, mutators ...RequestMutator) (*WSGAAATestAAAServerResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAATestAAAServerResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemAaaTest, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAATestAAAServerResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
