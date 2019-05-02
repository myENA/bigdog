package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGConnectivityToolsService struct {
    c *Client
}

func NewWSGConnectivityToolsService (c *Client) *WSGConnectivityToolsService {
    s := new(WSGConnectivityToolsService)
    s.c = c
    return s
}

