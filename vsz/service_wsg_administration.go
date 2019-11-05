package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAdministrationService struct {
	apiClient *APIClient
}

func NewWSGAdministrationService(c *APIClient) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	serv := new(WSGAdministrationService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddAdminaaa
//
// Use this API command to create a new Admin AAA server
//
// Request Body:
//	 - body *administration.CreateAdminAAAServer
func (s *WSGAdministrationService) AddAdminaaa(ctx context.Context, body *administration.CreateAdminAAAServer) (*common.CreateResult, error) {
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

// DeleteAdminaaaById
//
// Use this API command to delete an existing Admin AAA server
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) DeleteAdminaaaById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
//
// Query Parameters:
// - qType string
//		- required
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, qType string) (*administration.RetrieveAdminAAAServerList, error) {
}

// FindAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
//
// Path Parameters:
// - pId string
//		- required
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

// UpdateAdminaaaById
//
// Use this API command to modify an existing Admin AAA server
//
// Request Body:
//	 - body *administration.ModifyAdminAAAServer
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) UpdateAdminaaaById(ctx context.Context, body *administration.ModifyAdminAAAServer, pId string) (*common.EmptyResult, error) {
}

// UpdateLicenseServer
//
// Use this API command to update license server configuration.
//
// Request Body:
//	 - body *administration.ModfiyLicenseServer
func (s *WSGAdministrationService) UpdateLicenseServer(ctx context.Context, body *administration.ModfiyLicenseServer) error {
}

// UpdateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context) error {
}
