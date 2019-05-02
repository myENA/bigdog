package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGSMSGatewayService struct {
    c *Client
}

func NewWSGSMSGatewayService (c *Client) *WSGSMSGatewayService {
    s := new(WSGSMSGatewayService)
    s.c = c
    return s
}

