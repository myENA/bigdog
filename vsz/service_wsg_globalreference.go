package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGGlobalreferenceService struct {
	client *Client
}

func NewWSGGlobalreferenceService(client *Client) *WSGGlobalreferenceService {
	s := new(WSGGlobalreferenceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGGlobalreferenceService() *WSGGlobalreferenceService {
	serv := new(WSGGlobalreferenceService)
	serv.client = ss.client
	return serv
}

// FindGlobalSettingsFriendlyNameLang
//
// Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
func (s *WSGGlobalreferenceService) FindGlobalSettingsFriendlyNameLang(ctx context.Context) (*system.FriendlyNameLangList, error) {
}

// FindGlobalSettingsPortalLang
//
// Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
func (s *WSGGlobalreferenceService) FindGlobalSettingsPortalLang(ctx context.Context) (*system.PortalLangList, error) {
}
