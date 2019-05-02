package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/configbackup"
)

type SwitchMSwitchConfigurationService struct {
    c *Client
}

func NewSwitchMSwitchConfigurationService (c *Client) *SwitchMSwitchConfigurationService {
    s := new(SwitchMSwitchConfigurationService)
    s.c = c
    return s
}

