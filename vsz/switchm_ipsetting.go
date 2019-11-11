package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMIPSettingService struct {
	apiClient *APIClient
}

func NewSwitchMIPSettingService(c *APIClient) *SwitchMIPSettingService {
	s := new(SwitchMIPSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMIPSettingService() *SwitchMIPSettingService {
	return NewSwitchMIPSettingService(ss.apiClient)
}

// AddIpConfigs
//
// Use this API command to Create IP Config.
//
// Request Body:
//	 - body *SwitchMIpConfigCreate
func (s *SwitchMIPSettingService) AddIpConfigs(ctx context.Context, body *SwitchMIpConfigCreate) (SwitchMIpConfigCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteIpConfigs
//
// Use this API command to Delete IP Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMIPSettingService) DeleteIpConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteIpConfigsById
//
// Use this API command to Delete IP Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMIPSettingService) DeleteIpConfigsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIpConfigs
//
// Use this API command to Retrieve IP Config List.
func (s *SwitchMIPSettingService) FindIpConfigs(ctx context.Context) (*SwitchMIpConfigList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIpConfigsById
//
// Use this API command to Retrieve IP Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMIPSettingService) FindIpConfigsById(ctx context.Context, pId string) (*SwitchMIpConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindIpConfigsByQueryCriteria
//
// Use this API command to Retrieve IP Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMIPSettingService) FindIpConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMIpConfigList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateIpConfigsById
//
// Use this API command to Update IP Config.
//
// Request Body:
//	 - body *SwitchMIpConfigModify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMIPSettingService) UpdateIpConfigsById(ctx context.Context, body *SwitchMIpConfigModify, pId string) (*SwitchMIpConfigEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
