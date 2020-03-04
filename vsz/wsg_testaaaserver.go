package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGTestAAAServerService struct {
	apiClient *APIClient
}

func NewWSGTestAAAServerService(c *APIClient) *WSGTestAAAServerService {
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
func (s *WSGTestAAAServerService) AddSystemAaaTest(ctx context.Context, body *WSGAAATestAuthenticationServer) (*WSGAAATestAAAServerResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddSystemAaaTest), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAAATestAAAServerResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
