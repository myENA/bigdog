package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
	"gopkg.in/go-playground/validator.v9"
)

type WSGSCGUserService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGSCGUserService(c *APIClient) *WSGSCGUserService {
	s := new(WSGSCGUserService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGSCGUserService() *WSGSCGUserService {
	return NewWSGSCGUserService(ss.apiClient)
}

// AddUsers
//
// Add SCG user.
//
// Request Body:
//	 - body *scguser.CreateScgUser
func (s *WSGSCGUserService) AddUsers(ctx context.Context, body *scguser.CreateScgUser) (*scguser.ScgUserAuditId, error) {
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

// DeleteUsers
//
// Delete multiple SCG user.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGSCGUserService) DeleteUsers(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
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

// DeleteUsersByUserId
//
// Delete SCG user.
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGSCGUserService) DeleteUsersByUserId(ctx context.Context, pUserId string) (*scguser.ScgUserAuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUsersByQueryCriteria
//
// Query SCG users.
//
// Request Body:
//	 - body *scguser.QueryCriteria
func (s *WSGSCGUserService) FindUsersByQueryCriteria(ctx context.Context, body *scguser.QueryCriteria) (*scguser.ScgUserList, error) {
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

// FindUsersByUserId
//
// Get SCG user.
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGSCGUserService) FindUsersByUserId(ctx context.Context, pUserId string) (*scguser.GetScgUser, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateUsersByUserId
//
// Update SCG user.
//
// Request Body:
//	 - body *scguser.ModifyScgUser
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGSCGUserService) PartialUpdateUsersByUserId(ctx context.Context, body *scguser.ModifyScgUser, pUserId string) (*scguser.ScgUserAuditId, error) {
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
