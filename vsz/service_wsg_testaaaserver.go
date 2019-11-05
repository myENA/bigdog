package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
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
	serv := new(WSGTestAAAServerService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddSystemAaaTest
//
// Use this API command to test AAA server.
//
// Request Body:
//	 - body *aaa.TestAuthenticationServer
func (s *WSGTestAAAServerService) AddSystemAaaTest(ctx context.Context, body *aaa.TestAuthenticationServer) (*aaa.TestAAAServerResult, error) {
}
