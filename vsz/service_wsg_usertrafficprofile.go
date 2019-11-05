package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGUserTrafficProfileService struct {
	apiClient *APIClient
}

func NewWSGUserTrafficProfileService(c *APIClient) *WSGUserTrafficProfileService {
	s := new(WSGUserTrafficProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGUserTrafficProfileService() *WSGUserTrafficProfileService {
	serv := new(WSGUserTrafficProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesUtp
//
// Use this API command to create a new user traffic profile.
//
// Request Body:
//	 - body *profile.CreateUserTrafficProfile
func (s *WSGUserTrafficProfileService) AddProfilesUtp(ctx context.Context, body *profile.CreateUserTrafficProfile) (*common.CreateResult, error) {
}

// AddProfilesUtpCloneById
//
// Use this API command to copy a traffic profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) AddProfilesUtpCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
}

// DeleteProfilesUtp
//
// Use this API command to delete a list of traffic profile.
//
// Request Body:
//	 - body *profile.DeleteBulkUserTrafficProfile
func (s *WSGUserTrafficProfileService) DeleteProfilesUtp(ctx context.Context, body *profile.DeleteBulkUserTrafficProfile) (*common.EmptyResult, error) {
}

// DeleteProfilesUtpById
//
// Use this API command to delete an user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteProfilesUtpDownlinkRateLimitingById
//
// Use this API command to disable downlink rate limiting of user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpDownlinkRateLimitingById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteProfilesUtpUplinkRateLimitingById
//
// Use this API command to disable uplink rateLimiting of user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpUplinkRateLimitingById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesUtp
//
// Use this API command to retrieve a list of user traffic profile.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGUserTrafficProfileService) FindProfilesUtp(ctx context.Context, qIndex string, qListSize string) (*profile.ProfileList, error) {
}

// FindProfilesUtpById
//
// Use this API command to retrieve an user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) FindProfilesUtpById(ctx context.Context, pId string) (*profile.UserTrafficProfile, error) {
}

// FindProfilesUtpByQueryCriteria
//
// Use this API command to retrieve a list of User Traffic Profile by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGUserTrafficProfileService) FindProfilesUtpByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.UserTrafficProfileList, error) {
}

// PartialUpdateProfilesUtpById
//
// Use this API command to modify the basic information of user traffic profile.
//
// Request Body:
//	 - body *profile.ModifyUserTrafficProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) PartialUpdateProfilesUtpById(ctx context.Context, body *profile.ModifyUserTrafficProfile, pId string) (*common.EmptyResult, error) {
}
