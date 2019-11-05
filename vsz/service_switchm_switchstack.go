package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/stack"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchStackService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchStackService(c *APIClient) *SwitchMSwitchStackService {
	s := new(SwitchMSwitchStackService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchStackService() *SwitchMSwitchStackService {
	return NewSwitchMSwitchStackService(ss.apiClient)
}

// AddStack
//
// Use this API command to create a stack configuration.
//
// Request Body:
//	 - body stack.StackConfigList
func (s *SwitchMSwitchStackService) AddStack(ctx context.Context, body stack.StackConfigList) (stack.AuditIdList, error) {
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

// FindStackBySwitchId
//
// Use this API command to retrieve a stack configuration configured via SZ.
//
// Path Parameters:
// - pSwitchId string
//		- required
func (s *SwitchMSwitchStackService) FindStackBySwitchId(ctx context.Context, pSwitchId string) (*stack.StackConfig, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindStackMemberBySerialNumber
//
// Use this API command to retrieve the member of switches in a stack.
//
// Path Parameters:
// - pSerialNumber string
//		- required
func (s *SwitchMSwitchStackService) FindStackMemberBySerialNumber(ctx context.Context, pSerialNumber string) (*stack.List, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
