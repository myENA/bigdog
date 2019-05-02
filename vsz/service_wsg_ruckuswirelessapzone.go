package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
)

type WSGRuckusWirelessAPZoneService struct {
    c *Client
}

func NewWSGRuckusWirelessAPZoneService (c *Client) *WSGRuckusWirelessAPZoneService {
    s := new(WSGRuckusWirelessAPZoneService)
    s.c = c
    return s
}

