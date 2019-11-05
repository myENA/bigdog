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

type WSGPrecedenceProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGPrecedenceProfileService(c *APIClient) *WSGPrecedenceProfileService {
	s := new(WSGPrecedenceProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGPrecedenceProfileService() *WSGPrecedenceProfileService {
	return NewWSGPrecedenceProfileService(ss.apiClient)
}

// AddPrecedence
//
// Use this API command to create Precedence Profile.
//
// Request Body:
//	 - body *profile.CreatePrecedenceProfile
func (s *WSGPrecedenceProfileService) AddPrecedence(ctx context.Context, body *profile.CreatePrecedenceProfile) (*common.CreateResult, error) {
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

// DeletePrecedence
//
// Use this API command to Bulk Delete Precedence Profile.
//
// Request Body:
//	 - body *profile.DeleteBulkPrecedenceProfile
func (s *WSGPrecedenceProfileService) DeletePrecedence(ctx context.Context, body *profile.DeleteBulkPrecedenceProfile) (*common.EmptyResult, error) {
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

// DeletePrecedenceById
//
// Use this API command to Delete Precedence Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGPrecedenceProfileService) DeletePrecedenceById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindPrecedence
//
// Use this API command to Get Precedence Profile list.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGPrecedenceProfileService) FindPrecedence(ctx context.Context, qIndex string, qListSize string) (*profile.PrecedenceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindPrecedenceById
//
// Use this API command to Get Precedence Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGPrecedenceProfileService) FindPrecedenceById(ctx context.Context, pId string) (*profile.CreatePrecedenceProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGPrecedenceProfileService) FindPrecedenceByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.PrecedenceList, error) {
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

// PartialUpdatePrecedenceById
//
// Use this API command to Modify Precedence Profile by profile's ID.
//
// Request Body:
//	 - body *profile.UpdatePrecedenceProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGPrecedenceProfileService) PartialUpdatePrecedenceById(ctx context.Context, body *profile.UpdatePrecedenceProfile, pId string) (*common.EmptyResult, error) {
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
