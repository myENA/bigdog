package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGBonjourFencingPolicyService struct {
	apiClient *VSZClient
}

func NewWSGBonjourFencingPolicyService(c *VSZClient) *WSGBonjourFencingPolicyService {
	s := new(WSGBonjourFencingPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourFencingPolicyService() *WSGBonjourFencingPolicyService {
	return NewWSGBonjourFencingPolicyService(ss.apiClient)
}

// AddRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to create Bonjour Fencing Policy.
//
// Operation ID: addRkszonesBonjourFencingPolicyByZoneId
// Operation path: /rkszones/{zoneId}/bonjourFencingPolicy
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateBonjourFencingPolicy
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) AddRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, body *WSGProfileCreateBonjourFencingPolicy, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesBonjourFencingPolicyByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesBonjourFencingPolicy
//
// Use this API command to delete Bulk Bonjour Fencing Policy.
//
// Operation ID: deleteRkszonesBonjourFencingPolicy
// Operation path: /rkszones/bonjourFencingPolicy
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesBonjourFencingPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesBonjourFencingPolicyById
//
// Use this API command to delete Bonjour Fencing Policy.
//
// Operation ID: deleteRkszonesBonjourFencingPolicyById
// Operation path: /rkszones/bonjourFencingPolicy/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesBonjourFencingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApsBonjourFencingStatisticByApMac
//
// Use this API command to get Bonjour Fencing Statistic per AP.
//
// Operation ID: findApsBonjourFencingStatisticByApMac
// Operation path: /aps/{apMac}/bonjourFencingStatistic
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGBonjourFencingPolicyService) FindApsBonjourFencingStatisticByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGProfileBonjourFencingStatisticAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBonjourFencingStatisticAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsBonjourFencingStatisticByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBonjourFencingStatisticAPIResponse), err
}

// FindRkszonesBonjourFencingPolicyById
//
// Use this API command to retrieve Bonjour Fencing Policy.
//
// Operation ID: findRkszonesBonjourFencingPolicyById
// Operation path: /rkszones/{zoneId}/bonjourFencingPolicy/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileBonjourFencingPolicyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBonjourFencingPolicyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesBonjourFencingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBonjourFencingPolicyAPIResponse), err
}

// FindRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Operation ID: findRkszonesBonjourFencingPolicyByZoneId
// Operation path: /rkszones/{zoneId}/bonjourFencingPolicy
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileBonjourFencingPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBonjourFencingPolicyListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesBonjourFencingPolicyByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBonjourFencingPolicyListAPIResponse), err
}

// FindServicesBonjourFencingPolicyByQueryCriteria
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Operation ID: findServicesBonjourFencingPolicyByQueryCriteria
// Operation path: /query/services/bonjourFencingPolicy
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBonjourFencingPolicyService) FindServicesBonjourFencingPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileBonjourFencingPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBonjourFencingPolicyListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesBonjourFencingPolicyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBonjourFencingPolicyListAPIResponse), err
}

// PartialUpdateRkszonesBonjourFencingPolicyById
//
// Use this API command to modify the configuration of Bonjour Fencing Policy.
//
// Operation ID: partialUpdateRkszonesBonjourFencingPolicyById
// Operation path: /rkszones/{zoneId}/bonjourFencingPolicy/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyBonjourFencingPolicy
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) PartialUpdateRkszonesBonjourFencingPolicyById(ctx context.Context, body *WSGProfileModifyBonjourFencingPolicy, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesBonjourFencingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
