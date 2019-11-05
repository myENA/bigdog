package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aaaservers"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchAAAServersService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchAAAServersService(c *APIClient) *SwitchMSwitchAAAServersService {
	s := new(SwitchMSwitchAAAServersService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchAAAServersService() *SwitchMSwitchAAAServersService {
	return NewSwitchMSwitchAAAServersService(ss.apiClient)
}

// AddAaaServersAdmin
//
// Use this API command to create a new AAA server.
//
// Request Body:
//	 - body *aaaservers.CreateAdminAAAServer
func (s *SwitchMSwitchAAAServersService) AddAaaServersAdmin(ctx context.Context, body *aaaservers.CreateAdminAAAServer) (*common.CreateResult, error) {
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

// DeleteAaaServersAdmin
//
// Use this API command to delete AAA Servers.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchAAAServersService) DeleteAaaServersAdmin(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteAaaServersAdminById
//
// Use this API command to delete a AAA server.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAAAServersService) DeleteAaaServersAdminById(ctx context.Context, pId string) (*aaaservers.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAaaServersAdmin
//
// Use this API command to retrieve a list of AAA server.
func (s *SwitchMSwitchAAAServersService) FindAaaServersAdmin(ctx context.Context) (*aaaservers.AaaServersQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAaaServersAdminById
//
// Use this API command to retrieve a AAA server.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAAAServersService) FindAaaServersAdminById(ctx context.Context, pId string) (*aaaservers.AAAServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAaaServersAdminById
//
// Use this API command to modify the basic information on a AAA server by complete attributes.
//
// Request Body:
//	 - body *aaaservers.CreateAdminAAAServer
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAAAServersService) UpdateAaaServersAdminById(ctx context.Context, body *aaaservers.CreateAdminAAAServer, pId string) (*aaaservers.EmptyResult, error) {
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
