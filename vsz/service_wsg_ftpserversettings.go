package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGFtpServerSettingsService struct {
	client *Client
}

func NewWSGFtpServerSettingsService(client *Client) *WSGFtpServerSettingsService {
	s := new(WSGFtpServerSettingsService)
	s.client = client
	return s
}

func (ss *WSGService) WSGFtpServerSettingsService() *WSGFtpServerSettingsService {
	serv := new(WSGFtpServerSettingsService)
	serv.client = ss.client
	return serv
}

// AddFtps
//
// Add FTP server.
//
// Request Body:
//	 - body *system.Ftp
func (s *WSGFtpServerSettingsService) AddFtps(ctx context.Context, body *system.Ftp) error {
}

// FindFtpsByFtpId
//
// Retrieve information of specific FTP server.
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) FindFtpsByFtpId(ctx context.Context, pFtpId string) (*system.Ftp, error) {
}

// FindFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGFtpServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*system.FtpList, error) {
}

// FindFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Request Body:
//	 - body *system.Ftp
func (s *WSGFtpServerSettingsService) FindFtpsTest(ctx context.Context, body *system.Ftp) (*system.FtpTestResponse, error) {
}

// FindFtpsTestByFtpId
//
// Test ftp server of specific FTP server id.
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) FindFtpsTestByFtpId(ctx context.Context, pFtpId string) (*system.FtpTestResponse, error) {
}
