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

func (s *WSGEventandAlarmService) AddAlertAlarmList(ctx context.Context, body *common.QueryCriteriaSuperSet) (*alarmlist.AlarmQueryResultList, error) {
}

func (s *WSGEventandAlarmService) AddAlertAlarmSummary(ctx context.Context, body *common.QueryCriteriaSuperSet) (*alertsummary.AlarmSummary, error) {
}

func (s *WSGEventandAlarmService) AddAlertEventList(ctx context.Context, body *common.QueryCriteriaSuperSet) (*eventlist.EventQueryResultList, error) {
}

func (s *WSGEventandAlarmService) AddAlertEventSummary(ctx context.Context, body *common.QueryCriteriaSuperSet) (*alertsummary.EventSummary, error) {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmAck(ctx context.Context, body *alert.AckBulkAlarms) (*common.EmptyResult, error) {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmAckByAlarmID(ctx context.Context, pAlarmID string) error {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmClear(ctx context.Context, body *alert.ClearBulkAlarms) (*common.EmptyResult, error) {
}

func (s *WSGEventandAlarmService) UpdateAlertAlarmClearByAlarmID(ctx context.Context, pAlarmID string) error {
}
