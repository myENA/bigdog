package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	return NewWSGUserTrafficProfileService(ss.apiClient)
}

// AddProfilesUtp
//
// Use this API command to create a new user traffic profile.
//
// Request Body:
//	 - body *WSGProfileCreateUserTrafficProfile
func (s *WSGUserTrafficProfileService) AddProfilesUtp(ctx context.Context, body *WSGProfileCreateUserTrafficProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddProfilesUtpCloneById
//
// Use this API command to copy a traffic profile.
//
// Request Body:
//	 - body *WSGProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) AddProfilesUtpCloneById(ctx context.Context, body *WSGProfileCloneRequest, pId string) (*WSGProfileCloneResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesUtp
//
// Use this API command to delete a list of traffic profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkUserTrafficProfile
func (s *WSGUserTrafficProfileService) DeleteProfilesUtp(ctx context.Context, body *WSGProfileDeleteBulkUserTrafficProfile) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesUtpById
//
// Use this API command to delete an user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesUtpDownlinkRateLimitingById
//
// Use this API command to disable downlink rate limiting of user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpDownlinkRateLimitingById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesUtpUplinkRateLimitingById
//
// Use this API command to disable uplink rateLimiting of user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) DeleteProfilesUtpUplinkRateLimitingById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesUtp
//
// Use this API command to retrieve a list of user traffic profile.
//
// Query Parameters:
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGUserTrafficProfileService) FindProfilesUtp(ctx context.Context, qIndex string, qListSize string) (*WSGProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesUtpById
//
// Use this API command to retrieve an user traffic profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) FindProfilesUtpById(ctx context.Context, pId string) (*WSGProfileUserTrafficProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesUtpByQueryCriteria
//
// Use this API command to retrieve a list of User Traffic Profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGUserTrafficProfileService) FindProfilesUtpByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileUserTrafficProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesUtpById
//
// Use this API command to modify the basic information of user traffic profile.
//
// Request Body:
//	 - body *WSGProfileModifyUserTrafficProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGUserTrafficProfileService) PartialUpdateProfilesUtpById(ctx context.Context, body *WSGProfileModifyUserTrafficProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
