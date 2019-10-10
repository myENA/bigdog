package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/ipconfig"
)

type SwitchMSwitchIPSettingService struct {
    client *Client
}

func NewSwitchMSwitchIPSettingService (client *Client) *SwitchMSwitchIPSettingService {
    s := new(SwitchMSwitchIPSettingService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchIPSettingService () *SwitchMSwitchIPSettingService {
    serv := new(SwitchMSwitchIPSettingService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchIPSettingService) FindIpConfigs (ctx context.Context) (*ipconfig.List, error) {
}

func (s *SwitchMSwitchIPSettingService) FindIpConfigsById (ctx context.Context, id string) (*ipconfig.IPConfig, error) {
}

func (s *SwitchMSwitchIPSettingService) FindIpConfigsByQueryCriteria (ctx context.Context) (*ipconfig.List, error) {
}

