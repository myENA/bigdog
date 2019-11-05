package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/syslog"
)

type WSGSyslogServerService struct {
	apiClient *APIClient
}

func NewWSGSyslogServerService(c *APIClient) *WSGSyslogServerService {
	s := new(WSGSyslogServerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSyslogServerService() *WSGSyslogServerService {
	serv := new(WSGSyslogServerService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindSystemSyslog
//
// Retrieve syslog server sertting.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGSyslogServerService) FindSystemSyslog(ctx context.Context, qIndex string, qListSize string) (*syslog.SyslogServerSetting, error) {
}

// PartialUpdateSystemSyslog
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *syslog.ModifySyslogSettings
func (s *WSGSyslogServerService) PartialUpdateSystemSyslog(ctx context.Context, body *syslog.ModifySyslogSettings) error {
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
