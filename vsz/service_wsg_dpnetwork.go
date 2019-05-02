package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGDPNetworkService struct {
    client *Client
}

func NewWSGDPNetworkService (client *Client) *WSGDPNetworkService {
    s := new(WSGDPNetworkService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDPNetworkService () *WSGDPNetworkService {
    serv := new(WSGDPNetworkService)
    serv.client = ss.client
    return serv
}

