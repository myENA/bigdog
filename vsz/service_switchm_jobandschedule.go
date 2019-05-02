package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/job"
)

type SwitchMJobandScheduleService struct {
    c *Client
}

func NewSwitchMJobandScheduleService (c *Client) *SwitchMJobandScheduleService {
    s := new(SwitchMJobandScheduleService)
    s.c = c
    return s
}

