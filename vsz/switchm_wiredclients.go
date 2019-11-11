package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMWiredClientsService struct {
	apiClient *APIClient
}

func NewSwitchMWiredClientsService(c *APIClient) *SwitchMWiredClientsService {
	s := new(SwitchMWiredClientsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMWiredClientsService() *SwitchMWiredClientsService {
	return NewSwitchMWiredClientsService(ss.apiClient)
}

// AddSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClients(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchConnectedDevicesQueryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchConnectedAPsQueryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
