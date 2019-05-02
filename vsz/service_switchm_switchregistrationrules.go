package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/registration"
)

type SwitchMSwitchRegistrationRulesService struct {
    c *Client
}

func NewSwitchMSwitchRegistrationRulesService (c *Client) *SwitchMSwitchRegistrationRulesService {
    s := new(SwitchMSwitchRegistrationRulesService)
    s.c = c
    return s
}

