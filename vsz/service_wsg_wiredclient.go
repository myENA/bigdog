package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wiredclientquery"
)

type WSGWiredClientService struct {
    client *Client
}

func NewWSGWiredClientService (client *Client) *WSGWiredClientService {
    s := new(WSGWiredClientService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWiredClientService () *WSGWiredClientService {
    serv := new(WSGWiredClientService)
    serv.client = ss.client
    return serv
}

