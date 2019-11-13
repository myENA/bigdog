package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGLWAPPTOSCGService struct {
	apiClient *APIClient
}

func NewWSGLWAPPTOSCGService(c *APIClient) *WSGLWAPPTOSCGService {
	s := new(WSGLWAPPTOSCGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLWAPPTOSCGService() *WSGLWAPPTOSCGService {
	return NewWSGLWAPPTOSCGService(ss.apiClient)
}

// FindLwapp2scg
//
// Use this API command to retrieve Lwapp Config.
func (s *WSGLWAPPTOSCGService) FindLwapp2scg(ctx context.Context) (*WSGSystemLwapp2scgConfiguration, error) {
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
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scg(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Request Body:
//	 - body *WSGSystemModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList(ctx context.Context, body *WSGSystemModifyLwapp2scg) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
