package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGBlockClientService struct {
	apiClient *APIClient
}

func NewWSGBlockClientService(c *APIClient) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.apiClient = c
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
//	 - body *WSGProfileBulkBlockClient
func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *WSGProfileBulkBlockClient) (WSGProfileCreateResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddBlockClientByApMacByApMac
//
// Create a new Block Client by AP MAC.
//
// Request Body:
//	 - body *WSGProfileBlockClient
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *WSGProfileBlockClient, pApMac string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteBlockClient
//
// Delete Block Client List.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBlockClientService) DeleteBlockClient(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) DeleteBlockClientById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, pId string) (*WSGProfileBlockClient, error) {
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
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBlockClientList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, pZoneId string) (*WSGProfileBlockClientList, error) {
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
//	 - body *WSGProfileModifyBlockClient
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) PartialUpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *WSGProfileModifyBlockClient
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) UpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
