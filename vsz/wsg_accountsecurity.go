package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAccountSecurityService struct {
	apiClient *APIClient
}

func NewWSGAccountSecurityService(c *APIClient) *WSGAccountSecurityService {
	s := new(WSGAccountSecurityService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountSecurityService() *WSGAccountSecurityService {
	return NewWSGAccountSecurityService(ss.apiClient)
}

// AddAccountSecurity
//
// Use this API command to create the account security proile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileCreate
func (s *WSGAccountSecurityService) AddAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileCreate) (*WSGCommonCreateResultIdName, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteAccountSecurity
//
// Use this API command to selete the account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDeleteList
func (s *WSGAccountSecurityService) DeleteAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileDeleteList) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteAccountSecurityById
//
// Use this API command to delete the account security profile by id.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDelete
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) DeleteAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileDelete, pId string) (*WSGCommonCreateResultIdName, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAccountSecurity
//
// Use this API command to get account security profiles.
func (s *WSGAccountSecurityService) FindAccountSecurity(ctx context.Context) (*WSGAccountSecurityProfileProfileListResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAccountSecurityById
//
// Use this API command to retrieve the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileGetById
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) FindAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileGetById, pId string) (*WSGAccountSecurityProfileGetByIdResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) PartialUpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, pId string) (*WSGCommonCreateResultIdName, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) UpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, pId string) (*WSGCommonCreateResultIdName, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
