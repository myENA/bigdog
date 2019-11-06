package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/staticroute"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchStaticRouteService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchStaticRouteService(c *APIClient) *SwitchMSwitchStaticRouteService {
	s := new(SwitchMSwitchStaticRouteService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchStaticRouteService() *SwitchMSwitchStaticRouteService {
	return NewSwitchMSwitchStaticRouteService(ss.apiClient)
}

// AddStaticRoutes
//
// Use this API command to Create Static Route.
//
// Request Body:
//	 - body *staticroute.CreateStaticRoute
func (s *SwitchMSwitchStaticRouteService) AddStaticRoutes(ctx context.Context, body *staticroute.CreateStaticRoute) (*common.CreateResult, error) {
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

// DeleteStaticRoutes
//
// Use this API command to Delete Static Route by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchStaticRouteService) DeleteStaticRoutes(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteStaticRoutesById
//
// Use this API command to Delete Static Route.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchStaticRouteService) DeleteStaticRoutesById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindStaticRoutesById
//
// Use this API command to Retrieve Static Route.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchStaticRouteService) FindStaticRoutesById(ctx context.Context, pId string) (*staticroute.StaticRoute, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindStaticRoutesByQueryCriteria
//
// Use this API command to Retrieve Static Route list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchStaticRouteService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*staticroute.StaticRoutesQueryResult, error) {
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

// UpdateStaticRoutesById
//
// Use this API command to Update Static Route.
//
// Request Body:
//	 - body *staticroute.UpdateStaticRoute
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchStaticRouteService) UpdateStaticRoutesById(ctx context.Context, body *staticroute.UpdateStaticRoute, pId string) (*staticroute.EmptyResult, error) {
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
