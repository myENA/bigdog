package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/configbackup"
)

type SwitchMSwitchConfigurationService struct {
    client *Client
}

func NewSwitchMSwitchConfigurationService (client *Client) *SwitchMSwitchConfigurationService {
    s := new(SwitchMSwitchConfigurationService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationService () *SwitchMSwitchConfigurationService {
    serv := new(SwitchMSwitchConfigurationService)
    serv.client = ss.client
    return serv
}

