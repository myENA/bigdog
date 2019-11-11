package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMLAGSettingService struct {
	apiClient *APIClient
}

func NewSwitchMLAGSettingService(c *APIClient) *SwitchMLAGSettingService {
	s := new(SwitchMLAGSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMLAGSettingService() *SwitchMLAGSettingService {
	return NewSwitchMLAGSettingService(ss.apiClient)
}

// AddLagConfigs
//
// Use this API command to Create LAG Config.
//
// Request Body:
//	 - body *SwitchMLagConfigCreate
func (s *SwitchMLAGSettingService) AddLagConfigs(ctx context.Context, body *SwitchMLagConfigCreate) (SwitchMLagConfigCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMLAGSettingService) DeleteLagConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMLAGSettingService) DeleteLagConfigsById(ctx context.Context, pId string) error {
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
func (s *SwitchMLAGSettingService) FindLagConfigs(ctx context.Context) (*SwitchMLagConfigList, error) {
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
func (s *SwitchMLAGSettingService) FindLagConfigsById(ctx context.Context, pId string) (*SwitchMLagConfig, error) {
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
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMLagConfigList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Request Body:
//	 - body *SwitchMLagConfigModify
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *SwitchMLagConfigModify, pId string) (*SwitchMLagConfigEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
