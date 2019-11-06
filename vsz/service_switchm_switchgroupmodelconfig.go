package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/groupmodelconfig"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchGroupModelConfigService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchGroupModelConfigService(c *APIClient) *SwitchMSwitchGroupModelConfigService {
	s := new(SwitchMSwitchGroupModelConfigService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchGroupModelConfigService() *SwitchMSwitchGroupModelConfigService {
	return NewSwitchMSwitchGroupModelConfigService(ss.apiClient)
}

// FindGroupModelConfigsByQueryCriteria
//
// Use this API command to retrieve the list of group model configs.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchGroupModelConfigService) FindGroupModelConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*groupmodelconfig.GroupModelConfigQueryResult, error) {
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

// UpdateGroupModelConfigsByGroupId
//
// Use this API command to add or remove the model family of a group config.
//
// Request Body:
//	 - body *groupmodelconfig.SelectedIds
//
// Path Parameters:
// - pGroupId string
//		- required
func (s *SwitchMSwitchGroupModelConfigService) UpdateGroupModelConfigsByGroupId(ctx context.Context, body *groupmodelconfig.SelectedIds, pGroupId string) (*groupmodelconfig.UpdateGroupConfigResultList, error) {
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
