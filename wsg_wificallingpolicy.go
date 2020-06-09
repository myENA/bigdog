package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWiFiCallingPolicyService struct {
	apiClient *VSZClient
}

func NewWSGWiFiCallingPolicyService(c *VSZClient) *WSGWiFiCallingPolicyService {
	s := new(WSGWiFiCallingPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWiFiCallingPolicyService() *WSGWiFiCallingPolicyService {
	return NewWSGWiFiCallingPolicyService(ss.apiClient)
}

// AddWifiCallingWifiCallingPolicy
//
// Use this API command to Create Wi-Fi Calling Policy.
//
// Request Body:
//	 - body *WSGWIFICallingCreateWifiCallingPolicy
func (s *WSGWiFiCallingPolicyService) AddWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingCreateWifiCallingPolicy) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingDeleteBulk) (interface{}, *APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicyById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) FindWifiCallingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWIFICallingPolicyList, *APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicy(ctx context.Context, optionalParams map[string][]string) (*WSGWIFICallingPolicyList, *APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicyById(ctx context.Context, id string) (*WSGWIFICallingPolicy, *APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) PartialUpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string) (*APIResponseMeta, error) {
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
func (s *WSGWiFiCallingPolicyService) UpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string) (*APIResponseMeta, error) {
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
