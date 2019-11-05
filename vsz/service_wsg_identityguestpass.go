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

type WSGIdentityGuestPassService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGIdentityGuestPassService(c *APIClient) *WSGIdentityGuestPassService {
	s := new(WSGIdentityGuestPassService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGIdentityGuestPassService() *WSGIdentityGuestPassService {
	return NewWSGIdentityGuestPassService(ss.apiClient)
}

// AddIdentityGuestpassGenerate
//
// Use this API command to generate identity guest pass.
//
// Request Body:
//	 - body *identity.CreateIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate(ctx context.Context, body *identity.CreateIdentityGuestPass) (*common.CreateResult, error) {
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

// AddIdentityGuestpassList
//
// Use this API command to retrieve a list of identity guest pass.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList(ctx context.Context, body *identity.QueryCriteria) (*identity.IdentityGuestPassList, error) {
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

// AddIdentityGuestpassUpload
//
// Use this API command to upload identity guest pass csv file.
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// AddIdentityGuestpassUploadCommon
//
// Use this API command to update common identity guest pass settings.
//
// Request Body:
//	 - body *identity.ImportIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon(ctx context.Context, body *identity.ImportIdentityGuestPass) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// DeleteIdentityGuestpass
//
// Use this API command to delete multiple identity guest passes.
//
// Request Body:
//	 - body *identity.DeleteBulk
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpass(ctx context.Context, body *identity.DeleteBulk) error {
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

// DeleteIdentityGuestpassByUserId
//
// Use this API command to delete identity guest pass.
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId(ctx context.Context, pUserId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIdentityGuestpass
//
// Use this API command to retrieve a list of identity guest pass.
//
// Query Parameters:
// - qDisplayName string
// - qExpirationFrom string
// - qExpirationTo string
// - qGeneratedTimeFrom string
// - qGeneratedTimeTo string
// - qIndex string
// - qListSize string
// - qTimeZone string
// - qWlan string
func (s *WSGIdentityGuestPassService) FindIdentityGuestpass(ctx context.Context, qDisplayName string, qExpirationFrom string, qExpirationTo string, qGeneratedTimeFrom string, qGeneratedTimeTo string, qIndex string, qListSize string, qTimeZone string, qWlan string) (*identity.IdentityGuestPassList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
