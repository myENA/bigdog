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

// AddDnsConfig
//
// Use this API command to Create DNS Config.
//
// Request Body:
//	 - body *dnsconfig.CreateDnsConfig
func (s *SwitchMSwitchCommonSettingsService) AddDnsConfig(ctx context.Context, body *dnsconfig.CreateDnsConfig) (*common.CreateResult, error) {
}

// FindDnsConfigBySwitchGroupId
//
// Use this API command to Retrieve DNS Config.
func (s *SwitchMSwitchCommonSettingsService) FindDnsConfigBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*dnsconfig.DnsConfig, error) {
}
