package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/ipconfig"
)

type SwitchMSwitchIPSettingService struct {
	client *Client
}

func NewSwitchMSwitchIPSettingService(client *Client) *SwitchMSwitchIPSettingService {
	s := new(SwitchMSwitchIPSettingService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchIPSettingService() *SwitchMSwitchIPSettingService {
	serv := new(SwitchMSwitchIPSettingService)
	serv.client = ss.client
	return serv
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
