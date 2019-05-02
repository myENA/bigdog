package vsz

// API Version: v8_0

import (
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

