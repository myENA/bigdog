package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGGlobalreferenceService struct {
    client *Client
}

func NewWSGGlobalreferenceService (client *Client) *WSGGlobalreferenceService {
    s := new(WSGGlobalreferenceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGGlobalreferenceService () *WSGGlobalreferenceService {
    serv := new(WSGGlobalreferenceService)
    serv.client = ss.client
    return serv
}

func (s *WSGGlobalreferenceService) FindGlobalSettingsFriendlyNameLang (ctx context.Context) (*system.FriendlyNameLangList, error) {
}

func (s *WSGGlobalreferenceService) FindGlobalSettingsPortalLang (ctx context.Context) (*system.PortalLangList, error) {
}

