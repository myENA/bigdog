package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/eventmanagement"
)

type WSGEventManagementSettingService struct {
	client *Client
}

func NewWSGEventManagementSettingService(client *Client) *WSGEventManagementSettingService {
	s := new(WSGEventManagementSettingService)
	s.client = client
	return s
}

func (ss *WSGService) WSGEventManagementSettingService() *WSGEventManagementSettingService {
	serv := new(WSGEventManagementSettingService)
	serv.client = ss.client
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
