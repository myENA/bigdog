package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/vlanconfig"
)

type SwitchMSwitchVLANSettingService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchVLANSettingService(c *APIClient) *SwitchMSwitchVLANSettingService {
	s := new(SwitchMSwitchVLANSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchVLANSettingService() *SwitchMSwitchVLANSettingService {
	serv := new(SwitchMSwitchVLANSettingService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *vlanconfig.CreateVlanConfig
func (s *SwitchMSwitchVLANSettingService) AddVlans(ctx context.Context, body *vlanconfig.CreateVlanConfig) (*common.CreateResult, error) {
}

// DeleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchVLANSettingService) DeleteVlans(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVLANSettingService) DeleteVlansById(ctx context.Context, pId string) error {
}

// FindVlans
//
// Use this API command to Retrieve the VLAN Config List.
func (s *SwitchMSwitchVLANSettingService) FindVlans(ctx context.Context) (*vlanconfig.VlanConfigQueryResult, error) {
}

// FindVlansById
//
// Use this API command to Retrieve the VLAN Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVLANSettingService) FindVlansById(ctx context.Context, pId string) (*vlanconfig.VlanConfig, error) {
}

// FindVlansByQueryCriteria
//
// Use this API command to Retrieve the VLAN Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchVLANSettingService) FindVlansByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*vlanconfig.VlanConfigQueryResult, error) {
}

// UpdateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *vlanconfig.UpdateVlanConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVLANSettingService) UpdateVlansById(ctx context.Context, body *vlanconfig.UpdateVlanConfig, pId string) (*vlanconfig.EmptyResult, error) {
}
