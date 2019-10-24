package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/portsettings"
)

type SwitchMSwitchPortSettingsService struct {
	client *Client
}

func NewSwitchMSwitchPortSettingsService(client *Client) *SwitchMSwitchPortSettingsService {
	s := new(SwitchMSwitchPortSettingsService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortSettingsService() *SwitchMSwitchPortSettingsService {
	serv := new(SwitchMSwitchPortSettingsService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchPortSettingsService) AddPortSettingsBulk(ctx context.Context, body *portsettings.CreateBulk) (*portsettings.EmptyResult, error) {
}

func (s *SwitchMSwitchPortSettingsService) FindPortSettings(ctx context.Context) (*portsettings.PortSettingsQueryResult, error) {
}

func (s *SwitchMSwitchPortSettingsService) FindPortSettingsById(ctx context.Context, pId string) (*portsettings.PortSettings, error) {
}

func (s *SwitchMSwitchPortSettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*portsettings.PortSettingsQueryResult, error) {
}
