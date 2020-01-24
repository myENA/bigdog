package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGAPRegistrationRulesService struct {
	apiClient *APIClient
}

func NewWSGAPRegistrationRulesService(c *APIClient) *WSGAPRegistrationRulesService {
	s := new(WSGAPRegistrationRulesService)
	s.apiClient = c
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
//	 - body *WSGAPRulesCreateApRule
func (s *WSGAPRegistrationRulesService) AddApRules(ctx context.Context, body *WSGAPRulesCreateApRule) (*WSGCommonCreateResult, error) {
	var (
		resp *WSGCommonCreateResult
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddApRules, true)
}

// DeleteApRulesById
//
// Use this API command to delete AP Registration Rules profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) DeleteApRulesById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteApRulesById, true)
}

// FindApRules
//
// Use this API command to retrieve a list of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context) (*WSGAPRulesApRuleList, error) {
	var (
		resp *WSGAPRulesApRuleList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApRules, true)
}

// FindApRulesById
//
// Use this API command to retrieve AP Registration Rules profile by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, id string) (*WSGAPRulesApRuleConfiguration, error) {
	var (
		resp *WSGAPRulesApRuleConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApRulesById, true)
}

// FindApRulesPriorityDownById
//
// Use this API command to move Priority Down of AP Registration Rules profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApRulesPriorityDownById, true)
}

// FindApRulesPriorityUpById
//
// Use this API command to move Priority Up of AP Registration Rules profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApRulesPriorityUpById, true)
}

// PartialUpdateApRulesById
//
// Use this API command to modify the basic information of AP Registration Rules profile.
//
// Request Body:
//	 - body *WSGAPRulesModifyApRule
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) PartialUpdateApRulesById(ctx context.Context, body *WSGAPRulesModifyApRule, id string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateApRulesById, true)
}
