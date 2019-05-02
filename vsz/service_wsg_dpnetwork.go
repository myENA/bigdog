package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGDPNetworkService struct {
    c *Client
}

func NewWSGDPNetworkService (c *Client) *WSGDPNetworkService {
    s := new(WSGDPNetworkService)
    s.c = c
    return s
}

