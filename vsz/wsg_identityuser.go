package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGIdentityUserService struct {
	apiClient *APIClient
}

func NewWSGIdentityUserService(c *APIClient) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.apiClient = c
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
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityUserList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddIdentityUsers
//
// Use this API command to create identity user.
//
// Request Body:
//	 - body *WSGIdentityCreateUser
func (s *WSGIdentityUserService) AddIdentityUsers(ctx context.Context, body *WSGIdentityCreateUser) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserService) DeleteIdentityUsers(ctx context.Context, body *WSGIdentityDeleteBulk) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteIdentityUsersById
//
// Use this API command to delete identity user.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserService) DeleteIdentityUsersById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Query Parameters:
// - qCreatedOnFrom string
//		- nullable
// - qCreatedOnTo string
//		- nullable
// - qDisplayName string
//		- nullable
// - qEmail string
//		- nullable
// - qFirstName string
//		- nullable
// - qIndex string
//		- nullable
// - qIsDisabled string
//		- nullable
// - qLastName string
//		- nullable
// - qListSize string
//		- nullable
// - qPhone string
//		- nullable
// - qTimeZone string
//		- nullable
// - qUserName string
//		- nullable
// - qUserSource string
//		- nullable
// - qUserType string
//		- nullable
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, qCreatedOnFrom string, qCreatedOnTo string, qDisplayName string, qEmail string, qFirstName string, qIndex string, qIsDisabled string, qLastName string, qListSize string, qPhone string, qTimeZone string, qUserName string, qUserSource string, qUserType string) (*WSGIdentityUserList, error) {
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
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context) (*WSGIdentityAaaServerList, error) {
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
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, pId string) (*WSGIdentityUserConfiguration, error) {
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
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context) (*WSGIdentityCountryList, error) {
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
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context) (*WSGIdentityPackageList, error) {
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
//	 - body *WSGIdentityModifyUser
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserService) PartialUpdateIdentityUsersById(ctx context.Context, body *WSGIdentityModifyUser, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
