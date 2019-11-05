package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/ipconfig"
)

type SwitchMSwitchIPSettingService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchIPSettingService(c *APIClient) *SwitchMSwitchIPSettingService {
	s := new(SwitchMSwitchIPSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchIPSettingService() *SwitchMSwitchIPSettingService {
	serv := new(SwitchMSwitchIPSettingService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddIpConfigs
//
// Use this API command to Create IP Config.
//
// Request Body:
//	 - body *ipconfig.Create
func (s *SwitchMSwitchIPSettingService) AddIpConfigs(ctx context.Context, body *ipconfig.Create) (ipconfig.CreateResult, error) {
}

// DeleteIpConfigs
//
// Use this API command to Delete IP Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchIPSettingService) DeleteIpConfigs(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteIpConfigsById
//
// Use this API command to Delete IP Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchIPSettingService) DeleteIpConfigsById(ctx context.Context, pId string) error {
}

// FindIpConfigs
//
// Use this API command to Retrieve IP Config List.
func (s *SwitchMSwitchIPSettingService) FindIpConfigs(ctx context.Context) (*ipconfig.List, error) {
}

// FindIpConfigsById
//
// Use this API command to Retrieve IP Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchIPSettingService) FindIpConfigsById(ctx context.Context, pId string) (*ipconfig.IpConfig, error) {
}

// FindIpConfigsByQueryCriteria
//
// Use this API command to Retrieve IP Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchIPSettingService) FindIpConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*ipconfig.List, error) {
}

// UpdateIpConfigsById
//
// Use this API command to Update IP Config.
//
// Request Body:
//	 - body *ipconfig.Modify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchIPSettingService) UpdateIpConfigsById(ctx context.Context, body *ipconfig.Modify, pId string) (*ipconfig.EmptyResult, error) {
}
