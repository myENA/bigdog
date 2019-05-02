package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMTopReportsService struct {
    client *Client
}

func NewSwitchMTopReportsService (client *Client) *SwitchMTopReportsService {
    s := new(SwitchMTopReportsService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMTopReportsService () *SwitchMTopReportsService {
    serv := new(SwitchMTopReportsService)
    serv.client = ss.client
    return serv
}

