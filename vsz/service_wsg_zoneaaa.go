package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZoneAAAService struct {
    c *Client
}

func NewWSGZoneAAAService (c *Client) *WSGZoneAAAService {
    s := new(WSGZoneAAAService)
    s.c = c
    return s
}

