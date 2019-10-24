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

func (s *WSGFtpServerSettingsService) AddFtps(ctx context.Context, body *system.Ftp) error {
}

func (s *WSGFtpServerSettingsService) FindFtpsByFtpId(ctx context.Context, pFtpId string) (*system.Ftp, error) {
}

func (s *WSGFtpServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*system.FtpList, error) {
}

func (s *WSGFtpServerSettingsService) FindFtpsTest(ctx context.Context, body *system.Ftp) (*system.FtpTestResponse, error) {
}

func (s *WSGFtpServerSettingsService) FindFtpsTestByFtpId(ctx context.Context, pFtpId string) (*system.FtpTestResponse, error) {
}
