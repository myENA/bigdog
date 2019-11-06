package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/portsettings"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchPortSettingsService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchPortSettingsService(c *APIClient) *SwitchMSwitchPortSettingsService {
	s := new(SwitchMSwitchPortSettingsService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortSettingsService() *SwitchMSwitchPortSettingsService {
	return NewSwitchMSwitchPortSettingsService(ss.apiClient)
}

// AddPortSettingsBulk
//
// Use this API command to Bulk update the port setting
//
// Request Body:
//	 - body *portsettings.CreateBulk
func (s *SwitchMSwitchPortSettingsService) AddPortSettingsBulk(ctx context.Context, body *portsettings.CreateBulk) (*portsettings.EmptyResult, error) {
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

// FindPortSettings
//
// Use this API command to Retrieve all Port Settings list.
func (s *SwitchMSwitchPortSettingsService) FindPortSettings(ctx context.Context) (*portsettings.PortSettingsQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindPortSettingsById
//
// Use this API command to Retrieve Port Settings.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchPortSettingsService) FindPortSettingsById(ctx context.Context, pId string) (*portsettings.PortSettings, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindPortSettingsByQueryCriteria
//
// Use this API command to Retrieve Port Settings list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchPortSettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*portsettings.PortSettingsQueryResult, error) {
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

// UpdatePortSettingsById
//
// Use this API command to Update Port Settings.
//
// Request Body:
//	 - body *portsettings.UpdatePortSettings
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchPortSettingsService) UpdatePortSettingsById(ctx context.Context, body *portsettings.UpdatePortSettings, pId string) (*portsettings.EmptyResult, error) {
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
