package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/specificsettings"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchSpecificSettingsService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchSpecificSettingsService(c *APIClient) *SwitchMSwitchSpecificSettingsService {
	s := new(SwitchMSwitchSpecificSettingsService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchSpecificSettingsService() *SwitchMSwitchSpecificSettingsService {
	return NewSwitchMSwitchSpecificSettingsService(ss.apiClient)
}

// DeleteSpecificSettingsById
//
// Use this API command to Delete Specific Settings.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) DeleteSpecificSettingsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSpecificSettings
//
// Use this API command to Retrieve all Specific Setting list.
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettings(ctx context.Context) (*specificsettings.SpecificSettingsAllResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSpecificSettingsById
//
// Use this API command to Retrieve Specific Settings.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, pId string) (*specificsettings.SpecificSettings, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateSpecificSettingsById
//
// Use this API command to Update Specific Settings.
//
// Request Body:
//	 - body *specificsettings.UpdateSpecificSettings
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) UpdateSpecificSettingsById(ctx context.Context, body *specificsettings.UpdateSpecificSettings, pId string) (*specificsettings.EmptyResult, error) {
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
