package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
)

type WSGRuckusWirelessAPZoneService struct {
    client *Client
}

func NewWSGRuckusWirelessAPZoneService (client *Client) *WSGRuckusWirelessAPZoneService {
    s := new(WSGRuckusWirelessAPZoneService)
    s.client = client
    return s
}

func (ss *WSGService) WSGRuckusWirelessAPZoneService () *WSGRuckusWirelessAPZoneService {
    serv := new(WSGRuckusWirelessAPZoneService)
    serv.client = ss.client
    return serv
}

