package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchPortsService struct {
    client *Client
}

func NewSwitchMSwitchPortsService (client *Client) *SwitchMSwitchPortsService {
    s := new(SwitchMSwitchPortsService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchPortsService () *SwitchMSwitchPortsService {
    serv := new(SwitchMSwitchPortsService)
    serv.client = ss.client
    return serv
}

