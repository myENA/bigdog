package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
	"gopkg.in/go-playground/validator.v9"
)

type WSGGlobalreferenceService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGGlobalreferenceService(c *APIClient) *WSGGlobalreferenceService {
	s := new(WSGGlobalreferenceService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGGlobalreferenceService() *WSGGlobalreferenceService {
	return NewWSGGlobalreferenceService(ss.apiClient)
}

// FindGlobalSettingsFriendlyNameLang
//
// Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
func (s *WSGGlobalreferenceService) FindGlobalSettingsFriendlyNameLang(ctx context.Context) (*system.FriendlyNameLangList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindGlobalSettingsPortalLang
//
// Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
func (s *WSGGlobalreferenceService) FindGlobalSettingsPortalLang(ctx context.Context) (*system.PortalLangList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
