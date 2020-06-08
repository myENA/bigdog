package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWificallingpolicyService struct {
	apiClient *VSZClient
}

func NewWSGWificallingpolicyService(c *VSZClient) *WSGWificallingpolicyService {
	s := new(WSGWificallingpolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWificallingpolicyService() *WSGWificallingpolicyService {
	return NewWSGWificallingpolicyService(ss.apiClient)
}

// AddWifiCallingWifiCallingPolicy
//
// Use this API command to Create Wi-Fi Calling Policy.
//
// Request Body:
//	 - body *WSGWIFICallingCreateWifiCallingPolicy
func (s *WSGWificallingpolicyService) AddWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingCreateWifiCallingPolicy) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddWifiCallingWifiCallingPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteWifiCallingWifiCallingPolicy
//
// Use this API command to Delete bulk Wi-Fi Calling policies.
//
// Request Body:
//	 - body *WSGWIFICallingDeleteBulk
func (s *WSGWificallingpolicyService) DeleteWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingDeleteBulk) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteWifiCallingWifiCallingPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteWifiCallingWifiCallingPolicyById
//
// Use this API command to Delete a Wi-Fi Calling policy by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWificallingpolicyService) DeleteWifiCallingWifiCallingPolicyById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteWifiCallingWifiCallingPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindWifiCallingByQueryCriteria
//
// Use this API command to Query Wi-Fi Calling Policy List.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWificallingpolicyService) FindWifiCallingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWIFICallingPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWIFICallingPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindWifiCallingByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWIFICallingPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindWifiCallingWifiCallingPolicy
//
// Use this API command to Retrieve List of Wi-Fi Calling Policy.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWificallingpolicyService) FindWifiCallingWifiCallingPolicy(ctx context.Context, optionalParams map[string][]string) (*WSGWIFICallingPolicyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWIFICallingPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindWifiCallingWifiCallingPolicy, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWIFICallingPolicyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindWifiCallingWifiCallingPolicyById
//
// Use this API command to Retrieve Wi-Fi Calling Policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWificallingpolicyService) FindWifiCallingWifiCallingPolicyById(ctx context.Context, id string) (*WSGWIFICallingPolicy, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWIFICallingPolicy
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindWifiCallingWifiCallingPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWIFICallingPolicy()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify a Wi-Fi Calling policy.
//
// Request Body:
//	 - body *WSGWIFICallingModifyWifiCallingPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWificallingpolicyService) PartialUpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateWifiCallingWifiCallingPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify Entire Wi-Fi Calling policy.
//
// Request Body:
//	 - body *WSGWIFICallingModifyWifiCallingPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGWificallingpolicyService) UpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateWifiCallingWifiCallingPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
