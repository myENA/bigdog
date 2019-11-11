package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGIdentityUserRoleService struct {
	apiClient *APIClient
}

func NewWSGIdentityUserRoleService(c *APIClient) *WSGIdentityUserRoleService {
	s := new(WSGIdentityUserRoleService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityUserRoleService() *WSGIdentityUserRoleService {
	return NewWSGIdentityUserRoleService(ss.apiClient)
}

// AddIdentityUserrole
//
// Use this API command to create identity user role.
//
// Request Body:
//	 - body *WSGIdentityCreateIdentityUserRole
func (s *WSGIdentityUserRoleService) AddIdentityUserrole(ctx context.Context, body *WSGIdentityCreateIdentityUserRole) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteIdentityUserrole
//
// Use this API command to delete multiple identity user roles.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserRoleService) DeleteIdentityUserrole(ctx context.Context, body *WSGIdentityDeleteBulk) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteIdentityUserroleById
//
// Use this API command to delete identity user role.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserRoleService) DeleteIdentityUserroleById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUserrole
//
// Use this API command to retrieve a list of identity user role.
func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context) (*WSGIdentityList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityUserroleById
//
// Use this API command to retrieve identity user role by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserRoleService) FindIdentityUserroleById(ctx context.Context, pId string) (*WSGIdentityUserRole, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateIdentityUserroleById
//
// Use this API command to modify the basic information of identity user role.
//
// Request Body:
//	 - body *WSGIdentityModifyIdentityUserRole
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserRoleService) PartialUpdateIdentityUserroleById(ctx context.Context, body *WSGIdentityModifyIdentityUserRole, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
