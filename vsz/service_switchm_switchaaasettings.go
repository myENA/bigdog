package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aaasettings"
)

type SwitchMSwitchAAASettingsService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchAAASettingsService(c *APIClient) *SwitchMSwitchAAASettingsService {
	s := new(SwitchMSwitchAAASettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchAAASettingsService() *SwitchMSwitchAAASettingsService {
	serv := new(SwitchMSwitchAAASettingsService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindAaaSettings
//
// Use this API command to retrieve the AAA settings.
func (s *SwitchMSwitchAAASettingsService) FindAaaSettings(ctx context.Context) (*aaasettings.AaaSettings, error) {
}

// UpdateAaaSettings
//
// Use this API command to modify the AAA settings.
//
// Request Body:
//	 - body *aaasettings.AaaSettings
func (s *SwitchMSwitchAAASettingsService) UpdateAaaSettings(ctx context.Context, body *aaasettings.AaaSettings) (*aaasettings.EmptyResult, error) {
}
