package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/job"
)

type SwitchMJobandScheduleService struct {
    client *Client
}

func NewSwitchMJobandScheduleService (client *Client) *SwitchMJobandScheduleService {
    s := new(SwitchMJobandScheduleService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMJobandScheduleService () *SwitchMJobandScheduleService {
    serv := new(SwitchMJobandScheduleService)
    serv.client = ss.client
    return serv
}

