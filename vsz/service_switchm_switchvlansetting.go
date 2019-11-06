package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/vlanconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchVLANSettingService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchVLANSettingService(c *APIClient) *SwitchMSwitchVLANSettingService {
	s := new(SwitchMSwitchVLANSettingService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchVLANSettingService() *SwitchMSwitchVLANSettingService {
	return NewSwitchMSwitchVLANSettingService(ss.apiClient)
}

// AddVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *vlanconfig.CreateVlanConfig
func (s *SwitchMSwitchVLANSettingService) AddVlans(ctx context.Context, body *vlanconfig.CreateVlanConfig) (*common.CreateResult, error) {
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

// DeleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchVLANSettingService) DeleteVlans(ctx context.Context, body *common.BulkDeleteRequest) error {
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

// DeleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVLANSettingService) DeleteVlansById(ctx context.Context, pId string) error {
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
func (s *SwitchMSwitchVLANSettingService) FindVlans(ctx context.Context) (*vlanconfig.VlanConfigQueryResult, error) {
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
func (s *SwitchMSwitchVLANSettingService) FindVlansById(ctx context.Context, pId string) (*vlanconfig.VlanConfig, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchVLANSettingService) FindVlansByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*vlanconfig.VlanConfigQueryResult, error) {
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

// UpdateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *vlanconfig.UpdateVlanConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchVLANSettingService) UpdateVlansById(ctx context.Context, body *vlanconfig.UpdateVlanConfig, pId string) (*vlanconfig.EmptyResult, error) {
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
