package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/deploylog"
	"github.com/myENA/ruckus-client/vsz/types/switchm/deploylogitem"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchConfigurationHistoryService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchConfigurationHistoryService(c *APIClient) *SwitchMSwitchConfigurationHistoryService {
	s := new(SwitchMSwitchConfigurationHistoryService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationHistoryService() *SwitchMSwitchConfigurationHistoryService {
	return NewSwitchMSwitchConfigurationHistoryService(ss.apiClient)
}

// FindConfigurationHistory
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistory(ctx context.Context) (*deploylog.ConfigurationHistoryQueryResult, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*deploylog.ConfigurationHistoryQueryResult, error) {
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

// FindConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*deploylogitem.ConfigurationHistoryDetailQueryResult, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*deploylogitem.ConfigurationHistoryDetailQueryResult, error) {
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
