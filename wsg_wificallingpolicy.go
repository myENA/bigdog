package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Operation ID: addWifiCallingWifiCallingPolicy
// Operation path: /wifiCalling/wifiCallingPolicy
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWIFICallingCreateWifiCallingPolicy
func (s *WSGWiFiCallingPolicyService) AddWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingCreateWifiCallingPolicy, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddWifiCallingWifiCallingPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteWifiCallingWifiCallingPolicy
//
// Use this API command to Delete bulk Wi-Fi Calling policies.
//
// Operation ID: deleteWifiCallingWifiCallingPolicy
// Operation path: /wifiCalling/wifiCallingPolicy
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGWIFICallingDeleteBulk
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicy(ctx context.Context, body *WSGWIFICallingDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteWifiCallingWifiCallingPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteWifiCallingWifiCallingPolicyById
//
// Use this API command to Delete a Wi-Fi Calling policy by ID.
//
// Operation ID: deleteWifiCallingWifiCallingPolicyById
// Operation path: /wifiCalling/wifiCallingPolicy/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteWifiCallingWifiCallingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindWifiCallingByQueryCriteria
//
// Use this API command to Query Wi-Fi Calling Policy List.
//
// Operation ID: findWifiCallingByQueryCriteria
// Operation path: /wifiCalling/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWiFiCallingPolicyService) FindWifiCallingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWIFICallingPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGWIFICallingPolicyListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindWifiCallingByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWIFICallingPolicyListAPIResponse), err
}

// FindWifiCallingWifiCallingPolicy
//
// Use this API command to Retrieve List of Wi-Fi Calling Policy.
//
// Operation ID: findWifiCallingWifiCallingPolicy
// Operation path: /wifiCalling/wifiCallingPolicy
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicy(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGWIFICallingPolicyListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGWIFICallingPolicyListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindWifiCallingWifiCallingPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWIFICallingPolicyListAPIResponse), err
}

// FindWifiCallingWifiCallingPolicyById
//
// Use this API command to Retrieve Wi-Fi Calling Policy.
//
// Operation ID: findWifiCallingWifiCallingPolicyById
// Operation path: /wifiCalling/wifiCallingPolicy/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGWIFICallingPolicyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGWIFICallingPolicyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindWifiCallingWifiCallingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWIFICallingPolicyAPIResponse), err
}

// PartialUpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify a Wi-Fi Calling policy.
//
// Operation ID: partialUpdateWifiCallingWifiCallingPolicyById
// Operation path: /wifiCalling/wifiCallingPolicy/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGWIFICallingModifyWifiCallingPolicy
//
// Required parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) PartialUpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateWifiCallingWifiCallingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify Entire Wi-Fi Calling policy.
//
// Operation ID: updateWifiCallingWifiCallingPolicyById
// Operation path: /wifiCalling/wifiCallingPolicy/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGWIFICallingModifyWifiCallingPolicy
//
// Required parameters:
// - id string
//		- required
func (s *WSGWiFiCallingPolicyService) UpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *WSGWIFICallingModifyWifiCallingPolicy, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateWifiCallingWifiCallingPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
