package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGBlockClientService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGBlockClientService(c *APIClient) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGBlockClientService() *WSGBlockClientService {
	return NewWSGBlockClientService(ss.apiClient)
}

// AddBlockClient
//
// Create new Block Clients by list.
//
// Request Body:
//	 - body *profile.BulkBlockClient
func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *profile.BulkBlockClient) (profile.CreateResultList, error) {
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

// AddBlockClientByApMacByApMac
//
// Create a new Block Client by AP MAC.
//
// Request Body:
//	 - body *profile.BlockClient
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *profile.BlockClient, pApMac string) (*common.CreateResult, error) {
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

// DeleteBlockClient
//
// Delete Block Client List.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGBlockClientService) DeleteBlockClient(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
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

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) DeleteBlockClientById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, pId string) (*profile.BlockClient, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindBlockClientByQueryCriteria
//
// Retrieve a list of Block Client.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BlockClientList, error) {
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

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, pZoneId string) (*profile.BlockClientList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *profile.ModifyBlockClient
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) PartialUpdateBlockClientById(ctx context.Context, body *profile.ModifyBlockClient, pId string) (*common.EmptyResult, error) {
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

// UpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *profile.ModifyBlockClient
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) UpdateBlockClientById(ctx context.Context, body *profile.ModifyBlockClient, pId string) (*common.EmptyResult, error) {
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
