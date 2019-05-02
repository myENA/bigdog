package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
)

type WSGAccessPointAPPService struct {
    client *Client
}

func NewWSGAccessPointAPPService (client *Client) *WSGAccessPointAPPService {
    s := new(WSGAccessPointAPPService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccessPointAPPService () *WSGAccessPointAPPService {
    serv := new(WSGAccessPointAPPService)
    serv.client = ss.client
    return serv
}

