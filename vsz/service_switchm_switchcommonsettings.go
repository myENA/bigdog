package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/dnsconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchCommonSettingsService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchCommonSettingsService(c *APIClient) *SwitchMSwitchCommonSettingsService {
	s := new(SwitchMSwitchCommonSettingsService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchCommonSettingsService() *SwitchMSwitchCommonSettingsService {
	return NewSwitchMSwitchCommonSettingsService(ss.apiClient)
}

// AddDnsConfig
//
// Use this API command to Create DNS Config.
//
// Request Body:
//	 - body *dnsconfig.CreateDnsConfig
func (s *SwitchMSwitchCommonSettingsService) AddDnsConfig(ctx context.Context, body *dnsconfig.CreateDnsConfig) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteDnsConfigBySwitchGroupId
//
// Use this API command to Delete DNS Config.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchCommonSettingsService) DeleteDnsConfigBySwitchGroupId(ctx context.Context, pSwitchGroupId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindDnsConfigBySwitchGroupId
//
// Use this API command to Retrieve DNS Config.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchCommonSettingsService) FindDnsConfigBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*dnsconfig.DnsConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}
