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

func NewWSGEventandAlarmService (client *Client) *WSGEventandAlarmService {
    s := new(WSGEventandAlarmService)
    s.client = client
    return s
}

func (ss *WSGService) WSGEventandAlarmService () *WSGEventandAlarmService {
    serv := new(WSGEventandAlarmService)
    serv.client = ss.client
    return serv
}

func (s *WSGEventandAlarmService) AddAlertAlarmList (ctx context.Context) (*alarmlist.AlarmQueryResultList, error) {
}

func (s *WSGEventandAlarmService) AddAlertAlarmSummary (ctx context.Context) (*alertsummary.AlarmSummary, error) {
}

func (s *WSGEventandAlarmService) AddAlertEventList (ctx context.Context) (*eventlist.EventQueryResultList, error) {
}

func (s *WSGEventandAlarmService) AddAlertEventSummary (ctx context.Context) (*alertsummary.EventSummary, error) {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmAck (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmAckByAlarmID (ctx context.Context, alarmID string) error {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmClear (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmClearByAlarmID (ctx context.Context, alarmID string) error {
}

