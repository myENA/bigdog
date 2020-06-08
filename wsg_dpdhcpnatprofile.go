package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDpdhcpnatprofileService struct {
	apiClient *VSZClient
}

func NewWSGDpdhcpnatprofileService(c *VSZClient) *WSGDpdhcpnatprofileService {
	s := new(WSGDpdhcpnatprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDpdhcpnatprofileService() *WSGDpdhcpnatprofileService {
	return NewWSGDpdhcpnatprofileService(ss.apiClient)
}

// AddDpProfileSettings
//
// Use this API command to create DP DHCP & NAT profile setting.
//
// Request Body:
//	 - body *WSGDPProfileSettingBO
func (s *WSGDpdhcpnatprofileService) AddDpProfileSettings(ctx context.Context, body *WSGDPProfileSettingBO) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteDpProfileSettings
//
// Use this API command to delete DP DHCP & NAT profile settings.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDpdhcpnatprofileService) DeleteDpProfileSettings(ctx context.Context, body *WSGDPProfileBulkDelete) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *WSGDpdhcpnatprofileService) DeleteDpProfileSettingsByDpKey(ctx context.Context, dpKey string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("dpKey", dpKey)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindDpProfileSettings
//
// Use this API command to retrieve DP DHCP & NAT profile setting list.
func (s *WSGDpdhcpnatprofileService) FindDpProfileSettings(ctx context.Context) (*WSGDPProfileSettingBOList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *WSGDpdhcpnatprofileService) FindDpProfileSettingsByDpKey(ctx context.Context, dpKey string) (*WSGDPProfileSettingBO, *APIResponseMeta, error) {
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
	req.SetPathParameter("dpKey", dpKey)
	httpResp, err = s.apiClient.Do(ctx, req)
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
func (s *WSGDpdhcpnatprofileService) UpdateDpProfileSettingsByDpKey(ctx context.Context, body *WSGDPProfileSettingBO, dpKey string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("dpKey", dpKey)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
