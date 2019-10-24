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

// AddCloneConfigurationByGroup
//
// Use this API command to Clone Switch Group Config.
//
// Request Body:
//	 - body *group.CloneConfigByGroup
func (s *SwitchMSwitchConfigurationCloneService) AddCloneConfigurationByGroup(ctx context.Context, body *group.CloneConfigByGroup) error {
}

// UpdateCloneConfiguration
//
// Use this API command to Clone Switch Config.
//
// Request Body:
//	 - body *group.CloneConfigBySwitch
func (s *SwitchMSwitchConfigurationCloneService) UpdateCloneConfiguration(ctx context.Context, body *group.CloneConfigBySwitch) error {
}
