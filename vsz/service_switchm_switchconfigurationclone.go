package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchConfigurationCloneService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchConfigurationCloneService(c *APIClient) *SwitchMSwitchConfigurationCloneService {
	s := new(SwitchMSwitchConfigurationCloneService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationCloneService() *SwitchMSwitchConfigurationCloneService {
	return NewSwitchMSwitchConfigurationCloneService(ss.apiClient)
}

// AddCloneConfiguration
//
// Use this API command to Get Switch Config.
//
// Request Body:
//	 - body *group.GetConfigBySwitch
func (s *SwitchMSwitchConfigurationCloneService) AddCloneConfiguration(ctx context.Context, body *group.GetConfigBySwitch) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// AddCloneConfigurationByGroup
//
// Use this API command to Clone Switch Group Config.
//
// Request Body:
//	 - body *group.CloneConfigByGroup
func (s *SwitchMSwitchConfigurationCloneService) AddCloneConfigurationByGroup(ctx context.Context, body *group.CloneConfigByGroup) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// UpdateCloneConfiguration
//
// Use this API command to Clone Switch Config.
//
// Request Body:
//	 - body *group.CloneConfigBySwitch
func (s *SwitchMSwitchConfigurationCloneService) UpdateCloneConfiguration(ctx context.Context, body *group.CloneConfigBySwitch) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}
