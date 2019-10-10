package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGHotspot20IdentityProviderProfileService struct {
	client *Client
}

func NewWSGHotspot20IdentityProviderProfileService(client *Client) *WSGHotspot20IdentityProviderProfileService {
	s := new(WSGHotspot20IdentityProviderProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGHotspot20IdentityProviderProfileService() *WSGHotspot20IdentityProviderProfileService {
	serv := new(WSGHotspot20IdentityProviderProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersAccountingsById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGHotspot20IdentityProviderProfileService) DeleteProfilesHs20IdentityprovidersOsuById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20Identityproviders(ctx context.Context) (*profile.Hs20ProviderList, error) {
}

func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersById(ctx context.Context, id string) (*profile.Hs20Provider, error) {
}

func (s *WSGHotspot20IdentityProviderProfileService) FindProfilesHs20IdentityprovidersByQueryCriteria(ctx context.Context) (*profile.Hs20ProviderList, error) {
}
