package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/lagconfig"
)

type SwitchMSwitchLAGSettingService struct {
    client *Client
}

func NewSwitchMSwitchLAGSettingService (client *Client) *SwitchMSwitchLAGSettingService {
    s := new(SwitchMSwitchLAGSettingService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchLAGSettingService () *SwitchMSwitchLAGSettingService {
    serv := new(SwitchMSwitchLAGSettingService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchLAGSettingService) FindLagConfigs (ctx context.Context) (*lagconfig.List, error) {
}

func (s *SwitchMSwitchLAGSettingService) FindLagConfigsById (ctx context.Context, id string) (*lagconfig.LagConfig, error) {
}

func (s *SwitchMSwitchLAGSettingService) FindLagConfigsByQueryCriteria (ctx context.Context) (*lagconfig.List, error) {
}

