package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAccountingServiceService struct {
    client *Client
}

func NewWSGAccountingServiceService (client *Client) *WSGAccountingServiceService {
    s := new(WSGAccountingServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccountingServiceService () *WSGAccountingServiceService {
    serv := new(WSGAccountingServiceService)
    serv.client = ss.client
    return serv
}

