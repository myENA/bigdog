package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGEventandAlarmService struct {
	apiClient *APIClient
}

func NewWSGEventandAlarmService(c *APIClient) *WSGEventandAlarmService {
	s := new(WSGEventandAlarmService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEventandAlarmService() *WSGEventandAlarmService {
	return NewWSGEventandAlarmService(ss.apiClient)
}

// AddAlertAlarmList
//
// Query Alarms with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertAlarmList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAlarmListAlarmQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddAlertAlarmSummary
//
// Use this API command to retrieve the alarm summary with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertAlarmSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAlertSummaryAlarmSummary, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddAlertEventList
//
// Query Events with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertEventList(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGEventListEventQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddAlertEventSummary
//
// Use this API command to retrieve the event summary with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertEventSummary(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAlertSummaryEventSummary, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAlertAlarmAck
//
// Acknowledge multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *WSGAlertAckBulkAlarms
func (s *WSGEventandAlarmService) UpdateAlertAlarmAck(ctx context.Context, body *WSGAlertAckBulkAlarms) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAlertAlarmAckByAlarmID
//
// Acknowledge a single Alarm with provided alarmID.
//
// Path Parameters:
// - pAlarmID string
//		- required
func (s *WSGEventandAlarmService) UpdateAlertAlarmAckByAlarmID(ctx context.Context, pAlarmID string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAlertAlarmClear
//
// Clear multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *WSGAlertClearBulkAlarms
func (s *WSGEventandAlarmService) UpdateAlertAlarmClear(ctx context.Context, body *WSGAlertClearBulkAlarms) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAlertAlarmClearByAlarmID
//
// Clear a single Alarm with provided alarmID.
//
// Path Parameters:
// - pAlarmID string
//		- required
func (s *WSGEventandAlarmService) UpdateAlertAlarmClearByAlarmID(ctx context.Context, pAlarmID string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
