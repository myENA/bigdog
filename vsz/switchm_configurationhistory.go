package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMConfigurationHistoryService struct {
	apiClient *APIClient
}

func NewSwitchMConfigurationHistoryService(c *APIClient) *SwitchMConfigurationHistoryService {
	s := new(SwitchMConfigurationHistoryService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationHistoryService() *SwitchMConfigurationHistoryService {
	return NewSwitchMConfigurationHistoryService(ss.apiClient)
}

// FindConfigurationHistory
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistory(ctx context.Context) (*SwitchMDeployLogConfigurationHistoryQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindConfigurationHistoryByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeployLogConfigurationHistoryQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindConfigurationHistoryDetailByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
