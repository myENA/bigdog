package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/registration"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchRegistrationRulesService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchRegistrationRulesService(c *APIClient) *SwitchMSwitchRegistrationRulesService {
	s := new(SwitchMSwitchRegistrationRulesService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchRegistrationRulesService() *SwitchMSwitchRegistrationRulesService {
	return NewSwitchMSwitchRegistrationRulesService(ss.apiClient)
}

// AddRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Request Body:
//	 - body *registration.RegistrationRule
func (s *SwitchMSwitchRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *registration.RegistrationRule) (*registration.CreateResult, error) {
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

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body registration.RuleUUIDs
func (s *SwitchMSwitchRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body registration.RuleUUIDs) (*registration.DeleteMultipleResult, error) {
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

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, pId string) (*registration.DeleteResult, error) {
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
func (s *SwitchMSwitchRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*registration.RuleQueryResultList, error) {
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
//	 - body *registration.RegistrationRule
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *registration.RegistrationRule, pId string) (*registration.ModifyResult, error) {
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
