package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/alarmlist"
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

