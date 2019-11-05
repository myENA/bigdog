package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/specificsettings"
)

type SwitchMSwitchSpecificSettingsService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchSpecificSettingsService(c *APIClient) *SwitchMSwitchSpecificSettingsService {
	s := new(SwitchMSwitchSpecificSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchSpecificSettingsService() *SwitchMSwitchSpecificSettingsService {
	serv := new(SwitchMSwitchSpecificSettingsService)
	serv.apiClient = ss.apiClient
	return serv
}

// DeleteSpecificSettingsById
//
// Use this API command to Delete Specific Settings.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) DeleteSpecificSettingsById(ctx context.Context, pId string) error {
}

// FindSpecificSettings
//
// Use this API command to Retrieve all Specific Setting list.
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettings(ctx context.Context) (*specificsettings.SpecificSettingsAllResult, error) {
}

// FindSpecificSettingsById
//
// Use this API command to Retrieve Specific Settings.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, pId string) (*specificsettings.SpecificSettings, error) {
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
}
