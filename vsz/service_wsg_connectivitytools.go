package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGConnectivityToolsService struct {
    client *Client
}

func NewWSGConnectivityToolsService (client *Client) *WSGConnectivityToolsService {
    s := new(WSGConnectivityToolsService)
    s.client = client
    return s
}

func (ss *WSGService) WSGConnectivityToolsService () *WSGConnectivityToolsService {
    serv := new(WSGConnectivityToolsService)
    serv.client = ss.client
    return serv
}

