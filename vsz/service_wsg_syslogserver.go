package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/syslog"
	"gopkg.in/go-playground/validator.v9"
)

type WSGSyslogServerService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGSyslogServerService(c *APIClient) *WSGSyslogServerService {
	s := new(WSGSyslogServerService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGSyslogServerService() *WSGSyslogServerService {
	return NewWSGSyslogServerService(ss.apiClient)
}

// FindSystemSyslog
//
// Retrieve syslog server sertting.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGSyslogServerService) FindSystemSyslog(ctx context.Context, qIndex string, qListSize string) (*syslog.SyslogServerSetting, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateSystemSyslog
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *syslog.ModifySyslogSettings
func (s *WSGSyslogServerService) PartialUpdateSystemSyslog(ctx context.Context, body *syslog.ModifySyslogSettings) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// PartialUpdateSystemSyslogPrimaryServer
//
// Modify Primary Server of syslog.
//
// Request Body:
//	 - body *syslog.PrimaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPrimaryServer(ctx context.Context, body *syslog.PrimaryServer) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// PartialUpdateSystemSyslogPriority
//
// Modify Priority of syslog.
//
// Request Body:
//	 - body *syslog.Priority
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPriority(ctx context.Context, body *syslog.Priority) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// PartialUpdateSystemSyslogSecondaryServer
//
// Modify Secondary Server of syslog.
//
// Request Body:
//	 - body *syslog.SecondaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogSecondaryServer(ctx context.Context, body *syslog.SecondaryServer) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}
