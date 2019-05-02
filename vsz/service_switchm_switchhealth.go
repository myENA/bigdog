package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/health"
)

type SwitchMSwitchHealthService struct {
    c *Client
}

func NewSwitchMSwitchHealthService (c *Client) *SwitchMSwitchHealthService {
    s := new(SwitchMSwitchHealthService)
    s.c = c
    return s
}

