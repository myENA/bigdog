package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/health"
)

type SwitchMSwitchHealthService struct {
    client *Client
}

func NewSwitchMSwitchHealthService (client *Client) *SwitchMSwitchHealthService {
    s := new(SwitchMSwitchHealthService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchHealthService () *SwitchMSwitchHealthService {
    serv := new(SwitchMSwitchHealthService)
    serv.client = ss.client
    return serv
}

