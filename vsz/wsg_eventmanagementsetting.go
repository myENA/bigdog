package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGEventManagementSettingService struct {
	apiClient *APIClient
}

func NewWSGEventManagementSettingService(c *APIClient) *WSGEventManagementSettingService {
	s := new(WSGEventManagementSettingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEventManagementSettingService() *WSGEventManagementSettingService {
	return NewWSGEventManagementSettingService(ss.apiClient)
}

// FindRkszonesEventEmailSettingsByZoneId
//
// Get Event E-mail Setting of Zone Override.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventEmailSettingsByZoneId(ctx context.Context, zoneId string) (*WSGEventManagementEventEmailSetting, error) {
	var (
		req      *APIRequest
		resp     *WSGEventManagementEventEmailSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesEventEmailSettingsByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEventManagementEventEmailSetting()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesEventNotificationSettingsByZoneId
//
// Get Event Notification Setting of Zone Override.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventNotificationSettingsByZoneId(ctx context.Context, zoneId string) (*WSGEventManagementEventDataResponse, error) {
	var (
		req      *APIRequest
		resp     *WSGEventManagementEventDataResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesEventNotificationSettingsByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEventManagementEventDataResponse()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateRkszonesEventEmailSettingsByZoneId
//
// Modify Event E-mail Setting of Zone Override.
//
// Request Body:
//	 - body *WSGEventManagementEventEmailSetting
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventEmailSettingsByZoneId(ctx context.Context, body *WSGEventManagementEventEmailSetting, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesEventEmailSettingsByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// UpdateRkszonesEventNotificationSettingsByZoneId
//
// Modify Event Notification Setting of Zone Override.
//
// Request Body:
//	 - body WSGEventManagementEventSettingList
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventNotificationSettingsByZoneId(ctx context.Context, body WSGEventManagementEventSettingList, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesEventNotificationSettingsByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
