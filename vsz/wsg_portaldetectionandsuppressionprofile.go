package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGPortalDetectionandSuppressionProfileService struct {
	apiClient *APIClient
}

func NewWSGPortalDetectionandSuppressionProfileService(c *APIClient) *WSGPortalDetectionandSuppressionProfileService {
	s := new(WSGPortalDetectionandSuppressionProfileService)
	s.apiClient = c
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
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) AddRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to delete multiple portal detection and suppression profiles.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) DeleteRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesById(ctx context.Context, pId string, pZoneId string) (*WSGPortalDetectionProfile, error) {
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
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGPortalDetectionProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalDetectionProfilesByZoneId
//
// Use this API command to get portal detection and suppression profile list.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByZoneId(ctx context.Context, pZoneId string) (*WSGPortalDetectionProfileList, error) {
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
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) PartialUpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesPortalDetectionProfilesById
//
// Use this API command to modify portal detection and suppression profile by profile's ID.
//
// Request Body:
//	 - body *WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGPortalDetectionandSuppressionProfileService) UpdateRkszonesPortalDetectionProfilesById(ctx context.Context, body *WSGPortalDetectionProfileCreatePortalDetectionProfile, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
