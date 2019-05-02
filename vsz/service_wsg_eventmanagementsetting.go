package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/eventmanagement"
)

type WSGEventManagementSettingService struct {
    client *Client
}

func NewWSGEventManagementSettingService (client *Client) *WSGEventManagementSettingService {
    s := new(WSGEventManagementSettingService)
    s.client = client
    return s
}

func (ss *WSGService) WSGEventManagementSettingService () *WSGEventManagementSettingService {
    serv := new(WSGEventManagementSettingService)
    serv.client = ss.client
    return serv
}

func (s *WSGEventManagementSettingService) FindRkszonesEventEmailSettingsByZoneId (ctx context.Context, zoneId string) (eventmanagement.EventEmailSetting, error) {
}

func (s *WSGEventManagementSettingService) FindRkszonesEventNotificationSettingsByZoneId (ctx context.Context, zoneId string) (eventmanagement.EventDataResponse, error) {
}
