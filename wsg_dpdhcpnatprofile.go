package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGDPDHCPNATProfileService struct {
	apiClient *VSZClient
}

func NewWSGDPDHCPNATProfileService(c *VSZClient) *WSGDPDHCPNATProfileService {
	s := new(WSGDPDHCPNATProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPDHCPNATProfileService() *WSGDPDHCPNATProfileService {
	return NewWSGDPDHCPNATProfileService(ss.apiClient)
}

// AddDpProfileSettings
//
// Use this API command to create DP DHCP & NAT profile setting.
//
// Operation ID: addDpProfileSettings
// Operation path: /dpProfileSettings
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDPProfileSettingBO
func (s *WSGDPDHCPNATProfileService) AddDpProfileSettings(ctx context.Context, body *WSGDPProfileSettingBO, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpProfileSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteDpProfileSettings
//
// Use this API command to delete DP DHCP & NAT profile settings.
//
// Operation ID: deleteDpProfileSettings
// Operation path: /dpProfileSettings
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettings(ctx context.Context, body *WSGDPProfileBulkDelete, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpProfileSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteDpProfileSettingsByDpKey
//
// Use this API command to delete DP DHCP & NAT profile setting.
//
// Operation ID: deleteDpProfileSettingsByDpKey
// Operation path: /dpProfileSettings/{dpKey}
// Success code: 204 (No Content)
//
// Required parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettingsByDpKey(ctx context.Context, dpKey string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpProfileSettingsByDpKey, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("dpKey", dpKey)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindDpProfileSettings
//
// Use this API command to retrieve DP DHCP & NAT profile setting list.
//
// Operation ID: findDpProfileSettings
// Operation path: /dpProfileSettings
// Success code: 200 (OK)
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettings(ctx context.Context, mutators ...RequestMutator) (*WSGDPProfileSettingBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileSettingBOListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpProfileSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGDPProfileSettingBOListAPIResponse), err
}

// FindDpProfileSettingsByDpKey
//
// Use this API command to retrieve DP DHCP & NAT profile setting.
//
// Operation ID: findDpProfileSettingsByDpKey
// Operation path: /dpProfileSettings/{dpKey}
// Success code: 200 (OK)
//
// Required parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettingsByDpKey(ctx context.Context, dpKey string, mutators ...RequestMutator) (*WSGDPProfileSettingBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileSettingBOAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpProfileSettingsByDpKey, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("dpKey", dpKey)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGDPProfileSettingBOAPIResponse), err
}

// UpdateDpProfileSettingsByDpKey
//
// Use this API command to modify DP DHCP & NAT profile setting.
//
// Operation ID: updateDpProfileSettingsByDpKey
// Operation path: /dpProfileSettings/{dpKey}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDPProfileSettingBO
//
// Required parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) UpdateDpProfileSettingsByDpKey(ctx context.Context, body *WSGDPProfileSettingBO, dpKey string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpProfileSettingsByDpKey, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("dpKey", dpKey)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
