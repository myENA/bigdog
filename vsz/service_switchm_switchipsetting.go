package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/ipconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchIPSettingService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchIPSettingService(c *APIClient) *SwitchMSwitchIPSettingService {
	s := new(SwitchMSwitchIPSettingService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchIPSettingService() *SwitchMSwitchIPSettingService {
	return NewSwitchMSwitchIPSettingService(ss.apiClient)
}

// AddIpConfigs
//
// Use this API command to Create IP Config.
//
// Request Body:
//	 - body *ipconfig.Create
func (s *SwitchMSwitchIPSettingService) AddIpConfigs(ctx context.Context, body *ipconfig.Create) (ipconfig.CreateResult, error) {
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

// DeleteIpConfigs
//
// Use this API command to Delete IP Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchIPSettingService) DeleteIpConfigs(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteIpConfigsById
//
// Use this API command to Delete IP Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchIPSettingService) DeleteIpConfigsById(ctx context.Context, pId string) error {
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
func (s *SwitchMSwitchIPSettingService) FindIpConfigs(ctx context.Context) (*ipconfig.List, error) {
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
func (s *SwitchMSwitchIPSettingService) FindIpConfigsById(ctx context.Context, pId string) (*ipconfig.IpConfig, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchIPSettingService) FindIpConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*ipconfig.List, error) {
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

// UpdateIpConfigsById
//
// Use this API command to Update IP Config.
//
// Request Body:
//	 - body *ipconfig.Modify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchIPSettingService) UpdateIpConfigsById(ctx context.Context, body *ipconfig.Modify, pId string) (*ipconfig.EmptyResult, error) {
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
