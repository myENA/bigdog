package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/veconfig"
)

type SwitchMSwitchVESettingService struct {
	client *Client
}

func NewSwitchMSwitchVESettingService(client *Client) *SwitchMSwitchVESettingService {
	s := new(SwitchMSwitchVESettingService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchVESettingService() *SwitchMSwitchVESettingService {
	serv := new(SwitchMSwitchVESettingService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchVESettingService) FindVeConfigs(ctx context.Context) (*veconfig.List, error) {
}

func (s *SwitchMSwitchVESettingService) FindVeConfigsById(ctx context.Context, id string) (*veconfig.VeConfig, error) {
}

func (s *SwitchMSwitchVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context) (*veconfig.List, error) {
}
