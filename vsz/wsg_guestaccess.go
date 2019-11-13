package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGGuestAccessService struct {
	apiClient *APIClient
}

func NewWSGGuestAccessService(c *APIClient) *WSGGuestAccessService {
	s := new(WSGGuestAccessService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGuestAccessService() *WSGGuestAccessService {
	return NewWSGGuestAccessService(ss.apiClient)
}

// AddRkszonesPortalsGuestByZoneId
//
// Use this API command to create new guest access of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceCreateGuestAccess
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) AddRkszonesPortalsGuestByZoneId(ctx context.Context, body *WSGPortalServiceCreateGuestAccess, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalsGuestById
//
// Use this API command to delete guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalsGuestRedirectById
//
// Use this API command to set redirect to the URL that user intends to visit on guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestRedirectById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalsGuestSmsGatewayById
//
// Use this API command to disable SMS gateway on guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestSmsGatewayById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsGuestById
//
// Use this API command to retrieve guest access of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) FindRkszonesPortalsGuestById(ctx context.Context, pId string, pZoneId string) (*WSGPortalServiceGuestAccess, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsGuestByZoneId
//
// Use this API command to retrieve a list of guest access of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) FindRkszonesPortalsGuestByZoneId(ctx context.Context, pZoneId string) (*WSGPortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesPortalsGuestById
//
// Use this API command to modify the basic information on guest access of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceModifyGuestAccess
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGGuestAccessService) PartialUpdateRkszonesPortalsGuestById(ctx context.Context, body *WSGPortalServiceModifyGuestAccess, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
