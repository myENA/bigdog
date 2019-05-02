package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
)

type SwitchMSwitchGroupService struct {
    client *Client
}

func NewSwitchMSwitchGroupService (client *Client) *SwitchMSwitchGroupService {
    s := new(SwitchMSwitchGroupService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchGroupService () *SwitchMSwitchGroupService {
    serv := new(SwitchMSwitchGroupService)
    serv.client = ss.client
    return serv
}

