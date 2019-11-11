package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGFtpServerSettingsService struct {
	apiClient *APIClient
}

func NewWSGFtpServerSettingsService(c *APIClient) *WSGFtpServerSettingsService {
	s := new(WSGFtpServerSettingsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFtpServerSettingsService() *WSGFtpServerSettingsService {
	return NewWSGFtpServerSettingsService(ss.apiClient)
}

// AddFtps
//
// Add FTP server.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFtpServerSettingsService) AddFtps(ctx context.Context, body *WSGSystemFtp) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteFtps
//
// Remove FTP servers.
//
// Request Body:
//	 - body *WSGSystemDeleteBulkFtp
func (s *WSGFtpServerSettingsService) DeleteFtps(ctx context.Context, body *WSGSystemDeleteBulkFtp) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteFtpsByFtpId
//
// Remove FTP server.
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) DeleteFtpsByFtpId(ctx context.Context, pFtpId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindFtpsByFtpId
//
// Retrieve information of specific FTP server.
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) FindFtpsByFtpId(ctx context.Context, pFtpId string) (*WSGSystemFtp, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFtpServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemFtpList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFtpServerSettingsService) FindFtpsTest(ctx context.Context, body *WSGSystemFtp) (*WSGSystemFtpTestResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindFtpsTestByFtpId
//
// Test ftp server of specific FTP server id.
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) FindFtpsTestByFtpId(ctx context.Context, pFtpId string) (*WSGSystemFtpTestResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateFtpsByFtpId
//
// Update FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) PartialUpdateFtpsByFtpId(ctx context.Context, body *WSGSystemFtp, pFtpId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
