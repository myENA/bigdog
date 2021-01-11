package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGRogueClassificationPolicyService struct {
	apiClient *VSZClient
}

func NewWSGRogueClassificationPolicyService(c *VSZClient) *WSGRogueClassificationPolicyService {
	s := new(WSGRogueClassificationPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRogueClassificationPolicyService() *WSGRogueClassificationPolicyService {
	return NewWSGRogueClassificationPolicyService(ss.apiClient)
}

// AddRkszonesRogueApPoliciesByZoneId
//
// Use this API command to create rogue AP policy.
//
// Operation ID: addRkszonesRogueApPoliciesByZoneId
// Operation path: /rkszones/{zoneId}/rogueApPolicies
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateRogueApPolicy
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) AddRkszonesRogueApPoliciesByZoneId(ctx context.Context, body *WSGProfileCreateRogueApPolicy, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesRogueApPoliciesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesRogueApPoliciesById
//
// Use this API command to delete rogue AP policy.
//
// Operation ID: deleteRkszonesRogueApPoliciesById
// Operation path: /rkszones/{zoneId}/rogueApPolicies/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) DeleteRkszonesRogueApPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesRogueApPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesRogueApPoliciesByZoneId
//
// Use this API command to delete bulk rogue AP policy.
//
// Operation ID: deleteRkszonesRogueApPoliciesByZoneId
// Operation path: /rkszones/{zoneId}/rogueApPolicies
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) DeleteRkszonesRogueApPoliciesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesRogueApPoliciesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesRogueApPoliciesById
//
// Use this API command to retrieve rogue AP policy.
//
// Operation ID: findRkszonesRogueApPoliciesById
// Operation path: /rkszones/{zoneId}/rogueApPolicies/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileRogueApPolicyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileRogueApPolicyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileRogueApPolicyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesRogueApPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileRogueApPolicyAPIResponse), err
}

// FindRkszonesRogueApPoliciesByZoneId
//
// Use this API command to retrieve a list of rogue AP policy.
//
// Operation ID: findRkszonesRogueApPoliciesByZoneId
// Operation path: /rkszones/{zoneId}/rogueApPolicies
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileRogueApPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileRogueApPolicyListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileRogueApPolicyListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesRogueApPoliciesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileRogueApPolicyListAPIResponse), err
}

// PartialUpdateRkszonesRogueApPoliciesById
//
// Use this API command to modify rogue AP policy.
//
// Operation ID: partialUpdateRkszonesRogueApPoliciesById
// Operation path: /rkszones/{zoneId}/rogueApPolicies/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileUpdateRogueApPolicy
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) PartialUpdateRkszonesRogueApPoliciesById(ctx context.Context, body *WSGProfileUpdateRogueApPolicy, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesRogueApPoliciesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
