package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
	"gopkg.in/go-playground/validator.v9"
)

type WSGUploadStatisticstoFTPService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGUploadStatisticstoFTPService(c *APIClient) *WSGUploadStatisticstoFTPService {
	s := new(WSGUploadStatisticstoFTPService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGUploadStatisticstoFTPService() *WSGUploadStatisticstoFTPService {
	return NewWSGUploadStatisticstoFTPService(ss.apiClient)
}

// FindGlobalSettingsStatsFtp
//
// Use this API command to retrieve the uploading statistical data to FTP server setting.
func (s *WSGUploadStatisticstoFTPService) FindGlobalSettingsStatsFtp(ctx context.Context) (*system.FtpGlobalSetting, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateGlobalSettingsStatsFtp
//
// Use this API command to modify the setting of uploading statistical data to FTP server.
//
// Request Body:
//	 - body *system.FtpGlobalSetting
func (s *WSGUploadStatisticstoFTPService) PartialUpdateGlobalSettingsStatsFtp(ctx context.Context, body *system.FtpGlobalSetting) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
