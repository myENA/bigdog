package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
)

type SwitchMSwitchGroupService struct {
    c *Client
}

func NewSwitchMSwitchGroupService (c *Client) *SwitchMSwitchGroupService {
    s := new(SwitchMSwitchGroupService)
    s.c = c
    return s
}

