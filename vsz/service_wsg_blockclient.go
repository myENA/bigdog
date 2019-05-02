package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBlockClientService struct {
    client *Client
}

func NewWSGBlockClientService (client *Client) *WSGBlockClientService {
    s := new(WSGBlockClientService)
    s.client = client
    return s
}

func (ss *WSGService) WSGBlockClientService () *WSGBlockClientService {
    serv := new(WSGBlockClientService)
    serv.client = ss.client
    return serv
}

