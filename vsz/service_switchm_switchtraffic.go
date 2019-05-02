package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/traffic"
)

type SwitchMSwitchTrafficService struct {
    client *Client
}

func NewSwitchMSwitchTrafficService (client *Client) *SwitchMSwitchTrafficService {
    s := new(SwitchMSwitchTrafficService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchTrafficService () *SwitchMSwitchTrafficService {
    serv := new(SwitchMSwitchTrafficService)
    serv.client = ss.client
    return serv
}

