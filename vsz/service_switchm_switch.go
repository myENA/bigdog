package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchService(c *APIClient) *SwitchMSwitchService {
	s := new(SwitchMSwitchService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchService() *SwitchMSwitchService {
	return NewSwitchMSwitchService(ss.apiClient)
}

// AddSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitch(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchQueryResultList, error) {
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

// AddSwitchSnmpSyncedSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone and SNMP synced.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchSnmpSyncedSwitch(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchQueryResultList, error) {
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

// AddSwitchViewDetails
//
// Use this API command to retrieve switch and port details for the selected Switch/SwitchGroup/Domain.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchViewDetails(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.StackMemberQueryResult, error) {
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

// DeleteSwitch
//
// Use this API command to delete multiple switches managed by SmartZone
//
// Request Body:
//	 - body switchmswitch.SwitchIdList
func (s *SwitchMSwitchService) DeleteSwitch(ctx context.Context, body switchmswitch.SwitchIdList) (*switchmswitch.DeleteSwitchesResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteSwitchById
//
// Use this API command to delete a switch managed by SmartZone.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchService) DeleteSwitchById(ctx context.Context, pId string) (*switchmswitch.AuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSwitchById
//
// Use this API command to retrieve a switch status.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchService) FindSwitchById(ctx context.Context, pId string) (*switchmswitch.NetworkSwitch, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSwitchFirmwareBySwitchId
//
// Use this API command to get a list of firmware update history.
//
// Path Parameters:
// - pSwitchId string
//		- required
func (s *SwitchMSwitchService) FindSwitchFirmwareBySwitchId(ctx context.Context, pSwitchId string) (*switchmswitch.FirmwareHistoryQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateSwitchMoveByDestinationSwitchGroupId
//
// Use this API command to move a list of switches to a switch group.
//
// Request Body:
//	 - body switchmswitch.SwitchIdList
//
// Path Parameters:
// - pDestinationSwitchGroupId string
//		- required
func (s *SwitchMSwitchService) UpdateSwitchMoveByDestinationSwitchGroupId(ctx context.Context, body switchmswitch.SwitchIdList, pDestinationSwitchGroupId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return errors.New("body cannot be empty")
	}
}
