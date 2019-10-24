package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/portcapacity"
)

type SwitchMSwitchPortCapacityService struct {
	client *Client
}

func NewSwitchMSwitchPortCapacityService(client *Client) *SwitchMSwitchPortCapacityService {
	s := new(SwitchMSwitchPortCapacityService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortCapacityService() *SwitchMSwitchPortCapacityService {
	serv := new(SwitchMSwitchPortCapacityService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchPortCapacityService) FindPortCapacity(ctx context.Context) (*portcapacity.Result, error) {
}
