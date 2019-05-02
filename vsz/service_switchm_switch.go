package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchService struct {
    c *Client
}

func NewSwitchMSwitchService (c *Client) *SwitchMSwitchService {
    s := new(SwitchMSwitchService)
    s.c = c
    return s
}

