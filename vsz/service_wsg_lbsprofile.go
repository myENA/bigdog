package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGLBSprofileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGLBSprofileService(c *APIClient) *WSGLBSprofileService {
	s := new(WSGLBSprofileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGLBSprofileService() *WSGLBSprofileService {
	return NewWSGLBSprofileService(ss.apiClient)
}

// AddProfilesLbs
//
// Create LBS profile.
//
// Request Body:
//	 - body *profile.LbsProfile
func (s *WSGLBSprofileService) AddProfilesLbs(ctx context.Context, body *profile.LbsProfile) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteProfilesLbs
//
// Delete multiple LBS profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGLBSprofileService) DeleteProfilesLbs(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteProfilesLbsById
//
// Delete LBS profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGLBSprofileService) DeleteProfilesLbsById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesLbsById
//
// Retrieve LBS profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGLBSprofileService) FindProfilesLbsById(ctx context.Context, pId string) (*profile.LbsProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesLbsByQueryCriteria
//
// Query LBS profiles.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGLBSprofileService) FindProfilesLbsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.LbsProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
