package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/gdpr"
	"gopkg.in/go-playground/validator.v9"
)

type WSGGDPRService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGGDPRService(c *APIClient) *WSGGDPRService {
	s := new(WSGGDPRService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGGDPRService() *WSGGDPRService {
	return NewWSGGDPRService(ss.apiClient)
}

// AddGdprReport
//
// Use this API command to execute a client-related data search or delete task and upload a report to FTP. Also use this API to check task progress or to interrupt it.
//
// Request Body:
//	 - body *gdpr.Report
func (s *WSGGDPRService) AddGdprReport(ctx context.Context, body *gdpr.Report) error {
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
