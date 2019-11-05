package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/devicepolicy"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDevicePolicyService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGDevicePolicyService(c *APIClient) *WSGDevicePolicyService {
	s := new(WSGDevicePolicyService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGDevicePolicyService() *WSGDevicePolicyService {
	return NewWSGDevicePolicyService(ss.apiClient)
}

// AddRkszonesDevicePolicyByZoneId
//
// Create a new Device Policy Porfile.
//
// Request Body:
//	 - body *devicepolicy.CreateDevicePolicy
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) AddRkszonesDevicePolicyByZoneId(ctx context.Context, body *devicepolicy.CreateDevicePolicy, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesDevicePolicyById
//
// Delete Device Policy Porfile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) DeleteRkszonesDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDevicePolicyById
//
// Retrieve a Device Policy Porfile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*devicepolicy.DevicePolicyPorfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesDevicePolicyByZoneId
//
// Retrieve a list of Device Policy Porfiles within a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*devicepolicy.PorfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesDevicePolicyById
//
// Modify a specific Device Policy Porfile.
//
// Request Body:
//	 - body *devicepolicy.ModifyDevicePolicy
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) PartialUpdateRkszonesDevicePolicyById(ctx context.Context, body *devicepolicy.ModifyDevicePolicy, pId string, pZoneId string) (*common.EmptyResult, error) {
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
