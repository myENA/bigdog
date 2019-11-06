package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aclconfig"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchAccessControlListService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchAccessControlListService(c *APIClient) *SwitchMSwitchAccessControlListService {
	s := new(SwitchMSwitchAccessControlListService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchAccessControlListService() *SwitchMSwitchAccessControlListService {
	return NewSwitchMSwitchAccessControlListService(ss.apiClient)
}

// AddAccessControls
//
// Use this API command to Create the Access Control Config.
//
// Request Body:
//	 - body *aclconfig.CreateACLConfig
func (s *SwitchMSwitchAccessControlListService) AddAccessControls(ctx context.Context, body *aclconfig.CreateACLConfig) (*common.CreateResult, error) {
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

// DeleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchAccessControlListService) DeleteAccessControls(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteAccessControlsById
//
// Use this API command to Delete the Access Control Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAccessControlListService) DeleteAccessControlsById(ctx context.Context, pId string) error {
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
func (s *SwitchMSwitchAccessControlListService) FindAccessControlsById(ctx context.Context, pId string) (*aclconfig.ACLConfig, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aclconfig.ACLConfigsQueryResult, error) {
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

// UpdateAccessControlsById
//
// Use this API command to Update the Access Control Config.
//
// Request Body:
//	 - body *aclconfig.UpdateACLConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAccessControlListService) UpdateAccessControlsById(ctx context.Context, body *aclconfig.UpdateACLConfig, pId string) (*aclconfig.EmptyResult, error) {
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
