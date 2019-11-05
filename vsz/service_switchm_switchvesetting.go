package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/veconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchVESettingService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchVESettingService(c *APIClient) *SwitchMSwitchVESettingService {
	s := new(SwitchMSwitchVESettingService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchVESettingService() *SwitchMSwitchVESettingService {
	return NewSwitchMSwitchVESettingService(ss.apiClient)
}

// AddVeConfigs
//
// Use this API command to Create VE Config.
//
// Request Body:
//	 - body *veconfig.Create
func (s *SwitchMSwitchVESettingService) AddVeConfigs(ctx context.Context, body *veconfig.Create) (veconfig.CreateResult, error) {
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

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchVESettingService) DeleteVeConfigs(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVESettingService) DeleteVeConfigsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVeConfigs
//
// Use this API command to Retrieve VE Config List.
func (s *SwitchMSwitchVESettingService) FindVeConfigs(ctx context.Context) (*veconfig.List, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVeConfigsById
//
// Use this API command to Retrieve VE Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVESettingService) FindVeConfigsById(ctx context.Context, pId string) (*veconfig.VeConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVeConfigsByQueryCriteria
//
// Use this API command to Retrieve VE Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*veconfig.List, error) {
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

// UpdateVeConfigsById
//
// Use this API command to Update VE Config.
//
// Request Body:
//	 - body *veconfig.Modify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVESettingService) UpdateVeConfigsById(ctx context.Context, body *veconfig.Modify, pId string) (*veconfig.EmptyResult, error) {
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
