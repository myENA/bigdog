package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGUploadStatisticstoFTPService struct {
	client *Client
}

func NewWSGUploadStatisticstoFTPService(client *Client) *WSGUploadStatisticstoFTPService {
	s := new(WSGUploadStatisticstoFTPService)
	s.client = client
	return s
}

func (ss *WSGService) WSGUploadStatisticstoFTPService() *WSGUploadStatisticstoFTPService {
	serv := new(WSGUploadStatisticstoFTPService)
	serv.client = ss.client
	return serv
}

// FindGlobalSettingsStatsFtp
//
// Use this API command to retrieve the uploading statistical data to FTP server setting.
func (s *WSGUploadStatisticstoFTPService) FindGlobalSettingsStatsFtp(ctx context.Context) (*system.FtpGlobalSetting, error) {
}
