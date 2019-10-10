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

func (s *SwitchMSwitchVLANSettingService) FindVlans(ctx context.Context) (*vlanconfig.VlanConfigQueryResult, error) {
}

func (s *SwitchMSwitchVLANSettingService) FindVlansById(ctx context.Context, id string) (*vlanconfig.VlanConfig, error) {
}

func (s *SwitchMSwitchVLANSettingService) FindVlansByQueryCriteria(ctx context.Context) (*vlanconfig.VlanConfigQueryResult, error) {
}
