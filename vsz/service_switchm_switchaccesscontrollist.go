package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aclconfig"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
)

type SwitchMSwitchAccessControlListService struct {
	client *Client
}

func NewSwitchMSwitchAccessControlListService(client *Client) *SwitchMSwitchAccessControlListService {
	s := new(SwitchMSwitchAccessControlListService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchAccessControlListService() *SwitchMSwitchAccessControlListService {
	serv := new(SwitchMSwitchAccessControlListService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchAccessControlListService) AddAccessControls(ctx context.Context, body *aclconfig.CreateACLConfig) (*common.CreateResult, error) {
}

func (s *SwitchMSwitchAccessControlListService) FindAccessControlsById(ctx context.Context, pId string) (*aclconfig.ACLConfig, error) {
}

func (s *SwitchMSwitchAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aclconfig.ACLConfigsQueryResult, error) {
}
