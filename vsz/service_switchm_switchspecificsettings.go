package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/specificsettings"
)

type SwitchMSwitchSpecificSettingsService struct {
	client *Client
}

func NewSwitchMSwitchSpecificSettingsService(client *Client) *SwitchMSwitchSpecificSettingsService {
	s := new(SwitchMSwitchSpecificSettingsService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchSpecificSettingsService() *SwitchMSwitchSpecificSettingsService {
	serv := new(SwitchMSwitchSpecificSettingsService)
	serv.client = ss.client
	return serv
}

// FindSpecificSettings
//
// Use this API command to Retrieve all Specific Setting list.
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettings(ctx context.Context) (*specificsettings.SpecificSettingsAllResult, error) {
}

// FindSpecificSettingsById
//
// Use this API command to Retrieve Specific Settings.
func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, pId string) (*specificsettings.SpecificSettings, error) {
}
