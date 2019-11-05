package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"gopkg.in/go-playground/validator.v9"
)

type WSGApplicationLogAndStatusService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGApplicationLogAndStatusService(c *APIClient) *WSGApplicationLogAndStatusService {
	s := new(WSGApplicationLogAndStatusService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGApplicationLogAndStatusService() *WSGApplicationLogAndStatusService {
	return NewWSGApplicationLogAndStatusService(ss.apiClient)
}

// FindApplicationsByBladeUUID
//
// Use this API command to retrieve a list of application log and status.
//
// Path Parameters:
// - pBladeUUID string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, pBladeUUID string, qIndex string, qListSize string) (*administration.ApplicationLogAndStatusList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApplicationsDownloadByBladeUUID
//
// Use this API command to download logs of the application.
//
// Path Parameters:
// - pBladeUUID string
//		- required
//
// Query Parameters:
// - qAppName string
//		- required
// - qLogFileName string
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, pBladeUUID string, qAppName string, qLogFileName string) (json.RawMessage, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApplicationsDownloadsnapByBladeUUID
//
// Use this API command to download snapshot logs.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, pBladeUUID string) (json.RawMessage, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Request Body:
//	 - body *administration.ModifyLogLevel
func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *administration.ModifyLogLevel) error {
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
