package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGLBSprofileService struct {
	apiClient *APIClient
}

func NewWSGLBSprofileService(c *APIClient) *WSGLBSprofileService {
	s := new(WSGLBSprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLBSprofileService() *WSGLBSprofileService {
	serv := new(WSGLBSprofileService)
	serv.apiClient = ss.apiClient
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

// DeleteProfilesLbs
//
// Delete multiple LBS profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGLBSprofileService) DeleteProfilesLbs(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteProfilesLbsById
//
// Delete LBS profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGLBSprofileService) DeleteProfilesLbsById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesLbsById
//
// Retrieve LBS profile.
//
// Path Parameters:
// - pId string
//		- required
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

// PartialUpdateProfilesLbsById
//
// Update LBS profile.
//
// Request Body:
//	 - body *profile.LbsProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGLBSprofileService) PartialUpdateProfilesLbsById(ctx context.Context, body *profile.LbsProfile, pId string) (*common.EmptyResult, error) {
}
