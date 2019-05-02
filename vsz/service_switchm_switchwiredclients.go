package vsz

// API Version: v8_0

import (
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

