package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/dnsconfig"
)

type SwitchMSwitchCommonSettingsService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchCommonSettingsService(c *APIClient) *SwitchMSwitchCommonSettingsService {
	s := new(SwitchMSwitchCommonSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchCommonSettingsService() *SwitchMSwitchCommonSettingsService {
	serv := new(SwitchMSwitchCommonSettingsService)
	serv.apiClient = ss.apiClient
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

// DeleteDnsConfigBySwitchGroupId
//
// Use this API command to Delete DNS Config.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchCommonSettingsService) DeleteDnsConfigBySwitchGroupId(ctx context.Context, pSwitchGroupId string) error {
}

// FindDnsConfigBySwitchGroupId
//
// Use this API command to Retrieve DNS Config.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchCommonSettingsService) FindDnsConfigBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*dnsconfig.DnsConfig, error) {
}

// UpdateDnsConfigBySwitchGroupId
//
// Use this API command to Update DNS Config.
//
// Request Body:
//	 - body *dnsconfig.UpdateDnsConfig
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchCommonSettingsService) UpdateDnsConfigBySwitchGroupId(ctx context.Context, body *dnsconfig.UpdateDnsConfig, pSwitchGroupId string) (*dnsconfig.EmptyResult, error) {
}
