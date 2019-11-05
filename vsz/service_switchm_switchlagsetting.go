package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/lagconfig"
)

type SwitchMSwitchLAGSettingService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchLAGSettingService(c *APIClient) *SwitchMSwitchLAGSettingService {
	s := new(SwitchMSwitchLAGSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchLAGSettingService() *SwitchMSwitchLAGSettingService {
	serv := new(SwitchMSwitchLAGSettingService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddLagConfigs
//
// Use this API command to Create LAG Config.
//
// Request Body:
//	 - body *lagconfig.Create
func (s *SwitchMSwitchLAGSettingService) AddLagConfigs(ctx context.Context, body *lagconfig.Create) (lagconfig.CreateResult, error) {
}

// DeleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchLAGSettingService) DeleteLagConfigs(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchLAGSettingService) DeleteLagConfigsById(ctx context.Context, pId string) error {
}

// FindLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
func (s *SwitchMSwitchLAGSettingService) FindLagConfigs(ctx context.Context) (*lagconfig.List, error) {
}

// FindLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchLAGSettingService) FindLagConfigsById(ctx context.Context, pId string) (*lagconfig.LagConfig, error) {
}

// FindLagConfigsByQueryCriteria
//
// Use this API command to Retrieve LAG Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*lagconfig.List, error) {
}

// UpdateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Request Body:
//	 - body *lagconfig.Modify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *lagconfig.Modify, pId string) (*lagconfig.EmptyResult, error) {
}
