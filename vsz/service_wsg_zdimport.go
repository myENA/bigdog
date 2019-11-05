package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZDImportService struct {
	apiClient *APIClient
}

func NewWSGZDImportService(c *APIClient) *WSGZDImportService {
	s := new(WSGZDImportService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZDImportService() *WSGZDImportService {
	serv := new(WSGZDImportService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddZdImportConnectZD
//
// Connect to ZD.
//
// Request Body:
//	 - body *administration.ZdImport
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *administration.ZdImport) (*common.EmptyResult, error) {
}

// AddZdImportMigrate
//
// Migrate ZD to SCG.
//
// Request Body:
//	 - body *administration.ZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *administration.ZdImport) (*common.EmptyResult, error) {
}

// FindZdImportGetZDAPs
//
// Get ZD AP.
//
// Query Parameters:
// - qIp string
//		- required
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, qIp string) (*administration.ZdAPList, error) {
}

// FindZdImportStatus
//
// Get Migrate Status.
//
// Query Parameters:
// - qDetails string
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, qDetails string) (*administration.ZdImportStatus, error) {
}
