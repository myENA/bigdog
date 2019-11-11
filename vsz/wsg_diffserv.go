package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDiffServService struct {
	apiClient *APIClient
}

func NewWSGDiffServService(c *APIClient) *WSGDiffServService {
	s := new(WSGDiffServService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDiffServService() *WSGDiffServService {
	return NewWSGDiffServService(ss.apiClient)
}

// AddRkszonesDiffservByZoneId
//
// Use this API command to create DiffServ profile.
//
// Request Body:
//	 - body *WSGZoneCreateDiffServProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDiffServService) AddRkszonesDiffservByZoneId(ctx context.Context, body *WSGZoneCreateDiffServProfile, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesDiffservById
//
// Use this API command to delete DiffServ profile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDiffServService) DeleteRkszonesDiffservById(ctx context.Context, pId string, pZoneId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDiffservById
//
// Use this API command to retrieve DiffServ profile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservById(ctx context.Context, pId string, pZoneId string) (*WSGZoneDiffServConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDiffservByZoneId
//
// Use this API command to retrieve a list of DiffServ profile.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservByZoneId(ctx context.Context, pZoneId string) (*WSGZoneDiffServList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesDiffservById
//
// Use this API command to modify the basic information of DiffServ profile.
//
// Request Body:
//	 - body *WSGZoneModifyDiffServProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDiffServService) PartialUpdateRkszonesDiffservById(ctx context.Context, body *WSGZoneModifyDiffServProfile, pId string, pZoneId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
