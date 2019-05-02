package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/stack"
)

type SwitchMSwitchStackService struct {
    c *Client
}

func NewSwitchMSwitchStackService (c *Client) *SwitchMSwitchStackService {
    s := new(SwitchMSwitchStackService)
    s.c = c
    return s
}

