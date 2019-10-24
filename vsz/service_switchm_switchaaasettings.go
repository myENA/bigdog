package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aaasettings"
)

type SwitchMSwitchAAASettingsService struct {
	client *Client
}

func NewSwitchMSwitchAAASettingsService(client *Client) *SwitchMSwitchAAASettingsService {
	s := new(SwitchMSwitchAAASettingsService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchAAASettingsService() *SwitchMSwitchAAASettingsService {
	serv := new(SwitchMSwitchAAASettingsService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchAAASettingsService) FindAaaSettings(ctx context.Context) (*aaasettings.AaaSettings, error) {
}
