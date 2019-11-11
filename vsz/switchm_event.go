package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMEventService struct {
	apiClient *APIClient
}

func NewSwitchMEventService(c *APIClient) *SwitchMEventService {
	s := new(SwitchMEventService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMEventService() *SwitchMEventService {
	return NewSwitchMEventService(ss.apiClient)
}

// AddCustomEvent
//
// Use this API command to create a new text pattern event config
//
// Request Body:
//	 - body *SwitchMEventConfigAddEventConfig
func (s *SwitchMEventService) AddCustomEvent(ctx context.Context, body *SwitchMEventConfigAddEventConfig) (*SwitchMEventConfigAddEventConfigResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMEventService) DeleteCustomEventById(ctx context.Context, pId string) (*SwitchMEventConfigDeleteEventConfigResult, error) {
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
func (s *SwitchMEventService) FindCustomEvent(ctx context.Context) (*SwitchMEventConfigGetEventConfigList, error) {
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
func (s *SwitchMEventService) FindCustomEventById(ctx context.Context, pId string) (*SwitchMEventConfig, error) {
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
//	 - body *SwitchMEventConfigUpdateEventConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMEventService) UpdateCustomEventById(ctx context.Context, body *SwitchMEventConfigUpdateEventConfig, pId string) (*SwitchMEventConfigUpdateEventConfigResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
