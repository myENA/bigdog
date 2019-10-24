package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sessionmanagement"
)

type WSGSessionManagementService struct {
	client *Client
}

func NewWSGSessionManagementService(client *Client) *WSGSessionManagementService {
	s := new(WSGSessionManagementService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSessionManagementService() *WSGSessionManagementService {
	serv := new(WSGSessionManagementService)
	serv.client = ss.client
	return serv
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context) (*sessionmanagement.RuckusSessions, error) {
}
