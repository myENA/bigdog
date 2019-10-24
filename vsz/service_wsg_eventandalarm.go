package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/alarmlist"
	"github.com/myENA/ruckus-client/vsz/types/wsg/alert"
	"github.com/myENA/ruckus-client/vsz/types/wsg/alertsummary"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/eventlist"
)

type WSGEventandAlarmService struct {
	client *Client
}

func NewWSGEventandAlarmService(client *Client) *WSGEventandAlarmService {
	s := new(WSGEventandAlarmService)
	s.client = client
	return s
}

func (ss *WSGService) WSGEventandAlarmService() *WSGEventandAlarmService {
	serv := new(WSGEventandAlarmService)
	serv.client = ss.client
	return serv
}

// AddAlertAlarmList
//
// Query Alarms with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertAlarmList(ctx context.Context, body *common.QueryCriteriaSuperSet) (*alarmlist.AlarmQueryResultList, error) {
}

// AddAlertAlarmSummary
//
// Use this API command to retrieve the alarm summary with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertAlarmSummary(ctx context.Context, body *common.QueryCriteriaSuperSet) (*alertsummary.AlarmSummary, error) {
}

// AddAlertEventList
//
// Query Events with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertEventList(ctx context.Context, body *common.QueryCriteriaSuperSet) (*eventlist.EventQueryResultList, error) {
}

// AddAlertEventSummary
//
// Use this API command to retrieve the event summary with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGEventandAlarmService) AddAlertEventSummary(ctx context.Context, body *common.QueryCriteriaSuperSet) (*alertsummary.EventSummary, error) {
}

// UpdateAlertAlarmAck
//
// Acknowledge multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *alert.AckBulkAlarms
func (s *WSGEventandAlarmService) UpdateAlertAlarmAck(ctx context.Context, body *alert.AckBulkAlarms) (*common.EmptyResult, error) {
}

// UpdateAlertAlarmAckByAlarmID
//
// Acknowledge a single Alarm with provided alarmID.
func (s *WSGEventandAlarmService) UpdateAlertAlarmAckByAlarmID(ctx context.Context, pAlarmID string) error {
}

// UpdateAlertAlarmClear
//
// Clear multiple Alarms with provided alarmIDs.
//
// Request Body:
//	 - body *alert.ClearBulkAlarms
func (s *WSGEventandAlarmService) UpdateAlertAlarmClear(ctx context.Context, body *alert.ClearBulkAlarms) (*common.EmptyResult, error) {
}

// UpdateAlertAlarmClearByAlarmID
//
// Clear a single Alarm with provided alarmID.
func (s *WSGEventandAlarmService) UpdateAlertAlarmClearByAlarmID(ctx context.Context, pAlarmID string) error {
}
