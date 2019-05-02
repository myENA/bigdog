package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZoneAAAService struct {
    client *Client
}

func NewWSGZoneAAAService (client *Client) *WSGZoneAAAService {
    s := new(WSGZoneAAAService)
    s.client = client
    return s
}

func (ss *WSGService) WSGZoneAAAService () *WSGZoneAAAService {
    serv := new(WSGZoneAAAService)
    serv.client = ss.client
    return serv
}

