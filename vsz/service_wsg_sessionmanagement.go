package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sessionmanagement"
	"gopkg.in/go-playground/validator.v9"
)

type WSGSessionManagementService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGSessionManagementService(c *APIClient) *WSGSessionManagementService {
	s := new(WSGSessionManagementService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGSessionManagementService() *WSGSessionManagementService {
	return NewWSGSessionManagementService(ss.apiClient)
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context) (*sessionmanagement.RuckusSessions, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
