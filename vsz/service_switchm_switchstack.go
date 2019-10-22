package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/stack"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchStackService struct {
    client *Client
}

func NewSwitchMSwitchStackService (client *Client) *SwitchMSwitchStackService {
    s := new(SwitchMSwitchStackService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchStackService () *SwitchMSwitchStackService {
    serv := new(SwitchMSwitchStackService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchStackService) AddStack (ctx context.Context) (stack.AuditIdList, error) {
}

func (s *SwitchMSwitchStackService) FindStackBySwitchId (ctx context.Context, switchId string) (*stack.StackConfig, error) {
}

func (s *SwitchMSwitchStackService) FindStackMemberBySerialNumber (ctx context.Context, serialNumber string) (*stack.List, error) {
}

