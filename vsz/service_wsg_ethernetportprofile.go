package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ethernetport"
	"gopkg.in/go-playground/validator.v9"
)

type WSGEthernetPortProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGEthernetPortProfileService(c *APIClient) *WSGEthernetPortProfileService {
	s := new(WSGEthernetPortProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGEthernetPortProfileService() *WSGEthernetPortProfileService {
	return NewWSGEthernetPortProfileService(ss.apiClient)
}

// AddRkszonesProfileEthernetPortByZoneId
//
// Create a new Ethernet Port Porfile.
//
// Request Body:
//	 - body *ethernetport.CreateEthernetPortProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEthernetPortProfileService) AddRkszonesProfileEthernetPortByZoneId(ctx context.Context, body *ethernetport.CreateEthernetPortProfile, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesProfileEthernetPortById
//
// Delete Ethernet Port Porfile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGEthernetPortProfileService) DeleteRkszonesProfileEthernetPortById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesProfileEthernetPortById
//
// Retrieve a Ethernet Port Porfile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById(ctx context.Context, pId string, pZoneId string) (*ethernetport.EthernetPortProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesProfileEthernetPortByZoneId
//
// Retrieve a list of Ethernet Port Porfiles within a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*ethernetport.ProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesProfileEthernetPortById
//
// Modify a specific Ethernet Port Porfile.
//
// Request Body:
//	 - body *ethernetport.ModifyEthernetPortProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGEthernetPortProfileService) PartialUpdateRkszonesProfileEthernetPortById(ctx context.Context, body *ethernetport.ModifyEthernetPortProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
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
