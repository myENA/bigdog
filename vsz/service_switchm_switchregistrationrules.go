package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/registration"
)

type SwitchMSwitchRegistrationRulesService struct {
    client *Client
}

func NewSwitchMSwitchRegistrationRulesService (client *Client) *SwitchMSwitchRegistrationRulesService {
    s := new(SwitchMSwitchRegistrationRulesService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchRegistrationRulesService () *SwitchMSwitchRegistrationRulesService {
    serv := new(SwitchMSwitchRegistrationRulesService)
    serv.client = ss.client
    return serv
}

