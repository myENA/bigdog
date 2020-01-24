package vsz

// API Version: v8_1

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
		req  *APIRequest
		resp *WSGEventManagementEventEmailSetting
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesEventEmailSettingsByZoneId, true)
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
		req  *APIRequest
		resp *WSGEventManagementEventDataResponse
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesEventNotificationSettingsByZoneId, true)
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
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
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
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
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
}
