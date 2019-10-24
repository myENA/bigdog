package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/lagconfig"
)

type SwitchMSwitchLAGSettingService struct {
	client *Client
}

func NewSwitchMSwitchLAGSettingService(client *Client) *SwitchMSwitchLAGSettingService {
	s := new(SwitchMSwitchLAGSettingService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchLAGSettingService() *SwitchMSwitchLAGSettingService {
	serv := new(SwitchMSwitchLAGSettingService)
	serv.client = ss.client
	return serv
}

// FindLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
func (s *SwitchMSwitchLAGSettingService) FindLagConfigs(ctx context.Context) (*lagconfig.List, error) {
}

// FindLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
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
