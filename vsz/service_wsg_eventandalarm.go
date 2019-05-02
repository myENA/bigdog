package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/alarmlist"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/eventlist"
)

type WSGEventandAlarmService struct {
    c *Client
}

func NewWSGEventandAlarmService (c *Client) *WSGEventandAlarmService {
    s := new(WSGEventandAlarmService)
    s.c = c
    return s
}

