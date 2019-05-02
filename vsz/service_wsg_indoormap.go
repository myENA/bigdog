package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
)

type WSGIndoorMapService struct {
    client *Client
}

func NewWSGIndoorMapService (client *Client) *WSGIndoorMapService {
    s := new(WSGIndoorMapService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIndoorMapService () *WSGIndoorMapService {
    serv := new(WSGIndoorMapService)
    serv.client = ss.client
    return serv
}

