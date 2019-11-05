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

type WSGClientIsolationWhitelistService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGClientIsolationWhitelistService(c *APIClient) *WSGClientIsolationWhitelistService {
	s := new(WSGClientIsolationWhitelistService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService() *WSGClientIsolationWhitelistService {
	return NewWSGClientIsolationWhitelistService(ss.apiClient)
}

// AddRkszonesClientIsolationWhitelistByZoneId
//
// Create a new ClientIsolationWhitelist.
//
// Request Body:
//	 - body *profile.CreateClientIsolationWhitelist
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) AddRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, body *profile.CreateClientIsolationWhitelist, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesClientIsolationWhitelist
//
// Use this API command to delete Bulk Client Isolation Whitelist.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
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

// DeleteRkszonesClientIsolationWhitelistById
//
// Delete a Client Isolation Whitelist.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesClientIsolationWhitelistById
//
// Retrieve an Client Isolation Whitelist.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById(ctx context.Context, pId string, pZoneId string) (*profile.ClientIsolationWhitelist, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesClientIsolationWhitelistByZoneId
//
// Retrieve a list of Client Isolation Whitelist.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, pZoneId string) (*profile.ClientIsolationWhitelistArray, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesClientIsolationWhitelistByQueryCriteria
//
// Retrieve a list of Client Isolation Whitelist.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.ClientIsolationWhitelistArray, error) {
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

// PartialUpdateRkszonesClientIsolationWhitelistById
//
// Modify a specific Client Isolation Whitelist basic.
//
// Request Body:
//	 - body *profile.ModifyClientIsolationWhitelist
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) PartialUpdateRkszonesClientIsolationWhitelistById(ctx context.Context, body *profile.ModifyClientIsolationWhitelist, pId string, pZoneId string) (*common.EmptyResult, error) {
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
