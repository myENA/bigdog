package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchGroupService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchGroupService(c *APIClient) *SwitchMSwitchGroupService {
	s := new(SwitchMSwitchGroupService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchGroupService() *SwitchMSwitchGroupService {
	return NewSwitchMSwitchGroupService(ss.apiClient)
}

// AddGroup
//
// Use this API command to create a new switch group under an existing domain or switch group.
//
// Request Body:
//	 - body *group.SwitchGroup
func (s *SwitchMSwitchGroupService) AddGroup(ctx context.Context, body *group.SwitchGroup) (*group.AuditId, error) {
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

// DeleteGroupBySwitchGroupId
//
// Use this API command to delete a switch group.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) DeleteGroupBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*group.DeleteSwitchGroupResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindGroupBySwitchGroupId
//
// Use this API command to retrieve switch group detail.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) FindGroupBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*group.SwitchGroupQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindGroupIdsByDomainByDomainId
//
// Use this API command to retrieve the switch groups by domain ID.
//
// Path Parameters:
// - pDomainId string
//		- required
func (s *SwitchMSwitchGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, pDomainId string) (*group.GroupsByIdsQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name or description.
//
// Request Body:
//	 - body *group.UpdateSwitchGroup
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) PartialUpdateGroupBySwitchGroupId(ctx context.Context, body *group.UpdateSwitchGroup, pSwitchGroupId string) (*group.UpdateSwitchGroupResult, error) {
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
