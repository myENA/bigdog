package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMPortsService struct {
	apiClient *APIClient
}

func NewSwitchMPortsService(c *APIClient) *SwitchMPortsService {
	s := new(SwitchMPortsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortsService() *SwitchMPortsService {
	return NewSwitchMPortsService(ss.apiClient)
}

// AddSwitchPortsDetails
//
// Use this API command to retrieve all the switch ports and its details currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsDetails(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchPortDetailsQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddSwitchPortsSummary
//
// Use this API command to retrieve ports summary based on status, speed of a switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsSummary(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchPortsSummaryQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
