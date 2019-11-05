package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aprules"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAPRegistrationRulesService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGAPRegistrationRulesService(c *APIClient) *WSGAPRegistrationRulesService {
	s := new(WSGAPRegistrationRulesService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGAPRegistrationRulesService() *WSGAPRegistrationRulesService {
	return NewWSGAPRegistrationRulesService(ss.apiClient)
}

// AddApRules
//
// Use this API command to create AP Registration Rules profile.
//
// Request Body:
//	 - body *aprules.CreateApRule
func (s *WSGAPRegistrationRulesService) AddApRules(ctx context.Context, body *aprules.CreateApRule) (*common.CreateResult, error) {
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

// DeleteApRulesById
//
// Use this API command to delete AP Registration Rules profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) DeleteApRulesById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApRules
//
// Use this API command to retrieve a list of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context) (*aprules.ApRuleList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApRulesById
//
// Use this API command to retrieve AP Registration Rules profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, pId string) (*aprules.ApRuleConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApRulesPriorityDownById
//
// Use this API command to move Priority Down of AP Registration Rules profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApRulesPriorityUpById
//
// Use this API command to move Priority Up of AP Registration Rules profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateApRulesById
//
// Use this API command to modify the basic information of AP Registration Rules profile.
//
// Request Body:
//	 - body *aprules.ModifyApRule
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) PartialUpdateApRulesById(ctx context.Context, body *aprules.ModifyApRule, pId string) (*common.EmptyResult, error) {
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
