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

type WSGRealTimeLocationServiceProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGRealTimeLocationServiceProfileService(c *APIClient) *WSGRealTimeLocationServiceProfileService {
	s := new(WSGRealTimeLocationServiceProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService() *WSGRealTimeLocationServiceProfileService {
	return NewWSGRealTimeLocationServiceProfileService(ss.apiClient)
}

// AddRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to create RTLS Profile.
//
// Request Body:
//	 - body *profile.CreateRtlsProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) AddRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, body *profile.CreateRtlsProfile, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesRealTimeLocationServiceById
//
// Use this API command to Delete RTLS Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) DeleteRkszonesRealTimeLocationServiceById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesRealTimeLocationServiceById
//
// Use this API command to Get RTLS Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceById(ctx context.Context, pId string, pZoneId string) (*profile.CreateRtlsProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to Get RTLS Profile by zone ID.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, pZoneId string) (*profile.RtlsProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesRealTimeLocationServiceById
//
// Use this API command to Modify RTLS Profile by profile's ID.
//
// Request Body:
//	 - body *profile.UpdateRtlsProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) UpdateRkszonesRealTimeLocationServiceById(ctx context.Context, body *profile.UpdateRtlsProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
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
