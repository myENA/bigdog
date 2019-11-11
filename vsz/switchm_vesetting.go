package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMVESettingService struct {
	apiClient *APIClient
}

func NewSwitchMVESettingService(c *APIClient) *SwitchMVESettingService {
	s := new(SwitchMVESettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVESettingService() *SwitchMVESettingService {
	return NewSwitchMVESettingService(ss.apiClient)
}

// AddVeConfigs
//
// Use this API command to Create VE Config.
//
// Request Body:
//	 - body *SwitchMVeConfigCreate
func (s *SwitchMVESettingService) AddVeConfigs(ctx context.Context, body *SwitchMVeConfigCreate) (SwitchMVeConfigCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteVeConfigs
//
// Use this API command to Delete VE Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVESettingService) DeleteVeConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteVeConfigsById
//
// Use this API command to Delete VE Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMVESettingService) DeleteVeConfigsById(ctx context.Context, pId string) error {
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
func (s *SwitchMVESettingService) FindVeConfigs(ctx context.Context) (*SwitchMVeConfigList, error) {
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
func (s *SwitchMVESettingService) FindVeConfigsById(ctx context.Context, pId string) (*SwitchMVeConfig, error) {
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
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVESettingService) FindVeConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVeConfigList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateVeConfigsById
//
// Use this API command to Update VE Config.
//
// Request Body:
//	 - body *SwitchMVeConfigModify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMVESettingService) UpdateVeConfigsById(ctx context.Context, body *SwitchMVeConfigModify, pId string) (*SwitchMVeConfigEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
