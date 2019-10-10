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

func (s *SwitchMSwitchAAAServersService) FindAaaServersAdmin(ctx context.Context) (*aaaservers.AaaServersQueryResult, error) {
}

func (s *SwitchMSwitchAAAServersService) FindAaaServersAdminById(ctx context.Context, id string) (*aaaservers.AAAServer, error) {
}
