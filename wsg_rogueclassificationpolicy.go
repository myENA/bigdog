package bigdog

// API Version: v9_0

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
// Request Body:
//	 - body *WSGProfileCreateRogueApPolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) AddRkszonesRogueApPoliciesByZoneId(ctx context.Context, body *WSGProfileCreateRogueApPolicy, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesRogueApPoliciesByZoneId, true)
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

// DeleteRkszonesRogueApPoliciesById
//
// Use this API command to delete rogue AP policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) DeleteRkszonesRogueApPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRogueApPoliciesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesRogueApPoliciesByZoneId
//
// Use this API command to delete bulk rogue AP policy.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) DeleteRkszonesRogueApPoliciesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRogueApPoliciesByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesRogueApPoliciesById
//
// Use this API command to retrieve rogue AP policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileRogueApPolicy, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRogueApPolicy
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRogueApPoliciesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileRogueApPolicy()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesRogueApPoliciesByZoneId
//
// Use this API command to retrieve a list of rogue AP policy.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileRogueApPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRogueApPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRogueApPoliciesByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileRogueApPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesRogueApPoliciesById
//
// Use this API command to modify rogue AP policy.
//
// Request Body:
//	 - body *WSGProfileUpdateRogueApPolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) PartialUpdateRkszonesRogueApPoliciesById(ctx context.Context, body *WSGProfileUpdateRogueApPolicy, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesRogueApPoliciesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
