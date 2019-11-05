package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portaldetectionprofile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGPortalDetectionandSuppressionProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGPortalDetectionandSuppressionProfileService(c *APIClient) *WSGPortalDetectionandSuppressionProfileService {
	s := new(WSGPortalDetectionandSuppressionProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGPortalDetectionandSuppressionProfileService() *WSGPortalDetectionandSuppressionProfileService {
	return NewWSGPortalDetectionandSuppressionProfileService(ss.apiClient)
}

// AddRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to create portal detection and suppression profile.
//
// Request Body:
//	 - body *portaldetectionprofile.CreatePortalDetectionProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) AddRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *portaldetectionprofile.CreatePortalDetectionProfile, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesPortalDetectionProfilesById
//
// Use this API command to delete portal detection and suppression profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to delete multiple portal detection and suppression profiles.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *common.BulkDeleteRequest, pZoneId string) (*common.EmptyResult, error) {
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

// FindRkszonesPortalDetectionProfilesById
//
// Use this API command to get portal detection and suppression profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesById(ctx context.Context, pId string, pZoneId string) (*portaldetectionprofile.PortalDetectionProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalDetectionProfilesByQueryCriteria
//
// Query portal detection and suppression profile with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*portaldetectionprofile.PortalDetectionProfileList, error) {
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

// FindRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to get portal detection and suppression profile list.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, pZoneId string) (*portaldetectionprofile.PortalDetectionProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Request Body:
//	 - body *portaldetectionprofile.CreatePortalDetectionProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) PartialUpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *portaldetectionprofile.CreatePortalDetectionProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
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

// UpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Request Body:
//	 - body *portaldetectionprofile.CreatePortalDetectionProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) UpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *portaldetectionprofile.CreatePortalDetectionProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
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
