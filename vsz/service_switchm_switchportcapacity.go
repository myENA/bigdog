package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/portcapacity"
)

type SwitchMSwitchPortCapacityService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchPortCapacityService(c *APIClient) *SwitchMSwitchPortCapacityService {
	s := new(SwitchMSwitchPortCapacityService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortCapacityService() *SwitchMSwitchPortCapacityService {
	serv := new(SwitchMSwitchPortCapacityService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindPortCapacity
//
// Use this API command to Retrieve Switch Port Capacity List.
func (s *SwitchMSwitchPortCapacityService) FindPortCapacity(ctx context.Context) (*portcapacity.Result, error) {
}
