package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGSoftGRETunnelProfileService struct {
	apiClient *APIClient
}

func NewWSGSoftGRETunnelProfileService(c *APIClient) *WSGSoftGRETunnelProfileService {
	s := new(WSGSoftGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSoftGRETunnelProfileService() *WSGSoftGRETunnelProfileService {
	return NewWSGSoftGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelSoftgre
//
// Use this API command to create SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileCreateSoftGREProfile
func (s *WSGSoftGRETunnelProfileService) AddProfilesTunnelSoftgre(ctx context.Context, body *WSGProfileCreateSoftGREProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTunnelSoftgre
//
// Use this API command to delete multiple SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTunnelSoftgreById
//
// Use this API command to delete SoftGRE tunnel profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgreById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelSoftgre
//
// Use this API command to retrieve a list of SoftGRE tunnel profile.
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgre(ctx context.Context) (*WSGProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelSoftgreById
//
// Use this API command to retrieve SoftGRE tunnel profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreById(ctx context.Context, pId string) (*WSGProfileSoftGREProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelSoftgreByQueryCriteria
//
// Use this API command to query a list of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileSoftGREProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesTunnelSoftgreById
//
// Use this API command to modify the basic information of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileModifySoftGREProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSoftGRETunnelProfileService) PartialUpdateProfilesTunnelSoftgreById(ctx context.Context, body *WSGProfileModifySoftGREProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
