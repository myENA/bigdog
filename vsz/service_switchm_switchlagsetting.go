package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/lagconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchLAGSettingService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchLAGSettingService(c *APIClient) *SwitchMSwitchLAGSettingService {
	s := new(SwitchMSwitchLAGSettingService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchLAGSettingService() *SwitchMSwitchLAGSettingService {
	return NewSwitchMSwitchLAGSettingService(ss.apiClient)
}

// AddLagConfigs
//
// Use this API command to Create LAG Config.
//
// Request Body:
//	 - body *lagconfig.Create
func (s *SwitchMSwitchLAGSettingService) AddLagConfigs(ctx context.Context, body *lagconfig.Create) (lagconfig.CreateResult, error) {
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

// DeleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchLAGSettingService) DeleteLagConfigs(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchLAGSettingService) DeleteLagConfigsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
func (s *SwitchMSwitchLAGSettingService) FindLagConfigs(ctx context.Context) (*lagconfig.List, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchLAGSettingService) FindLagConfigsById(ctx context.Context, pId string) (*lagconfig.LagConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLagConfigsByQueryCriteria
//
// Use this API command to Retrieve LAG Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*lagconfig.List, error) {
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

// UpdateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Request Body:
//	 - body *lagconfig.Modify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *lagconfig.Modify, pId string) (*lagconfig.EmptyResult, error) {
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
