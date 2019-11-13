package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGUploadStatisticstoFTPService struct {
	apiClient *APIClient
}

func NewWSGUploadStatisticstoFTPService(c *APIClient) *WSGUploadStatisticstoFTPService {
	s := new(WSGUploadStatisticstoFTPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGUploadStatisticstoFTPService() *WSGUploadStatisticstoFTPService {
	return NewWSGUploadStatisticstoFTPService(ss.apiClient)
}

// FindGlobalSettingsStatsFtp
//
// Use this API command to retrieve the uploading statistical data to FTP server setting.
func (s *WSGUploadStatisticstoFTPService) FindGlobalSettingsStatsFtp(ctx context.Context) (*WSGSystemFtpGlobalSetting, error) {
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
//	 - body *WSGSystemFtpGlobalSetting
func (s *WSGUploadStatisticstoFTPService) PartialUpdateGlobalSettingsStatsFtp(ctx context.Context, body *WSGSystemFtpGlobalSetting) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
