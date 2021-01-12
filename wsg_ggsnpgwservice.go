package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGGGSNPGWServiceService struct {
	apiClient *VSZClient
}

func NewWSGGGSNPGWServiceService(c *VSZClient) *WSGGGSNPGWServiceService {
	s := new(WSGGGSNPGWServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGGSNPGWServiceService() *WSGGGSNPGWServiceService {
	return NewWSGGGSNPGWServiceService(ss.apiClient)
}

// DeleteServicesGgsnDnsServerList
//
// Use this API command to Disable the dns server list of GGSN/PGW.
//
// Operation ID: deleteServicesGgsnDnsServerList
// Operation path: /services/ggsn/dnsServerList
// Success code: 204 (No Content)
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnDnsServerList(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesGgsnDnsServerList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesGgsnGgsnList
//
// Use this API command to disable the ggsn server list of GGSN/PGW.
//
// Operation ID: deleteServicesGgsnGgsnList
// Operation path: /services/ggsn/ggsnList
// Success code: 204 (No Content)
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnGgsnList(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesGgsnGgsnList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindServicesGgsn
//
// Use this API command to retrieve GGSN/PGW setting.
//
// Operation ID: findServicesGgsn
// Operation path: /services/ggsn
// Success code: 200 (OK)
func (s *WSGGGSNPGWServiceService) FindServicesGgsn(ctx context.Context, mutators ...RequestMutator) (*WSGServiceGgsnConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceGgsnConfigAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesGgsn, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceGgsnConfigAPIResponse), err
}

// PartialUpdateServicesGgsn
//
// Use this API command to modify the setting of GGSN/PGW.
//
// Operation ID: partialUpdateServicesGgsn
// Operation path: /services/ggsn
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceGgsnConfig
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsn(ctx context.Context, body *WSGServiceGgsnConfig, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesGgsn, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesGgsnDnsServerList
//
// Use this API command to modify the dns server list of GGSN/PGW.
//
// Operation ID: partialUpdateServicesGgsnDnsServerList
// Operation path: /services/ggsn/dnsServerList
// Success code: 204 (No Content)
//
// Request body:
//	 - body WSGServiceDnsServerList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnDnsServerList(ctx context.Context, body WSGServiceDnsServerList, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesGgsnDnsServerList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesGgsnGgsnList
//
// Use this API command to modify the ggsn server list of GGSN/PGW.
//
// Operation ID: partialUpdateServicesGgsnGgsnList
// Operation path: /services/ggsn/ggsnList
// Success code: 204 (No Content)
//
// Request body:
//	 - body WSGServiceGgsnList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGgsnList(ctx context.Context, body WSGServiceGgsnList, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesGgsnGgsnList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesGgsnGtpSettings
//
// Use this API command to modify the gtp setting of GGSN/PGW.
//
// Operation ID: partialUpdateServicesGgsnGtpSettings
// Operation path: /services/ggsn/gtpSettings
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceGtpSettings
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGtpSettings(ctx context.Context, body *WSGServiceGtpSettings, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesGgsnGtpSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
