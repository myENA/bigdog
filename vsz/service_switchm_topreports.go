package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMTopReportsService struct {
    c *Client
}

func NewSwitchMTopReportsService (c *Client) *SwitchMTopReportsService {
    s := new(SwitchMTopReportsService)
    s.c = c
    return s
}

