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

func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettings(ctx context.Context) (*specificsettings.SpecificSettingsAllResult, error) {
}

func (s *SwitchMSwitchSpecificSettingsService) FindSpecificSettingsById(ctx context.Context, id string) (*specificsettings.SpecificSettings, error) {
}
