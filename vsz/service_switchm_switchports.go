package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchPortsService struct {
	client *Client
}

func NewSwitchMSwitchPortsService(client *Client) *SwitchMSwitchPortsService {
	s := new(SwitchMSwitchPortsService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortsService() *SwitchMSwitchPortsService {
	serv := new(SwitchMSwitchPortsService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchPortsService) AddSwitchPortsDetails(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.PortDetailsQueryResultList, error) {
}

func (s *SwitchMSwitchPortsService) AddSwitchPortsSummary(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchPortsSummaryQueryResultList, error) {
}
