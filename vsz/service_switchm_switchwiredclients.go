package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchWiredClientsService struct {
    c *Client
}

func NewSwitchMSwitchWiredClientsService (c *Client) *SwitchMSwitchWiredClientsService {
    s := new(SwitchMSwitchWiredClientsService)
    s.c = c
    return s
}

