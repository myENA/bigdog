package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/stack"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchStackService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchStackService(c *APIClient) *SwitchMSwitchStackService {
	s := new(SwitchMSwitchStackService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchStackService() *SwitchMSwitchStackService {
	serv := new(SwitchMSwitchStackService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddStack
//
// Use this API command to create a stack configuration.
//
// Request Body:
//	 - body stack.StackConfigList
func (s *SwitchMSwitchStackService) AddStack(ctx context.Context, body stack.StackConfigList) (stack.AuditIdList, error) {
}

// FindStackBySwitchId
//
// Use this API command to retrieve a stack configuration configured via SZ.
//
// Path Parameters:
// - pSwitchId string
//		- required
func (s *SwitchMSwitchStackService) FindStackBySwitchId(ctx context.Context, pSwitchId string) (*stack.StackConfig, error) {
}

// FindStackMemberBySerialNumber
//
// Use this API command to retrieve the member of switches in a stack.
//
// Path Parameters:
// - pSerialNumber string
//		- required
func (s *SwitchMSwitchStackService) FindStackMemberBySerialNumber(ctx context.Context, pSerialNumber string) (*stack.List, error) {
}
