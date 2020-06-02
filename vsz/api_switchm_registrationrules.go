package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMRegistrationrulesService struct {
	apiClient *APIClient
}

func NewSwitchMRegistrationrulesService(c *APIClient) *SwitchMRegistrationrulesService {
	s := new(SwitchMRegistrationrulesService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMRegistrationrulesService() *SwitchMRegistrationrulesService {
	return NewSwitchMRegistrationrulesService(ss.apiClient)
}

// AddRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRule
func (s *SwitchMRegistrationrulesService) AddRegistrationRules(ctx context.Context, body *SwitchMRegistrationRule) (*SwitchMRegistrationCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddRegistrationRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMRegistrationRuleUUIDs
func (s *SwitchMRegistrationrulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRuleUUIDs) (*SwitchMRegistrationDeleteMultipleResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationDeleteMultipleResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationrulesService) DeleteRegistrationRulesById(ctx context.Context, id string) (*SwitchMRegistrationDeleteResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationDeleteResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationrulesService) FindRegistrationRules(ctx context.Context) (*SwitchMRegistrationRuleQueryResultList, *APIResponseMeta, error) {
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *SwitchMRegistrationrulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRule, id string) (*SwitchMRegistrationModifyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateRegistrationRulesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationModifyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
