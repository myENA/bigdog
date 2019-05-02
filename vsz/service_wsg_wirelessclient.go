package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clientquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGWirelessClientService struct {
    client *Client
}

func NewWSGWirelessClientService (client *Client) *WSGWirelessClientService {
    s := new(WSGWirelessClientService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWirelessClientService () *WSGWirelessClientService {
    serv := new(WSGWirelessClientService)
    serv.client = ss.client
    return serv
}

