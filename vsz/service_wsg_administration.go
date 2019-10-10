package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
)

type WSGAdministrationService struct {
	client *Client
}

func NewWSGAdministrationService(client *Client) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	serv := new(WSGAdministrationService)
	serv.client = ss.client
	return serv
}

func (s *WSGAdministrationService) AddRestart(ctx context.Context) error {
}

func (s *WSGAdministrationService) AddShutdown(ctx context.Context) error {
}

func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context) (*administration.RetrieveAdminAAAServerList, error) {
}

func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, id string) (*administration.RetrieveAdminAAAServer, error) {
}

func (s *WSGAdministrationService) FindLicenses(ctx context.Context) (*administration.LicensesList, error) {
}

func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context) (*administration.LicenseServer, error) {
}

func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context) (*administration.LicensesSummaryList, error) {
}

func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context) (*administration.LicensesSyncLogsList, error) {
}

func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context) error {
}
