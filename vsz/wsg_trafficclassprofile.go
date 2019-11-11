package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGTrafficClassProfileService struct {
	apiClient *APIClient
}

func NewWSGTrafficClassProfileService(c *APIClient) *WSGTrafficClassProfileService {
	s := new(WSGTrafficClassProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTrafficClassProfileService() *WSGTrafficClassProfileService {
	return NewWSGTrafficClassProfileService(ss.apiClient)
}

// AddRkszonesTrafficClassProfileByZoneId
//
// Use this API command to create a new Traffic Class Profile of a zone.
//
// Request Body:
//	 - body *WSGProfileCreateTrafficClassProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) AddRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *WSGProfileCreateTrafficClassProfile, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesTrafficClassProfileById
//
// Use this API command to delete a Traffic Class Profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesTrafficClassProfileByZoneId
//
// Use this API command to bulk delete Traffic Class Profiles of a zone.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesTrafficClassProfileById
//
// Use this API command to retrieve a Traffic Class Profile of zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileById(ctx context.Context, pId string, pZoneId string) (*WSGCommonTrafficClassProfileRef, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesTrafficClassProfileByZoneId
//
// Use this API command to retrieve a list of Traffic Class Profile of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileByZoneId(ctx context.Context, pZoneId string) (*WSGProfileTrafficClassProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesTrafficClassProfileByQueryCriteria
//
// Retrieve a list of Traffic Class Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGTrafficClassProfileService) FindServicesTrafficClassProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileTrafficClassProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesTrafficClassProfileById
//
// Use this API command to modify Traffic Class Profile of a zone.
//
// Request Body:
//	 - body *WSGProfileCreateTrafficClassProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) PartialUpdateRkszonesTrafficClassProfileById(ctx context.Context, body *WSGProfileCreateTrafficClassProfile, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
