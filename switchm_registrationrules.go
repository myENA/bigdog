package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMRegistrationRulesService struct {
	apiClient *VSZClient
}

func NewSwitchMRegistrationRulesService(c *VSZClient) *SwitchMRegistrationRulesService {
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
//	 - body *SwitchMSwitchRegistrationRulesRegistrationRule
func (s *SwitchMRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *SwitchMSwitchRegistrationRulesRegistrationRule) (*SwitchMSwitchRegistrationRulesCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchRegistrationRulesCreateResult
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
	resp = NewSwitchMSwitchRegistrationRulesCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMSwitchRegistrationRulesRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMSwitchRegistrationRulesRuleUUIDs) (*SwitchMSwitchRegistrationRulesDeleteMultipleResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchRegistrationRulesDeleteMultipleResult
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
	resp = NewSwitchMSwitchRegistrationRulesDeleteMultipleResult()
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
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, id string) (*SwitchMSwitchRegistrationRulesDeleteResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchRegistrationRulesDeleteResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchRegistrationRulesDeleteResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*SwitchMSwitchRegistrationRulesRuleQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchRegistrationRulesRuleQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindRegistrationRules, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchRegistrationRulesRuleQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Request Body:
//	 - body *SwitchMSwitchRegistrationRulesRegistrationRule
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMSwitchRegistrationRulesRegistrationRule, id string) (*SwitchMSwitchRegistrationRulesModifyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchRegistrationRulesModifyResult
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
	resp = NewSwitchMSwitchRegistrationRulesModifyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
