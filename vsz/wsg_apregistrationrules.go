package vsz

// API Version: v9_0

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
func (s *WSGAPRegistrationRulesService) AddApRules(ctx context.Context, body *WSGAPRulesCreateApRule) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteApRulesById
//
// Use this API command to delete AP Registration Rules profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) DeleteApRulesById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApRulesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApRules
//
// Use this API command to retrieve a list of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context) (*WSGAPRulesApRuleList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPRulesApRuleList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApRules, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPRulesApRuleList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindApRulesById
//
// Use this API command to retrieve AP Registration Rules profile by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, id string) (*WSGAPRulesApRuleConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPRulesApRuleConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApRulesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPRulesApRuleConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindApRulesPriorityDownById
//
// Use this API command to move Priority Down of AP Registration Rules profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApRulesPriorityDownById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApRulesPriorityUpById
//
// Use this API command to move Priority Up of AP Registration Rules profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApRulesPriorityUpById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateApRulesById
//
// Use this API command to modify the configuration of AP Registration Rules profile.
//
// Request Body:
//	 - body *WSGAPRulesModifyApRule
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) PartialUpdateApRulesById(ctx context.Context, body *WSGAPRulesModifyApRule, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateApRulesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
