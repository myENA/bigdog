package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
)

type SwitchMSwitchConfigurationCloneService struct {
	client *Client
}

func NewSwitchMSwitchConfigurationCloneService(client *Client) *SwitchMSwitchConfigurationCloneService {
	s := new(SwitchMSwitchConfigurationCloneService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationCloneService() *SwitchMSwitchConfigurationCloneService {
	serv := new(SwitchMSwitchConfigurationCloneService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchConfigurationCloneService) AddCloneConfigurationByGroup(ctx context.Context) error {
}

func (s *SwitchMSwitchConfigurationCloneService) UpdateCloneConfiguration(ctx context.Context) error {
}
