package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wiredclientquery"
)

type WSGWiredClientService struct {
    c *Client
}

func NewWSGWiredClientService (c *Client) *WSGWiredClientService {
    s := new(WSGWiredClientService)
    s.c = c
    return s
}

