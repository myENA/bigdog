package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Use this API command to enable/disable bonjour gateway policy.
//
// Operation ID: addRkszonesBonjourGatewayEnableByZoneId
// Operation path: /rkszones/{zoneId}/bonjourGateway/enable
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGZoneModifyBonjourGatewayEnable
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayEnableByZoneId(ctx context.Context, body *WSGZoneModifyBonjourGatewayEnable, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesBonjourGatewayEnableByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to create bonjour gateway policy.
//
// Operation ID: addRkszonesBonjourGatewayPoliciesByZoneId
// Operation path: /rkszones/{zoneId}/bonjourGateway/policies
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGZoneCreateBonjourGatewayPolicy
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, body *WSGZoneCreateBonjourGatewayPolicy, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesBonjourGatewayPoliciesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesBonjourGatewayPoliciesById
//
// Use this API command to delete bonjour gateway policy.
//
// Operation ID: deleteRkszonesBonjourGatewayPoliciesById
// Operation path: /rkszones/{zoneId}/bonjourGateway/policies/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) DeleteRkszonesBonjourGatewayPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesBonjourGatewayPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesBonjourGatewayPoliciesById
//
// Use this API command to retrieve bonjour gateway policy.
//
// Operation ID: findRkszonesBonjourGatewayPoliciesById
// Operation path: /rkszones/{zoneId}/bonjourGateway/policies/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGZoneBonjourGatewayPolicyConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGZoneBonjourGatewayPolicyConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesBonjourGatewayPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneBonjourGatewayPolicyConfigurationAPIResponse), err
}

// FindRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to retrieve a list of bonjour gateway policies.
//
// Operation ID: findRkszonesBonjourGatewayPoliciesByZoneId
// Operation path: /rkszones/{zoneId}/bonjourGateway/policies
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGZoneBonjourGatewayPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGZoneBonjourGatewayPolicyListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesBonjourGatewayPoliciesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneBonjourGatewayPolicyListAPIResponse), err
}

// FindServicesBonjourPolicyByQueryCriteria
//
// Query bonjourPolicy Profiles with specified filters.
//
// Operation ID: findServicesBonjourPolicyByQueryCriteria
// Operation path: /query/services/bonjourPolicy
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBonjourGatewayPoliciesService) FindServicesBonjourPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesBonjourPolicyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateRkszonesBonjourGatewayPoliciesById
//
// Use this API command to modify the configuration of bonjour gateway policy.
//
// Operation ID: partialUpdateRkszonesBonjourGatewayPoliciesById
// Operation path: /rkszones/{zoneId}/bonjourGateway/policies/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGZoneModifyBonjourGatewayPolicy
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) PartialUpdateRkszonesBonjourGatewayPoliciesById(ctx context.Context, body *WSGZoneModifyBonjourGatewayPolicy, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesBonjourGatewayPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
