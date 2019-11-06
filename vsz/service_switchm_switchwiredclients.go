package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchWiredClientsService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchWiredClientsService(c *APIClient) *SwitchMSwitchWiredClientsService {
	s := new(SwitchMSwitchWiredClientsService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchWiredClientsService() *SwitchMSwitchWiredClientsService {
	return NewSwitchMSwitchWiredClientsService(ss.apiClient)
}

// AddSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchWiredClientsService) AddSwitchClients(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.ConnectedDevicesQueryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.ConnectedAPsQueryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
