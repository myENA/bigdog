package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/syslog"
)

type WSGSyslogServerService struct {
    c *Client
}

func NewWSGSyslogServerService (c *Client) *WSGSyslogServerService {
    s := new(WSGSyslogServerService)
    s.c = c
    return s
}

