package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/eventconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchEventService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchEventService(c *APIClient) *SwitchMSwitchEventService {
	s := new(SwitchMSwitchEventService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchEventService() *SwitchMSwitchEventService {
	return NewSwitchMSwitchEventService(ss.apiClient)
}

// AddCustomEvent
//
// Use this API command to create a new text pattern event config
//
// Request Body:
//	 - body *eventconfig.AddEventConfig
func (s *SwitchMSwitchEventService) AddCustomEvent(ctx context.Context, body *eventconfig.AddEventConfig) (*eventconfig.AddEventConfigResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchEventService) DeleteCustomEventById(ctx context.Context, pId string) (*eventconfig.DeleteEventConfigResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMSwitchEventService) FindCustomEvent(ctx context.Context) (*eventconfig.GetEventConfigList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchEventService) FindCustomEventById(ctx context.Context, pId string) (*eventconfig.EventConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Request Body:
//	 - body *eventconfig.UpdateEventConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchEventService) UpdateCustomEventById(ctx context.Context, body *eventconfig.UpdateEventConfig, pId string) (*eventconfig.UpdateEventConfigResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}
