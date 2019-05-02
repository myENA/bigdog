package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clientquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGWirelessClientService struct {
    c *Client
}

func NewWSGWirelessClientService (c *Client) *WSGWirelessClientService {
    s := new(WSGWirelessClientService)
    s.c = c
    return s
}

