package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGWechatService struct {
    client *Client
}

func NewWSGWechatService (client *Client) *WSGWechatService {
    s := new(WSGWechatService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWechatService () *WSGWechatService {
    serv := new(WSGWechatService)
    serv.client = ss.client
    return serv
}

