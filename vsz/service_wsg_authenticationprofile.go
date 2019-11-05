package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAuthenticationProfileService struct {
	apiClient *APIClient
}

func NewWSGAuthenticationProfileService(c *APIClient) *WSGAuthenticationProfileService {
	s := new(WSGAuthenticationProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationProfileService() *WSGAuthenticationProfileService {
	serv := new(WSGAuthenticationProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesAuth
//
// Use this API command to create a new authentication profile.
//
// Request Body:
//	 - body *profile.CreateAuthenticationProfile
func (s *WSGAuthenticationProfileService) AddProfilesAuth(ctx context.Context, body *profile.CreateAuthenticationProfile) (*common.CreateResult, error) {
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

// DeleteProfilesAuth
//
// Use this API command to delete a list of authentication profile.
//
// Request Body:
//	 - body *profile.DeleteBulkAuthenticationProfile
func (s *WSGAuthenticationProfileService) DeleteProfilesAuth(ctx context.Context, body *profile.DeleteBulkAuthenticationProfile) (*common.EmptyResult, error) {
}

// DeleteProfilesAuthById
//
// Use this API command to delete an authentication profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationProfileService) DeleteProfilesAuthById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateProfilesAuthById
//
// Use this API command to modify the basic information of an authentication profile.
//
// Request Body:
//	 - body *profile.ModifyAuthenticationProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationProfileService) PartialUpdateProfilesAuthById(ctx context.Context, body *profile.ModifyAuthenticationProfile, pId string) (*common.EmptyResult, error) {
}
