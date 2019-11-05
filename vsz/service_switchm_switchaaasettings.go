package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aaasettings"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchAAASettingsService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchAAASettingsService(c *APIClient) *SwitchMSwitchAAASettingsService {
	s := new(SwitchMSwitchAAASettingsService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchAAASettingsService() *SwitchMSwitchAAASettingsService {
	return NewSwitchMSwitchAAASettingsService(ss.apiClient)
}

// FindAaaSettings
//
// Use this API command to retrieve the AAA settings.
func (s *SwitchMSwitchAAASettingsService) FindAaaSettings(ctx context.Context) (*aaasettings.AaaSettings, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAaaSettings
//
// Use this API command to modify the AAA settings.
//
// Request Body:
//	 - body *aaasettings.AaaSettings
func (s *SwitchMSwitchAAASettingsService) UpdateAaaSettings(ctx context.Context, body *aaasettings.AaaSettings) (*aaasettings.EmptyResult, error) {
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
