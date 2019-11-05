package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
	"gopkg.in/go-playground/validator.v9"
)

type WSGTestAAAServerService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGTestAAAServerService(c *APIClient) *WSGTestAAAServerService {
	s := new(WSGTestAAAServerService)
	s.apiClient = c
	s.validate = validator.New()
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
//	 - body *aaa.TestAuthenticationServer
func (s *WSGTestAAAServerService) AddSystemAaaTest(ctx context.Context, body *aaa.TestAuthenticationServer) (*aaa.TestAAAServerResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
