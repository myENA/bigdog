package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchWiredClientsService struct {
    client *Client
}

func NewSwitchMSwitchWiredClientsService (client *Client) *SwitchMSwitchWiredClientsService {
    s := new(SwitchMSwitchWiredClientsService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchWiredClientsService () *SwitchMSwitchWiredClientsService {
    serv := new(SwitchMSwitchWiredClientsService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchWiredClientsService) AddSwitchClients (ctx context.Context) (*switchmswitch.ConnectedDevicesQueryList, error) {
}

func (s *SwitchMSwitchWiredClientsService) AddSwitchClientsAp (ctx context.Context) (*switchmswitch.ConnectedAPsQueryList, error) {
}

