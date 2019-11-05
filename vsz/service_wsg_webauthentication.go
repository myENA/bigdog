package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
	"gopkg.in/go-playground/validator.v9"
)

type WSGWebAuthenticationService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGWebAuthenticationService(c *APIClient) *WSGWebAuthenticationService {
	s := new(WSGWebAuthenticationService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGWebAuthenticationService() *WSGWebAuthenticationService {
	return NewWSGWebAuthenticationService(ss.apiClient)
}

// AddRkszonesPortalsWebauthByZoneId
//
// Use this API command to create a new web authentication of a zone.
//
// Request Body:
//	 - body *portalservice.CreateWebAuthentication
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWebAuthenticationService) AddRkszonesPortalsWebauthByZoneId(ctx context.Context, body *portalservice.CreateWebAuthentication, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesPortalsWebauthById
//
// Use this API command to delete an web authentication of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalsWebauthRedirectById
//
// Use this API command to set redirect to the URL that user intends to visit on web authentication of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthRedirectById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsWebauthById
//
// Use this API command to retrieve a web authentication of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthById(ctx context.Context, pId string, pZoneId string) (*portalservice.WebAuthentication, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsWebauthByZoneId
//
// Use this API command to retrieve a list of web authentication of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesPortalsWebauthById
//
// Use this API command to modify the basic information on web authentication of a zone.
//
// Request Body:
//	 - body *portalservice.ModifyWebAuthentication
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWebAuthenticationService) PartialUpdateRkszonesPortalsWebauthById(ctx context.Context, body *portalservice.ModifyWebAuthentication, pId string, pZoneId string) (*common.EmptyResult, error) {
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
