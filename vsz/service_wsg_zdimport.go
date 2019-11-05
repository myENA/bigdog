package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"gopkg.in/go-playground/validator.v9"
)

type WSGZDImportService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGZDImportService(c *APIClient) *WSGZDImportService {
	s := new(WSGZDImportService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGZDImportService() *WSGZDImportService {
	return NewWSGZDImportService(ss.apiClient)
}

// AddZdImportConnectZD
//
// Connect to ZD.
//
// Request Body:
//	 - body *administration.ZdImport
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *administration.ZdImport) (*common.EmptyResult, error) {
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

// AddZdImportMigrate
//
// Migrate ZD to SCG.
//
// Request Body:
//	 - body *administration.ZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *administration.ZdImport) (*common.EmptyResult, error) {
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

// FindZdImportGetZDAPs
//
// Get ZD AP.
//
// Query Parameters:
// - qIp string
//		- required
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, qIp string) (*administration.ZdAPList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindZdImportStatus
//
// Get Migrate Status.
//
// Query Parameters:
// - qDetails string
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, qDetails string) (*administration.ZdImportStatus, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
