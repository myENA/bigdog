package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGBridgeService struct {
	apiClient *APIClient
}

func NewWSGBridgeService(c *APIClient) *WSGBridgeService {
	s := new(WSGBridgeService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBridgeService() *WSGBridgeService {
	return NewWSGBridgeService(ss.apiClient)
}

// AddProfilesBridge
//
// Use this API command to create Bridge profile.
//
// Request Body:
//	 - body *WSGProfileCreateBridgeProfile
func (s *WSGBridgeService) AddProfilesBridge(ctx context.Context, body *WSGProfileCreateBridgeProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesBridge
//
// Use this API command to delete multiple bridge profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBridgeService) DeleteProfilesBridge(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesBridgeById
//
// Use this API command to delete Bridge profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) DeleteProfilesBridgeById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesBridge
//
// Use this API command to retrieve a list of Bridge profile.
func (s *WSGBridgeService) FindProfilesBridge(ctx context.Context) (*WSGProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesBridgeById
//
// Use this API command to retrieve Bridge profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) FindProfilesBridgeById(ctx context.Context, pId string) (*WSGProfileBridgeProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesBridgeByQueryCriteria
//
// Use this API command to query a list of Bridge profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBridgeService) FindProfilesBridgeByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBridgeProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesBridgeById
//
// Use this API command to modify the basic information of Bridge profile.
//
// Request Body:
//	 - body *WSGProfileModifyBridgeProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) PartialUpdateProfilesBridgeById(ctx context.Context, body *WSGProfileModifyBridgeProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
