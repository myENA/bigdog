package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMAccessControlListService struct {
	apiClient *APIClient
}

func NewSwitchMAccessControlListService(c *APIClient) *SwitchMAccessControlListService {
	s := new(SwitchMAccessControlListService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAccessControlListService() *SwitchMAccessControlListService {
	return NewSwitchMAccessControlListService(ss.apiClient)
}

// AddAccessControls
//
// Use this API command to Create the Access Control Config.
//
// Request Body:
//	 - body *SwitchMACLConfigCreateACLConfig
func (s *SwitchMAccessControlListService) AddAccessControls(ctx context.Context, body *SwitchMACLConfigCreateACLConfig) (*SwitchMCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMAccessControlListService) DeleteAccessControls(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteAccessControlsById
//
// Use this API command to Delete the Access Control Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMAccessControlListService) DeleteAccessControlsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAccessControlsById
//
// Use this API command to Retrieve the Access Control Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMAccessControlListService) FindAccessControlsById(ctx context.Context, pId string) (*SwitchMACLConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAccessControlsByQueryCriteria
//
// Use this API command to Retrieve the Access Control Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMACLConfigsQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAccessControlsById
//
// Use this API command to Update the Access Control Config.
//
// Request Body:
//	 - body *SwitchMACLConfigUpdateACLConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMAccessControlListService) UpdateAccessControlsById(ctx context.Context, body *SwitchMACLConfigUpdateACLConfig, pId string) (*SwitchMACLConfigEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
