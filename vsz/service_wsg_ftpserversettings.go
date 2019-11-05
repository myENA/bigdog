package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
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
	serv := new(WSGFtpServerSettingsService)
	serv.apiClient = ss.apiClient
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

// DeleteFtps
//
// Remove FTP servers.
//
// Request Body:
//	 - body *system.DeleteBulkFtp
func (s *WSGFtpServerSettingsService) DeleteFtps(ctx context.Context, body *system.DeleteBulkFtp) error {
}

// DeleteFtpsByFtpId
//
// Remove FTP server.
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) DeleteFtpsByFtpId(ctx context.Context, pFtpId string) error {
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

// PartialUpdateFtpsByFtpId
//
// Update FTP server settings.
//
// Request Body:
//	 - body *system.Ftp
//
// Path Parameters:
// - pFtpId string
//		- required
func (s *WSGFtpServerSettingsService) PartialUpdateFtpsByFtpId(ctx context.Context, body *system.Ftp, pFtpId string) error {
}
