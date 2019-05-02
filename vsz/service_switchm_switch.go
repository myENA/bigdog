package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchService struct {
    client *Client
}

func NewSwitchMSwitchService (client *Client) *SwitchMSwitchService {
    s := new(SwitchMSwitchService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchService () *SwitchMSwitchService {
    serv := new(SwitchMSwitchService)
    serv.client = ss.client
    return serv
}

