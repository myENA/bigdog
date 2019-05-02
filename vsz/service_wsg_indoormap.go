package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
)

type WSGIndoorMapService struct {
    c *Client
}

func NewWSGIndoorMapService (c *Client) *WSGIndoorMapService {
    s := new(WSGIndoorMapService)
    s.c = c
    return s
}

