package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/traffic"
)

type SwitchMSwitchTrafficService struct {
    c *Client
}

func NewSwitchMSwitchTrafficService (c *Client) *SwitchMSwitchTrafficService {
    s := new(SwitchMSwitchTrafficService)
    s.c = c
    return s
}

