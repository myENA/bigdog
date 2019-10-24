package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZDImportService struct {
	client *Client
}

func NewWSGZDImportService(client *Client) *WSGZDImportService {
	s := new(WSGZDImportService)
	s.client = client
	return s
}

func (ss *WSGService) WSGZDImportService() *WSGZDImportService {
	serv := new(WSGZDImportService)
	serv.client = ss.client
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
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, qIp string) (*administration.ZdAPList, error) {
}

// FindZdImportStatus
//
// Get Migrate Status.
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, qDetails string) (*administration.ZdImportStatus, error) {
}
