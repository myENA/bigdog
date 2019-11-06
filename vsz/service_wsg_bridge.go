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

type WSGBridgeService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGBridgeService(c *APIClient) *WSGBridgeService {
	s := new(WSGBridgeService)
	s.apiClient = c
	s.validate = validator.New()
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
//	 - body *profile.CreateBridgeProfile
func (s *WSGBridgeService) AddProfilesBridge(ctx context.Context, body *profile.CreateBridgeProfile) (*common.CreateResult, error) {
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

// DeleteProfilesBridge
//
// Use this API command to delete multiple bridge profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGBridgeService) DeleteProfilesBridge(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
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
func (s *WSGBridgeService) FindProfilesBridge(ctx context.Context) (*profile.ProfileList, error) {
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
func (s *WSGBridgeService) FindProfilesBridgeById(ctx context.Context, pId string) (*profile.BridgeProfile, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGBridgeService) FindProfilesBridgeByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BridgeProfileList, error) {
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

// PartialUpdateProfilesBridgeById
//
// Use this API command to modify the basic information of Bridge profile.
//
// Request Body:
//	 - body *profile.ModifyBridgeProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) PartialUpdateProfilesBridgeById(ctx context.Context, body *profile.ModifyBridgeProfile, pId string) (*common.EmptyResult, error) {
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
