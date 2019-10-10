package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/dnsconfig"
)

type SwitchMSwitchCommonSettingsService struct {
	client *Client
}

func NewSwitchMSwitchCommonSettingsService(client *Client) *SwitchMSwitchCommonSettingsService {
	s := new(SwitchMSwitchCommonSettingsService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchCommonSettingsService() *SwitchMSwitchCommonSettingsService {
	serv := new(SwitchMSwitchCommonSettingsService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchCommonSettingsService) AddDnsConfig(ctx context.Context) (*common.CreateResult, error) {
}

func (s *SwitchMSwitchCommonSettingsService) FindDnsConfigBySwitchGroupId(ctx context.Context, switchGroupId string) (*dnsconfig.DnsConfig, error) {
}
