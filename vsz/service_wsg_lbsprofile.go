package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGLBSprofileService struct {
	client *Client
}

func NewWSGLBSprofileService(client *Client) *WSGLBSprofileService {
	s := new(WSGLBSprofileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGLBSprofileService() *WSGLBSprofileService {
	serv := new(WSGLBSprofileService)
	serv.client = ss.client
	return serv
}

// AddProfilesLbs
//
// Create LBS profile.
//
// Request Body:
//	 - body *profile.LbsProfile
func (s *WSGLBSprofileService) AddProfilesLbs(ctx context.Context, body *profile.LbsProfile) (*common.CreateResult, error) {
}

// FindProfilesLbsById
//
// Retrieve LBS profile.
func (s *WSGLBSprofileService) FindProfilesLbsById(ctx context.Context, pId string) (*profile.LbsProfile, error) {
}

// FindProfilesLbsByQueryCriteria
//
// Query LBS profiles.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGLBSprofileService) FindProfilesLbsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.LbsProfileList, error) {
}
