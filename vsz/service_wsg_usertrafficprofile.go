package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGUserTrafficProfileService struct {
	client *Client
}

func NewWSGUserTrafficProfileService(client *Client) *WSGUserTrafficProfileService {
	s := new(WSGUserTrafficProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGUserTrafficProfileService() *WSGUserTrafficProfileService {
	serv := new(WSGUserTrafficProfileService)
	serv.client = ss.client
	return serv
}

// AddProfilesUtpCloneById
//
// Use this API command to copy a traffic profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
func (s *WSGUserTrafficProfileService) AddProfilesUtpCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
}

// DeleteProfilesUtpDownlinkRateLimitingById
//
// Use this API command to disable downlink rate limiting of user traffic profile.
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpDownlinkRateLimitingById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteProfilesUtpUplinkRateLimitingById
//
// Use this API command to disable uplink rateLimiting of user traffic profile.
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpUplinkRateLimitingById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesUtp
//
// Use this API command to retrieve a list of user traffic profile.
func (s *WSGUserTrafficProfileService) FindProfilesUtp(ctx context.Context, qIndex string, qListSize string) (*profile.ProfileList, error) {
}

// FindProfilesUtpById
//
// Use this API command to retrieve an user traffic profile.
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
