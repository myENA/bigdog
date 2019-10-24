package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/syslog"
)

type WSGSyslogServerService struct {
	client *Client
}

func NewWSGSyslogServerService(client *Client) *WSGSyslogServerService {
	s := new(WSGSyslogServerService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSyslogServerService() *WSGSyslogServerService {
	serv := new(WSGSyslogServerService)
	serv.client = ss.client
	return serv
}

// FindSystemSyslog
//
// Retrieve syslog server sertting.
func (s *WSGSyslogServerService) FindSystemSyslog(ctx context.Context, qIndex string, qListSize string) (*syslog.SyslogServerSetting, error) {
}

// PartialUpdateSystemSyslogPrimaryServer
//
// Modify Primary Server of syslog.
//
// Request Body:
//	 - body *syslog.PrimaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPrimaryServer(ctx context.Context, body *syslog.PrimaryServer) error {
}

// PartialUpdateSystemSyslogPriority
//
// Modify Priority of syslog.
//
// Request Body:
//	 - body *syslog.Priority
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPriority(ctx context.Context, body *syslog.Priority) error {
}

// PartialUpdateSystemSyslogSecondaryServer
//
// Modify Secondary Server of syslog.
//
// Request Body:
//	 - body *syslog.SecondaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogSecondaryServer(ctx context.Context, body *syslog.SecondaryServer) error {
}
