package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
	"gopkg.in/go-playground/validator.v9"
)

type WSGLWAPPTOSCGService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGLWAPPTOSCGService(c *APIClient) *WSGLWAPPTOSCGService {
	s := new(WSGLWAPPTOSCGService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGLWAPPTOSCGService() *WSGLWAPPTOSCGService {
	return NewWSGLWAPPTOSCGService(ss.apiClient)
}

// FindLwapp2scg
//
// Use this API command to retrieve Lwapp Config.
func (s *WSGLWAPPTOSCGService) FindLwapp2scg(ctx context.Context) (*system.Lwapp2scgConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateLwapp2scg
//
// Use this API command to modify the basic information of the Lwapp Config.
//
// Request Body:
//	 - body *system.ModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scg(ctx context.Context, body *system.ModifyLwapp2scg) (*common.EmptyResult, error) {
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

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Request Body:
//	 - body *system.ModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList(ctx context.Context, body *system.ModifyLwapp2scg) (*common.EmptyResult, error) {
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
