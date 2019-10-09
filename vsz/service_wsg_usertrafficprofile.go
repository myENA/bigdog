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

func NewWSGUserTrafficProfileService (client *Client) *WSGUserTrafficProfileService {
    s := new(WSGUserTrafficProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGUserTrafficProfileService () *WSGUserTrafficProfileService {
    serv := new(WSGUserTrafficProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGUserTrafficProfileService) AddProfilesUtpCloneById (ctx context.Context, id string) (*profile.ProfileCloneResponse, error) {
}

func (s *WSGUserTrafficProfileService) DeleteProfilesUtpDownlinkRateLimitingById (ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGUserTrafficProfileService) DeleteProfilesUtpUplinkRateLimitingById (ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGUserTrafficProfileService) FindProfilesUtp (ctx context.Context) (*profile.ProfileList, error) {
}

func (s *WSGUserTrafficProfileService) FindProfilesUtpById (ctx context.Context, id string) (*profile.UserTrafficProfile, error) {
}

func (s *WSGUserTrafficProfileService) FindProfilesUtpByQueryCriteria (ctx context.Context) (*profile.UserTrafficProfileList, error) {
}

