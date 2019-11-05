package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
	"gopkg.in/go-playground/validator.v9"
)

type WSGIdentitySubscriptionPackageService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGIdentitySubscriptionPackageService(c *APIClient) *WSGIdentitySubscriptionPackageService {
	s := new(WSGIdentitySubscriptionPackageService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGIdentitySubscriptionPackageService() *WSGIdentitySubscriptionPackageService {
	return NewWSGIdentitySubscriptionPackageService(ss.apiClient)
}

// AddIdentityPackageList
//
// Use this API command to retrieve a list of subscription package.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackageList(ctx context.Context, body *identity.QueryCriteria) (*identity.SubscriptionPackageList, error) {
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

// AddIdentityPackages
//
// Use this API command to create subscription package.
//
// Request Body:
//	 - body *identity.CreateSubscriptionPackage
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackages(ctx context.Context, body *identity.CreateSubscriptionPackage) (*common.CreateResult, error) {
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

// DeleteIdentityPackages
//
// Use this API command to delete multiple subscription packages.
//
// Request Body:
//	 - body *identity.DeleteBulk
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackages(ctx context.Context, body *identity.DeleteBulk) error {
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

// DeleteIdentityPackagesById
//
// Use this API command to delete subscription package.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackagesById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages(ctx context.Context) (*identity.SubscriptionPackageList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityPackagesById
//
// Use this API command to retrieve subscription package.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById(ctx context.Context, pId string) (*identity.SubscriptionPackage, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateIdentityPackagesById
//
// Use this API command to modify the basic information of subscription package.
//
// Request Body:
//	 - body *identity.ModifySubscriptionPackage
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentitySubscriptionPackageService) PartialUpdateIdentityPackagesById(ctx context.Context, body *identity.ModifySubscriptionPackage, pId string) (*common.EmptyResult, error) {
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
