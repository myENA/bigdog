package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGEventmanagementsettingService struct {
	apiClient *VSZClient
}

func NewWSGEventmanagementsettingService(c *VSZClient) *WSGEventmanagementsettingService {
	s := new(WSGEventmanagementsettingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEventmanagementsettingService() *WSGEventmanagementsettingService {
	return NewWSGEventmanagementsettingService(ss.apiClient)
}

// FindRkszonesEventEmailSettingsByZoneId
//
// Get Event E-mail Setting of Zone Override.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEventmanagementsettingService) FindRkszonesEventEmailSettingsByZoneId(ctx context.Context, zoneId string) (*WSGEventManagementEventEmailSetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGEventManagementEventEmailSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesEventEmailSettingsByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEventManagementEventEmailSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesEventNotificationSettingsByZoneId
//
// Get Event Notification Setting of Zone Override.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEventmanagementsettingService) FindRkszonesEventNotificationSettingsByZoneId(ctx context.Context, zoneId string) (*WSGEventManagementEventDataResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGEventManagementEventDataResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesEventNotificationSettingsByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEventManagementEventDataResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGEventmanagementsettingService) UpdateRkszonesEventEmailSettingsByZoneId(ctx context.Context, body *WSGEventManagementEventEmailSetting, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesEventEmailSettingsByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGEventmanagementsettingService) UpdateRkszonesEventNotificationSettingsByZoneId(ctx context.Context, body WSGEventManagementEventSettingList, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesEventNotificationSettingsByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
