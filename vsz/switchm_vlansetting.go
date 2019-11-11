package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMVLANSettingService struct {
	apiClient *APIClient
}

func NewSwitchMVLANSettingService(c *APIClient) *SwitchMVLANSettingService {
	s := new(SwitchMVLANSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVLANSettingService() *SwitchMVLANSettingService {
	return NewSwitchMVLANSettingService(ss.apiClient)
}

// AddVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVlanConfigCreateVlanConfig
func (s *SwitchMVLANSettingService) AddVlans(ctx context.Context, body *SwitchMVlanConfigCreateVlanConfig) (*SwitchMCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVLANSettingService) DeleteVlans(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMVLANSettingService) DeleteVlansById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVlans
//
// Use this API command to Retrieve the VLAN Config List.
func (s *SwitchMVLANSettingService) FindVlans(ctx context.Context) (*SwitchMVlanConfigQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVlansById
//
// Use this API command to Retrieve the VLAN Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMVLANSettingService) FindVlansById(ctx context.Context, pId string) (*SwitchMVlanConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVlansByQueryCriteria
//
// Use this API command to Retrieve the VLAN Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVLANSettingService) FindVlansByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVlanConfigQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVlanConfigUpdateVlanConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMVLANSettingService) UpdateVlansById(ctx context.Context, body *SwitchMVlanConfigUpdateVlanConfig, pId string) (*SwitchMVlanConfigEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
