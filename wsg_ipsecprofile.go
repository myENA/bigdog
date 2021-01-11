package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGIPSECProfileService struct {
	apiClient *VSZClient
}

func NewWSGIPSECProfileService(c *VSZClient) *WSGIPSECProfileService {
	s := new(WSGIPSECProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIPSECProfileService() *WSGIPSECProfileService {
	return NewWSGIPSECProfileService(ss.apiClient)
}

// AddProfilesTunnelIpsec
//
// Create a new ipsec.
//
// Operation ID: addProfilesTunnelIpsec
// Operation path: /profiles/tunnel/ipsec
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateIpsecProfile
func (s *WSGIPSECProfileService) AddProfilesTunnelIpsec(ctx context.Context, body *WSGProfileCreateIpsecProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesTunnelIpsec, true)
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

// DeleteProfilesTunnelIpsec
//
// Delete multiple ipsec.
//
// Operation ID: deleteProfilesTunnelIpsec
// Operation path: /profiles/tunnel/ipsec
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsec(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsec, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesTunnelIpsecById
//
// Delete a ipsec.
//
// Operation ID: deleteProfilesTunnelIpsecById
// Operation path: /profiles/tunnel/ipsec/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsecById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsecById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesTunnelIpsec
//
// Retrieve a list of IPSEC.
//
// Operation ID: findProfilesTunnelIpsec
// Operation path: /profiles/tunnel/ipsec
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsec(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTunnelIpsec, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesTunnelIpsecById
//
// Retrieve a IPSEC.
//
// Operation ID: findProfilesTunnelIpsecById
// Operation path: /profiles/tunnel/ipsec/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileIpsecProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileIpsecProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileIpsecProfileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTunnelIpsecById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileIpsecProfileAPIResponse), err
}

// FindProfilesTunnelIpsecByQueryCriteria
//
// Query a list of IPSEC.
//
// Operation ID: findProfilesTunnelIpsecByQueryCriteria
// Operation path: /profiles/tunnel/ipsec/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileIpsecProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileIpsecProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileIpsecProfileListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesTunnelIpsecByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileIpsecProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileIpsecProfileListAPIResponse), err
}

// PartialUpdateProfilesTunnelIpsecById
//
// Modify a specific ipsec basic.
//
// Operation ID: partialUpdateProfilesTunnelIpsecById
// Operation path: /profiles/tunnel/ipsec/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyIpsecProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) PartialUpdateProfilesTunnelIpsecById(ctx context.Context, body *WSGProfileModifyIpsecProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelIpsecById, true)
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
