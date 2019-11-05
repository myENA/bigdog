package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sessionmanagement"
)

type WSGSessionManagementService struct {
	apiClient *APIClient
}

func NewWSGSessionManagementService(c *APIClient) *WSGSessionManagementService {
	s := new(WSGSessionManagementService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSessionManagementService() *WSGSessionManagementService {
	serv := new(WSGSessionManagementService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context) (*sessionmanagement.RuckusSessions, error) {
}
