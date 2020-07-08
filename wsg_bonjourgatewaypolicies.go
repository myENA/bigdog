package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGBonjourGatewayPoliciesService struct {
	apiClient *VSZClient
}

func NewWSGBonjourGatewayPoliciesService(c *VSZClient) *WSGBonjourGatewayPoliciesService {
	s := new(WSGBonjourGatewayPoliciesService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourGatewayPoliciesService() *WSGBonjourGatewayPoliciesService {
	return NewWSGBonjourGatewayPoliciesService(ss.apiClient)
}

// AddRkszonesBonjourGatewayEnableByZoneId
//
// Operation ID: addRkszonesBonjourGatewayEnableByZoneId
//
// Use this API command to enable/disable bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneModifyBonjourGatewayEnable
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayEnableByZoneId(ctx context.Context, body *WSGZoneModifyBonjourGatewayEnable, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesBonjourGatewayEnableByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddRkszonesBonjourGatewayPoliciesByZoneId
//
// Operation ID: addRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to create bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneCreateBonjourGatewayPolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, body *WSGZoneCreateBonjourGatewayPolicy, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesBonjourGatewayPoliciesByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesBonjourGatewayPoliciesById
//
// Operation ID: deleteRkszonesBonjourGatewayPoliciesById
//
// Use this API command to delete bonjour gateway policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) DeleteRkszonesBonjourGatewayPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBonjourGatewayPoliciesById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesBonjourGatewayPoliciesById
//
// Operation ID: findRkszonesBonjourGatewayPoliciesById
//
// Use this API command to retrieve bonjour gateway policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGZoneBonjourGatewayPolicyConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGZoneBonjourGatewayPolicyConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourGatewayPoliciesById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGZoneBonjourGatewayPolicyConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesBonjourGatewayPoliciesByZoneId
//
// Operation ID: findRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to retrieve a list of bonjour gateway policies.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGZoneBonjourGatewayPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGZoneBonjourGatewayPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourGatewayPoliciesByZoneId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGZoneBonjourGatewayPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesBonjourPolicyByQueryCriteria
//
// Operation ID: findServicesBonjourPolicyByQueryCriteria
//
// Query bonjourPolicy Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBonjourGatewayPoliciesService) FindServicesBonjourPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesBonjourPolicyByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesBonjourGatewayPoliciesById
//
// Operation ID: partialUpdateRkszonesBonjourGatewayPoliciesById
//
// Use this API command to modify the configuration of bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneModifyBonjourGatewayPolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) PartialUpdateRkszonesBonjourGatewayPoliciesById(ctx context.Context, body *WSGZoneModifyBonjourGatewayPolicy, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesBonjourGatewayPoliciesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
