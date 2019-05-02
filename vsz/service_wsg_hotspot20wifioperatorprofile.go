package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGHotspot20WiFiOperatorProfileService struct {
    c *Client
}

func NewWSGHotspot20WiFiOperatorProfileService (c *Client) *WSGHotspot20WiFiOperatorProfileService {
    s := new(WSGHotspot20WiFiOperatorProfileService)
    s.c = c
    return s
}

