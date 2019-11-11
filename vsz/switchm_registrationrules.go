package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMRegistrationRulesService struct {
	apiClient *APIClient
}

func NewSwitchMRegistrationRulesService(c *APIClient) *SwitchMRegistrationRulesService {
	s := new(SwitchMRegistrationRulesService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMRegistrationRulesService() *SwitchMRegistrationRulesService {
	return NewSwitchMRegistrationRulesService(ss.apiClient)
}

// AddRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRule
func (s *SwitchMRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *SwitchMRegistrationRule) (*SwitchMRegistrationCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMRegistrationRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRuleUUIDs) (*SwitchMRegistrationDeleteMultipleResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, pId string) (*SwitchMRegistrationDeleteResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*SwitchMRegistrationRuleQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRule
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRule, pId string) (*SwitchMRegistrationModifyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
