package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGRuckusGRETunnelProfileService struct {
	apiClient *APIClient
}

func NewWSGRuckusGRETunnelProfileService(c *APIClient) *WSGRuckusGRETunnelProfileService {
	s := new(WSGRuckusGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRuckusGRETunnelProfileService() *WSGRuckusGRETunnelProfileService {
	return NewWSGRuckusGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelRuckusgre
//
// Use this API command to create RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileCreateRuckusGREProfile
func (s *WSGRuckusGRETunnelProfileService) AddProfilesTunnelRuckusgre(ctx context.Context, body *WSGProfileCreateRuckusGREProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTunnelRuckusgre
//
// Use this API command to delete multiple RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTunnelRuckusgreById
//
// Use this API command to delete RuckusGRE tunnel profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgreById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelRuckusgre
//
// Use this API command to retrieve a list of RuckusGRE tunnel profile.
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgre(ctx context.Context) (*WSGProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelRuckusgreById
//
// Use this API command to retrieve RuckusGRE tunnel profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreById(ctx context.Context, pId string) (*WSGProfileRuckusGREProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelRuckusgreByQueryCriteria
//
// Use this API command to query a list of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileRuckusGREProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesTunnelRuckusgreById
//
// Use this API command to modify the basic information of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileModifyRuckusGREProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusGRETunnelProfileService) PartialUpdateProfilesTunnelRuckusgreById(ctx context.Context, body *WSGProfileModifyRuckusGREProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
