package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
)

type WSGGlobalreferenceService struct {
	apiClient *APIClient
}

func NewWSGGlobalreferenceService(c *APIClient) *WSGGlobalreferenceService {
	s := new(WSGGlobalreferenceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGlobalreferenceService() *WSGGlobalreferenceService {
	return NewWSGGlobalreferenceService(ss.apiClient)
}

// FindGlobalSettingsFriendlyNameLang
//
// Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
func (s *WSGGlobalreferenceService) FindGlobalSettingsFriendlyNameLang(ctx context.Context) (*WSGSystemFriendlyNameLangList, error) {
	var (
		resp *WSGSystemFriendlyNameLangList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindGlobalSettingsPortalLang
//
// Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
func (s *WSGGlobalreferenceService) FindGlobalSettingsPortalLang(ctx context.Context) (*WSGSystemPortalLangList, error) {
	var (
		resp *WSGSystemPortalLangList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}
