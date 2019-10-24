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

// AddRestart
//
// Use this API command to restart the controller.
func (s *WSGAdministrationService) AddRestart(ctx context.Context) error {
}

// AddShutdown
//
// Use this API command to shut down the controller.
func (s *WSGAdministrationService) AddShutdown(ctx context.Context) error {
}

// FindAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, qType string) (*administration.RetrieveAdminAAAServerList, error) {
}

// FindAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, pId string) (*administration.RetrieveAdminAAAServer, error) {
}

// FindLicenses
//
// Use this API command to get all licenses currently assign in SCG.
func (s *WSGAdministrationService) FindLicenses(ctx context.Context) (*administration.LicensesList, error) {
}

// FindLicenseServer
//
// Use this API command to get license server configuration.
func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context) (*administration.LicenseServer, error) {
}

// FindLicensesSummary
//
// Use this API command to get licenses summary information.
func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context) (*administration.LicensesSummaryList, error) {
}

// FindLicensesSyncLogs
//
// Use this API command to get licenses synchronize logs.
func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context) (*administration.LicensesSyncLogsList, error) {
}

// UpdateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context) error {
}
