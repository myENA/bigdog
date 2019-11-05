package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/portcapacity"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchPortCapacityService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchPortCapacityService(c *APIClient) *SwitchMSwitchPortCapacityService {
	s := new(SwitchMSwitchPortCapacityService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortCapacityService() *SwitchMSwitchPortCapacityService {
	return NewSwitchMSwitchPortCapacityService(ss.apiClient)
}

// FindPortCapacity
//
// Use this API command to Retrieve Switch Port Capacity List.
func (s *SwitchMSwitchPortCapacityService) FindPortCapacity(ctx context.Context) (*portcapacity.Result, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
