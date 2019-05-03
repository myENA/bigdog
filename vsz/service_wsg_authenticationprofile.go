package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAuthenticationProfileService struct {
    client *Client
}

func NewWSGAuthenticationProfileService (client *Client) *WSGAuthenticationProfileService {
    s := new(WSGAuthenticationProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAuthenticationProfileService () *WSGAuthenticationProfileService {
    serv := new(WSGAuthenticationProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGAuthenticationProfileService) AddProfilesAuthCloneById (ctx context.Context, id string) (*profile.ProfileCloneResponse, error) {
}

func (s *WSGAuthenticationProfileService) FindProfilesAuth (ctx context.Context) (*profile.AuthenticationProfileList, error) {
}

func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthorizationList (ctx context.Context) (*profile.BaseServiceInfoList, error) {
}

func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthServiceListByQueryCriteria (ctx context.Context) (*profile.BaseServiceInfoList, error) {
}

func (s *WSGAuthenticationProfileService) FindProfilesAuthById (ctx context.Context, id string) (*profile.AuthenticationProfile, error) {
}

func (s *WSGAuthenticationProfileService) FindProfilesAuthByQueryCriteria (ctx context.Context) (*profile.AuthenticationProfileList, error) {
}

