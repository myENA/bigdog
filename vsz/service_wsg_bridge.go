package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBridgeService struct {
    client *Client
}

func NewWSGBridgeService (client *Client) *WSGBridgeService {
    s := new(WSGBridgeService)
    s.client = client
    return s
}

func (ss *WSGService) WSGBridgeService () *WSGBridgeService {
    serv := new(WSGBridgeService)
    serv.client = ss.client
    return serv
}

