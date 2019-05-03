package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/stack"
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

func (s *SwitchMSwitchStackService) FindStackMemberBySerialNumber (ctx context.Context, serialNumber string) (*stack.List, error) {
}

