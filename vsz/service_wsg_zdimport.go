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

func NewWSGZDImportService (client *Client) *WSGZDImportService {
    s := new(WSGZDImportService)
    s.client = client
    return s
}

func (ss *WSGService) WSGZDImportService () *WSGZDImportService {
    serv := new(WSGZDImportService)
    serv.client = ss.client
    return serv
}

func (s *WSGZDImportService) AddZdImportConnectZD (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGZDImportService) AddZdImportMigrate (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGZDImportService) FindZdImportGetZDAPs (ctx context.Context) (*administration.ZdAPList, error) {
}

func (s *WSGZDImportService) FindZdImportStatus (ctx context.Context) (*administration.ZdImportStatus, error) {
}

