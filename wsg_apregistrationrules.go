package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGAPRegistrationRulesService struct {
	apiClient *VSZClient
}

func NewWSGAPRegistrationRulesService(c *VSZClient) *WSGAPRegistrationRulesService {
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
// Operation ID: addApRules
// Operation path: /apRules
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAPRulesCreateApRule
func (s *WSGAPRegistrationRulesService) AddApRules(ctx context.Context, body *WSGAPRulesCreateApRule, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteApRulesById
//
// Use this API command to delete AP Registration Rules profile.
//
// Operation ID: deleteApRulesById
// Operation path: /apRules/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) DeleteApRulesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApRules
//
// Use this API command to retrieve a list of AP Registration Rules profile.
//
// Operation ID: findApRules
// Operation path: /apRules
// Success code: 200 (OK)
func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context, mutators ...RequestMutator) (*WSGAPRulesApRuleListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPRulesApRuleListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPRulesApRuleListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPRulesApRuleListAPIResponse), err
}

// FindApRulesById
//
// Use this API command to retrieve AP Registration Rules profile by ID.
//
// Operation ID: findApRulesById
// Operation path: /apRules/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAPRulesApRuleConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPRulesApRuleConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPRulesApRuleConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPRulesApRuleConfigurationAPIResponse), err
}

// FindApRulesPriorityDownById
//
// Use this API command to move Priority Down of AP Registration Rules profile.
//
// Operation ID: findApRulesPriorityDownById
// Operation path: /apRules/priorityDown/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApRulesPriorityDownById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApRulesPriorityUpById
//
// Use this API command to move Priority Up of AP Registration Rules profile.
//
// Operation ID: findApRulesPriorityUpById
// Operation path: /apRules/priorityUp/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApRulesPriorityUpById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateApRulesById
//
// Use this API command to modify the configuration of AP Registration Rules profile.
//
// Operation ID: partialUpdateApRulesById
// Operation path: /apRules/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAPRulesModifyApRule
//
// Required parameters:
// - id string
//		- required
func (s *WSGAPRegistrationRulesService) PartialUpdateApRulesById(ctx context.Context, body *WSGAPRulesModifyApRule, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateApRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
