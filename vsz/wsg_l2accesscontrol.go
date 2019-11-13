package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGL2AccessControlService struct {
	apiClient *APIClient
}

func NewWSGL2AccessControlService(c *APIClient) *WSGL2AccessControlService {
	s := new(WSGL2AccessControlService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlService() *WSGL2AccessControlService {
	return NewWSGL2AccessControlService(ss.apiClient)
}

// AddRkszonesL2ACLByZoneId
//
// Create a new L2 Access Control.
//
// Request Body:
//	 - body *WSGPortalServiceCreateL2ACL
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) AddRkszonesL2ACLByZoneId(ctx context.Context, body *WSGPortalServiceCreateL2ACL, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesL2ACLById
//
// Delete an L2 Access Control.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) DeleteRkszonesL2ACLById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesL2ACLById
//
// Retrieve an L2 Access Control.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) FindRkszonesL2ACLById(ctx context.Context, pId string, pZoneId string) (*WSGPortalServiceL2ACL, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesL2ACLByZoneId
//
// Retrieve a list of L2 Access Control.
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
func (s *WSGL2AccessControlService) FindRkszonesL2ACLByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*WSGPortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesL2ACLById
//
// Modify a specific L2 Access Control basic.
//
// Request Body:
//	 - body *WSGPortalServiceModifyL2ACL
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) PartialUpdateRkszonesL2ACLById(ctx context.Context, body *WSGPortalServiceModifyL2ACL, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
