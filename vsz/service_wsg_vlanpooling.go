package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vlanpooling"
)

type WSGVlanPoolingService struct {
    client *Client
}

func NewWSGVlanPoolingService (client *Client) *WSGVlanPoolingService {
    s := new(WSGVlanPoolingService)
    s.client = client
    return s
}

func (ss *WSGService) WSGVlanPoolingService () *WSGVlanPoolingService {
    serv := new(WSGVlanPoolingService)
    serv.client = ss.client
    return serv
}

