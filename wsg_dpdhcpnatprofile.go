package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
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
// Request Body:
//	 - body *WSGDPProfileSettingBO
func (s *WSGDPDHCPNATProfileService) AddDpProfileSettings(ctx context.Context, body *WSGDPProfileSettingBO, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpProfileSettings, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteDpProfileSettings
//
// Use this API command to delete DP DHCP & NAT profile settings.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettings(ctx context.Context, body *WSGDPProfileBulkDelete, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpProfileSettings, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteDpProfileSettingsByDpKey
//
// Use this API command to delete DP DHCP & NAT profile setting.
//
// Required Parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettingsByDpKey(ctx context.Context, dpKey string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpProfileSettingsByDpKey, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("dpKey", dpKey)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindDpProfileSettings
//
// Use this API command to retrieve DP DHCP & NAT profile setting list.
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettings(ctx context.Context, mutators ...RequestMutator) (*WSGDPProfileSettingBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileSettingBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpProfileSettings, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileSettingBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpProfileSettingsByDpKey
//
// Use this API command to retrieve DP DHCP & NAT profile setting.
//
// Required Parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettingsByDpKey(ctx context.Context, dpKey string, mutators ...RequestMutator) (*WSGDPProfileSettingBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileSettingBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpProfileSettingsByDpKey, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("dpKey", dpKey)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileSettingBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpProfileSettingsByDpKey
//
// Use this API command to modify DP DHCP & NAT profile setting.
//
// Request Body:
//	 - body *WSGDPProfileSettingBO
//
// Required Parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) UpdateDpProfileSettingsByDpKey(ctx context.Context, body *WSGDPProfileSettingBO, dpKey string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpProfileSettingsByDpKey, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("dpKey", dpKey)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
