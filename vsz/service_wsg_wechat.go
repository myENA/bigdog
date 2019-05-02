package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGWechatService struct {
    c *Client
}

func NewWSGWechatService (c *Client) *WSGWechatService {
    s := new(WSGWechatService)
    s.c = c
    return s
}

