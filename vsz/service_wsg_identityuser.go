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

type WSGIdentityUserService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGIdentityUserService(c *APIClient) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGIdentityUserService() *WSGIdentityUserService {
	return NewWSGIdentityUserService(ss.apiClient)
}

// AddIdentityUserList
//
// Use this API command to retrieve a list of identity user.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *identity.QueryCriteria) (*identity.UserList, error) {
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

// AddIdentityUsers
//
// Use this API command to create identity user.
//
// Request Body:
//	 - body *identity.CreateUser
func (s *WSGIdentityUserService) AddIdentityUsers(ctx context.Context, body *identity.CreateUser) (*common.CreateResult, error) {
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

// DeleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Request Body:
//	 - body *identity.DeleteBulk
func (s *WSGIdentityUserService) DeleteIdentityUsers(ctx context.Context, body *identity.DeleteBulk) error {
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

// DeleteIdentityUsersById
//
// Use this API command to delete identity user.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserService) DeleteIdentityUsersById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Query Parameters:
// - qCreatedOnFrom string
// - qCreatedOnTo string
// - qDisplayName string
// - qEmail string
// - qFirstName string
// - qIndex string
// - qIsDisabled string
// - qLastName string
// - qListSize string
// - qPhone string
// - qTimeZone string
// - qUserName string
// - qUserSource string
// - qUserType string
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, qCreatedOnFrom string, qCreatedOnTo string, qDisplayName string, qEmail string, qFirstName string, qIndex string, qIsDisabled string, qLastName string, qListSize string, qPhone string, qTimeZone string, qUserName string, qUserSource string, qUserType string) (*identity.UserList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context) (*identity.AaaServerList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, pId string) (*identity.UserConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context) (*identity.CountryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context) (*identity.PackageList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateIdentityUsersById
//
// Use this API command to modify the basic information of identity user.
//
// Request Body:
//	 - body *identity.ModifyUser
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserService) PartialUpdateIdentityUsersById(ctx context.Context, body *identity.ModifyUser, pId string) (*common.EmptyResult, error) {
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
