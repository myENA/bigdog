package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGTestAAAServerService struct {
    client *Client
}

func NewWSGTestAAAServerService (client *Client) *WSGTestAAAServerService {
    s := new(WSGTestAAAServerService)
    s.client = client
    return s
}

func (ss *WSGService) WSGTestAAAServerService () *WSGTestAAAServerService {
    serv := new(WSGTestAAAServerService)
    serv.client = ss.client
    return serv
}

