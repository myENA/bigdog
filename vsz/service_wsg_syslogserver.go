package vsz

// API Version: v8_0

import (
	"context"
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

func (s *WSGSyslogServerService) FindSystemSyslog (ctx context.Context) (*syslog.SyslogServerSetting, error) {
}

func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPrimaryServer (ctx context.Context) error {
}

func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPriority (ctx context.Context) error {
}

func (s *WSGSyslogServerService) PartialUpdateSystemSyslogSecondaryServer (ctx context.Context) error {
}

