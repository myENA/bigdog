package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/eventmanagement"
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
	serv := new(WSGEventManagementSettingService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindRkszonesEventEmailSettingsByZoneId
//
// Get Event E-mail Setting of Zone Override.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventEmailSettingsByZoneId(ctx context.Context, pZoneId string) (*eventmanagement.EventEmailSetting, error) {
}

// FindRkszonesEventNotificationSettingsByZoneId
//
// Get Event Notification Setting of Zone Override.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventNotificationSettingsByZoneId(ctx context.Context, pZoneId string) (*eventmanagement.EventDataResponse, error) {
}

// UpdateRkszonesEventEmailSettingsByZoneId
//
// Modify Event E-mail Setting of Zone Override.
//
// Request Body:
//	 - body *eventmanagement.EventEmailSetting
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventEmailSettingsByZoneId(ctx context.Context, body *eventmanagement.EventEmailSetting, pZoneId string) (*common.EmptyResult, error) {
}

// UpdateRkszonesEventNotificationSettingsByZoneId
//
// Modify Event Notification Setting of Zone Override.
//
// Request Body:
//	 - body eventmanagement.EventSettingList
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventNotificationSettingsByZoneId(ctx context.Context, body eventmanagement.EventSettingList, pZoneId string) (*common.EmptyResult, error) {
}
