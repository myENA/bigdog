package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAuthenticationProfileService struct {
	client *Client
}

func NewWSGAuthenticationProfileService(client *Client) *WSGAuthenticationProfileService {
	s := new(WSGAuthenticationProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAuthenticationProfileService() *WSGAuthenticationProfileService {
	serv := new(WSGAuthenticationProfileService)
	serv.client = ss.client
	return serv
}

// AddProfilesAuthCloneById
//
// Use this API command to clone an authentication profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationProfileService) AddProfilesAuthCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
}

// FindProfilesAuth
//
// Use this API command to retrieve a list of authentication profiles.
func (s *WSGAuthenticationProfileService) FindProfilesAuth(ctx context.Context) (*profile.AuthenticationProfileList, error) {
}

// FindProfilesAuthAuthorizationList
//
// Use this API command to retrieve a list of authorization profiles.
//
// Query Parameters:
// - qType string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthorizationList(ctx context.Context, qType string) (*profile.BaseServiceInfoList, error) {
}

// FindProfilesAuthAuthServiceListByQueryCriteria
//
// Use this API command to retrieve a list of authentication service.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthServiceListByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BaseServiceInfoList, error) {
}

// FindProfilesAuthById
//
// Use this API command to retrieve an authentication profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthById(ctx context.Context, pId string) (*profile.AuthenticationProfile, error) {
}

// FindProfilesAuthByQueryCriteria
//
// Use this API command to retrieve a list of authentication profiles by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.AuthenticationProfileList, error) {
}
