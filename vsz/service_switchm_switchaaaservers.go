package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aaaservers"
)

type SwitchMSwitchAAAServersService struct {
	client *Client
}

func NewSwitchMSwitchAAAServersService(client *Client) *SwitchMSwitchAAAServersService {
	s := new(SwitchMSwitchAAAServersService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchAAAServersService() *SwitchMSwitchAAAServersService {
	serv := new(SwitchMSwitchAAAServersService)
	serv.client = ss.client
	return serv
}

// FindAaaServersAdmin
//
// Use this API command to retrieve a list of AAA server.
func (s *SwitchMSwitchAAAServersService) FindAaaServersAdmin(ctx context.Context) (*aaaservers.AaaServersQueryResult, error) {
}

// FindAaaServersAdminById
//
// Use this API command to retrieve a AAA server.
func (s *SwitchMSwitchAAAServersService) FindAaaServersAdminById(ctx context.Context, pId string) (*aaaservers.AAAServer, error) {
}
