package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/syslog"
)

type WSGSyslogServerService struct {
    client *Client
}

func NewWSGSyslogServerService (client *Client) *WSGSyslogServerService {
    s := new(WSGSyslogServerService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSyslogServerService () *WSGSyslogServerService {
    serv := new(WSGSyslogServerService)
    serv.client = ss.client
    return serv
}

