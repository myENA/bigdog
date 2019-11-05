package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sci"
	"gopkg.in/go-playground/validator.v9"
)

type WSGSCIService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGSCIService(c *APIClient) *WSGSCIService {
	s := new(WSGSCIService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	return NewWSGSCIService(ss.apiClient)
}

// AddSciSciEventCode
//
// Use this API command to modify SciAcceptedEventCodes.
//
// Request Body:
//	 - body *sci.ModifyEventCode
func (s *WSGSCIService) AddSciSciEventCode(ctx context.Context, body *sci.ModifyEventCode) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// AddSciSciProfile
//
// Use this API command to create sciProfile.
//
// Request Body:
//	 - body *sci.CreateSciProfile
func (s *WSGSCIService) AddSciSciProfile(ctx context.Context, body *sci.CreateSciProfile) (*common.CreateResult, error) {
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

// AddSciSciProfileSciPriority
//
// Use this API command to modify sciPriorities.
//
// Request Body:
//	 - body *sci.ModifySciPriorityList
func (s *WSGSCIService) AddSciSciProfileSciPriority(ctx context.Context, body *sci.ModifySciPriorityList) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// DeleteSciSciProfile
//
// Use this API command to delete sciProfile list.
//
// Request Body:
//	 - body *sci.DeleteSciProfileList
func (s *WSGSCIService) DeleteSciSciProfile(ctx context.Context, body *sci.DeleteSciProfileList) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// DeleteSciSciProfileById
//
// Use this API command to delete sciProfile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSCIService) DeleteSciSciProfileById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSciSciEventCode
//
// Use this API command to retrieve SciAcceptedEventCodes.
func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context) (*sci.SciEventCode, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSciSciProfile
//
// Use this API command to retrieve sciProfile list.
func (s *WSGSCIService) FindSciSciProfile(ctx context.Context) (*sci.SciProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSciSciProfileById
//
// Use this API command to retrieve sciProfile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, pId string) (*sci.SciProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateSciSciEnabled
//
// Use this API command to modify SCI settings is enabled or not.
//
// Request Body:
//	 - body *sci.ModifySciEnabled
func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context, body *sci.ModifySciEnabled) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// PartialUpdateSciSciProfileById
//
// Use this API command to modify sciProfile.
//
// Request Body:
//	 - body *sci.ModifySciProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSCIService) PartialUpdateSciSciProfileById(ctx context.Context, body *sci.ModifySciProfile, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}
