package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventEmailSettingsByZoneId(ctx context.Context, pZoneId string) (*WSGEventManagementEventEmailSetting, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesEventNotificationSettingsByZoneId
//
// Get Event Notification Setting of Zone Override.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) FindRkszonesEventNotificationSettingsByZoneId(ctx context.Context, pZoneId string) (*WSGEventManagementEventDataResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesEventEmailSettingsByZoneId
//
// Modify Event E-mail Setting of Zone Override.
//
// Request Body:
//	 - body *WSGEventManagementEventEmailSetting
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventEmailSettingsByZoneId(ctx context.Context, body *WSGEventManagementEventEmailSetting, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesEventNotificationSettingsByZoneId
//
// Modify Event Notification Setting of Zone Override.
//
// Request Body:
//	 - body WSGEventManagementEventSettingList
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGEventManagementSettingService) UpdateRkszonesEventNotificationSettingsByZoneId(ctx context.Context, body WSGEventManagementEventSettingList, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
