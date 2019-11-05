package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/veconfig"
)

type SwitchMSwitchVESettingService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchVESettingService(c *APIClient) *SwitchMSwitchVESettingService {
	s := new(SwitchMSwitchVESettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchVESettingService() *SwitchMSwitchVESettingService {
	serv := new(SwitchMSwitchVESettingService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddVeConfigs
//
// Use this API command to Create VE Config.
//
// Request Body:
//	 - body *veconfig.Create
func (s *SwitchMSwitchVESettingService) AddVeConfigs(ctx context.Context, body *veconfig.Create) (veconfig.CreateResult, error) {
}

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchVESettingService) DeleteVeConfigs(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVESettingService) DeleteVeConfigsById(ctx context.Context, pId string) error {
}

// FindVeConfigs
//
// Use this API command to Retrieve VE Config List.
func (s *SwitchMSwitchVESettingService) FindVeConfigs(ctx context.Context) (*veconfig.List, error) {
}

// FindVeConfigsById
//
// Use this API command to Retrieve VE Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVESettingService) FindVeConfigsById(ctx context.Context, pId string) (*veconfig.VeConfig, error) {
}

// FindVeConfigsByQueryCriteria
//
// Use this API command to Retrieve VE Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*veconfig.List, error) {
}

// UpdateVeConfigsById
//
// Use this API command to Update VE Config.
//
// Request Body:
//	 - body *veconfig.Modify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVESettingService) UpdateVeConfigsById(ctx context.Context, body *veconfig.Modify, pId string) (*veconfig.EmptyResult, error) {
}
