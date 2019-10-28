package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/vlanconfig"
)

type SwitchMSwitchVLANSettingService struct {
	client *Client
}

func NewSwitchMSwitchVLANSettingService(client *Client) *SwitchMSwitchVLANSettingService {
	s := new(SwitchMSwitchVLANSettingService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchVLANSettingService() *SwitchMSwitchVLANSettingService {
	serv := new(SwitchMSwitchVLANSettingService)
	serv.client = ss.client
	return serv
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
