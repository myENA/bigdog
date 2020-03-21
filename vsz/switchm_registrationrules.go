package vsz

// API Version: v9_0

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
func (s *SwitchMRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *SwitchMRegistrationRule) (*SwitchMRegistrationCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddRegistrationRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMRegistrationRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRuleUUIDs) (*SwitchMRegistrationDeleteMultipleResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationDeleteMultipleResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationDeleteMultipleResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, id string) (*SwitchMRegistrationDeleteResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationDeleteResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationDeleteResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*SwitchMRegistrationRuleQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationRuleQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindRegistrationRules, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationRuleQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRule, id string) (*SwitchMRegistrationModifyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationModifyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateRegistrationRulesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationModifyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
