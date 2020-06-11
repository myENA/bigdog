package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGEventManagementSettingService struct {
	apiClient *VSZClient
}

func NewWSGEventManagementSettingService(c *VSZClient) *WSGEventManagementSettingService {
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
func (s *WSGEventManagementSettingService) FindRkszonesEventEmailSettingsByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGEventManagementEventEmailSetting, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGEventManagementSettingService) FindRkszonesEventNotificationSettingsByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGEventManagementEventDataResponse, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGEventManagementSettingService) UpdateRkszonesEventEmailSettingsByZoneId(ctx context.Context, body *WSGEventManagementEventEmailSetting, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGEventManagementSettingService) UpdateRkszonesEventNotificationSettingsByZoneId(ctx context.Context, body WSGEventManagementEventSettingList, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
