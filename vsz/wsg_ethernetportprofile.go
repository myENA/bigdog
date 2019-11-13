package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGEthernetPortProfileService struct {
	apiClient *APIClient
}

func NewWSGEthernetPortProfileService(c *APIClient) *WSGEthernetPortProfileService {
	s := new(WSGEthernetPortProfileService)
	s.apiClient = c
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
//	 - body *WSGEthernetPortCreateEthernetPortProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEthernetPortProfileService) AddRkszonesProfileEthernetPortByZoneId(ctx context.Context, body *WSGEthernetPortCreateEthernetPortProfile, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGEthernetPortProfileService) DeleteRkszonesProfileEthernetPortById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById(ctx context.Context, pId string, pZoneId string) (*WSGEthernetPortProfile, error) {
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
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*WSGEthernetPortProfileList, error) {
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
//	 - body *WSGEthernetPortModifyEthernetPortProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGEthernetPortProfileService) PartialUpdateRkszonesProfileEthernetPortById(ctx context.Context, body *WSGEthernetPortModifyEthernetPortProfile, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
