package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchService struct {
	client *Client
}

func NewSwitchMSwitchService(client *Client) *SwitchMSwitchService {
	s := new(SwitchMSwitchService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchService() *SwitchMSwitchService {
	serv := new(SwitchMSwitchService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchService) AddSwitch(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchQueryResultList, error) {
}

func (s *SwitchMSwitchService) AddSwitchSnmpSyncedSwitch(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchQueryResultList, error) {
}

func (s *SwitchMSwitchService) AddSwitchViewDetails(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.StackMemberQueryResult, error) {
}

func (s *SwitchMSwitchService) FindSwitchById(ctx context.Context, pId string) (*switchmswitch.NetworkSwitch, error) {
}

func (s *SwitchMSwitchService) FindSwitchFirmwareBySwitchId(ctx context.Context, pSwitchId string) (*switchmswitch.FirmwareHistoryQueryResultList, error) {
}

func (s *SwitchMSwitchService) UpdateSwitchMoveByDestinationSwitchGroupId(ctx context.Context, body switchmswitch.SwitchIdList, pDestinationSwitchGroupId string) error {
}
