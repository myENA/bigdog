package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGTestaaaserverService struct {
	apiClient *APIClient
}

func NewWSGTestaaaserverService(c *APIClient) *WSGTestaaaserverService {
	s := new(WSGTestaaaserverService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTestaaaserverService() *WSGTestaaaserverService {
	return NewWSGTestaaaserverService(ss.apiClient)
}

// AddSystemAaaTest
//
// Use this API command to test AAA server.
//
// Request Body:
//	 - body *WSGAAATestAuthenticationServer
func (s *WSGTestaaaserverService) AddSystemAaaTest(ctx context.Context, body *WSGAAATestAuthenticationServer) (*WSGAAATestAAAServerResult, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAAATestAAAServerResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
