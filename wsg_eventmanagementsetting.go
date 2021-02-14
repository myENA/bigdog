package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Operation ID: findRkszonesEventEmailSettingsByZoneId
// Operation path: /rkszones/{zoneId}/eventEmailSettings
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventEmailSettingsByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGEventManagementEventEmailSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGEventManagementEventEmailSettingAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesEventEmailSettingsByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGEventManagementEventEmailSettingAPIResponse), err
}

// FindRkszonesEventNotificationSettingsByZoneId
//
// Get Event Notification Setting of Zone Override.
//
// Operation ID: findRkszonesEventNotificationSettingsByZoneId
// Operation path: /rkszones/{zoneId}/eventNotificationSettings
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventNotificationSettingsByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGEventManagementEventDataResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGEventManagementEventDataResponseAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesEventNotificationSettingsByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGEventManagementEventDataResponseAPIResponse), err
}

// UpdateRkszonesEventEmailSettingsByZoneId
//
// Modify Event E-mail Setting of Zone Override.
//
// Operation ID: updateRkszonesEventEmailSettingsByZoneId
// Operation path: /rkszones/{zoneId}/eventEmailSettings
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGEventManagementEventEmailSetting
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventEmailSettingsByZoneId(ctx context.Context, body *WSGEventManagementEventEmailSetting, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesEventEmailSettingsByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesEventNotificationSettingsByZoneId
//
// Modify Event Notification Setting of Zone Override.
//
// Operation ID: updateRkszonesEventNotificationSettingsByZoneId
// Operation path: /rkszones/{zoneId}/eventNotificationSettings
// Success code: 204 (No Content)
//
// Request body:
//	 - body WSGEventManagementEventSettingList
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventNotificationSettingsByZoneId(ctx context.Context, body WSGEventManagementEventSettingList, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesEventNotificationSettingsByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
