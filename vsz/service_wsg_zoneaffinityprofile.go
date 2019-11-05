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

type WSGZoneAffinityProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGZoneAffinityProfileService(c *APIClient) *WSGZoneAffinityProfileService {
	s := new(WSGZoneAffinityProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGZoneAffinityProfileService() *WSGZoneAffinityProfileService {
	return NewWSGZoneAffinityProfileService(ss.apiClient)
}

// AddProfilesZoneAffinity
//
// Use this API command to create zone affinity profile.
//
// Request Body:
//	 - body *profile.CreateZoneAffinityProfile
func (s *WSGZoneAffinityProfileService) AddProfilesZoneAffinity(ctx context.Context, body *profile.CreateZoneAffinityProfile) (*common.CreateResult, error) {
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

// DeleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGZoneAffinityProfileService) DeleteProfilesZoneAffinityById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesZoneAffinity
//
// Use this API command to get all zone affinity profiles.
//
// Query Parameters:
// - qVdpId string
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, qVdpId string) (*profile.ZoneAffinityProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, pId string) (*profile.ReturnZoneAffinityProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesZoneAffinityById
//
// Use this API command to modify zone affinity profile.
//
// Request Body:
//	 - body *profile.ModifyZoneAffinityProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGZoneAffinityProfileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *profile.ModifyZoneAffinityProfile, pId string) (*common.EmptyResult, error) {
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
