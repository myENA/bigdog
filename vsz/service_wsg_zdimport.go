package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZDImportService struct {
    c *Client
}

func NewWSGZDImportService (c *Client) *WSGZDImportService {
    s := new(WSGZDImportService)
    s.c = c
    return s
}

