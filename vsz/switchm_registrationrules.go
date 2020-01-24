package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	var (
		resp *SwitchMRegistrationCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddRegistrationRules, true)
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMRegistrationRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRuleUUIDs) (*SwitchMRegistrationDeleteMultipleResult, error) {
	var (
		resp *SwitchMRegistrationDeleteMultipleResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRules, true)
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, id string) (*SwitchMRegistrationDeleteResult, error) {
	var (
		resp *SwitchMRegistrationDeleteResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*SwitchMRegistrationRuleQueryResultList, error) {
	var (
		resp *SwitchMRegistrationRuleQueryResultList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindRegistrationRules, true)
}

// UpdateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRule
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRule, id string) (*SwitchMRegistrationModifyResult, error) {
	var (
		resp *SwitchMRegistrationModifyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateRegistrationRulesById, true)
}
