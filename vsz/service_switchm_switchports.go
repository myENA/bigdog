package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchPortsService struct {
    c *Client
}

func NewSwitchMSwitchPortsService (c *Client) *SwitchMSwitchPortsService {
    s := new(SwitchMSwitchPortsService)
    s.c = c
    return s
}

