package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/veconfig"
)

type SwitchMSwitchVESettingService struct {
	client *Client
}

func NewSwitchMSwitchVESettingService(client *Client) *SwitchMSwitchVESettingService {
	s := new(SwitchMSwitchVESettingService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchVESettingService() *SwitchMSwitchVESettingService {
	serv := new(SwitchMSwitchVESettingService)
	serv.client = ss.client
	return serv
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
