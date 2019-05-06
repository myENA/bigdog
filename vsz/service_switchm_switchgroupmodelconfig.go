package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/groupmodelconfig"
)

type SwitchMSwitchGroupModelConfigService struct {
    client *Client
}

func NewSwitchMSwitchGroupModelConfigService (client *Client) *SwitchMSwitchGroupModelConfigService {
    s := new(SwitchMSwitchGroupModelConfigService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchGroupModelConfigService () *SwitchMSwitchGroupModelConfigService {
    serv := new(SwitchMSwitchGroupModelConfigService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchGroupModelConfigService) FindGroupModelConfigsByQueryCriteria (ctx context.Context) (*groupmodelconfig.GroupModelConfigQueryResult, error) {
}

func (s *SwitchMSwitchGroupModelConfigService) UpdateGroupModelConfigsByGroupId (ctx context.Context, groupId string) (*groupmodelconfig.UpdateGroupConfigResultList, error) {
}

