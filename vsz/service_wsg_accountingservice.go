package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAccountingServiceService struct {
    c *Client
}

func NewWSGAccountingServiceService (c *Client) *WSGAccountingServiceService {
    s := new(WSGAccountingServiceService)
    s.c = c
    return s
}

