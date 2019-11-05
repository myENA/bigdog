package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
)

type SwitchMSwitchConfigurationCloneService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchConfigurationCloneService(c *APIClient) *SwitchMSwitchConfigurationCloneService {
	s := new(SwitchMSwitchConfigurationCloneService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationCloneService() *SwitchMSwitchConfigurationCloneService {
	serv := new(SwitchMSwitchConfigurationCloneService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddCloneConfiguration
//
// Use this API command to Get Switch Config.
//
// Request Body:
//	 - body *group.GetConfigBySwitch
func (s *SwitchMSwitchConfigurationCloneService) AddCloneConfiguration(ctx context.Context, body *group.GetConfigBySwitch) error {
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
